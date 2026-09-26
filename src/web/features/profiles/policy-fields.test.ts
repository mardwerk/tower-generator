// The Profile rules editor: the Active Ability label, the descriptor's
// controls and a save round trip through a scripted Profiles API.
import { createElement } from 'react';
import { renderToStaticMarkup } from 'react-dom/server';
import type {
  DesignPolicy,
  MechanicsDefinition,
  ProfilesState,
  UnitProfile,
} from '../../api/contract.js';
import { LabApiError } from '../../api/client.js';
import { DesignPolicyEditor } from './design-policy.js';
import {
  activeAbilityLabel,
  numberProblem,
  policyErrors,
  policyFields,
  policyKeys,
  policySummary,
  readPolicy,
  setPolicyValue,
  writePolicy,
  type PolicyKey,
} from './policy-fields.js';
import { editedProfile, savedRules } from './profiles.js';

function equal(actual: unknown, expected: unknown, what: string) {
  const [a, b] = [JSON.stringify(actual), JSON.stringify(expected)];
  if (a !== b) throw new Error(`${what}:\n  got      ${a}\n  expected ${b}`);
}

/** The bundled Default's design policy. */
const defaultPolicy: DesignPolicy = {
  version: '1',
  distinctPathSpecializations: false,
  distinctFirstUpgrades: true,
  distinctCapstones: true,
  preserveEarlyAttackIdentity: true,
  maxManualAbilityPaths: 1,
  manualAbilityPath: 'path2',
  tier5Uniqueness: 'one-per-player-unit-type-and-path',
};

const definition = (policy: DesignPolicy | undefined) =>
  ({
    version: '2',
    id: 'default-td',
    revision: 'test',
    label: 'Default',
    balanceStatus: 'unbalanced',
    progression: { paths: 3 },
    rules: { manualBoostUnlockTier: 4 },
    profile: {
      currency: 'Money',
      ...(policy ? { designPolicy: policy } : {}),
      maxStatValue: 1000000,
      maxChangesPerTier: 3,
      earlyTierMaxChanges: 2,
    },
  }) as MechanicsDefinition;

const text = (policy: DesignPolicy | undefined) => JSON.stringify(definition(policy), null, 2);

function policyOf(definitionText: string): DesignPolicy {
  const read = readPolicy(definitionText);
  if (read.state !== 'policy') throw new Error(`no policy: ${read.state}`);
  return read.policy;
}

/**
 * A scripted `profiles/save` that checks the policy fields as the server's
 * schema does and answers with its messages, one per line.
 */
function scriptedProfiles() {
  const saved: UnitProfile[] = [];
  const bodies: unknown[] = [];
  const prefix = 'mechanicsDefinition.profile.designPolicy.';
  return {
    saved,
    bodies,
    async call(endpoint: string, body: { profile: UnitProfile }): Promise<ProfilesState> {
      if (endpoint !== 'profiles/save') throw new Error(`unexpected ${endpoint}`);
      bodies.push(structuredClone(body));
      const policy = body.profile.mechanicsDefinition.profile.designPolicy!;
      const problems: string[] = [];
      if (policy.maxManualAbilityPaths > 3)
        problems.push(`${prefix}maxManualAbilityPaths: Too big: expected number to be <=3`);
      const multiplier = policy.minTier5SpecialtyMultiplier;
      if (multiplier !== undefined && multiplier <= 1)
        problems.push(`${prefix}minTier5SpecialtyMultiplier: Too small: expected number to be >1`);
      if (problems.length) throw new LabApiError(problems.join('\n'));
      saved.push(structuredClone(body.profile));
      return {
        directory: 'profiles',
        profiles: saved.map((profile) => ({
          profile,
          builtIn: false,
          progression: { paths: [], maxActivePaths: 2 } as unknown as never,
        })),
      };
    },
  };
}

