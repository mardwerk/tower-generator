# Anime-to-tower-defense generation approaches and a reference unit specification

## 1. Thread identity and overview

Proposed title: **Anime-to-tower-defense generation approaches and a reference unit specification**.

Provisional reference key: **UG-anime-td-systems**. No assigned thread ID or original chat title is visible. The assistant's delivery and HTML used the artifact label "UnitForge"; that is not an approved rename of Unit-Generator. The HTML's "RESEARCH / 001" is not treated as a conversation ID. This title and key identify this handover only.

All IDs below are scoped to `UG-anime-td-systems/`. For example, D03 means `UG-anime-td-systems/D03`. H identifies conversation/material sources, S historical external sources, N newly inspected verification sources, D decisions/proposals, E evidence, and Q open questions.

### Reviewed material and limits

| ID | Material and locator | Coverage |
|---|---|---|
| H01 | Original user request, "input an Anime Character name and get an TowerDefense Unit out" and "There are no restrictions on the underlying 'System'." | Original message. No message timestamp or link is exposed. |
| H02 | Assistant delivery, "Build a grounded unit-design workflow that evolves into a validated unit compiler" and "not a live arbitrary-character AI generation service." | Original assistant message that summarizes its work. Not owner acceptance or a substitute for missing execution logs. |
| H03 | `anime-unit-forge.html`, especially `#approaches`, `#architecture`, `#luffy`, `#validation`, `#delivery`, and `#sources`. | Text, embedded data and code inspected. Self-dated 23 September 2026. |
| H04 | `unit-forge-reference-package.zip`, all nine contained files. | Archive extracted and inspected; code and schema rerun. File locators below refer to its `unit-forge/` directory. |
| H05 | Five distinct screenshots: desktop hero/comparison/lab and mobile hero/Luffy. | Static layouts inspected. A repeated comparison attachment is a duplicate, not another test. |
| H06 | Current handover request, "Treat this as context supplied now, not historical evidence." | Current product context only. |
| H07 | `ug-handover-inspection/verification.json` and `node-verification.txt`. | New verification record, 24 September 2026. Includes file hashes and environment versions. |

The intermediate research conversation, original web-tool transcripts, raw browser-test logs, and any owner response approving the delivered architecture are not available. Other project-thread summaries are not treated as this conversation's original history. The artifact date does not establish the original messages' dates.

The starting question was whether a standalone system could turn a character name, exemplified by "Monkey D. Luffy", into a TD unit concept detailed enough to implement. The user requested unrestricted investigation of modern AI approaches and an HTML presentation, without importing other project knowledge.

The thread contributed an approach comparison, a distinction between concept/specification/importable content, an explicit reference game contract, and a worked unit with executable checks. It reached an **assistant recommendation, not an owner-approved architecture or a demonstrated arbitrary-character generator**.

### Current context, not historical decisions

H06 places Unit-Generator within Mardwerk, with Towerright as the private SaaS. The intended public component is standalone FOSS with a useful configurable default. Basic generation must not require private Towerright components; provider, hosting and inference costs remain separate choices. Towerright's special "tuning" has no established technical definition.

H06 also limits automatic output commitments: executable game code, visual assets and proven balance are not assumed. The approach remains unresolved; a functioning prototype comes before a possible paper. Towerright-first development followed by extraction is a proposed sequence, not acceptance of H03's compiler architecture. Better results than direct LLM prompting with a wiki entry remain a hypothesis.

## 2. Ideas and approaches explored

### Alternatives for generation and control

All six rows below are assistant-developed alternatives in H03 `#approaches` and embedded `report-data.approaches`. None received visible owner approval. Their numerical rankings are subjective planning scores, not experimental findings.

