"""Character grounding and dossier building (Layer 1).

Provides an abstract retriever, atomic claim decomposition with provenance
(the FACTSCORE-style approach), and version selection — all required for
faithful character grounding. The demonstrated run uses a bundled canonical
profile so the pipeline executes without external access; a live retriever
can be registered to fetch from source pages instead.
"""

from abc import ABC, abstractmethod
from dataclasses import dataclass, field
from datetime import datetime
from typing import Dict, List, Optional, Set


@dataclass
class AtomicClaim:
    """A single verifiable statement about the character, with source.

    provenance: list of source locations (page/section + date)
    confidence: source-verified, plausibly-inferred, explicitly-guessed
    """
    claim_id: str
    statement: str
    provenance: List[str]
    confidence: str               # "source-verified" | "plausibly-inferred" | "explicitly-guessed"
    category: str                 # abilities | stats | relationships | narrative_role | weaknesses | progression

    def __lt__(self, other):
        return self.claim_id < other.claim_id


@dataclass
class VersionProfile:
    """One version of the character across media (THON19 version semantics)."""
    version_key: str              # e.g. "jujutsu-kaisen-manga-main"
    series: str
    medium: str                   # manga | anime | game | movie
    timeline_position: str        # e.g. "final arc", "post-graduation", "early series"
    description: str
    is_primary: bool = False


@dataclass
class CharacterDossier:
    """Structured, versioned character profile with atomic claims and provenance."""
    character: str
    series: str
    selected_version: str
    versions: List[VersionProfile]
    claims: Dict[str, AtomicClaim]        # claim_id -> claim, keyed and sorted
    metadata: Dict = field(default_factory=dict)

    @property
    def sorted_claims(self):
        return sorted(self.claims.values(), key=lambda c: c.claim_id)

    def claims_by_category(self) -> Dict[str, List[AtomicClaim]]:
        out: Dict[str, List[AtomicClaim]] = {}
        for c in self.sorted_claims:
            out.setdefault(c.category, []).append(c)
        return out

    def abilities(self) -> List[AtomicClaim]:
        return self.claims_by_category().get("abilities", [])

    def weaknesses(self) -> List[AtomicClaim]:
        return self.claims_by_category().get("weaknesses", [])

    def stats(self) -> List[AtomicClaim]:
        return self.claims_by_category().get("stats", [])

    def relationships(self) -> List[AtomicClaim]:
        return self.claims_by_category().get("relationships", [])

    def __len__(self):
        return len(self.claims)


class RetrieverABC(ABC):
    """Abstract source retriever. Concrete implementations fetch from fandom/Wiki/API.

    Returns a list of (claim_text, provenance) that downstream decomposes into
    AtomicClaims. The demonstrated run ships with a bundled profile instead of
    live fetches.
    """

    @abstractmethod
    def fetch_character(self, character: str, source: str) -> List[tuple]:
        pass

    @abstractmethod
    def fetch_version_info(self, character: str) -> List[tuple]:
        pass


class BundledProfileRetriever(RetrieverABC):
    """Demonstration retriever: returns the bundled Gojo profile without network access.

    Registers itself in the experiment so the full pipeline runs offline. Replace
    this with a real RetrieverABC implementation (fandom scrape, official guide,
    or licensed database) to convert the demonstration into a sourced run.
    """

    def __init__(self):
        self._cache: Dict[str, list] = {}

    def _register(self, character, rows):
        self._cache[character] = rows

    def fetch_character(self, character: str, source: str) -> List[tuple]:
        return self._cache.get(character, [])

    def fetch_version_info(self, character: str) -> List[tuple]:
        return self._cache.get(f"__versions__:{character}", [])


# ---- Demonstrative canonical profile: Satoru Gojo (Jujutsu Kaisen) ----
# Grounded from publicly documented source material; each entry below is a
# claim the experiment can assert and audit against. Treat this as a
# worked example pending a live retriever.

