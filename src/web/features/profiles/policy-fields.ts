// The Profile rules editor's descriptor: one entry per typed design policy
// field, with its label, control, help text and allowed values. It holds UI
// text only; the Engine owns every rule, check and prompt.
import type { DesignPolicy, PathKey } from '../../api/contract.js';

/** Fields every version 1 policy sets to one fixed literal; the editor shows no control for them. */
export const fixedPolicyKeys = [
  'version',
  'tier5Uniqueness',
] as const satisfies readonly (keyof DesignPolicy)[];
export type FixedPolicyKey = (typeof fixedPolicyKeys)[number];

/** The editable fields: every `DesignPolicy` field except the fixed literals. */
export type PolicyKey = Exclude<keyof DesignPolicy, FixedPolicyKey>;
export type PolicyValue = DesignPolicy[PolicyKey];

/** One allowed value of a choice; `undefined` means the field is absent. */
export interface PolicyChoice {
  id: string;
  value: PathKey | null | undefined;
  label: string;
}

export type PolicyControl =
  /** An optional toggle is off when absent; switching it off removes the field. */
  | { kind: 'toggle'; optional: boolean }
  | {
      kind: 'number';
      integer: boolean;
      min: number;
      /** The value must be above `min`, not equal to it. */
      minExclusive: boolean;
      max: number;
      /** Set for an optional number: what an empty field, which removes it, means. */
      empty?: string;
    }
  | { kind: 'choice'; choices: readonly PolicyChoice[] };

export interface PolicyField {
  label: string;
  help: string;
  control: PolicyControl;
}

const toggle = (optional: boolean): PolicyControl => ({ kind: 'toggle', optional });

export const pathNames: Record<PathKey, string> = {
  path1: 'Top',
  path2: 'Middle',
  path3: 'Bottom',
};

/**
 * One entry per editable field, in display order. Typed as a `Record` over
 * `PolicyKey`, so typecheck fails when a `DesignPolicy` field has no entry.
 */
export const policyFields: Record<PolicyKey, PolicyField> = {
  distinctPathSpecializations: {
    label: 'Distinct specializations',
    help: 'Each path declares a different specialization.',
    control: toggle(false),
  },
  distinctFirstUpgrades: {
    label: 'Distinct first purchases',
    help: 'The resolved 1-0-0, 0-1-0 and 0-0-1 builds differ.',
    control: toggle(false),
  },
  distinctEarlyBenefits: {
    label: 'Distinct early benefits',
    help: "Two paths' first two purchases together must differ in what they improve or unlock, whatever their order, names, prices or amounts; lowers do not count.",
    control: toggle(true),
  },
  exclusiveEarlyBenefits: {
    label: 'Exclusive early benefits',
    help: "What one path's first two purchases improve or unlock, no other path's first two purchases improve or unlock; a path may repeat its own, and later purchases may improve it. It includes Distinct early benefits.",
    control: toggle(true),
  },
  distinctCapstones: {
    label: 'Distinct capstones',
    help: 'The pure 5-0-0, 0-5-0 and 0-0-5 builds differ, ignoring names and prices.',
    control: toggle(false),
  },
  preserveEarlyAttackIdentity: {
    label: 'Early attack identity',
    help: "Each path's first and second purchases keep the base attack's form.",
    control: toggle(true),
  },
  requireTier3BehaviorChange: {
    label: 'Behavior at the third purchase',
    help: 'Every third purchase adds a supported behavior or access, not only larger numbers. One whose only new capability is a proposed mechanic is reported as a design gap.',
    control: toggle(true),
  },
  requireTier5BehaviorChange: {
    label: 'Behavior at the fifth purchase',
    help: 'Every fifth purchase adds a supported behavior or access, not only larger numbers. One whose only new capability is a proposed mechanic is reported as a design gap.',
    control: toggle(true),
  },
  minTier5SpecialtyMultiplier: {
    label: 'Capstone multiplier',
    help: "The fifth purchase multiplies its path's specialty metric over the fourth by at least this much.",
    control: {
      kind: 'number',
      integer: false,
      min: 1,
      minExclusive: true,
      max: 20,
      empty: 'No minimum',
    },
  },
  manualAbilityPath: {
    label: 'Active Ability path',
    help: 'Which path may unlock an Active Ability, an ability the player activates.',
    control: {
      kind: 'choice',
      choices: [
        { id: 'any', value: undefined, label: 'Any path' },
        { id: 'none', value: null, label: 'No path' },
        { id: 'path1', value: 'path1', label: 'Top path only' },
        { id: 'path2', value: 'path2', label: 'Middle path only' },
        { id: 'path3', value: 'path3', label: 'Bottom path only' },
      ],
    },
  },
  maxManualAbilityPaths: {
    label: 'Active Ability paths',
    help: 'How many paths may unlock an Active Ability; 0 means none.',
    control: { kind: 'number', integer: true, min: 0, minExclusive: false, max: 3 },
  },
};

export const policyKeys = Object.keys(policyFields) as PolicyKey[];