| Approach | Purpose, mechanism and expected benefit | Significant dependency or cost | Historical treatment |
|---|---|---|---|
| Curated templates and rules | Map a known character to an authored archetype and existing mechanics. Gives repeatable, inexpensive output for a limited roster. | Human mapping work grows with coverage; unknown names remain unsupported. | Proposed baseline and fallback, not a complete long-tail solution. |
| Single LLM with structured input/output | Supply a game contract, schema and approved examples; generate one UnitSpec and validate it. Fast way to test the output and character-to-mechanic translation. | May invent facts, contradict mechanics or produce arbitrary stats despite valid JSON. | Proposed cheap comparison baseline and prototype route. |
| Grounded structured workflow | Resolve character and era, retrieve permitted evidence, extract claims, design a constrained kit, validate the result. Intended to improve provenance and missing-evidence handling. | Source curation, scope filtering and validation work; factual improvement does not establish balance. | Assistant's preferred first implementation. |
| Unit compiler with simulation | Translate a proposed kit into typed mechanic data, validate dependencies, test numbers, then map to an engine. Intended to reduce implementation ambiguity. | Substantial mechanic-runtime, simulator and importer work; new mechanics need implementation. | Assistant's longer-term target, not an implemented compiler. |
| Multi-agent system | Divide difficult research, design, critique or coding into specialist tasks under budgets and stop conditions. | More calls, coordination and failure modes; agreement is not independent evidence. | Conditional expansion after an experiment shows value. |
| Domain-tuned model or preference learning | Learn from accepted units, rejected alternatives and designer edits to improve recurring design/style choices. | Requires suitable data, stable rules and held-out tests; can reinforce errors or sameness. | Later optimization after diagnosing a repeated failure. No training occurred. |

### Separate responsibilities, not competing generation algorithms

The evidence plan used identity, aliases, variant and spoiler scope before design. It separated supported character facts, gameplay adaptations, project assumptions and unresolved claims. Curated dossiers and ordinary lookup were the starting option; embeddings, reranking, pgvector and a knowledge graph were conditional scale/relationship tools, not v1 necessities. Sources included official material, metadata APIs and community references with different evidential roles. [H03 `#architecture`]

Cloud, local and hybrid deployment were independent choices. The report discussed hosted APIs, Ollama/vLLM, private local data, permitted cloud jobs, credentials, licensing, hardware and running cost. No model, provider or deployment mode was selected. Standalone was not established to mean fully offline. [H03 `#approaches`, `#tools`]

The proposed contract/validation layer joined `GameContract`, a mechanic registry and `UnitSpec`; shared data would drive numbers, tooltips and exports. Unsupported mechanics would return an explicit dependency request rather than a fabricated engine API. A conventional application state machine preceded optional LangGraph orchestration. These were proposals; the supplied fixture is not that general system. [H03 `#architecture`]

Assets were a later branch: procedural or rigged animation with AI-assisted images, or optional 3D generation, rigging and audio. Engine mappings covered Unity, Godot, Unreal, Roblox and a custom web engine, but none was built. Simulation, bounded parameter search and human playtesting addressed different questions from creative generation. Reinforcement learning was only a passing conditional suggestion. [H03 `#tools`, `#delivery`]

## 3. Decision record

Approval and verification are independent. A created file or passing test does not make its architecture an accepted product decision.