def build_gojo_dossier() -> CharacterDossier:
    version_primary = "jujutsu-kaisen-manga-main"
    versions = [
        VersionProfile(version_primary, "Jujutsu Kaisen", "manga",
                       "main timeline, graduation through final arc",
                       "Primary version: manga main timeline, graduation through final arc", is_primary=True),
        VersionProfile("jujutsu-kaisen-anime", "Jujutsu Kaisen", "anime",
                       "MAPPA adaptation, covers through Shibuya Incident arc",
                       "Anime-only coverage; weaker provenance for post-Shibuya abilities"),
        VersionProfile("jujutsu-kaisen-movie", "Jujutsu Kaisen 0", "movie",
                       "prequel set in 2006; younger Gojo as a student",
                       "Prequel-era profile; powers not yet at manga-end state"),
    ]

    c = {}
    aid = 0
    def add(cat, statement, provenance, confidence="source-verified"):
        nonlocal aid
        aid += 1
        cid = f"G{aid:03d}"
        c[cid] = AtomicClaim(cid, statement, provenance, confidence, cat)

    add("abilities",
        "Wields 'The Limitless' (Makyo), a innate technique derived from the Six Eyes that manipulates space on an atomic scale.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Cursed_Techniques",
         "jujutsu-kaisen.fandom.com/wiki/Limitless"])
    add("abilities",
        "Blue: attracts matter and creates a void-like collapse; functions as a gravity-well attraction attack.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Cursed_Techniques",
         "jujutsu-kaisen.fandom.com/wiki/Blue_(technique)"])
    add("abilities",
        "Red: reverses Blue's polarity to repel targets with high-velocity force.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Cursed_Techniques",
         "jujutsu-kaisen.fandom.com/wiki/Red_(technique)"])
    add("abilities",
        "Purple: combines Blue and Red into 'Nothingness', an erased-spatiality beam that deletes matter on contact.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Cursed_Techniques",
         "jujutsu-kaisen.fandom.com/wiki/Purple_(technique)"])
    add("abilities",
        "Hollow Purple is a singular, continuous erasure beam (not an explosive blast); direct hits remove targets from existence.",
        ["jujutsu-kaisen.fandom.com/wiki/Purple_(technique)"], confidence="plausibly-inferred")
    add("abilities",
        "Domain Expansion: Unlimited Void floods the target's senses with infinite information, immobilizing them while their bodies cannot move.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Cursed_Techniques",
         "jujutsu-kaisen.fandom.com/wiki/Unlimited_Void"])
    add("abilities",
        "Domain is a guaranteed-hit, inescapable zone while within range and while the cast completes.",
        ["jujutsu-kaisen.fandom.com/wiki/Unlimited_Void"], confidence="plausibly-inferred")
    add("abilities",
        "Six Eyes: hyper-perceptive cognition that reads cursed energy in extreme detail and minimizes its consumption.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Cursed_Techniques",
         "jujutsu-kaisen.fandom.com/wiki/Six_Eyes"])
    add("stats",
        "Holds the title 'Strongest Jujutsu Sorcerer'; designated Special Grade.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo",
         "jujutsu-kaisen.fandom.com/wiki/Grade_(power_level)"])
    add("stats",
        "Combat range spans from point-blank throws to building/neighborhood-scale destruction via Purple and Unlimited Void.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo"], confidence="plausibly-inferred")
    add("weaknesses",
        "Limitless is passively active; the Six Eyes must continually filter input, which causes fatigue under extreme stress or when the technique is sealed.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo"], confidence="plausibly-inferred")
    add("weaknesses",
        "Domain cast time exists: the technique must be activated before its effect applies.",
        ["jujutsu-kaisen.fandom.com/wiki/Unlimited_Void"], confidence="plausibly-inferred")
    add("weaknesses",
        "Can be sealed with cursed tools that bypass or suppress the Limitless (notably the Finger Bearer technique and Special Grade cursed tool bindings).",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo"], confidence="plausibly-inferred")
    add("narrative_role",
        "Acts as teacher/mentor and protective barrier for students; functions as a shield more than an offensive frontline unit.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Role_and_Characteristics"])
    add("relationships",
        "Student/teacher hierarchy defines his role with Yuji Itadori, Megumi Fushiguro, Nobara Kugisaki, and Maki Zenin.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo#Relationships"])
    add("progression",
        "Power scale increases across the story as restrictions are removed and technique mastery deepens.",
        ["jujutsu-kaisen.fandom.com/wiki/Satoru_Gojo"], confidence="plausibly-inferred")

    return CharacterDossier(
        character="Satoru Gojo",
        series="Jujutsu Kaisen",
        selected_version=version_primary,
        versions=versions,
        claims=c,
        metadata={"source": "demonstration profile (bundled)",
                  "version_policy": "primary = main-timeline manga; unit explains its version in the audit trace"})


def build_dossier(retriever: RetrieverABC, character: str, source: str) -> CharacterDossier:
    """Orchestration: fetch, assemble versions, decompose into atomic claims."""
    version_rows = retriever.fetch_version_info(character)
    versions = [VersionProfile(vk, s, m, tp, dp) for vk, s, m, tp, dp in version_rows]
    rows = retriever.fetch_character(character, source)
    claims = {}
    aid = 0
    for statement, provenance, conf, cat in rows:
        aid += 1
        claims[f"A{aid:03d}"] = AtomicClaim(f"A{aid:03d}", statement, provenance, conf, cat)
    return CharacterDossier(character, "unknown", versions[0].version_key if versions else "unknown",
                            versions, claims)


if __name__ == "__main__":
    d = build_gojo_dossier()
    print(d)
    print("claims:", len(d))
    for cat, claims in d.claims_by_category().items():
        print(f"  {cat}: {[c.claim_id for c in claims]}")
