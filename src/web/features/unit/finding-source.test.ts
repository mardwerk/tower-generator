// Who produced a Finding, as the unit view labels it.
import { findingSource } from './finding-source.js';

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
};