| ID and statement | Status and actor | Recorded rationale, alternatives and dependencies |
|---|---|---|
| D01. Research unrestricted system approaches for name-to-unit generation and deliver an HTML research site. | **Accepted**, as task scope explicitly requested by the owner in H01. | The owner wanted a standalone project and a directly useful unit concept. No implementation architecture, engine or detailed output contract was selected. |
| D02. Use saved game context; without the actual game rules, label the example as a reference-game specification. | **Proposed**, assistant; instantiated in the demonstration. H03 `#verdict`: "Name-only does not mean context-free." | Avoid hiding assumptions about combat, economy and supported mechanics. The reference profile is an alternative to pretending to know the user's game, not an owner-approved ruleset. |
| D03. Start with a grounded workflow and evolve toward a validated compiler/simulator. | **Proposed**, assistant. H02 and H03 `#verdict`: "Build a grounded unit-design workflow first." | Claimed benefits were better grounding and explicit execution semantics. Competing options are the six approaches above. Depends on an adequate contract, evidence and implemented mechanics; no comparative evidence establishes superiority. |
| D04. For engine-oriented output, generate at authoring time, use trusted runtime code, and flag unsupported mechanics explicitly. | **Proposed**, assistant. H03 `#verdict`, `#architecture`, `NEEDS_NEW_PRIMITIVE`. | Avoid unchecked scripts and live model decisions about combat. Requires tested runtime components. It does not establish that concept-only output must be restricted to already implemented mechanics. |
| D05. Delay complex agents and model training; add them only for a demonstrated problem or benefit. | **Proposed**, a deferral suggested by the assistant, not an owner decision. H03 `report-data.approaches`, `#validation`. | Agents add cost/coordination; training needs consistent examples. The stated reconsideration conditions were a beneficial ablation for agents and a repeated failure not solved by prompting, retrieval or constraints for training. |
| D06. Exclude a rubber-defense passive and later transformations from the worked Luffy kit. | **Rejected**, for this example only, by the assistant. H03 `#unit-overview`; H04 `luffy.unit.json`, `provenance.designDecisions`. | Defense had no function because reference towers cannot take damage. Later transformations were outside the example's scope. The report explicitly allowed a different variant for damageable towers. These are not product-wide bans. |

No later owner approval, supersession or reversal is visible. D02 to D05 remain proposals even where the report uses decisive language such as "recommended target". Reconsideration conditions beyond those above were not recorded.

The main unresolved scope tension is H01's "EVERYTHING"/direct implementation aspiration versus H06's narrower default output boundary. H02 acknowledged missing integration, assets and balance, but that acknowledgment was not owner approval of a revised completion contract. Q01 preserves the issue.

## 4. Evidence, research and artifacts

### Created and inspected

| ID | Artifact/result | What it establishes and what it does not |
|---|---|---|
| E01 | H03 interactive HTML; 34 source entries, six approach patterns, eight proposed stages, comparison and cost controls. | A research presentation and local demonstrator exist. Architecture scores, cost assumptions and phases are assistant analysis, not measured provider performance. |
| E02 | H04 `luffy.unit.json`, `game.contract.json`. | One worked unit and an explicit reference ruleset exist. They are not a generation dataset, real-game contract or evidence of arbitrary-character generation. |
| E03 | H04 `unit.schema.json`, `reference-combat.js`, `reference-tests.js`; H07 rerun. | Draft 2020-12 schema validation passed and **37/37 original checks passed again** under Node v22.16.0. These are fixture regression checks, not a benchmark of generated units. |
| E04 | H04 `README.md`, `VALIDATION.md`; H05 screenshots. | Original validation reports Node/Chromium checks, browser controls, keyboard behavior and no document overflow at 320, 390, 768, 1024 and 1440 pixels. Static screenshots were inspected now; browser interactions were not rerun. |
| E05 | H04 `sources.json`. | Preserves the original labels, URLs, claimed access scopes and 23 September 2026 check date. Current primary-source verification is recorded separately in section 7. |
| E06 | H03 `#validation`, proposed comparison plan. | A plan exists, but no benchmark dataset or results were delivered. Its numbers are proposals, not achieved quality. |

The archive has nine files. Its HTML matches the standalone attachment; all seven embedded export payloads match their corresponding packaged files. H07 records hashes. The ZIP contains **no LICENSE file**, live generator, actual provider prompt set, general mechanic compiler/registry, engine adapter, production assets or training dataset. Research prose describes these possibilities, but it does not implement them.

### Reference unit details worth preserving

E02 uses `onepiece.luffy.eastblue`, unit version `0.1.0`, and `td-reference-1.0.0`. It has six linear tiers, 20 combat ticks per second, visible-ground targeting, invulnerable stationary towers and one copy per owner. Placement costs 450 reference credits; total maximum investment is 6,900 and maximum sale refund 4,830. These are invented design values, not balance findings. [H04 JSON identity/economy/tiers]

