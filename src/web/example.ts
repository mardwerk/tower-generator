import { readFileSync, readdirSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { join } from 'node:path';
import type { Example, TowerState } from './types.js';

export const defaults = fileURLToPath(new URL('../../default/', import.meta.url));

export function readJSON(path: string): unknown {
  return JSON.parse(readFileSync(path, 'utf8'));
}

type Model = Record<string, unknown>;

function models(value: unknown, kind: string): Model[] {
  if (Array.isArray(value)) return value.flatMap((child) => models(child, kind));
  if (!value || typeof value !== 'object') return [];
  const model = value as Model;
  const type = String(model.$type ?? '')
    .split(',')[0]
    ?.split('.')
    .at(-1);
  return [
    ...(type === kind ? [model] : []),
    ...Object.values(model).flatMap((child) => models(child, kind)),
  ];
}

export function readExample(): Example {
  const design = readJSON(join(defaults, 'usopp.json')) as Example['design'];
  const directory = join(defaults, 'game-data/Towers', design.id);
  const states = readdirSync(directory)
    .filter((name) => name.endsWith('.json'))
    .map((name): TowerState => {
      const tower = readJSON(join(directory, name)) as Model;
      const weapon = models(tower, 'WeaponModel')[0]!;
      const projectile = weapon.projectile as Model;
      const contact = models(projectile, 'CreateProjectileOnContactModel')[0];
      const hit = (contact?.projectile ?? projectile) as Model;
      const damage = models(hit, 'DamageModel')[0]!;
      const travel = models(projectile, 'TravelStraitModel')[0]!;
      const emission = weapon.emission as Model;
      return {
        name: String(tower.name),
        tiers: tower.tiers as number[],
        cost: Number(tower.cost),
        damage: Number(damage.damage),
        pierce: Number(hit.pierce),
        blastRadius: contact ? Number(hit.radius) : 0,
        attack: contact
          ? 'explosion'
          : Number(emission.count ?? 1) > 1
            ? 'spread'
            : Number(travel.speed) >= 400
              ? 'sniper'
              : 'pellet',
        range: Number(tower.range),
        interval: Number(weapon.rate),
        projectiles: Number(emission.count ?? 1),
        camo: models(tower, 'FilterInvisibleModel').every((filter) => !filter.isActive),
      };
    });
  return {
    design,
    states,
    checks: readJSON(join(defaults, 'checks.json')) as Example['checks'],
    source: readJSON(join(defaults, 'source.json')) as Example['source'],
  };
}
