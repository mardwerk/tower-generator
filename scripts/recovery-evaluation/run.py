#!/usr/bin/env python3
"""Compare paid-step recovery after forced process restarts, using no paid service."""

import argparse
from collections import Counter
from datetime import datetime, timezone
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
import json
import os
from pathlib import Path
import platform
import signal
import socket
import subprocess
import sys
import tempfile
import threading
import time
import uuid

ROOT = Path(__file__).resolve().parent
CASE_TIMEOUT = 120
COMMAND_TIMEOUT = 15


class FakeProvider:
    """Every accepted call records one independent charge before responding."""

    def __init__(self, ledger_path):
        self.records = []
        self.lock = threading.Lock()
        self.ledger = ledger_path.open('a', encoding='utf-8')
        provider = self

        class Handler(BaseHTTPRequestHandler):
            def do_POST(self):
                try:
                    self.connection.settimeout(10)
                    length = int(self.headers.get('Content-Length', '0'))
                    if self.path != '/paid-step' or not 0 < length <= 4096:
                        raise ValueError('Expected a bounded POST /paid-step')
                    request = json.loads(self.rfile.read(length))
                    operation = request['operation_id']
                    step = request['step']
                    if not isinstance(operation, str) or not operation or step not in ('identity', 'synthesis'):
                        raise ValueError('Invalid operation_id or step')
                    with provider.lock:
                        identities = [r for r in provider.records if r['operation_id'] == operation and r['step'] == 'identity']
                        if step == 'synthesis' and (not identities or request.get('input') != identities[-1]['result']):
                            raise ValueError('Synthesis must consume the actual identity result')
                        charge = 1 + sum(r['operation_id'] == operation and r['step'] == step for r in provider.records)
                        result = (f'{operation}:identity:charge-{charge}:{uuid.uuid4()}' if step == 'identity'
                                  else f'{request["input"]}:synthesis:charge-{charge}')
                        record = {'operation_id': operation, 'step': step, 'charge': charge, 'result': result}
                        if step == 'synthesis':
                            record['input'] = request['input']
                        provider.ledger.write(json.dumps(record) + '\n')
                        provider.ledger.flush()
                        os.fsync(provider.ledger.fileno())
                        provider.records.append(record)
                    body = json.dumps({'operation_id': operation, 'step': step, 'result': result}).encode()
                    self.send_response(200)
                    self.send_header('Content-Type', 'application/json')
                    self.send_header('Content-Length', str(len(body)))
                    self.end_headers()
                    self.wfile.write(body)
                except (ValueError, KeyError, TypeError, json.JSONDecodeError):
                    self.send_error(400)
                except (BrokenPipeError, ConnectionResetError):
                    pass

            def log_message(self, *_):
                pass

        self.server = ThreadingHTTPServer(('127.0.0.1', 0), Handler)
        self.server.daemon_threads = True
        self.server.timeout = 1
        self.thread = threading.Thread(target=self.server.serve_forever, daemon=True)
        self.thread.start()
        self.url = f'http://127.0.0.1:{self.server.server_port}'

    def observations(self, operation):
        with self.lock:
            records = [dict(r) for r in self.records if r['operation_id'] == operation]
        counts = Counter(r['step'] for r in records)
        return {'identity': counts['identity'], 'synthesis': counts['synthesis']}, records

    def close(self):
        self.server.shutdown()
        self.server.server_close()
        self.thread.join(timeout=5)
        self.ledger.close()


class Processes:
    def __init__(self, directory):
        self.directory = directory
        self.items = []

    def start(self, command, label, cwd=None):
        path = self.directory / f'{len(self.items)}-{label}.log'
        stream = path.open('wb')
        try:
            process = subprocess.Popen(command, cwd=cwd, stdout=stream, stderr=subprocess.STDOUT,
                                       start_new_session=os.name == 'posix')
        finally:
            stream.close()
        item = {'process': process, 'log': path, 'label': label}
        self.items.append(item)
        return item

    @staticmethod
    def kill(item):
        process = item['process']
        if process.poll() is None:
            try:
                if os.name == 'posix':
                    os.killpg(process.pid, signal.SIGKILL)
                else:
                    process.kill()
            except ProcessLookupError:
                pass
        process.wait(timeout=10)
        return {'pid': process.pid, 'exit_code': process.returncode, 'forced': process.returncode != 0}

    def close(self):
        for item in reversed(self.items):
            self.kill(item)