The basic attack has a four-tick windup and can miss without retargeting. Every fourth committed basic becomes a 120-degree sweep after tier 2; it replaces the basic rather than adding a free hit. The tier-3 active has six pulses, fixed 60-degree aim, a 24-tick action lock and 600-tick cooldown. The specification delays first use for 30 seconds and preserves cooldown through sale/rebuy; the supplied harness does not implement the full sale/upgrade command system. [H03 `#unit-mechanics`]

The lab uses immortal stationary targets and normally starts the active ready at time zero. That differs from first unlock in the contract and is documented and tested. It does not establish pathing, overkill, waves, economy, native integration or balance. The schema and runtime are specific to this six-tier, three-ability example, not a demonstrated general unit language. [H04 README, core header, schema description]

Character provenance contains only two limited claims tied to S27 and S28. "East Blue-inspired" is not a completed scene-level canon audit. The exact kit, rarity, progression and numbers are gameplay adaptations. Production assets are explicitly `not_created`; publication is disabled. The proposed 2D art alone would contain 432 frames, approximately 60.8 MiB of raw pixels, before atlas overhead. An art specification is not delivered art. [H03 `#unit-assets`; H04 provenance/readiness/presentation]

### Proposed experiments and unsupported conclusions

E06 proposed 40 cases: 20 varied characters, 10 alias/ambiguity/variant cases, five obscure/noncombat cases and five nonexistent/adversarial inputs. It proposed three trials per case, comparisons of one-shot/grounded/compiler systems, and one-change-at-a-time critic or agent experiments. Proposed targets included 100% mandatory export gates, zero accepted wrong-identity/forbidden-era cases, at least 95% audited factual-support precision and median human recognizability of at least 4/5. **None was measured.**

The claims that the compiler is strongest, that grounding improves this product's results, or that the system will outperform wiki-plus-prompt are unproven here. E03 establishes only the delivered fixture's tested behavior. External documentation establishes described interfaces or approaches, not Unit-Generator's quality, novelty or economics.

## 5. Recommended direction and scope

### Historical recommendation

The assistant recommended D03, supported by D02 and D04, with D05 delaying additional complexity. The longer-term eight-stage proposal ran from identity/evidence through design, compilation, simulation, assets, engine export and publication review. **No direction was settled by the owner.** [H02; H03 `#architecture`, `#delivery`]

### Current recommendation, awaiting owner approval

D07 is a **new proposal for this handover**, not a recovered decision: build the smallest text/specification pilot that tests whether saved game context, traceable evidence and validation reduce the owner's work. Add a component only when it addresses an observed failure or a chosen output promise.

The smallest supported workflow is a name plus a saved public game profile, lookup in a small reviewed character set, one model-backed generation step, schema and consistency checks, bounded correction, and editable HTML/JSON output with facts, adaptations and unresolved items separated. An unknown identity or insufficient evidence should produce an explicit incomplete result rather than fabricated certainty. Requested mechanics outside current game support should remain labeled proposals, not be presented as executable features. This uses the developed workflow idea without committing to the full compiler target.

Build or investigate now: the pilot output contract, the configurable default profile, a small reviewed evidence set, one usable provider route, generation and validation code, and a record of designer edits. Retain a direct-prompt baseline with comparable inputs to test usefulness rather than assume it. The complete E06 research study is not necessary to make v1 function.

Defer automatic asset production, multiple engine adapters, full balance simulation, autonomous agent teams, training and reinforcement learning. Add one engine test only if Q01 makes engine import part of the pilot's success criterion.

A Towerright-first pilot can follow this sequence without private runtime dependencies: keep generation inputs, public defaults and basic execution independently runnable, while Towerright supplies optional private configuration or data. This is a proposed boundary, not an established implementation. Do not interpret "tuned" as fine-tuned weights by default. Public configuration, dependency documentation and an explicit FOSS license remain work to complete, not properties proven by E01 to E04.

## 6. Open questions, conflicts and risks

