// The Profiles tab's vocabulary facts list the properties that accept bonus
// damage by name, each once and apart from the others.
import { createElement } from 'react';
import { renderToStaticMarkup } from 'react-dom/server';
import type { Vocabulary } from '../../api/contract.js';
import { VocabularyFacts } from './profiles.js';

function bonusRow(vocabulary: Vocabulary): string {
  const html = renderToStaticMarkup(createElement(VocabularyFacts, { vocabulary, ceiling: 1e6 }));
  const match = /<dt[^>]*>Bonus damage against<\/dt><dd[^>]*>([^<]*)<\/dd>/.exec(html);
  if (match?.[1] === undefined) throw new Error(`no bonus damage row in ${html}`);
  return match[1];
}

const vocabulary: Vocabulary = {
  enemyProperties: [
    { id: 'lead', name: 'Lead', description: '' },
    { id: 'hardened', name: 'Hardened', description: '' },
    { id: 'blimp', name: 'Blimp', description: '' },
  ],
  damageTypes: [{ id: 'sharp', name: 'Sharp', description: '', ineffectiveAgainst: ['lead'] }],
  targeting: [{ id: 'first', name: 'First', description: '' }],
  detection: [],
  statusEffects: [],
};

export const tests: Record<string, () => Promise<void>> = {
  'the Default lists Hardened and Blimp as separate bonus damage properties': async () => {
    const row = bonusRow({ ...vocabulary, bonusDamageProperties: ['hardened', 'blimp'] });
    if (row !== 'Hardened, Blimp') throw new Error(`bonus damage against: ${row}`);
  },
  'a vocabulary without bonus damage properties shows none': async () => {
    const row = bonusRow(vocabulary);
    if (row !== 'none') throw new Error(`bonus damage against: ${row}`);
  },
};
