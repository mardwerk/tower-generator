// The Profile rules editor: controls for the design policy's typed fields,
// described by policy-fields.ts, and their read-only summary.
import { useEffect, useId, useRef, useState } from 'react';
import type { DesignPolicy } from '../../api/contract.js';
import { fieldControl, Input } from '../../ui/input.js';
import { cn } from '../../ui/utils.js';
import {
  activeAbilityLabel,
  numberProblem,
  policyFields,
  policyKeys,
  policySummary,
  readPolicy,
  setPolicyValue,
  undescribedPolicyKeys,
  writePolicy,
  type PolicyControl,
  type PolicyKey,
  type PolicyValue,
} from './policy-fields.js';

export type PolicyIssues = Partial<Record<PolicyKey, string[]>>;

/** A Profile's design policy, read-only. */
export function DesignPolicySummary({ policy }: { policy: DesignPolicy | undefined }) {
  return (
    <section aria-labelledby="profile-design-policy" className="my-4">
      <h3 id="profile-design-policy" className="mb-2 text-sm font-semibold">
        Design policy
      </h3>
      {policy ? (
        <dl className="grid gap-x-6 gap-y-2.5 text-[13px] sm:grid-cols-[max-content_1fr]">
          {policySummary(policy).map(([term, value]) => (
            <div className="contents" key={term}>
              <dt className="text-muted-foreground">{term}</dt>
              <dd className="[overflow-wrap:anywhere]">{value}</dd>
            </div>
          ))}
        </dl>
      ) : (
        <p className="text-[13px] text-muted-foreground">This Profile has no design policy.</p>
      )}
    </section>
  );
}

/**
 * Controls for the design policy inside the Definition's JSON text. The text
 * stays the one source: each change rewrites its `profile.designPolicy`.
 */
export function DesignPolicyEditor({
  definition,
  issues,
  onChange,
}: {
  definition: string;
  issues: PolicyIssues;
  onChange: (definition: string, key: PolicyKey) => void;
}) {
  const read = readPolicy(definition);
  const body = () => {
    if (read.state === 'invalid')
      return (
        <p className="text-[13px] text-muted-foreground">
          Fix the Definition JSON to edit the design policy.
        </p>
      );
    if (read.state === 'none')
      return (
        <p className="text-[13px] text-muted-foreground">
          This Profile has no design policy. Duplicate the Default to start from its rules.
        </p>
      );
    const policy = read.policy;
    const set = (key: PolicyKey, value: PolicyValue) =>
      onChange(writePolicy(definition, setPolicyValue(policy, key, value)), key);
    const others = undescribedPolicyKeys(policy);
    return (
      <>
        <div className="grid gap-x-6 sm:grid-cols-2">
          {policyKeys.map((key) => (
            <PolicyControlField
              key={key}
              name={key}
              value={policy[key]}
              issues={issues[key] ?? []}
              onChange={(value) => set(key, value)}
            />
          ))}
        </div>
        <p className="mt-1 text-[13px]">
          <span className="text-muted-foreground">Active Ability: </span>
          {activeAbilityLabel(policy)}
        </p>
        {others.length > 0 && (
          <p className="mt-1 text-xs text-muted-foreground">
            Also set, and edited only in the Definition JSON: {others.join(', ')}.
          </p>
        )}
      </>
    );
  };
  return (
    <fieldset className="my-3 min-w-0 rounded-lg border border-border px-4 pt-1 pb-3">
      <legend className="px-1 text-sm font-semibold">Design policy</legend>
      <p className="text-xs text-muted-foreground">
        The Engine checks these rules during generation. Editing them rewrites the Definition's{' '}
        <code>profile.designPolicy</code> below.
      </p>
      {body()}
    </fieldset>
  );
}

