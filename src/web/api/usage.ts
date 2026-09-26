/** Display of reported usage. Totals come only from provider reports. */
import type { AuthorResult, CheckedArtifact, DraftArtifact, PreparedRequest } from './contract.js';

type UsageArtifact = AuthorResult | CheckedArtifact | DraftArtifact | PreparedRequest;
type Run = DraftArtifact['run'];

export interface StageUsage {
  stage: 'Draft' | 'Review';
  run: Pick<Run, 'modelId' | 'usage'> | null;
  status?: 'completed' | 'unavailable';
}

export interface UsageTotal {
  value: number | null;
  partial: boolean;
}

export interface UsageSummary {
  stages: StageUsage[];
  cost: UsageTotal;
  tokens: UsageTotal;
}

/** Only this revision contributes. Estimates stay separate from reported charges. */
export function summarizeUsage(artifact: UsageArtifact): UsageSummary {
  const draft =
    artifact.kind === 'prepared'
      ? null
      : artifact.kind === 'checked'
        ? artifact.draft.run
        : artifact.kind === 'result'
          ? artifact.run.draft
          : artifact.run;
  const review = artifact.kind === 'result' ? artifact.run.review : null;
  const stages: StageUsage[] = [
    { stage: 'Draft', run: draft },
    { stage: 'Review', run: review },
  ];
  const completed = stages.filter(({ run }) => run !== null);
  return {
    stages,
    cost: total(completed.map(({ run }) => run?.usage?.costUsd ?? null)),
    tokens: total(completed.map(({ run }) => run?.usage?.totalTokens ?? null)),
  };
}

function total(values: (number | null)[]): UsageTotal {
  const reported = values.filter((value): value is number => value !== null);
  return {
    value: reported.length ? reported.reduce((sum, value) => sum + value, 0) : null,
    partial: reported.length < values.length,
  };
}

export function formatCost(value: number | null): string {
  if (value === null) return 'Unavailable';
  if (value > 0 && value < 0.00000001) return `$${value.toExponential(3)} USD`;
  return `$${value.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 8 })} USD`;
}
