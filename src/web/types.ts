export type Upgrade = {
  name: string;
  description: string;
  cost: number;
};

export type TowerState = {
  name: string;
  tiers: number[];
  cost: number;
  damage: number;
  pierce: number;
  range: number;
  interval: number;
  projectiles: number;
  camo: boolean;
};

export type CheckReport = {
  filesChecked: number;
  score: { points: number; complete: boolean };
  errors: { code: string; message: string; file: string; pointer: string }[];
  checker: { version: string };
  profile: { id: string; revision: string };
};

export type Example = {
  design: {
    id: string;
    name: string;
    description: string;
    paths: { name: string; upgrades: Upgrade[] }[];
  };
  states: TowerState[];
  checks: Record<'candidate' | 'missingState' | 'missingUpgrade', CheckReport>;
  source: { capture: string; build: string; commit: string };
};
