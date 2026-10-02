export type Classification = {
  delivery: string;
  functions: string[];
  status: 'suggested' | 'reviewed';
};
export type Ability = {
  id: string;
  kind: 'ability' | 'trait' | 'equipment';
  name: string;
  description: string;
  scope: string;
  limitations?: string;
  evidence: { source: string; passage: string }[];
  classification?: Classification;
};
export type WikiEntry = {
  key: string;
  path: string;
  metadata: {
    name: string;
    work: string;
    scope: string;
    aliases: string[];
    status: 'draft' | 'reviewed';
    reviewedAt?: string;
    abilities: Ability[];
    sources: string[];
  };
  body: string;
  markdown: string;
  revision: string;
  reviewStale: boolean;
  documents: {
    id: string;
    title: string;
    url: string | null;
    retrievedAt: string | null;
    access: 'supplied' | 'retrieved';
    excerpt?: boolean;
    passages: { id: string; text: string }[];
    retainedPassages?: { id: string; retrievedAt: string | null }[];
  }[];
};
export type WikiSummary = {
  aliases: string[];
  key: string;
  name: string;
  work: string;
  scope: string;
  status: 'draft' | 'reviewed';
  abilities: number;
  sources: number;
};
export type Categories = {
  version: number;
  delivery: Record<string, string>;
  functions: Record<string, string>;
};
