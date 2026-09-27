// Kit stat values as UnitLab shows them.
import { createElement } from 'react';
import { renderToStaticMarkup } from 'react-dom/server';
import { StatValues, statValue } from './kit-stats.js';

function equal(actual: string, expected: string, what: string) {
  if (actual !== expected) throw new Error(`${what}: got ${actual}, expected ${expected}`);
}

/** The rendered row whose title is the label, failing when none has it. */
function row(rows: string[], label: string): string {
  const found = rows.find((item) => item.includes(`title="${label}"`));
  if (found === undefined) throw new Error(`no row titled ${label}`);
  return found;
}

export const tests: Record<string, () => Promise<void>> = {
  'bonus damage reads as +N, like the unit sheet': async () => {
    equal(statValue('damage', 50, undefined, 'bonusDamage'), '+50', 'a new bonus');
    equal(statValue('damage', 12.5, undefined, 'bonusDamage'), '+12.5', 'a fractional bonus');
  },
  'Hardened and Blimp bonuses show apart, each with its own label': async () => {
    const html = renderToStaticMarkup(
      createElement(StatValues, {
        changes: [
          {
            key: 'bonusDamage.hardened',
            label: 'Damage against Hardened',
            kind: 'bonusDamage',
            after: 3,
          },
          {
            key: 'bonusDamage.blimp',
            label: 'Damage against Blimp',
            kind: 'bonusDamage',
            before: 1,
            after: 4,
            improvement: true,
          },
        ],
      }),
    );
    const rows = html.split('</li>').slice(0, -1);
    equal(String(rows.length), '2', 'bonus damage rows');
    const hardened = row(rows, 'Damage against Hardened');
    const blimp = row(rows, 'Damage against Blimp');
    if (!hardened.includes('+3') || hardened.includes('Blimp'))
      throw new Error(`the Hardened row reads ${hardened}`);
    if (!blimp.includes('+1') || !blimp.includes('+4') || blimp.includes('Hardened'))
      throw new Error(`the Blimp row reads ${blimp}`);
  },
  'other stats keep their units': async () => {
    equal(statValue('damage', 20), '20', 'damage');
    equal(statValue('intervalSeconds', 0.95), '0.95 s', 'an interval');
    equal(statValue('damage', 30, 'percent', 'moveSpeed'), '30%', 'a slow');
  },
};
