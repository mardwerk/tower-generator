import {
  ArrowRight,
  ChevronsLeft,
  Coins,
  Crosshair,
  Timer,
  Sword,
  Target,
  Layers,
  Expand,
  Snowflake,
  Flame,
  Zap,
  Eye,
  Sparkles,
  MoveUpRight,
  ShieldOff,
  type LucideIcon,
} from 'lucide-react';
import type { StatChange } from '../../api/contract.js';
import { cn } from '../../ui/utils.js';

export const statLabels = {
  damage: 'Damage',
  intervalSeconds: 'Attack interval',
  range: 'Range',
  pierce: 'Pierce',
  projectiles: 'Projectiles per attack',
  splashRadius: 'Splash radius',
  slowPercent: 'Slow',
  slowSeconds: 'Slow duration',
  burnDamagePerSecond: 'Burn damage per second',
  burnSeconds: 'Burn duration',
  stunSeconds: 'Stun duration',
  durationSeconds: 'Active duration',
  cooldownSeconds: 'Cooldown',
  damageMultiplier: 'Active damage bonus',
  intervalMultiplier: 'Active attack speed',
  rangeBonus: 'Active range bonus',
  camo: 'Camo detection',
  delivery: 'Delivery',
  damageType: 'Damage type',
  targeting: 'Targeting',
  ability: 'Active Ability',
  distribution: 'Volley targets',
  followUp: 'Secondary attack',
  activeFollowUp: 'During activation',
} as const;
export type StatKey = keyof typeof statLabels;

/**
 * A stat value with its unit. The server words an Active Ability's
 * multipliers as percentages ("+100%", "100% faster"), so they arrive as
 * text and show as they are.
 */
export function statValue(
  key: StatKey,
  value: string | number,
  unit?: string,
  kind?: StatChange['kind'],
): string {
  if (typeof value === 'string') return value;
  const number = Number(value.toFixed(4)).toString();
  if (kind === 'bonusDamage') return `+${number}`;
  if (unit === 's') return `${number} s`;
  if (unit === 'percent') return `${number}%`;
  if (unit) return `${number} ${unit}`;
  if (key.endsWith('Seconds')) return `${number} s`;
  if (key === 'slowPercent') return `${number}%`;
  return number;
}

const icons: Record<StatKey, LucideIcon> = {
  damage: Sword,
  intervalSeconds: Timer,
  range: Crosshair,
  pierce: Layers,
  projectiles: Target,
  splashRadius: Expand,
  slowPercent: Snowflake,
  slowSeconds: Snowflake,
  burnDamagePerSecond: Flame,
  burnSeconds: Flame,
  stunSeconds: Zap,
  durationSeconds: Timer,
  cooldownSeconds: Timer,
  damageMultiplier: Sword,
  intervalMultiplier: Timer,
  rangeBonus: Crosshair,
  camo: Eye,
  delivery: MoveUpRight,
  damageType: Sparkles,
  targeting: Target,
  ability: Zap,
  distribution: Target,
  followUp: Sparkles,
  activeFollowUp: Sparkles,
};

/** Profile-defined status effects and detection traits are drawn by kind. */
const kindIcons: Record<NonNullable<StatChange['kind']>, LucideIcon> = {
  moveSpeed: Snowflake,
  damageOverTime: Flame,
  disable: Zap,
  damageTaken: ShieldOff,
  knockback: ChevronsLeft,
  custom: Sparkles,
  detection: Eye,
  bonusDamage: Sword,
};
export function Cost({ value, currency }: { value: number; currency: string }) {
  return (
    <span
      className="kit-cost my-1.5 inline-flex items-center gap-1 text-[13px] text-warning tabular-nums"
      title={`${currency}: purchase cost`}
      aria-label={`${currency} ${value}`}
    >
      <Coins className="size-[15px]" aria-hidden="true" />
      <span>{value.toLocaleString('en-US')}</span>
    </span>
  );
}
/** Flags a change for the worse; the delta already says how much. */
function worseLabel(key: string): string {
  if (key === 'intervalSeconds' || key === 'intervalMultiplier') return 'Slower';
  if (key === 'cooldownSeconds') return 'Longer wait';
  return 'Reduced';
}

/** Each change as "before → after (delta)", as the unit sheet words it. */
export function StatValues({ changes }: { changes: StatChange[] }) {
  return (
    <ul className="kit-stats relative z-[2] mt-2 grid gap-1.5 text-xs">
      {changes.map((change) => {
        const key = change.key as StatKey;
        const Icon = (change.kind && kindIcons[change.kind]) ?? icons[key] ?? Sparkles;
        const label = change.label ?? statLabels[key] ?? change.key;
        return (
          <li
            key={change.key}
            className="flex flex-wrap items-center gap-1.5"
            title={label}
            data-improvement={change.improvement}
          >
            <Icon className="size-[15px] text-muted-foreground" aria-hidden="true" />
            <span className="text-muted-foreground">{label}</span>
            <span className="inline-flex flex-wrap items-center gap-1.5 tabular-nums">
              {change.before !== undefined && (
                <>
                  <span className="text-destructive">
                    <span className="sr-only">Previous: </span>
                    {statValue(key, change.before, change.unit, change.kind)}
                  </span>
                  <ArrowRight className="size-3" aria-hidden="true" />
                </>
              )}
              <span className={cn(change.before !== undefined && 'text-success')}>
                <span className="sr-only">{change.before === undefined ? '' : 'New: '}</span>
                {statValue(key, change.after, change.unit, change.kind)}
              </span>
              {change.delta && (
                <span className="text-muted-foreground">
                  <span className="sr-only">Change: </span>({change.delta})
                </span>
              )}
            </span>
            {change.improvement === false && (
              <span className="text-[11px] text-warning">{worseLabel(change.key)}</span>
            )}
          </li>
        );
      })}
    </ul>
  );
}