| ID/type | Issue and why it matters | Pilot impact and resolution |
|---|---|---|
| Q01. Product | Does success mean a useful concept, a detailed implementation specification, or import into an engine? H01's "EVERYTHING" and H06's default exclusions differ in scope. | **Blocks a meaningful completion criterion.** Owner approval of one sample output and explicit exclusions resolves it; no silent supersession. |
| Q02. Public/private boundary | What must the public default include, and what may remain private? "Tuned" could mean configuration, examples, retrieval data, weights or review policy. | **Blocks standalone acceptance**, not an isolated internal experiment. Record the boundary and demonstrate basic generation without Towerright credentials or private modules. |
| Q03. Technical feasibility | Can a supported, adaptable specification cover the chosen units without hardcoding Luffy's slots? What happens when a proposed mechanic lacks runtime support? | **Blocks promises of generality/executability**, not a bounded concept pilot. Test materially different kits and explicit unsupported-mechanic behavior. |
| Q04. Empirical research | Does the workflow improve usable output enough to justify extra work and cost over direct prompting? Does a critic or agent help? | **Does not block building a pilot; blocks superiority claims.** Compare source correctness, editing time, acceptance and cost using controlled inputs. E06 is an unapproved larger plan. |
| Q05. Product/data | Which character variants, spoiler limits and sources are allowed? Two limited Luffy claims do not establish broad canon coverage. | **Blocks broad name-only fidelity claims.** A reviewed initial roster and explicit identity/evidence failure states permit a smaller pilot. Source reuse terms need separate confirmation. |
| Q06. Technical/deployment | What provider or local model is the supported default, with what schema limits, hardware and cost? | **Blocks a runnable AI service once implementation starts.** Test the actual prompt/schema with one chosen route. FOSS and standalone do not decide offline operation or free inference. |
| Q07. Product/technical | Are engine import and balance required now? The current fixture omits full match behavior and has no roster comparisons. | **Conditional blocker only.** Either exclude these promises under Q01 or select one engine/runtime and representative gameplay checks. Passing E03 is insufficient. |
| Q08. Public/private and distribution | Can example characters, stored source material, models and assets be redistributed with the standalone project? The archive also lacks a software license. | **Blocks a cleared public release**, not preparation of a bounded experiment using permitted material. Resolve software licensing and source/character/asset permissions separately. S11, S25 and S26 do not themselves clear this package. |

## 7. Source register

### Register conventions

H01 to H07 above are the internal source register. S01 to S34 reproduce all original external entries from H03 `#sources` and H04 `sources.json`; S35 is the additional normative URI embedded in the schema. No standalone scholarly paper or external dataset was individually cited in the original 34-entry list.

The original register reports checks on 23 September 2026. It reports full-page review for most entries, but only search-index evidence for S10, S11, S27 and S28; documentation retrieval for S16, S18, S19 and S32; overview-only review for S26; and relevant-passage review for S33. Those are historical access claims, not reconstructed original browsing evidence.

The following verification describes **24 September 2026** access. "Metadata" means bibliographic identity or page title/introductory material only. "Sections" means relevant text was inspected, not a complete documentation set. Except where specified, the author is the named organization/maintainer, the source is living official documentation with no verified publication year, and no DOI/arXiv ID was recorded. Its ordinary URL is the identifier. Mutable pages were not pinned or archived at the original review date. Verification is not product approval.

S01. OpenAI, "Structured model outputs". Sections verified. Schema-constrained output and remaining value/factual errors. H03 #verdict/#tools; D03/D04. No task-performance test.

https://developers.openai.com/api/docs/guides/structured-outputs

S02. Anthropic, "Structured outputs". Sections verified. JSON outputs, strict tool inputs and schema restrictions. H03 #tools; D03/D04. Provider compatibility still needs implementation testing.

https://platform.claude.com/docs/en/build-with-claude/structured-outputs

S03. Google, "Structured outputs"; originally labeled "Gemini structured outputs". Sections verified. Valid JSON is not semantic correctness. H03 #tools; D03/D04. Current footer says 23 September 2026; the original entry recorded 17 September. No original page snapshot resolves that difference.

https://ai.google.dev/gemini-api/docs/structured-output