/** Which paths may unlock an Active Ability, as the Engine reads the two fields. */
export function activeAbilityLabel(
  policy: Pick<DesignPolicy, 'manualAbilityPath' | 'maxManualAbilityPaths'>,
): string {
  const { manualAbilityPath: path, maxManualAbilityPaths: max } = policy;
  if (path === null || max === 0) return 'No path';
  // A path the server would reject still reads as stored.
  if (path !== undefined)
    return path in pathNames ? `${pathNames[path]} path only` : `Only ${String(path)}`;
  return `Any path, at most ${max} ${max === 1 ? 'path' : 'paths'}`;
}

/** The problem with a number control's text, or null when it may be written. */
export function numberProblem(
  control: Extract<PolicyControl, { kind: 'number' }>,
  text: string,
): string | null {
  const trimmed = text.trim();
  if (trimmed === '' && control.empty) return null;
  const value = Number(trimmed);
  const lower = control.minExclusive ? 'above' : 'from';
  const upper = control.minExclusive ? 'and at most' : 'to';
  const bounds = `${lower} ${control.min} ${upper} ${control.max}`;
  const valid =
    trimmed !== '' &&
    Number.isFinite(value) &&
    (!control.integer || Number.isInteger(value)) &&
    (control.minExclusive ? value > control.min : value >= control.min) &&
    value <= control.max;
  if (valid) return null;
  const kind = control.integer ? 'a whole number' : 'a number';
  const empty = control.empty ? `, or leave it empty for ${control.empty.toLowerCase()}` : '';
  return `Enter ${kind} ${bounds}${empty}.`;
}

/** The value a control shows for a field. */
export function policyValueText(key: PolicyKey, policy: DesignPolicy): string {
  const control = policyFields[key].control;
  const value = policy[key];
  switch (control.kind) {
    case 'toggle':
      return value ? 'On' : 'Off';
    case 'number':
      return value === undefined ? (control.empty ?? 'Not set') : String(value);
    case 'choice':
      return control.choices.find((choice) => choice.value === value)?.label ?? String(value);
  }
}

/** Fields in a policy that are neither described nor fixed, for example from a newer server. */
export function undescribedPolicyKeys(policy: DesignPolicy): string[] {
  return Object.keys(policy).filter(
    (key) => !(key in policyFields) && !(fixedPolicyKeys as readonly string[]).includes(key),
  );
}

/**
 * A policy's rules as label and value rows for a read-only view. The two
 * Active Ability fields read as one row, as the Engine combines them.
 */
export function policySummary(policy: DesignPolicy): [string, string][] {
  const rows: [string, string][] = [];
  for (const key of policyKeys) {
    if (key === 'maxManualAbilityPaths') continue;
    if (key === 'manualAbilityPath') rows.push(['Active Ability', activeAbilityLabel(policy)]);
    else rows.push([policyFields[key].label, policyValueText(key, policy)]);
  }
  const others = undescribedPolicyKeys(policy);
  if (others.length) rows.push(['Other fields', others.join(', ')]);
  return rows;
}

/** A copy of the policy with one field set; `undefined`, or an optional toggle switched off, removes it. */
export function setPolicyValue(
  policy: DesignPolicy,
  key: PolicyKey,
  value: PolicyValue,
): DesignPolicy {
  const control = policyFields[key].control;
  const next: Record<string, unknown> = { ...policy };
  if (value === undefined || (control.kind === 'toggle' && control.optional && value === false))
    delete next[key];
  else next[key] = value;
  return next as unknown as DesignPolicy;
}

/** The design policy inside a mechanics Definition's JSON text. */
export type DefinitionPolicy =
  { state: 'invalid' } | { state: 'none' } | { state: 'policy'; policy: DesignPolicy };

export function readPolicy(definition: string): DefinitionPolicy {
  let parsed: unknown;
  try {
    parsed = JSON.parse(definition);
  } catch {
    return { state: 'invalid' };
  }
  const profile = (parsed as { profile?: unknown } | null)?.profile;
  if (!isObject(profile)) return { state: 'invalid' };
  const policy = profile.designPolicy;
  if (policy === undefined) return { state: 'none' };
  if (!isObject(policy)) return { state: 'invalid' };
  return { state: 'policy', policy: policy as unknown as DesignPolicy };
}

/** The Definition's JSON text with its design policy replaced; other fields keep their values and order. */
export function writePolicy(definition: string, policy: DesignPolicy): string {
  const parsed = JSON.parse(definition) as { profile: Record<string, unknown> };
  parsed.profile.designPolicy = policy;
  return JSON.stringify(parsed, null, 2);
}

function isObject(value: unknown): value is Record<string, unknown> {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

const errorPrefix = 'mechanicsDefinition.profile.designPolicy.';

/**
 * Splits a server validation message into the lines about each editable
 * field, without their path, and the remaining lines.
 */
export function policyErrors(message: string): {
  fields: Partial<Record<PolicyKey, string[]>>;
  rest: string[];
} {
  const fields: Partial<Record<PolicyKey, string[]>> = {};
  const rest: string[] = [];
  for (const line of message.split('\n').filter((line) => line.trim())) {
    const match = line.startsWith(errorPrefix)
      ? /^([A-Za-z0-9]+)[^:]*: (.*)$/.exec(line.slice(errorPrefix.length))
      : null;
    const key = match?.[1];
    if (match && key && key in policyFields) (fields[key as PolicyKey] ??= []).push(match[2]!);
    else rest.push(line);
  }
  return { fields, rest };
}
