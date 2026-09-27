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
  'a change shows both values and its delta, as the unit sheet words it': async () => {
    const html = renderToStaticMarkup(
      createElement(StatValues, {
        changes: [
          { key: 'damage', before: 1, after: 2, delta: '+1', improvement: true },
          {
            key: 'intervalSeconds',
            before: 0.95,
            after: 0.8075,
            delta: 'attacks 18% faster',
            improvement: true,
          },
        ],
      }),
    );
    const rows = html.split('</li>').slice(0, -1);
    const damage = row(rows, 'Damage');
    if (!/1<\/span>.*2<\/span>.*\(\+1\)/.test(damage)) throw new Error(`damage reads ${damage}`);
    const interval = row(rows, 'Attack interval');
    if (
      !interval.includes('0.95 s') ||
      !interval.includes('0.8075 s') ||
      !interval.includes('(attacks 18% faster)')
    )
      throw new Error(`the interval reads ${interval}`);
    if (interval.includes('Slower')) throw new Error(`a faster interval reads ${interval}`);
  },
  'Active Ability multipliers read as percentages, never ×': async () => {
    const html = renderToStaticMarkup(
      createElement(StatValues, {
        changes: [
          {
            key: 'damageMultiplier',
            before: '+100%',
            after: '+200%',
            delta: '+100 percentage points',
            improvement: true,
          },
          { key: 'intervalMultiplier', after: '100% faster' },
        ],
      }),
    );
    const rows = html.split('</li>').slice(0, -1);
    const damage = row(rows, 'Active damage bonus');
    if (
      !damage.includes('+100%') ||
      !damage.includes('+200%') ||
      !damage.includes('(+100 percentage points)')
    )
      throw new Error(`the damage bonus reads ${damage}`);
    if (!row(rows, 'Active attack speed').includes('100% faster'))
      throw new Error('the attack speed is not a percentage');
    if (html.includes('×')) throw new Error(`a multiplier reads as ×: ${html}`);
    equal(statValue('intervalMultiplier', 0.5), '0.5', 'a bare number is not rewritten');
  },
  'other stats keep their units': async () => {
    equal(statValue('damage', 20), '20', 'damage');
    equal(statValue('intervalSeconds', 0.95), '0.95 s', 'an interval');
    equal(statValue('damage', 30, 'percent', 'moveSpeed'), '30%', 'a slow');
  },
};
