import { useEffect, useState, type CSSProperties } from 'react';
import {
  Check,
  Code,
  Coins,
  Crosshair,
  Eye,
  Flame,
  Sprout,
  Layers,
  RotateCcw,
  Sword,
  Target,
  Timer,
} from 'lucide-react';
import { Badge } from '../ui/badge.js';
import { Alert } from '../ui/alert.js';
import { IconButton } from '../ui/icon-button.js';
import { Field } from '../ui/field.js';
import { Disclosure } from '../ui/disclosure.js';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '../ui/select.js';
import { cn } from '../ui/utils.js';
import { UnitWorkspace } from './workspace.js';
import type { CharacterSource, Example, TowerState } from '../types.js';
import { Button } from '../ui/button.js';
import { Textarea } from '../ui/input.js';
import { SourceEvidence } from './sources.js';
import { api } from './api.js';

const numbers = new Intl.NumberFormat('en-US', { maximumFractionDigits: 3 });
const views = [
  { value: 'upgrades', label: 'Upgrade paths', icon: Layers },
  { value: 'states', label: 'All states', icon: Crosshair },
  { value: 'data', label: 'Tower JSON', icon: Code },
];

export function Tower({
  draft,
  onDraft,
  source,
  onBusy,
}: {
  draft: string;
  onDraft: (value: string) => void;
  source?: CharacterSource;
  onBusy: (value: boolean) => void;
}) {
  const [example, setExample] = useState<Example>();
  const [error, setError] = useState('');
  const [revision, setRevision] = useState(0);
  const [tiers, setTiers] = useState([0, 0, 0]);
  const [tab, setTab] = useState('upgrades');
  const [raw, setRaw] = useState('');
  const [rawError, setRawError] = useState('');
  const [building, setBuilding] = useState(false);
  const [buildError, setBuildError] = useState('');
  const [status, setStatus] = useState('');

  useEffect(() => {
    const controller = new AbortController();
    setError('');
    fetch('/api/example', { signal: controller.signal })
      .then(async (response) => {
        const data = await response.json();
        if (!response.ok) throw new Error(data.error);
        setExample(data);
        if (!draft) onDraft(JSON.stringify(data.design, null, 2));
      })
      .catch((failure: Error) => {
        if (failure.name !== 'AbortError') setError(failure.message);
      });
    return () => controller.abort();
  }, [revision]);

  const selected = example?.states.find((state) =>
    state.tiers.every((tier, path) => tier === tiers[path]),
  );
  useEffect(() => {
    if (tab !== 'data' || !selected) return;
    const controller = new AbortController();
    setRaw('');
    setRawError('');
    fetch(`/api/tower?state=${encodeURIComponent(selected.name)}`, { signal: controller.signal })
      .then(async (response) => {
        const data = await response.json();
        if (!response.ok) throw new Error(data.error);
        setRaw(JSON.stringify(data, null, 2));
      })
      .catch((failure: Error) => {
        if (failure.name !== 'AbortError') setRawError(failure.message);
      });
    return () => controller.abort();
  }, [tab, selected?.name, revision]);

  const legal = (path: number, tier: number) => {
    const candidate = tiers.map((value, index) => (index === path ? tier : value));
    return (
      example?.states.some((state) =>
        state.tiers.every((value, index) => value === candidate[index]),
      ) ?? false
    );
  };
  const choose = (path: number, tier: number) =>
    setTiers(tiers.map((value, index) => (index === path ? tier : value)));
  const code = tiers.join('-');
  const appearance = example?.design.paths
    .flatMap((path, index) => path.upgrades.slice(0, tiers[index]))
    .filter((upgrade) => upgrade.appearance)
    .at(-1)?.appearance;

  const research = source ?? example?.characterSource;
  async function build() {
    setBuilding(true);
    onBusy(true);
    setBuildError('');
    setStatus('');
    try {
      const design = JSON.parse(draft) as Example['design'];
      if (research) design.sourceId = research.id;
      const result = await api<{ message: string }>('/api/build', design);
      onDraft(JSON.stringify(design, null, 2));
      setStatus(result.message);
      setRevision((value) => value + 1);
    } catch (failure) {
      setBuildError(failure instanceof Error ? failure.message : String(failure));
    } finally {
      setBuilding(false);
      onBusy(false);
    }
  }

  return (
    <>
      <UnitWorkspace
        gallery={
          example &&
          selected && (
            <>
              <div className="mb-4 flex items-center justify-between gap-2 text-xs">
                <span className="font-mono">{code}</span>
                <IconButton
                  size="icon-sm"
                  label="Reset to base state"
                  onClick={() => setTiers([0, 0, 0])}
                >
                  <RotateCcw />
                </IconButton>
              </div>
              {example.design.paths.map((path, index) => (
                <Field key={path.name} label={path.name}>
                  <Select
                    value={String(tiers[index])}
                    onValueChange={(value) => choose(index, Number(value))}
                  >
                    <SelectTrigger aria-label={`${path.name} upgrade`} size="sm">
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      {[0, 1, 2, 3, 4, 5].map((tier) => (
                        <SelectItem key={tier} value={String(tier)} disabled={!legal(index, tier)}>
                          {tier === 0 ? 'No upgrades' : String(tier)}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </Field>
              ))}
              <p className="my-4 text-xs text-muted-foreground">
                Up to two paths. The second stops at upgrade 2.
              </p>
              {research && (
                <Disclosure bare title="Character research">
                  <p className="mb-2 text-[13px]">{research.character.work}</p>
                  <p className="mb-3 text-xs text-muted-foreground">{research.character.scope}</p>
                  <SourceEvidence source={research} />
                </Disclosure>
              )}
              <Disclosure bare title="Dart reference">
                <p className="text-xs text-muted-foreground">
                  BTD6 {example.source.capture}, build {example.source.build}. Dart art and sounds
                  are placeholders.
                </p>
              </Disclosure>
            </>
          )
        }
        workflow={
          example && (
            <>
              <div className="mb-5 flex items-center justify-between gap-2 text-xs">
                <span className="text-muted-foreground">
                  Default Profile {example.checks.candidate.profile.revision}
                </span>
                <Badge variant={example.checks.candidate.score.complete ? 'success' : 'danger'}>
                  {example.checks.candidate.score.points}/100
                </Badge>
              </div>
              <Button
                variant="primary"
                className="mb-3 w-full"
                disabled={building || !draft}
                onClick={() => void build()}
              >
                {building ? 'Building…' : 'Build and check'}
              </Button>
              {buildError && <Alert>{buildError}</Alert>}
              {status && (
                <p role="status" className="mb-3 text-xs text-success">
                  {status}
                </p>
              )}
              <Button
                variant="ghost"
                size="sm"
                className="mb-5"
                disabled={building}
                onClick={() => setRevision((value) => value + 1)}
              >
                <RotateCcw /> Reload Tower
              </Button>
              <ul className="space-y-4 text-[13px]" aria-label="Validation results">
                <li className="flex items-start gap-2.5">
                  <Check className="mt-0.5 size-4 text-success" />
                  <div>
                    {example.checks.candidate.filesChecked} files checked
                    <p className="text-xs text-muted-foreground">Candidate passed.</p>
                  </div>
                </li>
                {(['missingState', 'missingUpgrade'] as const).map((key) => (
                  <li key={key} className="flex items-start gap-2.5">
                    <Check className="mt-0.5 size-4 text-success" />
                    <div>
                      {key === 'missingState' ? 'Missing state' : 'Missing upgrade'}
                      <p className="text-xs text-muted-foreground">
                        {example.checks[key].errors.length
                          ? 'Rejected as expected.'
                          : 'Not rejected.'}
                      </p>
                    </div>
                  </li>
                ))}
              </ul>
              <p className="my-5 text-xs text-muted-foreground">
                Declared data checks pass. Balance and gameplay are untested.
              </p>
              <Disclosure bare title="Check details">
                <p className="mb-3 text-xs text-muted-foreground">
                  Checker {example.checks.candidate.checker.version}
                </p>
                <p className="text-xs text-muted-foreground">
                  {example.states.length} ordinary states, 15 upgrades. No Paragon or Monkey
                  Knowledge.
                </p>
                <a
                  className="mt-3 inline-block text-xs"
                  href="https://github.com/mardwerk/tower-generator/issues/100"
                  target="_blank"
                  rel="noreferrer"
                >
                  Profile follow-up
                </a>
              </Disclosure>
            </>
          )
        }
      >
        <header className="mt-6 border-b border-border pb-2">
          <div className="flex items-center gap-3">
            <h1 className="text-[1.6rem] font-semibold tracking-tight">
              {example?.design.name ?? 'Usopp'}
            </h1>
            <Badge variant="warning">Draft</Badge>
          </div>
          {selected && (
            <div className="mt-3.5 mb-1.5 flex flex-wrap items-center gap-x-3 gap-y-1 text-xs">
              <span className="font-mono text-muted-foreground">{code}</span>
              <Cost value={selected.cost} />
            </div>
          )}
        </header>
        <div className="my-3 flex flex-wrap gap-1" role="group" aria-label="Tower view">
          {views.map(({ value, label, icon: Icon }) => (
            <Button
              key={value}
              size="xs"
              variant={tab === value ? 'secondary' : 'ghost'}
              aria-pressed={tab === value}
              onClick={() => setTab(value)}
            >
              <Icon />
              {label}
            </Button>
          ))}
          <Button
            size="xs"
            variant={tab === 'design' ? 'secondary' : 'ghost'}
            aria-pressed={tab === 'design'}
            onClick={() => setTab('design')}
          >
            Edit design
          </Button>
        </div>
        {error ? (
          <Alert>{error}</Alert>
        ) : !example || !selected ? (
          <p role="status" className="my-5 text-muted-foreground">
            Loading example…
          </p>
        ) : (
          <>
            <div
              className="my-5 rounded-lg border border-border bg-card p-[18px]"
              aria-label="Selected state"
            >
              <div className="mb-3 flex items-center gap-2 text-[13px]">
                {selected.attack === 'explosion' ? (
                  <Flame className="size-4 text-warning" />
                ) : selected.attack === 'spread' ? (
                  <Sprout className="size-4 text-success" />
                ) : (
                  <Crosshair className="size-4 text-muted-foreground" />
                )}
                <span>
                  {
                    {
                      pellet: 'Lead Star',
                      explosion: 'Exploding Star',
                      spread: 'Leaf shuriken',
                      sniper: 'Kabuto sniper shot',
                    }[selected.attack]
                  }
                </span>
              </div>
              <Stats state={selected} />
              {appearance && (
                <p className="mt-3 text-xs text-muted-foreground">
                  Planned appearance: {appearance}
                </p>
              )}
            </div>
            {tab === 'design' && (
              <Field label="Tower design">
                <Textarea
                  aria-label="Tower design"
                  className="min-h-[500px] font-mono text-xs"
                  value={draft}
                  disabled={building}
                  onChange={(event) => onDraft(event.target.value)}
                />
              </Field>
            )}
            {tab === 'upgrades' && (
              <div
                className="upgrade-paths grid gap-[18px]"
                style={{ '--path-count': 3, '--path-rows': 5 } as CSSProperties}
              >
                {example.design.paths.map((path, pathIndex) => (
                  <div key={path.name} className="path-section min-w-0" aria-label={path.name}>
                    {path.upgrades.map((upgrade, index) => {
                      const tier = index + 1;
                      const active = tier <= (tiers[pathIndex] ?? 0);
                      const upgradeCode = [0, 1, 2]
                        .map((position) => (position === pathIndex ? tier : 'x'))
                        .join('-');
                      return (
                        <button
                          key={upgrade.name}
                          aria-label={`${upgradeCode} ${upgrade.name}`}
                          aria-pressed={active}
                          disabled={!legal(pathIndex, tier)}
                          onClick={() => choose(pathIndex, tier)}
                          className={cn(
                            'tier-card grid w-full grid-cols-[auto_minmax(0,1fr)] gap-2.5 border-t border-border px-1.5 py-4 text-left transition-colors outline-none hover:bg-accent/50 focus-visible:ring-2 focus-visible:ring-ring/60 disabled:cursor-not-allowed disabled:opacity-35',
                            active && 'bg-accent/50',
                          )}
                          style={{ '--tier-row': tier } as CSSProperties}
                        >
                          <span className="mt-px grid h-[22px] min-w-[22px] place-items-center rounded-[5px] border border-border px-1 font-mono text-[11px] whitespace-nowrap text-muted-foreground">
                            {upgradeCode}
                          </span>
                          <span className="min-w-0">
                            <span className="flex items-center gap-2 text-[13px]">
                              {upgrade.name}
                              {active && <Check className="size-3 text-success" />}
                            </span>
                            <Cost value={upgrade.cost} />
                            <span className="my-2 block text-xs text-muted-foreground">
                              {upgrade.description}
                            </span>
                          </span>
                        </button>
                      );
                    })}
                  </div>
                ))}
              </div>
            )}
            {tab === 'states' && (
              <div className="overflow-x-auto">
                <table className="w-full min-w-[580px] border-collapse text-left text-xs">
                  <thead className="text-muted-foreground">
                    <tr>
                      {[
                        'State',
                        'Cost',
                        'Damage',
                        'Pierce',
                        'Range',
                        'Interval',
                        'Pellets',
                        'Camo',
                      ].map((label) => (
                        <th className="py-1.5 pr-3 font-medium" key={label}>
                          {label}
                        </th>
                      ))}
                    </tr>
                  </thead>
                  <tbody>
                    {[...example.states]
                      .sort((a, b) => a.tiers.join('').localeCompare(b.tiers.join('')))
                      .map((state) => (
                        <tr
                          key={state.name}
                          className={cn(
                            'border-t border-border',
                            selected.name === state.name && 'bg-accent',
                          )}
                        >
                          <td className="py-2 pr-3">
                            <button
                              className="font-mono text-link hover:underline"
                              onClick={() => setTiers(state.tiers)}
                            >
                              {state.tiers.join('-')}
                            </button>
                          </td>
                          {[
                            state.cost,
                            state.damage,
                            state.pierce,
                            state.range,
                            state.interval,
                            state.projectiles,
                          ].map((value, index) => (
                            <td className="py-2 pr-3 tabular-nums" key={index}>
                              {numbers.format(value)}
                            </td>
                          ))}
                          <td className="py-2">{state.camo ? 'Yes' : 'No'}</td>
                        </tr>
                      ))}
                  </tbody>
                </table>
              </div>
            )}
            {tab === 'data' && (
              <>
                <p className="my-3 font-mono text-xs text-muted-foreground">{selected.name}.json</p>
                {rawError ? (
                  <Alert>{rawError}</Alert>
                ) : (
                  <pre className="max-h-[600px] overflow-auto rounded-md bg-muted/60 p-3 font-mono text-xs leading-normal whitespace-pre-wrap [overflow-wrap:anywhere]">
                    {raw || 'Loading Tower data…'}
                  </pre>
                )}
              </>
            )}
          </>
        )}
      </UnitWorkspace>
    </>
  );
}

function Cost({ value }: { value: number }) {
  return (
    <span
      className="kit-cost my-1.5 inline-flex items-center gap-1 text-[13px] text-warning tabular-nums"
      title="Purchase cost"
    >
      <Coins className="size-[15px]" aria-hidden="true" />
      <span>{numbers.format(value)}</span>
    </span>
  );
}

function Stats({ state }: { state: TowerState }) {
  const stats = [
    { label: 'Damage', value: numbers.format(state.damage), icon: Sword },
    { label: 'Pierce', value: numbers.format(state.pierce), icon: Layers },
    { label: 'Range', value: numbers.format(state.range), icon: Crosshair },
    { label: 'Attack interval', value: `${numbers.format(state.interval)} s`, icon: Timer },
    { label: 'Projectiles', value: numbers.format(state.projectiles), icon: Target },
    { label: 'Camo detection', value: state.camo ? 'Yes' : 'No', icon: Eye },
  ];
  if (state.blastRadius)
    stats.push({ label: 'Blast radius', value: numbers.format(state.blastRadius), icon: Flame });
  return (
    <ul className="kit-stats grid gap-2 text-xs sm:grid-cols-2">
      {stats.map(({ label, value, icon: Icon }) => (
        <li key={label} className="flex flex-wrap items-center gap-1.5">
          <Icon className="size-[15px] text-muted-foreground" aria-hidden="true" />
          <span className="text-muted-foreground">{label}</span>
          <span className="tabular-nums">{value}</span>
        </li>
      ))}
    </ul>
  );
}
