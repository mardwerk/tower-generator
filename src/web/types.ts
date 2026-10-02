export type Upgrade = {
  name: string;
  description: string;
  cost: number;
  appearance?: string;
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
  attack: 'pellet' | 'explosion' | 'spread' | 'sniper';
  blastRadius: number;
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
    sourceId?: string;
    paths: { name: string; upgrades: Upgrade[] }[];
  };
  characterSource?: CharacterSource;
  states: TowerState[];
  checks: Record<'candidate' | 'missingState' | 'missingUpgrade', CheckReport>;
  source: { capture: string; build: string; commit: string };
};

export type CharacterSource = {
  formatVersion: 1;
  kind: 'character-source';
  id: string;
  character: { name: string; work: string; scope: string };
  documents: {
    id: string;
    title: string;
    url: string | null;
    resolvedUrl?: string;
    retrievedAt: string | null;
    access: 'supplied' | 'retrieved';
    excerpt?: boolean;
    text: string;
  }[];
  notes: string;
};
export type SourceEntry = {
  id: string;
  character: CharacterSource['character'];
  documents: number;
};