def command(arguments, cwd=None, timeout=COMMAND_TIMEOUT, allowed=(0,)):
    result = subprocess.run(arguments, cwd=cwd, stdout=subprocess.PIPE, stderr=subprocess.PIPE,
                            text=True, timeout=timeout)
    if result.returncode not in allowed:
        raise RuntimeError(f'{arguments[0]} exited {result.returncode}: {result.stderr[-2000:]} {result.stdout[-2000:]}')
    return result


def json_command(arguments, allowed=(0,)):
    result = command(arguments, allowed=allowed)
    # SDK diagnostics can precede the command's final machine-readable response.
    lines = result.stdout.strip().splitlines()
    for line in reversed(lines):
        try:
            payload = json.loads(line)
            if isinstance(payload, dict):
                return payload, result.returncode
        except json.JSONDecodeError:
            pass
    raise RuntimeError(f'No JSON outcome from {arguments[0]}: {result.stdout[-2000:]}')


def wait_until(check, deadline, description, process=None):
    while time.monotonic() < deadline:
        value = check()
        if value:
            return value
        if process and process['process'].poll() is not None:
            raise RuntimeError(f"{description}: process exited early: {process['log'].read_text()[-2000:]}")
        time.sleep(0.05)
    raise TimeoutError(f'Timed out waiting for {description}')


def free_port():
    with socket.socket() as listener:
        listener.bind(('127.0.0.1', 0))
        return listener.getsockname()[1]


def build(module, destination):
    command(['go', 'build', '-buildvcs=false', '-o', str(destination), '.'], cwd=ROOT / module, timeout=180)


def verify(case, provider, operation, expected, status):
    counts, ledger = provider.observations(operation)
    case.update(provider_calls=counts, provider_ledger=ledger)
    if counts != expected:
        raise AssertionError(f'{operation}: expected provider calls {expected}, observed {counts}')
    if case['outcome'].get('status', case['outcome'].get('state')) != status:
        raise AssertionError(f'{operation}: expected {status}, observed {case["outcome"]}')
    if status == 'completed':
        identity = next(r['result'] for r in reversed(ledger) if r['step'] == 'identity')
        synthesis = next(r for r in ledger if r['step'] == 'synthesis')
        if (case['outcome']['results']['identity'] != identity or synthesis['input'] != identity
                or case['outcome']['results']['synthesis'] != synthesis['result']):
            raise AssertionError(f'{operation}: synthesis did not consume the recovered identity result')
        case['synthesis_consumed_identity_result'] = True
        case['completed_checkpoint_reused'] = case['case'] == 'checkpoint'
    case['passed'] = True


def go_case(binary, provider, temporary, processes, failpoint):
    operation = f'go-{failpoint}'
    marker = temporary / f'{operation}.marker'
    store = temporary / f'{operation}-store'
    args = [str(binary), '--provider', provider.url, '--store', str(store), '--operation', operation]
    case = {'engine': 'go_checkpoint', 'case': failpoint, 'operation_id': operation, 'passed': False,
            'execution_policy': 'reuse_checkpoint_stop_on_uncertain_spend'}
    try:
        first = processes.start(args + ['--pause-after', failpoint, '--pause-step', 'identity',
                                       '--marker', str(marker)], operation)
        wait_until(marker.exists, time.monotonic() + CASE_TIMEOUT, 'Go failpoint', first)
        before, _ = provider.observations(operation)
        if before != {'identity': 1, 'synthesis': 0}:
            raise AssertionError(f'Unexpected calls before Go interruption: {before}')
        checkpoint_path = store / operation / 'identity.checkpoint.json'
        if (store / operation / 'synthesis.attempt.json').exists():
            raise AssertionError('Go synthesis was attempted before interruption')
        if failpoint == 'checkpoint':
            checkpoint = json.loads(checkpoint_path.read_text())
            _, ledger = provider.observations(operation)
            if checkpoint['operation_id'] != operation or checkpoint['step'] != 'identity' or checkpoint['result'] != ledger[0]['result']:
                raise AssertionError('Go checkpoint does not contain the actual identity response')
            case['before_restart_checkpoint'] = checkpoint
        elif checkpoint_path.exists():
            raise AssertionError('Unknown-spend failpoint unexpectedly has an identity checkpoint')
        case['interrupted_processes'] = [Processes.kill(first)]
        restarted = processes.start(args, operation + '-restart')
        restarted['process'].wait(timeout=CASE_TIMEOUT)
        case['restart_pid'] = restarted['process'].pid
        case['restart_exit_code'] = restarted['process'].returncode
        case['outcome'] = json.loads(restarted['log'].read_text().strip().splitlines()[-1])
        expected_status = 'completed' if failpoint == 'checkpoint' else 'spend_uncertain'
        expected_exit = 0 if failpoint == 'checkpoint' else 1
        if restarted['process'].returncode != expected_exit:
            raise AssertionError(f'Unexpected Go restart exit: {restarted["process"].returncode}')
        verify(case, provider, operation, {'identity': 1, 'synthesis': int(failpoint == 'checkpoint')}, expected_status)
    except (OSError, ValueError, KeyError, IndexError, RuntimeError, AssertionError, TimeoutError, subprocess.TimeoutExpired) as error:
        case['error'] = str(error)
        case['provider_calls'], case['provider_ledger'] = provider.observations(operation)
    finally:
        for item in processes.items:
            if item['label'].startswith(operation):
                Processes.kill(item)
    return case