S04. Erik S. and Barry Zhang, "Building effective agents". Anthropic Engineering, 19 December 2024; engineering article. Byline/date and relevant sections verified. Workflow/agent distinction and complexity trade-offs. H03 #approaches; D05. Practitioner guidance, not a Unit-Generator experiment; current page warns its tooling discussion has changed.

https://www.anthropic.com/engineering/building-effective-agents

S05. Mikaela Grace, Jeremy Hadfield, Rodrigo Olivares and Jiri De Jonghe, "Demystifying evals for AI agents". Anthropic Engineering, 9 January 2026; engineering article. Byline/date and relevant sections verified. Graders, repeated trials and outcome inspection. H03 #validation; E06/Q04. Does not validate the report's TD-specific thresholds.

https://www.anthropic.com/engineering/demystifying-evals-for-ai-agents

S06. LangChain, "LangGraph overview". Sections verified. Stateful orchestration, persistence and human review. H03 #architecture; D05. Optional infrastructure, not supplied mechanic semantics.

https://docs.langchain.com/oss/python/langgraph/overview

S07. Ollama, "Structured Outputs". Sections verified, including the local/Cloud support distinction. H03 #tools; Q06. No local model or hardware benchmark was run.

https://docs.ollama.com/capabilities/structured-outputs

S08. vLLM contributors, "Structured Outputs". Page identity verified; metadata inspected. H03 #tools; Q06. Serving option only; no configuration or model tested.

https://docs.vllm.ai/en/latest/features/structured_outputs/

S09. pgvector maintainers, "pgvector" repository; originally "pgvector repository". Software repository. README introduction inspected. H03 #architecture/#delivery; optional vector retrieval. No commit pinned or corpus benchmark conducted.

https://github.com/pgvector/pgvector

S10. AniList, "Introduction"; originally "API introduction". Official search-index metadata verified; direct retrieval failed. H03 #architecture; Q05. Identity/metadata candidate, not an inspected ability database.

https://docs.anilist.co/guide/introduction

S11. AniList, "Terms of Use"; originally "API Terms of Use". API terms. Official indexed passages inspected; full page unavailable. H03 #architecture; Q05/Q08. Index mentions commercial licensing conditions and prohibits mass collection. Do not reinterpret the original limited note as a blanket noncommercial-only rule; permissions remain unconfirmed.

https://docs.anilist.co/guide/terms-of-use

S12. OWASP Gen AI Security Project, "LLM01:2025 Prompt Injection". Versioned 2025 security guidance. Relevant sections verified. External-content handling and limited privileges. H03 #requirements/#delivery; D04. Guidance, not a security audit of the package.

https://genai.owasp.org/llmrisk/llm01-prompt-injection/

S13. ComfyUI, "ComfyUI Official Documentation"; originally "Official documentation". Overview inspected. H03 #tools; optional image workflows. No workflow, model license or custom-node code was tested.

https://docs.comfy.org/

S14. OpenAI, "Image generation"; originally "Image generation tool". Page identity verified; metadata inspected. H03 #tools; optional asset branch. No API trial or production asset evidence.

https://developers.openai.com/api/docs/guides/tools-image-generation

S15. Google, "Gemini image generation" [original register label]. Official URL retrieved; metadata only. H03 #tools; optional image branch. Exact page heading not confirmed in the returned extract; no images generated.

https://ai.google.dev/gemini-api/docs/image-generation

S16. Meshy, "Text to 3D API". Page identity verified; metadata inspected. H03 #tools; optional 3D branch. No model generation, mesh inspection or integration test.

https://docs.meshy.ai/en/api/text-to-3d

S17. Meshy, "Rigging API". Relevant sections verified, including qualified humanoid support. H03 #tools/#luffy; E02. Does not establish that elastic-limb rigging works.

https://docs.meshy.ai/en/api/rigging

S18. Meshy, "Animation API". Introductory/request sections inspected. H03 #tools; optional animation branch. Documentation describes rigged-asset animation, not a delivered Luffy animation.

https://docs.meshy.ai/en/api/animation

