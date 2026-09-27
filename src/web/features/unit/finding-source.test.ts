// Who produced a Finding, as the unit view labels it.
import type { Finding } from '../../api/contract.js';
import {
  findingSource,
  isReviewVerdict,
  omissionVerdictRule,
  pathIdentityVerdictRule,
  reviewVerdicts,
} from './finding-source.js';

function equal(actual: string, expected: string, what: string) {
  if (actual !== expected) throw new Error(`${what}: got ${actual}, expected ${expected}`);
}

export const tests: Record<string, () => Promise<void>> = {
  'a finding code could not read asks for a person': async () => {
    equal(
      findingSource({ method: 'model', rule: 'source-fit' }),
      'Model review',
      'a model finding',
    );
    equal(
      findingSource({ method: 'deterministic', rule: 'typed-mechanics' }),
      'Structural check',
      'a structural check',
    );
    equal(
      findingSource({ method: 'deterministic', rule: 'review-claim-unread' }),
      'Human review needed',
      'an unread review claim',
    );
  },
  'a review verdict is listed apart, a pass included': async () => {
    const findings: Pick<Finding, 'method' | 'rule' | 'outcome'>[] = [
      { method: 'deterministic', rule: 'typed-mechanics', outcome: 'pass' },
      { method: 'model', rule: omissionVerdictRule, outcome: 'fail' },
      { method: 'model', rule: pathIdentityVerdictRule, outcome: 'pass' },
      { method: 'model', rule: 'source-fit', outcome: 'fail' },
    ];
    equal(
      findings.map((finding) => String(isReviewVerdict(finding))).join(','),
      'false,true,true,false',
      'verdicts',
    );
    equal(
      reviewVerdicts(findings, omissionVerdictRule)
        .map((finding) => finding.outcome)
        .join(','),
      'fail',
      'omission verdicts',
    );
    equal(
      reviewVerdicts(findings, pathIdentityVerdictRule)
        .map((finding) => finding.outcome)
        .join(','),
      'pass',
      'third purchase verdicts',
    );
  },
};
