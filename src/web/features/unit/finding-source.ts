import type { Finding } from '../../api/contract.js';

/** The rule of the Finding code adds when it cannot read a model finding's claim. */
export const humanReviewRule = 'review-claim-unread';

/** Names who produced a Finding, or that a person must read it. */
export function findingSource(finding: Pick<Finding, 'method' | 'rule'>): string {
  if (finding.method === 'model') return 'Model review';
  if (finding.rule === humanReviewRule) return 'Human review needed';
  return 'Structural check';
}
