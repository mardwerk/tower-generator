"""The configured OpenRouter model, with bounded JSON calls and English search."""

import json
import os
import re
import shlex
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen

from src.research import PROJECT


def configuration():
    values = {}
    path = PROJECT / '.env'
    if path.exists():
        for number, line in enumerate(path.read_text(encoding='utf-8-sig').splitlines(), 1):
            line = line.strip().removeprefix('export ')
            if not line or line.startswith('#'):
                continue
            try:
                name, value = line.split('=', 1)
                if not re.fullmatch(r'[A-Za-z_][A-Za-z0-9_]*', name.strip()):
                    raise ValueError()
                parts = shlex.split(value, comments=True)
                values[name.strip()] = ' '.join(parts)
            except ValueError:
                raise ValueError(f'Invalid local .env format at line {number}') from None
    values.update(os.environ)
    return values


class Model:
    def __init__(self):
        settings = configuration()
        self.key = settings.get('OPENROUTER_API_KEY', '').strip()
        self.name = settings.get('OPENROUTER_MODEL', '').strip()
        self.reasoning = settings.get('OPENROUTER_REASONING', 'low').strip()
        if not self.key or not self.name:
            raise ValueError('Research needs OPENROUTER_API_KEY and OPENROUTER_MODEL in the environment or local .env')
        if not re.fullmatch(r'[a-zA-Z0-9._-]+/[a-zA-Z0-9._:/-]+', self.name) or self.key in self.name:
            raise ValueError('Invalid OpenRouter model configuration')
        if self.reasoning not in {'none', 'minimal', 'low', 'medium', 'high', 'xhigh'}:
            raise ValueError('Invalid OpenRouter reasoning configuration')
        self.usage = []

    def json(self, instructions, data, *, tokens, search_results=0):
        encoded = json.dumps(data, ensure_ascii=False)
        if len(encoded) + len(instructions) > 100_000:
            raise ValueError('Research model input exceeds the 100,000 character budget')
        payload = {'model': self.name, 'max_tokens': tokens,
                   'response_format': {'type': 'json_object'},
                   'messages': [{'role': 'system', 'content': instructions},
                                {'role': 'user', 'content': encoded}]}
        if self.reasoning != 'none':
            payload['reasoning'] = {'effort': self.reasoning}
        if search_results:
            payload['plugins'] = [{'id': 'web', 'engine': 'exa', 'max_results': search_results}]
        request = Request('https://openrouter.ai/api/v1/chat/completions', data=json.dumps(payload).encode(),
                          headers={'Authorization': 'Bearer ' + self.key, 'Content-Type': 'application/json',
                                   'X-OpenRouter-Title': 'Tower Generator research'})
        try:
            with urlopen(request, timeout=180) as response:
                raw = response.read(2_000_001)
            if len(raw) > 2_000_000:
                raise ValueError('Research model response exceeds its size limit')
            result = json.loads(raw)
        except HTTPError as error:
            raise ValueError(f'OpenRouter research request failed with HTTP {error.code}; saved Wiki files are unchanged') from None
        except (URLError, TimeoutError):
            raise ValueError('OpenRouter research request failed or timed out; saved Wiki files are unchanged') from None
        if result.get('error'):
            raise ValueError('OpenRouter returned a research error; saved Wiki files are unchanged')
        try:
            choice = result['choices'][0]
            if choice.get('finish_reason') == 'length':
                raise ValueError('Research reached its model output budget; saved Wiki files are unchanged')
            message = choice['message']
            content = message['content'].strip()
            content = re.sub(r'^```(?:json)?\s*|\s*```$', '', content)
            output = json.loads(content)
            if not isinstance(output, dict):
                raise ValueError()
        except (KeyError, IndexError, TypeError, AttributeError, json.JSONDecodeError):
            raise ValueError('Research model did not return a JSON object; saved Wiki files are unchanged') from None
        usage = result.get('usage', {})
        self.usage.append({k: usage[k] for k in ['prompt_tokens', 'completion_tokens', 'total_tokens', 'cost'] if k in usage})
        sources = [a['url_citation'] for a in message.get('annotations', [])
                   if a.get('type') == 'url_citation' and isinstance(a.get('url_citation'), dict)]
        return output, sources
