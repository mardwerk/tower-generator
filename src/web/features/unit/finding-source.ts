import type { Finding } from '../../api/contract.js';

/** The rule of the Finding code adds when it cannot read a model finding's claim. */
export const humanReviewRule = 'review-claim-unread';

/** The rules of the Findings that record the review's verdict on each whole-technique omission, each path's third and fifth purchase and each proposed mechanic of a purchase. */
export const omissionVerdictRule = 'omission-verdict';
export const pathIdentityVerdictRule = 'path-identity-verdict';
export const capstoneVerdictRule = 'capstone-verdict';
export const proposalVerdictRule = 'proposal-verdict';

/** Names who produced a Finding, or that a person must read it. */
export function findingSource(finding: Pick<Finding, 'method' | 'rule'>): string {
  if (finding.method === 'model') return 'Model review';
  if (finding.rule === humanReviewRule) return 'Human review needed';
  return 'Structural check';
}

/** Whether a Finding records a review verdict, which the report lists in its own section. */
export function isReviewVerdict(finding: Pick<Finding, 'method' | 'rule'>): boolean {
  return (
    finding.method === 'model' &&
    (finding.rule === omissionVerdictRule ||
      finding.rule === pathIdentityVerdictRule ||
      finding.rule === capstoneVerdictRule ||
      finding.rule === proposalVerdictRule)
  );
}

/** A Result's review verdicts under one rule, passes included, in recorded order. */
export function reviewVerdicts<T extends Pick<Finding, 'method' | 'rule'>>(
  findings: T[],
  rule: string,
): T[] {
  return findings.filter((finding) => finding.method === 'model' && finding.rule === rule);
}
