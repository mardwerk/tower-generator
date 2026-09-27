// A purchase's plan intent and the reserved techniques of a typed unit (#35).
import { createElement } from 'react';
import { renderToStaticMarkup } from 'react-dom/server';
import type { Ability, Purchase } from '../../api/contract.js';
import { ProposedMechanics, PurchasePlan, splitAbilities } from './purchase-plan.js';

function contains(markup: string, want: string, what: string) {
  if (!markup.includes(want)) throw new Error(`${what}: missing ${want}\n  in ${markup}`);
}

const purchase = (fields: Partial<Purchase>): Purchase => ({
  code: 'x-x-1',
  name: 'Boundman Force',
  cost: 120,
  effects: ['Raises damage from 1 to 2 (+1).'],
  text: '',
  ...fields,
});

const ability = (id: string, placement: Ability['placement']): Ability =>
  ({ id, name: id, placement, status: 'proposed', description: `${id} reason` }) as Ability;

export const tests = {
  async 'a purchase shows its planned technique and a name flag as plan intent'() {
    const markup = renderToStaticMarkup(
      createElement(PurchasePlan, {
        purchase: purchase({
          technique: 'Gum-Gum Pistol',
          baseTechnique: true,
          nameMentions: ['Gear 4 Boundman', 'Armament Haki'],
        }),
      }),
    );
    contains(
      markup,
      '<span class="font-medium">Plan</span> · adapts Gum-Gum Pistol (base attack)',
      'plan line',
    );
    contains(
      markup,
      'Review: the name suggests Gear 4 Boundman and Armament Haki, which the plan does not map to this purchase.',
      'name flag',
    );
  },
  async 'a purchase from a plan without typed techniques shows no plan line'() {
    const markup = renderToStaticMarkup(createElement(PurchasePlan, { purchase: purchase({}) }));
    if (markup !== '') throw new Error(`expected no plan markup, got ${markup}`);
  },
  async 'a purchase shows its proposed mechanics as not yet supported'() {
    const markup = renderToStaticMarkup(
      createElement(ProposedMechanics, {
        purchase: purchase({
          proposedMechanics: [
            {
              name: 'Boundman bounce',
              effect: 'Each punch bounces to one more enemy within 12 of its target.',
              sourceIds: ['source3:83'],
            },
          ],
        }),
      }),
    );
    contains(
      markup,
      '<span class="font-medium">Proposed (not yet supported)</span> · Boundman bounce: Each punch bounces to one more enemy within 12 of its target.',
      'proposed mechanic',
    );
    const none = renderToStaticMarkup(createElement(ProposedMechanics, { purchase: purchase({}) }));
    if (none !== '') throw new Error(`expected no proposed markup, got ${none}`);
  },
  async 'a typed unit lists reserved and omitted techniques apart from its other abilities'() {
    const abilities = [
      ability('path-2-active', 'upgrade'),
      ability('reserved-1', 'reserved'),
      ability('omitted-1', 'omitted'),
      ability('form', 'conditional'),
    ];
    const assigned = new Set(['path-2-active']);
    const typed = splitAbilities(abilities, assigned, true);
    const ids = (list: Ability[]) => list.map((entry) => entry.id).join(',');
    if (ids(typed.reserved) !== 'reserved-1,omitted-1' || ids(typed.remaining) !== 'form')
      throw new Error(`typed: ${ids(typed.reserved)} / ${ids(typed.remaining)}`);
    const prose = splitAbilities(abilities, assigned, false);
    if (ids(prose.reserved) !== '' || ids(prose.remaining) !== 'reserved-1,omitted-1,form')
      throw new Error(`prose: ${ids(prose.reserved)} / ${ids(prose.remaining)}`);
  },
};