S19. ElevenLabs, "Sound effects"; originally "Sound effects capabilities". Page identity verified; metadata inspected. H03 #tools; optional audio branch. No audio production or redistribution-rights review.

https://elevenlabs.io/docs/overview/capabilities/sound-effects

S20. Unity, "ScriptableObject"; originally "ScriptableObject manual". Unity 6.0 manual URL. Relevant data-asset sections inspected. H03 #delivery; Q07. Proposed mapping only; no Unity importer or compile test.

https://docs.unity3d.com/6000.0/Documentation/Manual/class-ScriptableObject.html

S21. Godot contributors, "Resources". Page identity verified; metadata inspected. H03 #delivery; Q07. The stable documentation URL is mutable; no Godot adapter tested.

https://docs.godotengine.org/en/stable/tutorials/scripting/resources.html

S22. Epic Games, "Data Assets in Unreal Engine". Page identity verified; metadata inspected. H03 #delivery; Q07. Supports a proposed data mapping, not an implemented Unreal integration.

https://dev.epicgames.com/documentation/en-us/unreal-engine/data-assets-in-unreal-engine

S23. Roblox, "ModuleScript"; originally "ModuleScript reference". Page identity verified; metadata inspected. H03 #delivery; Q07. No Luau export or server-combat integration tested.

https://create.roblox.com/docs/reference/engine/classes/ModuleScript

S24. Optuna contributors, "Optuna: A hyperparameter optimization framework"; originally "Official documentation". Documentation identity verified; metadata inspected. H03 #tools/#validation; optional bounded optimization. No optimization experiment or separately cited Optuna paper.

https://optuna.readthedocs.io/en/stable/

S25. German Federal Ministry of Justice / Federal Office of Justice, "Copyright Act, section 23" [original register label]. Statutory text/English translation. Original URL retrieval failed. H03 #requirements/#delivery; Q08. The German provision was checked through new N01. Neither establishes clearance for the character or package; no exact historical translation revision verified.

https://www.gesetze-im-internet.de/englisch_urhg/englisch_urhg.html

S26. U.S. Copyright Office, "Copyright and Artificial Intelligence". Official study/report overview; living page. Overview inspected; linked reports not analyzed. H03 #requirements/#delivery; Q08. Authorship and AI-policy source, not a character-use license.

https://www.copyright.gov/ai/

S27. Greg / ONE PIECE.com, "第35回：May the Fourth be with you!"; originally "Greg column: May the Fourth be with you!". Official-site column, 20 May 2015. Title/date and relevant indexed excerpt verified; full page returned 403. H03 #luffy; E02 claim C01. Supports limited rubber/unarmed-fighter inspiration, not an early-series canon audit; index also includes later-form context.

https://one-piece.com/greg/o20150520_0399/index.html

S28. "Luffy's First Gum Gum Pistol | One Piece". YouTube clip; original register attributes Crunchyroll and 31 December 2023. Indexed title/URL recovered; video not inspected and channel/date not independently reconfirmed. H03 #luffy; E02 claim C02. Preserve those attribution/date details as originally reported, not newly verified.

https://www.youtube.com/watch?v=yB9shgf-PJ0

S29. OpenAI, "Pricing"; originally "API pricing". Living pricing documentation. Page identity verified; metadata inspected. H03 #delivery; E01 cost calculator. No quote or measured cost adopted.

https://developers.openai.com/api/docs/pricing

S30. Anthropic, "Pricing"; originally "Claude API pricing". Living pricing documentation. Page identity verified; metadata inspected. H03 #delivery; E01 cost calculator. Original calculator rates remain illustrative.

https://platform.claude.com/docs/en/about-claude/pricing

S31. Google, "Gemini Developer API pricing"; originally "Gemini API pricing". Living pricing documentation. Page identity verified; metadata inspected. H03 #delivery; E01 cost calculator. No task cost or provider comparison measured.

https://ai.google.dev/gemini-api/docs/pricing

