// Kit stat values as UnitLab shows them.
import { statValue } from './kit-stats.js';

function equal(actual: string, expected: string, what: string) {
  if (actual !== expected) throw new Error(`${what}: got ${actual}, expected ${expected}`);
}

export const tests: Record<string, () => Promise<void>> = {
  'bonus damage reads as +N, like the unit sheet': async () => {
    equal(statValue('damage', 50, undefined, 'bonusDamage'), '+50', 'a new bonus');
    equal(statValue('damage', 12.5, undefined, 'bonusDamage'), '+12.5', 'a fractional bonus');
  },
  'other stats keep their units': async () => {
    equal(statValue('damage', 20), '20', 'damage');
    equal(statValue('intervalSeconds', 0.95), '0.95 s', 'an interval');
    equal(statValue('damage', 30, 'percent', 'moveSpeed'), '30%', 'a slow');
  },
};