function PolicyControlField({
  name,
  value,
  issues,
  onChange,
}: {
  name: PolicyKey;
  value: PolicyValue;
  issues: string[];
  onChange: (value: PolicyValue) => void;
}) {
  const field = policyFields[name];
  const id = useId();
  const help = `${id}-help`;
  const control = field.control;
  const problems = (problem: string | null) => [...(problem ? [problem] : []), ...issues];
  if (control.kind === 'toggle')
    return (
      <div className="my-2.5 flex flex-col gap-1 text-[13px]" data-policy-field={name}>
        <label className="flex items-center gap-2">
          <input
            type="checkbox"
            name={name}
            className="size-4 accent-primary"
            checked={value === true}
            aria-describedby={help}
            aria-invalid={issues.length > 0 || undefined}
            onChange={(event) => onChange(event.target.checked)}
          />
          <span>{field.label}</span>
        </label>
        <Messages id={help} help={field.help} problems={problems(null)} />
      </div>
    );
  if (control.kind === 'choice') {
    const selected = control.choices.find((choice) => choice.value === value);
    return (
      <label className="my-2.5 flex flex-col gap-1.5 text-[13px]" data-policy-field={name}>
        <span>{field.label}</span>
        <select
          name={name}
          className={cn(fieldControl, 'h-9 py-1')}
          value={selected?.id ?? ''}
          aria-describedby={help}
          aria-invalid={issues.length > 0 || undefined}
          onChange={(event) =>
            onChange(control.choices.find((choice) => choice.id === event.target.value)?.value)
          }
        >
          {!selected && <option value="">{String(value)}</option>}
          {control.choices.map((choice) => (
            <option key={choice.id} value={choice.id}>
              {choice.label}
            </option>
          ))}
        </select>
        <Messages id={help} help={field.help} problems={problems(null)} />
      </label>
    );
  }
  return (
    <NumberField
      name={name}
      label={field.label}
      help={field.help}
      control={control}
      value={typeof value === 'number' ? value : undefined}
      issues={issues}
      onChange={onChange}
    />
  );
}

/**
 * A number keeps its own text while typed and writes only a valid value. An
 * invalid entry blocks the form's submission with its problem.
 */
function NumberField({
  name,
  label,
  help,
  control,
  value,
  issues,
  onChange,
}: {
  name: PolicyKey;
  label: string;
  help: string;
  control: Extract<PolicyControl, { kind: 'number' }>;
  value: number | undefined;
  issues: string[];
  onChange: (value: number | undefined) => void;
}) {
  const id = useId();
  const [text, setText] = useState(value === undefined ? '' : String(value));
  const input = useRef<HTMLInputElement>(null);
  useEffect(() => {
    // Follow an edit of the Definition JSON; the typed text already shows a valid value.
    setText((current) =>
      (current.trim() === '' ? undefined : Number(current)) === value
        ? current
        : value === undefined
          ? ''
          : String(value),
    );
  }, [value]);
  const problem = numberProblem(control, text);
  useEffect(() => input.current?.setCustomValidity(problem ?? ''), [problem]);
  return (
    <label className="my-2.5 flex flex-col gap-1.5 text-[13px]" data-policy-field={name}>
      <span>{label}</span>
      <Input
        ref={input}
        name={name}
        inputMode={control.integer ? 'numeric' : 'decimal'}
        value={text}
        placeholder={control.empty}
        aria-describedby={`${id}-help`}
        aria-invalid={problem !== null || issues.length > 0 || undefined}
        onChange={(event) => {
          const next = event.target.value;
          setText(next);
          if (numberProblem(control, next) === null)
            onChange(next.trim() === '' ? undefined : Number(next));
        }}
      />
      <Messages
        id={`${id}-help`}
        help={help}
        problems={[...(problem ? [problem] : []), ...issues]}
      />
    </label>
  );
}

function Messages({ id, help, problems }: { id: string; help: string; problems: string[] }) {
  return (
    <span id={id} className="flex flex-col gap-0.5 text-xs">
      <span className="text-muted-foreground">{help}</span>
      {problems.map((problem) => (
        <span key={problem} className="text-destructive" role="alert">
          {problem}
        </span>
      ))}
    </span>
  );
}