S32. Meshy, "Pricing"; originally "API pricing". Living pricing documentation. Page identity verified; metadata inspected. H03 #delivery; optional asset costs. No service quote or asset-production cost validated.

https://docs.meshy.ai/en/api/pricing

S33. Prithvi Rajasekaran, "Harness design for long-running application development". Anthropic Engineering, 24 March 2026; engineering article. Byline/date and relevant sections verified. H03 #approaches; D05. It discusses simplification but also useful planner/generator/evaluator roles. It is not evidence that multi-agent systems are generally inferior, nor a TD-generation study.

https://www.anthropic.com/engineering/harness-design-long-running-apps

S34. Tencent Hunyuan3D Team, "Hunyuan3D 2.1" repository, 2025. Software/research repository; original label "Hunyuan3D 2.1 repository". README sections and citation metadata inspected. H03 #tools; optional local 3D branch. No weights, implementation or license audit; no latest/best claim. Associated paper metadata was newly checked as N03, not silently treated as historically read.

https://github.com/Tencent-Hunyuan/Hunyuan3D-2.1

S35. JSON Schema Draft 2020-12 meta-schema. Historical normative dependency in H04 `unit.schema.json`, `$schema`; outside the numbered 34-entry bibliography. The direct URI was not retrieved successfully, but N02 confirms its identity. Local schema and instance validation were rerun in E03.

https://json-schema.org/draft/2020-12/schema

### New sources found during verification

These are new verification references, not additions to the historical research findings.

N01. German statutory text, "Gesetz über Urheberrecht und verwandte Schutzrechte (Urheberrechtsgesetz), § 23 Bearbeitungen und Umgestaltungen". Official primary provision; full displayed section inspected on 24 September 2026. Confirms the provision associated with S25, including adaptation and sufficient-distance language. Historical revision not reconstructed; no project-specific legal conclusion.

https://www.gesetze-im-internet.de/urhg/__23.html

N02. Austin Wright, Henry Andrews, Ben Hutton and Greg Dennis, "Draft 2020-12". JSON Schema specification landing page; published 16 June 2022. Metadata and overview inspected, not the complete linked specifications. Identifies `draft-bhutton-json-schema-01`, `draft-bhutton-json-schema-validation-01` and the S35 meta-schema URI.

https://json-schema.org/draft/2020-12

N03. Team Hunyuan3D and 52 other credited authors, "Hunyuan3D 2.1: From Images to High-Fidelity 3D Assets with Production-Ready PBR Material". arXiv preprint, submitted 18 June 2025; arXiv:2506.15442, DOI 10.48550/arXiv.2506.15442. Metadata and abstract inspected; full text not inspected. Newly followed from S34 for bibliographic verification, not evidence for Unit-Generator quality. Full author names are retained in the accompanying verification register.

https://arxiv.org/abs/2506.15442

No broad new literature search was added. Uninspected papers, datasets and tools merely linked from these pages are not treated as sources used by the historical conversation.

## 8. Consolidation handoff

**Must not be lost:** D02's explicit assumptions; E02's distinction between character facts and adaptations; E03's reproducible but narrow verification; D04's unsupported-mechanic boundary; E06's unexecuted comparison plan; and the access limitations in S10, S11, S27 and S28. Preserve the single-LLM and template alternatives rather than recording D03 as a settled winner.

**Approval or reconciliation required:** Q01's output promise, Q02's public/private boundary, D03's architecture proposal and D07's smaller pilot scope. D05's deferrals are not permanent bans. E03's passing tests approve none of these decisions.

**Next concrete action:** resolve Q01 and Q02 in a short pilot contract using E02 as an example, not a mandatory schema. Then implement one real name-to-specification run under D07, including an explicit unknown-name result and a direct-prompt comparison. Preserve the exact output, validation errors, cost and owner edits.

**Later paper:** E06 offers comparison and ablation ideas, not a completed study. A future paper would need a suitable corpus, held-out cases, clear grading and measured results under Q04. This thread establishes neither scientific novelty nor an advantage over direct prompting; neither claim is required to finish a functioning v1.
