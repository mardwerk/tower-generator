// A purchase's plan intent, shown apart from its resolved effects, and the
// techniques no build grants (#35).
import type { Ability, Purchase } from '../../api/contract.js';

/** Reads ["a"], ["a", "b"] and ["a", "b", "c"] as "a", "a and b" and "a, b and c". */
function listText(items: string[]): string {
  return items.length < 2
    ? items.join('')
    : `${items.slice(0, -1).join(', ')} and ${items[items.length - 1]}`;
}

/**
 * The technique the plan maps a purchase to, the plan's adaptation and any
 * name that points to another technique. All three are plan intent; the
 * resolved effects follow separately. A purchase from a plan without typed
 * techniques shows only its adaptation, if any.
 */
export function PurchasePlan({ purchase }: { purchase: Purchase | undefined }) {
  if (!purchase) return null;
  const mentions = purchase.nameMentions ?? [];
  return (
    <>
      {purchase.technique && (
        <p className="mt-1 text-[13px] text-muted-foreground">
          <span className="font-medium">Plan</span> · adapts {purchase.technique}
          {purchase.baseTechnique && ' (base attack)'}
        </p>
      )}
      {purchase.adaptation && (
        <p className="mt-1 text-[13px] text-muted-foreground">{purchase.adaptation}</p>
      )}
      {mentions.length > 0 && (
        <p className="mt-1 text-[13px] text-warning">
          Review: the name suggests {listText(mentions)}, which the plan does not map to this
          purchase.
        </p>
      )}
    </>
  );
}

/**
 * Splits the abilities no purchase lists. A typed unit shows reserved and
 * omitted techniques on their own, since no build grants them; a unit
 * without typed purchases keeps them with its other abilities.
 */
export function splitAbilities(
  abilities: Ability[],
  assigned: Set<string>,
  typed: boolean,
): { reserved: Ability[]; remaining: Ability[] } {
  const unused = (ability: Ability) =>
    ability.placement === 'reserved' || ability.placement === 'omitted';
  return {
    reserved: typed ? abilities.filter(unused) : [],
    remaining: abilities.filter(
      (ability) => !assigned.has(ability.id) && !(typed && unused(ability)),
    ),
  };
}