export const tests: Record<string, () => Promise<void>> = {
  async 'the Active Ability label reads absent, null, a path and a zero maximum'() {
    const label = (manualAbilityPath: DesignPolicy['manualAbilityPath'] | 'absent', max: number) =>
      activeAbilityLabel(
        manualAbilityPath === 'absent'
          ? { maxManualAbilityPaths: max }
          : { manualAbilityPath, maxManualAbilityPaths: max },
      );
    equal(label('absent', 1), 'Any path, at most 1 path', 'absent, maximum 1');
    equal(label('absent', 2), 'Any path, at most 2 paths', 'absent, maximum 2');
    equal(label(null, 1), 'No path', 'null');
    equal(label('path2', 1), 'Middle path only', 'middle path');
    equal(label('path1', 3), 'Top path only', 'top path, maximum 3');
    equal(label('path3', 1), 'Bottom path only', 'bottom path');
    equal(label('path2', 0), 'No path', 'a path with maximum 0');
    equal(label('absent', 0), 'No path', 'absent with maximum 0');
    equal(label(null, 0), 'No path', 'null with maximum 0');
    equal(label('path4' as 'path1', 1), 'Only path4', 'a path the server rejects');
  },
  async 'the descriptor covers every editable field and excludes the fixed literals'() {
    equal(
      policyKeys,
      [
        'distinctPathSpecializations',
        'distinctFirstUpgrades',
        'distinctCapstones',
        'preserveEarlyAttackIdentity',
        'requireTier3BehaviorChange',
        'requireTier5BehaviorChange',
        'minTier5SpecialtyMultiplier',
        'manualAbilityPath',
        'maxManualAbilityPaths',
      ],
      'descriptor keys',
    );
    equal('version' in policyFields || 'tier5Uniqueness' in policyFields, false, 'fixed literals');
    const choices = policyFields.manualAbilityPath.control;
    equal(
      choices.kind === 'choice'
        ? choices.choices.map((choice) => (choice.value === undefined ? 'absent' : choice.value))
        : [],
      ['absent', null, 'path1', 'path2', 'path3'],
      'Active Ability path values',
    );
  },
  async 'each descriptor entry renders its control with the stored value'() {
    const markup = renderToStaticMarkup(
      createElement(DesignPolicyEditor, {
        definition: text({ ...defaultPolicy, requireTier3BehaviorChange: true }),
        issues: { minTier5SpecialtyMultiplier: ['Too small: expected number to be >1'] },
        onChange: () => undefined,
      }),
    );
    const element = (key: PolicyKey) => {
      const match = new RegExp(`<(input|select)[^>]*name="${key}"[^>]*>`).exec(markup);
      if (!match) throw new Error(`no control for ${key}`);
      return match[0];
    };
    for (const key of policyKeys) {
      if (!markup.includes(`data-policy-field="${key}"`)) throw new Error(`no field for ${key}`);
      if (!markup.includes(`>${policyFields[key].label}<`)) throw new Error(`no label for ${key}`);
      const control = policyFields[key].control;
      const tag = element(key);
      if (control.kind === 'toggle' && !tag.includes('type="checkbox"'))
        throw new Error(`${key} is not a checkbox`);
      if (control.kind === 'choice' && !tag.startsWith('<select'))
        throw new Error(`${key} is not a select`);
      if (control.kind === 'number' && !tag.includes('inputMode='))
        throw new Error(`${key} is not a number field`);
    }
    const checked = policyKeys.filter(
      (key) => policyFields[key].control.kind === 'toggle' && element(key).includes('checked'),
    );
    equal(
      checked,
      [
        'distinctFirstUpgrades',
        'distinctCapstones',
        'preserveEarlyAttackIdentity',
        'requireTier3BehaviorChange',
      ],
      'checked toggles',
    );
    equal(
      /<option value="([^"]+)" selected=""/.exec(markup)?.[1],
      'path2',
      'selected Active Ability path',
    );
    equal(element('maxManualAbilityPaths').includes('value="1"'), true, 'maximum paths');
    equal(element('minTier5SpecialtyMultiplier').includes('value=""'), true, 'no multiplier');
    equal(
      markup.includes('Too small: expected number to be &gt;1'),
      true,
      'server issue beside the multiplier',
    );
    equal(markup.includes('Middle path only'), true, 'Active Ability summary');
  },
  async 'the editor explains a Definition it cannot edit'() {
    const render = (definition: string) =>
      renderToStaticMarkup(
        createElement(DesignPolicyEditor, { definition, issues: {}, onChange: () => undefined }),
      );
    equal(render('{').includes('Fix the Definition JSON'), true, 'invalid JSON');
    equal(render(text(undefined)).includes('has no design policy'), true, 'no policy');
    equal(render(text(undefined)).includes('<input'), false, 'no controls without a policy');
  },
  async 'values are written into the Definition and optional toggles are removed'() {
    const start = text(defaultPolicy);
    equal(writePolicy(start, policyOf(start)), start, 'an untouched policy keeps its text');
    let policy = setPolicyValue(defaultPolicy, 'preserveEarlyAttackIdentity', false);
    equal('preserveEarlyAttackIdentity' in policy, false, 'optional toggle off removes it');
    policy = setPolicyValue(policy, 'distinctCapstones', false);
    equal(policy.distinctCapstones, false, 'required toggle off writes false');
    policy = setPolicyValue(policy, 'manualAbilityPath', undefined);
    equal('manualAbilityPath' in policy, false, 'Any path removes the path');
    policy = setPolicyValue(policy, 'manualAbilityPath', null);
    equal(policy.manualAbilityPath, null, 'No path writes null');
    const next = policyOf(writePolicy(start, policy));
    equal(next, policy, 'the written Definition holds the policy');
    equal(
      JSON.parse(writePolicy(start, policy)).profile.currency,
      'Money',
      'other Definition fields stay',
    );
  },
  async 'number fields accept only values within their bounds'() {
    const multiplier = policyFields.minTier5SpecialtyMultiplier.control;
    const maximum = policyFields.maxManualAbilityPaths.control;
    if (multiplier.kind !== 'number' || maximum.kind !== 'number') throw new Error('not numbers');
    equal(numberProblem(multiplier, ''), null, 'empty multiplier');
    equal(numberProblem(multiplier, '2.5'), null, 'multiplier 2.5');
    equal(numberProblem(multiplier, '20'), null, 'multiplier 20');
    equal(
      numberProblem(multiplier, '1'),
      'Enter a number above 1 and at most 20, or leave it empty for no minimum.',
      'multiplier 1',
    );
    equal(numberProblem(multiplier, 'x') !== null, true, 'multiplier text');
    equal(numberProblem(maximum, '0'), null, 'maximum 0');
    equal(numberProblem(maximum, '3'), null, 'maximum 3');
    equal(numberProblem(maximum, '4'), 'Enter a whole number from 0 to 3.', 'maximum 4');
    equal(numberProblem(maximum, '1.5') !== null, true, 'maximum 1.5');
    equal(numberProblem(maximum, '') !== null, true, 'empty maximum');
  },
  async 'server errors are placed beside their field'() {
    const prefix = 'mechanicsDefinition.profile.designPolicy';
    equal(
      policyErrors(
        [
          `${prefix}.minTier5SpecialtyMultiplier: Too small: expected number to be >1`,
          `${prefix}.manualAbilityPath: Invalid option: expected one of "path1"|"path2"|"path3"`,
          `${prefix}.version: Invalid input: expected "1"`,
          `${prefix}: Unrecognized key: "distinctEarlyBenefits"`,
          'id: Too small',
        ].join('\n'),
      ),
      {
        fields: {
          minTier5SpecialtyMultiplier: ['Too small: expected number to be >1'],
          manualAbilityPath: ['Invalid option: expected one of "path1"|"path2"|"path3"'],
        },
        rest: [
          `${prefix}.version: Invalid input: expected "1"`,
          `${prefix}: Unrecognized key: "distinctEarlyBenefits"`,
          'id: Too small',
        ],
      },
      'split errors',
    );
  },
  async 'the read-only summary combines the Active Ability fields'() {
    equal(
      policySummary({ ...defaultPolicy, minTier5SpecialtyMultiplier: 3 }),
      [
        ['Distinct specializations', 'Off'],
        ['Distinct first purchases', 'On'],
        ['Distinct capstones', 'On'],
        ['Early attack identity', 'On'],
        ['Behavior at the third purchase', 'Off'],
        ['Behavior at the fifth purchase', 'Off'],
        ['Capstone multiplier', '3'],
        ['Active Ability', 'Middle path only'],
      ],
      'summary',
    );
    const newer = { ...defaultPolicy, distinctEarlyBenefits: true } as DesignPolicy;
    equal(policySummary(newer).at(-1), ['Other fields', 'distinctEarlyBenefits'], 'other fields');
  },
  async 'an edited policy saves through profiles/save, and its errors come back scoped'() {
    const api = scriptedProfiles();
    const form = {
      id: 'luffy-td',
      name: 'Luffy rules',
      task: 'Make a unit.',
      rules: 'Rules text.',
      definition: text(defaultPolicy),
    };
    const edit = (key: PolicyKey, value: DesignPolicy[PolicyKey]) => {
      form.definition = writePolicy(
        form.definition,
        setPolicyValue(policyOf(form.definition), key, value),
      );
    };
    edit('manualAbilityPath', null);
    edit('maxManualAbilityPaths', 0);
    edit('minTier5SpecialtyMultiplier', 1);
    const rejected = editedProfile(form);
    if (typeof rejected === 'string') throw new Error(rejected);
    const message = await api.call('profiles/save', { profile: rejected }).then(
      () => '',
      (reason: unknown) => (reason as Error).message,
    );
    equal(
      policyErrors(message),
      {
        fields: { minTier5SpecialtyMultiplier: ['Too small: expected number to be >1'] },
        rest: [],
      },
      'scoped rejection',
    );
    edit('minTier5SpecialtyMultiplier', 3);
    const accepted = editedProfile(form);
    if (typeof accepted === 'string') throw new Error(accepted);
    const state = await api.call('profiles/save', { profile: accepted });
    equal(api.bodies.at(-1), { profile: accepted }, 'posted body');
    equal(
      state.profiles[0]!.profile,
      {
        schemaVersion: '2',
        kind: 'profile',
        id: 'luffy-td',
        name: 'Luffy rules',
        task: 'Make a unit.',
        rules: savedRules('luffy-td', 'Rules text.'),
        mechanicsDefinition: definition({
          ...defaultPolicy,
          maxManualAbilityPaths: 0,
          manualAbilityPath: null,
          minTier5SpecialtyMultiplier: 3,
        }),
      },
      'saved Profile',
    );
    equal(
      activeAbilityLabel(state.profiles[0]!.profile.mechanicsDefinition.profile.designPolicy!),
      'No path',
      'saved Active Ability',
    );
  },
};