def temporal_case(binary, cli, provider, temporary, processes, failpoint):
    operation = f'temporal-{failpoint}'
    marker = temporary / f'{operation}.marker'
    port = free_port()
    address = f'127.0.0.1:{port}'
    db = temporary / f'{operation}.db'
    server_args = [str(cli), 'server', 'start-dev', '--ip', '127.0.0.1', '--port', str(port),
                   '--db-filename', str(db), '--headless']
    common = ['--address', address, '--operation', operation]
    case = {'engine': 'temporal', 'case': failpoint, 'operation_id': operation, 'passed': False,
            'execution_policy': 'reuse_activity_history_retry_maximum_attempts_2',
            'activity_maximum_attempts': 2, 'persisted_dev_database': True}
    deadline = time.monotonic() + CASE_TIMEOUT
    try:
        def start_server(suffix):
            server = processes.start(server_args, operation + '-server' + suffix)

            def ready():
                try:
                    result = command([str(cli), '--address', address, 'operator', 'cluster', 'health'], timeout=3, allowed=(0, 1))
                    return result.returncode == 0
                except subprocess.TimeoutExpired:
                    return False
            wait_until(ready, deadline, 'Temporal Service readiness', server)
            return server

        def status():
            payload, _ = json_command([str(binary), 'status', *common])
            return payload

        server = start_server('')
        worker = processes.start([str(binary), 'worker', *common], operation + '-worker')
        json_command([str(binary), 'start', *common, '--provider', provider.url,
                      '--pause-after', failpoint, '--pause-step', 'identity', '--marker', str(marker)])
        if failpoint == 'checkpoint':
            def checkpoint():
                observed = status()
                return observed if observed.get('state') == 'waiting' and observed.get('step') == 'identity' and observed.get('checkpoint') else None
            case['before_restart'] = wait_until(checkpoint, deadline, 'Temporal recorded identity checkpoint', worker)
            history, _ = json_command([str(binary), 'history', *common])
            events = history['events']
            scheduled = [e for e in events if e['event_type'] == 'ActivityTaskScheduled']
            completed_events = [e for e in events if e['event_type'] == 'ActivityTaskCompleted']
            if len(scheduled) != 1 or len(completed_events) != 1:
                raise AssertionError(f'Checkpoint must have one scheduled and completed identity Activity: {history}')
            case['before_restart_history'] = history
        else:
            wait_until(marker.exists, deadline, 'Temporal provider-completion failpoint', worker)
            history, _ = json_command([str(binary), 'history', *common])
            events = history['events']
            if any(e['event_type'] == 'ActivityTaskCompleted' for e in events):
                raise AssertionError('Unknown-spend failpoint already has a recorded Activity completion')
            if sum(e['event_type'] == 'ActivityTaskScheduled' for e in events) != 1:
                raise AssertionError('Expected only the first identity Activity before interruption')
            case['before_restart_history'] = history
        before, _ = provider.observations(operation)
        if before != {'identity': 1, 'synthesis': 0}:
            raise AssertionError(f'Unexpected calls before Temporal interruption: {before}')
        case['interrupted_processes'] = [Processes.kill(worker), Processes.kill(server)]
        server = start_server('-restart')
        worker = processes.start([str(binary), 'worker', *common], operation + '-worker-restart')
        case['restart_pids'] = [worker['process'].pid, server['process'].pid]
        case['service_and_worker_restarted'] = True
        if failpoint == 'checkpoint':
            json_command([str(binary), 'continue', *common])

        def completed():
            observed = status()
            if observed.get('state') in ('failed', 'terminated', 'canceled', 'timed_out'):
                raise RuntimeError(f'Temporal operation ended: {observed}')
            return observed if observed.get('state') == 'completed' else None
        case['outcome'] = wait_until(completed, deadline, 'Temporal completion after restart', worker)
        expected = {'identity': 1 if failpoint == 'checkpoint' else 2, 'synthesis': 1}
        verify(case, provider, operation, expected, 'completed')
    except (OSError, ValueError, KeyError, IndexError, RuntimeError, AssertionError, TimeoutError, subprocess.TimeoutExpired) as error:
        case['error'] = str(error)
        case['provider_calls'], case['provider_ledger'] = provider.observations(operation)
    finally:
        for item in processes.items:
            if item['label'].startswith(operation):
                Processes.kill(item)
    return case


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--temporal-cli', type=Path, help='Path to the official Temporal CLI')
    parser.add_argument('--skip-temporal', action='store_true', help='Run Go only; mark Temporal as not validated')
    parser.add_argument('--output', type=Path, help='Also write the JSON report to this path; default is stdout only')
    args = parser.parse_args()
    if not args.skip_temporal and not args.temporal_cli:
        parser.error('--temporal-cli is required unless --skip-temporal is explicit')
    report = {'schema_version': 1, 'recorded_at': datetime.now(timezone.utc).isoformat(),
              'flags': {'skip_temporal': args.skip_temporal, 'case_timeout_seconds': CASE_TIMEOUT},
              'versions': {'python': platform.python_version(), 'platform': platform.platform(),
                           'go': command(['go', 'version']).stdout.strip()},
              'scope': {'paid_services_called': False, 'provider_idempotency': False,
                        'process_restart': True, 'power_loss': False, 'wiki_commit_recovery': False,
                        'performance_benchmark': False, 'architecture_selected': False},
              'temporal_validation': 'not_validated' if args.skip_temporal else 'attempted', 'cases': []}
    with tempfile.TemporaryDirectory(prefix='tower-recovery-') as location:
        temporary = Path(location)
        processes = Processes(temporary)
        provider = FakeProvider(temporary / 'provider.jsonl')
        try:
            go_binary = temporary / 'go-checkpoint'
            build('go-checkpoint', go_binary)
            report['cases'] += [go_case(go_binary, provider, temporary, processes, failpoint)
                                for failpoint in ('checkpoint', 'provider')]
            if not args.skip_temporal:
                cli = args.temporal_cli.expanduser().resolve()
                report['versions']['temporal_cli'] = command([str(cli), '--version']).stdout.strip()
                temporal_binary = temporary / 'temporal-runner'
                build('temporal', temporal_binary)
                report['versions']['temporal_go_sdk'] = command(['go', 'list', '-m', 'go.temporal.io/sdk'], cwd=ROOT / 'temporal').stdout.strip()
                report['cases'] += [temporal_case(temporal_binary, cli, provider, temporary, processes, failpoint)
                                    for failpoint in ('checkpoint', 'provider')]
                report['temporal_validation'] = 'passed' if all(c['passed'] for c in report['cases'] if c['engine'] == 'temporal') else 'failed'
        except (OSError, ValueError, RuntimeError, TimeoutError, subprocess.TimeoutExpired) as error:
            report['error'] = str(error)
        finally:
            processes.close()
            provider.close()
        report['passed'] = not report.get('error') and all(c['passed'] for c in report['cases'])
    output = json.dumps(report, indent=2) + '\n'
    if args.output:
        args.output.expanduser().write_text(output, encoding='utf-8')
    sys.stdout.write(output)
    return 0 if report['passed'] else 1


if __name__ == '__main__':
    sys.exit(main())
