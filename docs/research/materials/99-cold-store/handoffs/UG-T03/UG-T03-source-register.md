# UG-T03 source register

Prepared 2026-09-24. Internal consolidation record, not a bibliography approved for publication.

A source already appearing in an attachment or an identifiable historical browsing query remains an original-thread source even when its DOI or corrected title is recovered now. Only Part B is new literature. "Verified" concerns the stated bibliographic scope; it does not mean reproduced experiments. "Content inspected" records what was examined for this handover, not what an earlier assistant or attachment claimed to inspect. No code, dataset or model listed here was installed or benchmarked for Unit-Generator.

For the 38 attachment catalogue entries, original publication and access wording is preserved separately from this handover's verification. Current metadata can differ from a linked older version. An unverified publication venue is not upgraded merely because a current title is verified. Authors abbreviated with "et al." are not silently completed. Associated artifact links are research leads, not license decisions.

## A. Sources already present in this thread

### Original messages, files and direct research references

#### UG-T03-S01. Research sources for unit generation and game design

Author not stated in supplied file. Research check date stated as 2026-09-22. Supplied internal research catalogue. Not a publication.

Identifier: Supplied file GAME-DESIGN-RESEARCH.md.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Attachment, entire file; R1 to R38 and additional references.

Bibliographic record: Verified as supplied artifact; individual source claims need separate verification.

Content inspected: Full text. Read mounted file and supplied line-cited text.

Relevance: Original research basis; UG-T03-E01 and UG-T03-E09.

Limits: Curated review, not systematic review, implementation commitment, or reproduced experiments.

Associated artifacts: Original file retained in this conversation. Supplied snapshot SHA-256: 8b1611fa74f65009d3c71644220f74880a46d95457302888ce3bd02e4e7c27cc.

Record notes: Its assertions that papers were inspected are historical assertions, not inherited verification labels for this handover.

#### UG-T03-S02. Unit Generator context

Author not stated in supplied file. Undated. Supplied internal contributor contract. Not a publication.

Identifier: Supplied file CONTEXT.md.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Attachment, Shared terms; Unit terms; Generation and evaluation terms; closing implementation-scope paragraph.

Bibliographic record: Verified as supplied artifact, not as an audit of repository implementation.

Content inspected: Full text. Read mounted file and supplied line-cited text.

Relevance: Terminology, ownership and reported implementation limits; UG-T03-E01, UG-T03-D09, UG-T03-Q04.

Limits: Referenced source code and linked local documents were not supplied.

Associated artifacts: Foundation is explicitly private internal guidance, not included. Supplied snapshot SHA-256: fee3f4ebbcaf0f7c44013db8102ca98510676e8723a399e95ee323f7a946ee50.

Record notes: Historical responsibility wording must not become a Towerright dependency for standalone use.

#### UG-T03-S03. Original user research request

Project owner. No visible original-message timestamp. Original conversation message. Not a publication.

Identifier: Speaker plus quotation locator.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: User: "its rather a transformation then a generation" and "BEST CASE i would like to train an AI model of some form".

Bibliographic record: Verified in visible conversation.

Content inspected: Full text. Original message, not a project summary.

Relevance: Establishes source-to-unit adaptation question and training interest; UG-T03-D03.

Limits: Interest in training does not approve a checkpoint, training method, dataset size, or architecture.

#### UG-T03-S04. Owner clarification on unknown mechanics and library growth

Project owner. No visible original-message timestamp. Original conversation message. Not a publication.

Identifier: Speaker plus quotation locator.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: User: "we also dont know the 'entire' mechanic's"; "would be okay, as longs as the mechanic is well defined and actually implementable"; "taking the 10 Straw Hat units as the initial then continuing with Jujutsu Kaisen units".

Bibliographic record: Verified in visible conversation.

Content inspected: Full text. Original clarification.

Relevance: Open mechanic coverage and permitted downstream changes; UG-T03-D01 and UG-T03-D02.

Limits: Universal coverage is aspirational. Automatic promotion, precise curriculum, and engineering scope were not agreed.

#### UG-T03-S05. Initial assistant research answer

Assistant. No visible original-message timestamp. Original conversation answer. Not a publication.

Identifier: Opening: "post-training an existing generative model is a credible approach".

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Assistant answer to UG-T03-S03, sections Main finding through Bottom line.

Bibliographic record: Verified as an assistant statement, not independent evidence.

Content inspected: Full text. Visible substantive answer; underlying old tool outputs are mostly redacted.

Relevance: Constrained synthesis, training alternatives, evaluation and planning proposals; UG-T03-D03 to UG-T03-D07. UG-T03-E14 records the answer's non-exhaustive negative search finding.

Limits: No owner approval or Unit-Generator implementation follows from the recommendation.

#### UG-T03-S06. Revised assistant answer on an evolving mechanics library

Assistant. No visible original-message timestamp. Original conversation answer. Not a publication.

Identifier: Opening: "the mechanics library must be an evolving output".

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Assistant answer to UG-T03-S04, sections 1 to 7.

Bibliographic record: Verified as an assistant statement, not independent evidence.

Content inspected: Full text. Visible substantive answer.

Relevance: Intent before implementation, extension contracts, consolidation, external library and reuse evaluation; UG-T03-D05, UG-T03-D08 to UG-T03-D10.

Limits: No owner message approves its exact architecture. Broken citation for Library Learning Doesn't is recovered as UG-T03-S25.

#### UG-T03-S07. caveman SKILL.md

JuliusBrussee/caveman repository. Mutable main branch; date not pinned. Repository task/style guidance. Not research.

Identifier: skills/caveman/SKILL.md.

Direct URL:
https://raw.githubusercontent.com/JuliusBrussee/caveman/refs/heads/main/skills/caveman/SKILL.md

Origin and locator: User linked this with "LITE" in UG-T03-S03; old browsing trace opened it.

Bibliographic record: Verified current repository file.

Content inspected: Full text. Current 73-line raw file.

Relevance: Requested writing style, not support for the technical architecture.

Limits: Original commit and original retrieved contents are unavailable; current content may differ.

Record notes: Not an approved project Skill or generated unit prompt.

#### UG-T03-S08. Laya model family and root checkpoint model card

convaiinnovations; repository linked by model author. Mutable model card; no revision pinned. Model artifact and vendor documentation. Not a peer-reviewed benchmark.

Identifier: convaiinnovations/laya.

Direct URLs:
https://huggingface.co/convaiinnovations/laya
https://huggingface.co/convaiinnovations/laya/raw/main/README.md

Origin and locator: User UG-T03-S03; assistant UG-T03-S05, "Where Laya fits".

Bibliographic record: Verified current model-card record.

Content inspected: Relevant sections. Architecture, checkpoint table and author-reported benchmark table.

Relevance: Bounded option scorer considered against open-ended generation; UG-T03-E07 and UG-T03-D07.

Limits: Root 421M encoder/decision head and 512-token question budget must not be generalized to every family checkpoint. Base accuracy 0.362, majority 0.461 and specialized 0.766 concern the card's typed-decision benchmark, not TD design; not reproduced.

Associated artifacts: https://github.com/NandhaKishorM/laya

Record notes: Public model-card licensing is not a Unit-Generator license or compatibility decision.

#### UG-T03-S09. From LLM-Driven Trading Card Generation to Procedural Relatedness: A Pokémon Case Study

Johannes Pfau; Panagiotis Vrettis. 2026. Preprint. No peer-reviewed venue established here.

Identifier: arXiv:2604.27972v1.

Direct URLs:
https://arxiv.org/abs/2604.27972
https://arxiv.org/html/2604.27972v1

Origin and locator: UG-T03-S05, missing paper A.

Bibliographic record: Verified arXiv record.

Content inspected: Relevant sections. Methods 3.1 to 3.3 and study description/results.

Relevance: Closest description-to-representative-card example; UG-T03-E03.

Limits: Qwen3-14B with 993-card retrieval; 49 participants and 196 generated cards. Representation judgments are not executable semantics, gameplay balance, or evidence that a newly trained adaptation LLM was required.

Associated artifacts: 993-card data described by paper; not downloaded or reused.

#### UG-T03-S10. Open Role-Playing with Delta-Engines

Hongqiu Wu et al.; author-list discrepancy retained below. 2024 preprint; v5 2025. Preprint. No proceedings venue established here.

Identifier: arXiv:2408.05842v5.

Direct URLs:
https://arxiv.org/abs/2408.05842
https://arxiv.org/html/2408.05842v5

Origin and locator: UG-T03-S05, missing paper B; UG-T03-S06, executable witness discussion.

Bibliographic record: Verified title, version and record; author lists disagree between record and rendered text.

Content inspected: Relevant sections. Sections 3, 5 and 6, including data construction, training, execution checks and correctness assessment. Section 6.2 also inspected for accumulated-extension degradation; that additional finding was not claimed as a historical Unit-Generator result.

Relevance: Character description to incremental game code and specialist training; UG-T03-E04, UG-T03-D05 and UG-T03-D06.

Limits: 502 co-designed examples across 175 roles; CodeGemma-7B LoRA. Correctness uses GPT-4 on executable outputs. Execution does not prove balance or safety of general runtime modification.

Associated artifacts: https://github.com/gingasan/delta-engine/tree/main/free-pokemon ; https://huggingface.co/google/codegemma-7b-it

Record notes: Abstract record lists Hongqiu Wu, Zekai Xu, Tianyang Xu, Shize Wei, Yan Wang, Jiale Hong, Weiqi Wu, Hai Zhao. Rendered v5 lists Hongqiu Wu, Zekai Xu, Tianyang Xu, Yan Wang, Weiqi Wu, Jiale Hong, Linfeng Liu, Hai Zhao. Do not silently combine these lists.

#### UG-T03-S11. Game-O-Matic: Generating Videogames that Represent Ideas

Mike Treanor; Bryan Blackford; Michael Mateas; Ian Bogost. 2012. Published workshop paper. Procedural Content Generation workshop.

Identifier: DOI:10.1145/2538528.2538537.

Direct URL:
https://mtreanor.com/publications/Game-O-Matic-PCG2012.pdf

Origin and locator: UG-T03-S05, missing paper C.

Bibliographic record: Verified author text and publisher DOI record.

Content inspected: Relevant sections. Introduction, concept graphs, micro-rhetorics, recipe assembly and evaluation limitation; PDF pages 1 and 3 visually inspected.

Relevance: Relationship-to-mechanic interpretation; UG-T03-D05 and UG-T03-E10.

Limits: Human-authored mappings and bounded mechanics; formal representational user evaluation is future work in this paper.

#### UG-T03-S12. Chaos Cards: Creating Novel Digital Card Games through Grammatical Content Generation and Meta-Based Card Evaluation

Tiannan Chen; Stephen J. Guy. 2020. Published conference paper. AIIDE 2020.

Identifier: AAAI article 7430.

Direct URL:
https://cdn.aaai.org/ojs/7430/7430-52-10748-1-2-20200923.pdf

Origin and locator: UG-T03-S05, missing paper D and strength-predictor discussion.

Bibliographic record: Verified primary paper identity.

Content inspected: Relevant sections. Grammar construction, card strength and deck-context evaluation; first page and page 4 screenshots inspected.

Relevance: Executable grammar generation separate from learned strength evaluation; UG-T03-E10, UG-T03-D06.

Limits: Neural model predicts strength; it is not the mechanic generator. Neither strength nor grammar validity proves canon fidelity or preference.

#### UG-T03-S13. Latent Predictor Networks for Code Generation

Wang Ling; Phil Blunsom; Edward Grefenstette; Karl Moritz Hermann; Tomáš Kočiský; Fumin Wang; Andrew Senior. 2016. Published conference paper. ACL 2016.

Identifier: ACL Anthology P16-1057; arXiv:1603.06744.

Direct URLs:
https://aclanthology.org/P16-1057/
https://arxiv.org/html/1603.06744

Origin and locator: UG-T03-S05, missing paper E; earlier browsing trace.

Bibliographic record: Verified publisher bibliographic record.

Content inspected: Abstract only. Publisher abstract and metadata rechecked; original answer described fuller task details.

Relevance: Mechanic-description-to-code is distinct from choosing the character adaptation; UG-T03-E10.

Limits: Inputs already describe intended card effects. This is not wiki-to-game design supervision.

Associated artifacts: Card-description/code datasets described by source, not imported.

#### UG-T03-S14. Mystical Tutor: A Magic: The Gathering Design Assistant via Denoising Sequence-to-Sequence Learning

Adam Summerville; Michael Mateas. 2016. Published conference paper. AIIDE 12(1), pages 86 to 92 (2016).

Identifier: DOI:10.1609/aiide.v12i1.12851; AAAI article 12851.

Direct URL:
https://ojs.aaai.org/index.php/AIIDE/article/view/12851

Origin and locator: UG-T03-S05, missing paper F.

Bibliographic record: Verified publisher bibliographic record, including issue year and DOI.

Content inspected: Abstract only. Publisher abstract accessed; original answer explicitly could not inspect full text.

Relevance: Completion from partial specifications; UG-T03-D03 and UG-T03-E10.

Limits: Partial-card completion does not demonstrate faithful character adaptation or complete playable units.

#### UG-T03-S15. PICARD: Parsing Incrementally for Constrained Auto-Regressive Decoding from Language Models

Torsten Scholak; Nathan Schucher; Dzmitry Bahdanau. 2021. Published conference paper. EMNLP 2021.

Identifier: DOI:10.18653/v1/2021.emnlp-main.779.

Direct URL:
https://aclanthology.org/2021.emnlp-main.779/

Origin and locator: UG-T03-S05, enforce structure during generation.

Bibliographic record: Verified publisher record.

Content inspected: Abstract only. Publisher abstract and metadata.

Relevance: Parsing-constrained output as optional implementation technique; UG-T03-E08, UG-T03-D04.

Limits: Text-to-SQL evidence does not establish arbitrary game semantic validity or balance.

#### UG-T03-S16. XGrammar-2: Dynamic and Efficient Structured Generation Engine for Agentic LLMs

Linzhang Li; Yixin Dong; Guanjie Wang; Ziyi Xu; Alexander Jiang; Tianqi Chen. 2026. Preprint with publication identified in current record. Current arXiv record reports ACM CAIS 2026.

Identifier: arXiv:2601.04426; current v4; DOI:10.1145/3786335.3813124.

Direct URLs:
https://arxiv.org/abs/2601.04426
https://arxiv.org/html/2601.04426v1

Origin and locator: UG-T03-S05, structured generation; original trace opened v1.

Bibliographic record: Verified current arXiv record; original and current versions distinguished.

Content inspected: Abstract only. Current v4 record and abstract; old v1 content not reconstructed.

Relevance: Dynamic output structures; UG-T03-E08.

Limits: Generation-engine efficiency does not establish semantic correctness, taste, or Unit-Generator performance.

Associated artifacts: https://github.com/mlc-ai/xgrammar

Record notes: Earlier wording used "XGrammar 2". Indexed v2 and current record differ in title wording; preserve identifier and version rather than silently conflating versions.

#### UG-T03-S17. LoRA: Low-Rank Adaptation of Large Language Models

Edward J. Hu; Yelong Shen; Phillip Wallis; Zeyuan Allen-Zhu; Yuanzhi Li; Shean Wang; Lu Wang; Weizhu Chen. 2021 preprint. Research paper, preprint record consulted. Proceedings version not independently checked in this handover.

Identifier: arXiv:2106.09685v2.

Direct URL:
https://arxiv.org/abs/2106.09685

Origin and locator: UG-T03-S05, supervised fine-tuning; UG-T03-S06 references training precedent.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Method summary and metadata.

Relevance: Parameter-efficient training option, not a project commitment; UG-T03-D06.

Limits: Adapter technique does not supply training data or prove domain performance.

#### UG-T03-S18. QLoRA: Efficient Finetuning of Quantized LLMs

Tim Dettmers; Artidoro Pagnoni; Ari Holtzman; Luke Zettlemoyer. 2023. Research paper, preprint record consulted. Proceedings version not independently checked here.

Identifier: arXiv:2305.14314.

Direct URL:
https://arxiv.org/abs/2305.14314

Origin and locator: UG-T03-S05, supervised fine-tuning.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Method summary and metadata.

Relevance: Quantized-backbone adapter training; UG-T03-D06.

Limits: Published memory or quality findings are not a hardware budget or accuracy estimate for the proposed Unit task.

#### UG-T03-S19. Direct Preference Optimization: Your Language Model is Secretly a Reward Model

Rafael Rafailov; Archit Sharma; Eric Mitchell; Stefano Ermon; Christopher D. Manning; Chelsea Finn. 2023 preprint; record includes later revision. Research paper, preprint record consulted. Final proceedings record not independently checked here.

Identifier: arXiv:2305.18290.

Direct URL:
https://arxiv.org/abs/2305.18290

Origin and locator: UG-T03-S05, later preference training.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Method and task scope from abstract.

Relevance: Learning reviewed preferences over alternatives; UG-T03-D06.

Limits: No TD taste supervision or reliable preference labels supplied by the method itself.

#### UG-T03-S20. Distilling System 2 into System 1

Ping Yu; Jing Xu; Jason Weston; Ilia Kulikov. 2024. Preprint. No final venue established here.

Identifier: arXiv:2407.06023v3.

Direct URL:
https://arxiv.org/abs/2407.06023

Origin and locator: UG-T03-S05, faster System One generator.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Distillation scope and metadata.

Relevance: Train a smaller/cheaper process on filtered expensive results; UG-T03-D06.

Limits: Task-dependent evidence, not proof that reviewed Unit generation can be compressed without losses.

#### UG-T03-S21. DreamCoder: Growing generalizable, interpretable knowledge with wake-sleep Bayesian program learning

Kevin Ellis; Catherine Wong; Maxwell Nye; Mathias Sable-Meyer; Luc Cary; Lucas Morales; Luke Hewitt; Armando Solar-Lezama; Joshua B. Tenenbaum. 2020 linked preprint. Preprint version explicitly linked in thread. Later publication not silently substituted.

Identifier: arXiv:2006.08381v1.

Direct URLs:
https://arxiv.org/abs/2006.08381
https://arxiv.org/html/2006.08381

Origin and locator: UG-T03-S06, library learning references.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Record and abstract rechecked; original answer made bounded-transfer argument.

Relevance: Reusable abstractions from solved programs; UG-T03-E05, UG-T03-D08.

Limits: Learning compositions does not create missing Consumer operations. Later DreamCoder publication titles/venues require version-specific citation.

#### UG-T03-S22. LILO: Learning Interpretable Libraries by Compressing and Documenting Code

Gabriel Grand; Lionel Wong; Maddy Bowers; Theo X. Olausson; Muxin Liu; Joshua B. Tenenbaum; Jacob Andreas. 2023 preprint; 2024 publication. Published conference paper; camera-ready preprint consulted. ICLR 2024.

Identifier: arXiv:2310.19791v4.

Direct URLs:
https://arxiv.org/abs/2310.19791
https://arxiv.org/html/2310.19791v4

Origin and locator: UG-T03-S05 closing library lead; UG-T03-S06, reuse/refactoring.

Bibliographic record: Verified arXiv record and camera-ready status.

Content inspected: Relevant sections. Synthesis, Stitch compression, documentation, experimental domains and discussion.

Relevance: LLM synthesis, compression and documentation loop; UG-T03-E05, UG-T03-D08.

Limits: String editing, CLEVR and LOGO tasks, not arbitrary real-time game extensions or canon fidelity.

Associated artifacts: Stitch is separately registered as UG-T03-S23; no code executed.

#### UG-T03-S23. Top-Down Synthesis for Library Learning

Matthew Bowers; Theo X. Olausson; Lionel Wong; Gabriel Grand; Joshua B. Tenenbaum; Kevin Ellis; Armando Solar-Lezama. 2023; preprint 2022. Published conference journal article. Proceedings of the ACM on Programming Languages 7, POPL, article 41.

Identifier: DOI:10.1145/3571234; arXiv:2211.16605v2.

Direct URLs:
https://arxiv.org/abs/2211.16605
https://doi.org/10.1145/3571234

Origin and locator: UG-T03-S06, promotion and consolidation.

Bibliographic record: Verified primary bibliographic records.

Content inspected: Abstract only. Author/institutional and proceedings records; algorithm summary.

Relevance: Stitch extracts reusable program structure; UG-T03-E05, UG-T03-D08.

Limits: Syntactic/program abstraction is not evidence of equivalent character meaning.

Associated artifacts: Stitch tool. No implementation imported or tested.

Record notes: The linked paper record uses Matthew Bowers; related LILO record uses Maddy Bowers. Retain each source's recorded name.

#### UG-T03-S24. Voyager: An Open-Ended Embodied Agent with Large Language Models

Guanzhi Wang et al.. 2023 linked preprint. Preprint record consulted. No final venue independently established here.

Identifier: arXiv:2305.16291.

Direct URLs:
https://arxiv.org/abs/2305.16291
https://arxiv.org/html/2305.16291

Origin and locator: UG-T03-S06, do not retrain after every unit.

Bibliographic record: Verified arXiv record.

Content inspected: Relevant sections. Skill library and iterative interaction method.

Relevance: External executable skill library; UG-T03-E05, UG-T03-D10.

Limits: Works within Minecraft's existing capabilities. Skill accumulation is not invention of new physics or evidence of game-design quality.

#### UG-T03-S25. Library Learning Doesn't: The Curious Case of the Single-Use "Library"

Ian Berlot-Attwell; Frank Rudzicz; Xujie Si. 2024. Workshop paper and preprint. 4th MATH-AI Workshop at NeurIPS 2024.

Identifier: arXiv:2410.20274v1.

Direct URLs:
https://arxiv.org/abs/2410.20274
https://arxiv.org/html/2410.20274v1

Origin and locator: UG-T03-S06, "Test whether the growing library actually improves adaptation"; citation marker was corrupted.

Bibliographic record: Verified, recovered by title; not a newly introduced source.

Content inspected: Relevant sections. Reuse analysis and ablations for LEGO-Prover and TroVE.

Relevance: Counterevidence to assuming library size causes gains; UG-T03-E06 and UG-T03-Q11.

Limits: Mathematical reasoning systems, not a universal refutation of library learning.

Associated artifacts: https://github.com/ikb-a/llm_lib_learning_fails

Record notes: Original broken marker retained as a provenance defect; recovery does not make the original claim independently verified at the time.

#### UG-T03-S26. Qwen3-14B

Qwen. Mutable model card; release revision not pinned. Model artifact. Not a project benchmark.

Identifier: Qwen/Qwen3-14B.

Direct URL:
https://huggingface.co/Qwen/Qwen3-14B

Origin and locator: UG-T03-S05 named comparison candidate and UG-T03-S09 model choice.

Bibliographic record: Verified current publisher model card.

Content inspected: Relevant sections. Identity, architecture/use guidance; not benchmark reproduction.

Relevance: Possible generative checkpoint; UG-T03-D06.

Limits: No training, inference, hardware measurement or Unit-Generator comparison was performed.

Record notes: No license compatibility audit or checkpoint choice was approved.

#### UG-T03-S27. Qwen3-8B

Qwen. Mutable model card; release revision not pinned. Model artifact. Not a project benchmark.

Identifier: Qwen/Qwen3-8B.

Direct URL:
https://huggingface.co/Qwen/Qwen3-8B

Origin and locator: UG-T03-S05 named comparison candidate.

Bibliographic record: Verified current publisher model card.

Content inspected: Relevant sections. Identity, architecture/use guidance; not benchmark reproduction.

Relevance: Possible generative checkpoint; UG-T03-D06.

Limits: No training, inference, hardware measurement or Unit-Generator comparison was performed.

Record notes: No license compatibility audit or checkpoint choice was approved.

#### UG-T03-S28. Qwen3.5-9B

Qwen. Mutable model card; release revision not pinned. Model artifact. Not a project benchmark.

Identifier: Qwen/Qwen3.5-9B.

Direct URL:
https://huggingface.co/Qwen/Qwen3.5-9B

Origin and locator: Visible old browsing trace only; not selected in the substantive answer.

Bibliographic record: Verified current publisher model card.

Content inspected: Relevant sections. Identity, architecture/use guidance; not benchmark reproduction.

Relevance: Possible generative checkpoint; UG-T03-D06.

Limits: No training, inference, hardware measurement or Unit-Generator comparison was performed.

Record notes: No license compatibility audit or checkpoint choice was approved.

#### UG-T03-S29. TRL SFTTrainer documentation and PEFT integration

Hugging Face. Mutable documentation; version not pinned. Technical documentation. Not empirical TD research.

Identifier: TRL SFTTrainer.

Direct URL:
https://huggingface.co/docs/trl/sft_trainer

Origin and locator: UG-T03-S05, training implementation option.

Bibliographic record: Verified current official documentation.

Content inspected: Relevant sections. SFTTrainer, adapter and quantized configuration guidance.

Relevance: Possible SFT/adapter implementation; UG-T03-D06.

Limits: No version-pinned training script or working run was produced.

#### UG-T03-S30. Official ONE PIECE Robin character feature

ONE PIECE official website. URL contains 2021-02-06; original page metadata not currently recoverable. Character source material. Official franchise site.

Identifier: News item o20210206_12101.

Direct URL:
https://one-piece.com/news/o20210206_12101/index.html

Origin and locator: UG-T03-S06, Robin remote-action example.

Bibliographic record: Unverified in current pass: page fetch failed.

Content inspected: Not inspected. Original assistant attribution retained; no fresh page content available.

Relevance: Canon starting point cited for illustrative ActionOrigin proposal; UG-T03-E02.

Limits: Do not treat every proposed targeting, ownership or duration rule as canon. The ActionOrigin design was assistant invention.

#### UG-T03-S31. Copyright in Derivative Works and Compilations

U.S. Copyright Office. Current circular; original revision not pinned. Official legal information. Circular 14.

Identifier: Circular 14.

Direct URL:
https://www.copyright.gov/circs/circ14.pdf

Origin and locator: UG-T03-S05, final rights warning.

Bibliographic record: Verified official document.

Content inspected: Relevant sections. Definition/scope of derivative works and compilations.

Relevance: Separate rights question, UG-T03-Q08.

Limits: U.S. general guidance does not settle EU law, a specific character license, training-data rights, or commercial permission. No project legal conclusion adopted.

#### UG-T03-S32. System One models concepts

Typesafe. Mutable website. Vendor technical description. Not independent research.

Identifier: System One concepts.

Direct URL:
https://docs.typesafe.ai/concepts/system-one

Origin and locator: Visible original browsing trace following Laya/decision-model question.

Bibliographic record: Verified current vendor page.

Content inspected: Relevant sections. Bounded typed decisions description.

Relevance: Exploratory model-category reference; UG-T03-D07.

Limits: Vendor terminology and claims are not a demonstrated Unit-Generator approach.

#### UG-T03-S33. Introducing System One Models and Jev

Typesafe. Date not captured. Vendor blog. Not independent research.

Identifier: Vendor introduction.

Direct URL:
https://typesafe.ai/blog/introducing-system-one-models-and-jev

Origin and locator: Visible original browsing trace; not an adopted dependency.

Bibliographic record: Verified current vendor page.

Content inspected: Relevant sections. Introduction and stated capabilities.

Relevance: System One decision-model background; UG-T03-D07.

Limits: No independent calibration or TD generation evidence.

#### UG-T03-S34. Consistent Game Content Creation via Function Calling for Large Language Models

Roberto Gallotta; Antonios Liapis; Georgios N. Yannakakis. 2024. Published conference paper. IEEE Conference on Games 2024.

Identifier: IEEE article 10645599.

Direct URLs:
https://antoniosliapis.com/papers/consistent_game_content_creation_via_function_calling_for_large_language_models.pdf
https://antoniosliapis.com/projects/project_llmaker.php

Origin and locator: Visible original browsing trace; earlier answer planning, not its final main paper list.

Bibliographic record: Verified author/project bibliographic record.

Content inspected: Abstract only. Author-hosted abstract and project record.

Relevance: Typed backend operations for content authoring; UG-T03-E08.

Limits: Function calling does not ensure faithful or strategically useful design.

Associated artifacts: https://github.com/gallorob/llmaker_functioncalling

#### UG-T03-S35. Procedural Content Generation via Machine Learning (PCGML)

Adam Summerville et al.. 2017 preprint; v3 2018. Survey, preprint record consulted. Final journal version not independently checked here.

Identifier: arXiv:1702.00539v3.

Direct URL:
https://arxiv.org/abs/1702.00539

Origin and locator: Visible original trace, opened HTML v3.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Survey scope and metadata.

Relevance: Survey context for model/data/representation choices; UG-T03-D03.

Limits: Broad PCG survey, not a direct character-adaptation result.

#### UG-T03-S36. Neural Networks Models for Analyzing Magic: the Gathering Cards

Felipe Zilio; Marcelo O. R. Prates; Luís C. Lamb. 2018. Conference paper; preprint record consulted. ICONIP 2018 accepted, per arXiv record.

Identifier: arXiv:1810.03744v1.

Direct URLs:
https://arxiv.org/abs/1810.03744
https://ar5iv.labs.arxiv.org/html/1810.03744

Origin and locator: Visible original trace, ar5iv link opened.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Card classification and image-conditioned text objective.

Relevance: Card text/image relationship lead; UG-T03-D03.

Limits: No faithful wiki-to-Unit or executable-balance result established.

#### UG-T03-S37. Type-Constrained Code Generation with Language Models

Niels Mündler; Jingxuan He; Hao Wang; Koushik Sen; Dawn Song; Martin Vechev. 2025. Research paper, preprint record consulted. Related publisher DOI recorded; venue not independently checked.

Identifier: arXiv:2504.09246v2; related DOI:10.1145/3729274.

Direct URL:
https://arxiv.org/abs/2504.09246

Origin and locator: Visible original browsing trace.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Simply typed language and TypeScript task scope.

Relevance: Type-aware generation beyond JSON syntax; UG-T03-E08.

Limits: Typing/code benchmarks do not certify game semantics or usefulness.

#### UG-T03-S38. "XGrammar Flexible and Efficient Structured Generation Engine Large Language Models"

Not recovered in this handover. Not established. Unresolved title/query reference. Not established.

Identifier: Original search wording.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Visible original search query, distinct from later XGrammar-2 link.

Bibliographic record: Unverified: do not assign the XGrammar-2 paper to this query.

Content inspected: Not inspected. Query text only; targeted recovery did not yield a dependable matching record.

Relevance: Structured decoding research trail; UG-T03-E08.

Limits: No author, identifier, exact title or finding inferred.

Associated artifacts: Current shared repository is registered under UG-T03-S16.

#### UG-T03-S39. UrzaGPT: LoRA-Tuned Large Language Models for Card Selection in Collectible Card Games

Timo Bertram. 2025. Preprint. No venue established here.

Identifier: arXiv:2508.08382v1.

Direct URL:
https://arxiv.org/abs/2508.08382

Origin and locator: Visible original trace, opened HTML.

Bibliographic record: Verified arXiv record.

Content inspected: Abstract only. Draft-log training and card-selection evaluation.

Relevance: Specialist card-decision training lead; UG-T03-D07.

Limits: Card selection is not card invention or character adaptation; not a proven primary Unit generator.

#### UG-T03-S40. Distilling Step-by-Step! Outperforming Larger Language Models with Less Training Data and Smaller Model Sizes

Cheng-Yu Hsieh et al.. 2023. Published findings paper; preprint consulted. Findings of ACL 2023, per arXiv.

Identifier: arXiv:2305.02301v2.

Direct URL:
https://arxiv.org/abs/2305.02301

Origin and locator: Visible original trace; separate from UG-T03-S20.

Bibliographic record: Verified arXiv record and stated acceptance.

Content inspected: Abstract only. NLP benchmark scope and method summary.

Relevance: Small-model training/distillation alternative; UG-T03-D06.

Limits: Results across four NLP benchmarks are not Unit-Generator sample-efficiency evidence.

#### UG-T03-S41. RPGAgent: Driving Coherent Story-to-Play Generation with an LLM-Based Multi-Agent System

Author list not independently recovered from accessible primary text. 2026. Conference paper record found. CHI 2026.

Identifier: DOI:10.1145/3772318.3790326.

Direct URLs:
https://dl.acm.org/doi/10.1145/3772318.3790326
https://programs.sigchi.org/chi/2026/program/content/222237

Origin and locator: Visible original trace, paper/program links and title query.

Bibliographic record: Verified title/DOI in primary search record only; full metadata not verified.

Content inspected: Metadata only. Publisher/program search records; article fetch returned 403.

Relevance: Story-to-play orchestration lead, not adopted architecture; UG-T03-D05.

Limits: No methods or performance claim accepted from a snippet.

#### UG-T03-S42. "1001 Nights" / "game mechanics"

Not resolved. Not established. Unresolved research/game reference. Not established.

Identifier: Original search terms.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Visible original browsing query.

Bibliographic record: Unverified.

Content inspected: Not inspected. No dependable primary match recovered in this pass.

Relevance: Narrative-to-gameplay lead; no main claim depends on it.

Limits: Do not identify a particular version, author, year or paper by guesswork.

#### UG-T03-S43. Damage Optimization in Video Games: A Player-Driven Co-Creative Approach

Johannes Pfau; Manik Charan; Erica Kleinman; Magy Seif El-Nasr. 2024. Published conference paper. CHI 2024, article 368, 16 pages.

Identifier: DOI:10.1145/3613904.3642643.

Direct URLs:
https://doi.org/10.1145/3613904.3642643
https://research-portal.uu.nl/ws/portalfiles/portal/240419261/3613904.3642643.pdf

Origin and locator: Original trace query "Damage optimization in video games".

Bibliographic record: Verified institutional bibliographic record.

Content inspected: Metadata only. Utrecht author publications record.

Relevance: Player-directed numeric calibration lead; UG-T03-Q10.

Limits: Not used as evidence that generated Units are balanced. Underlying study not reviewed here.

#### UG-T03-S44. Progression Balancing × Baldur's Gate 3: Insights, Terms and Tools for Multi-Dimensional Video Game Balance

Johannes Pfau. 2025. Published conference paper. CHI 2025, article 468, 12 pages.

Identifier: DOI:10.1145/3706598.3713162.

Direct URLs:
https://doi.org/10.1145/3706598.3713162
https://research-portal.uu.nl/en/publications/progression-balancing-baldurs-gate-3-insights-terms-and-tools-for/

Origin and locator: Original trace queries "Progression Balancing" and "Baldur".

Bibliographic record: Verified institutional record.

Content inspected: Abstract only. Institutional abstract/metadata.

Relevance: Progression-aware calibration lead; UG-T03-Q10.

Limits: Different game and evaluation purpose; no numeric targets transferred.

#### UG-T03-S45. Author publication navigation pages

Eric Butler; Mike Treanor. Mutable websites. Bibliographic navigation sources. Not research results.

Identifier: Author websites.

Direct URLs:
https://www.ericbutler.net/publications
https://mtreanor.com

Origin and locator: Visible original browsing trace.

Bibliographic record: Verified as navigation references; no historical page snapshot.

Content inspected: Metadata only. Current identity/links where accessible.

Relevance: Recover paper identities, especially UG-T03-S11 and UG-T03-S93.

Limits: A publication listing is not a paper review.

#### UG-T03-S46. Unrecoverable opaque browsing references

Not attributable. Not established. Unresolved tool-provenance references. Not applicable.

Identifier: Opaque turn...view/search identifiers with redacted target output.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Visible historical tool calls include references whose output was redacted and whose URL/title cannot be recovered.

Bibliographic record: Unverified.

Content inspected: Not inspected. Only visible query/URL/title-bearing sources are registered individually.

Relevance: Coverage limitation; UG-T03-E01.

Limits: No fabricated metadata or findings. Generic search queries are retained as a coverage category, not counted as papers.

#### UG-T03-S47. Current owner-supplied product context for this handover

Project owner. Current handover request; no historical date assigned. Current product context. Not a historical thread decision.

Identifier: User handover request: CURRENT PRODUCT CONTEXT.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: Current request, explicitly distinguished from earlier conversation.

Bibliographic record: Verified in visible current request.

Content inspected: Full text. All owner-supplied context bullets.

Relevance: Controls the new v1 recommendation and public/private boundary; UG-T03-Q01, UG-T03-Q04, UG-T03-Q05 and UG-T03-Q08.

Limits: Must not be cited as proof that the earlier research conversation approved those choices.

#### UG-T03-S48. Referenced but unavailable repository-local documents and private guidance

Mardwerk, attribution from supplied context. Dates not established, except archive filename. Missing internal references. Not inspected.

Identifier: Local paths below.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: UG-T03-S01 project fit, archive and final indexes; UG-T03-S02 opening and closing paragraphs.

Bibliographic record: Unverified beyond references in supplied files.

Content inspected: Not inspected. Paths: unit-generator/docs/PRODUCT.md; docs/PRODUCT.md; docs/API.md; towerright/CONTEXT.md; unit-generator/research/archive/GENERATOR-SURVEY-2026-09-21.md; GENERATOR-REFERENCES.md; unit-generator/research/FRAMEWORK-SOURCES.md; Foundation guidance.

Relevance: Unverified code scope and disclosure boundary; UG-T03-Q04, UG-T03-Q08 and UG-T03-Q12.

Limits: No implied access to source code, other repositories or private Foundation. Relative paths are not downloadable attachments.

#### UG-T03-S49. Related Mardwerk repositories

Mardwerk. Mutable repositories. Repository pointers. Not empirical research.

Identifier: map-generator; .github.

Direct URLs:
https://github.com/mardwerk/map-generator
https://github.com/mardwerk/.github

Origin and locator: CONTEXT.md, final related-repositories line.

Bibliographic record: Unverified repository contents; pointers verified in supplied file.

Content inspected: Not inspected. Only supplied URLs.

Relevance: Ecosystem context only; neither is a standalone v1 dependency.

Limits: No source code, deployment, license or integration decision inspected.

#### UG-T03-S50. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R1.

Direct URL:
https://arxiv.org/html/2212.03387v1

Origin and locator: UG-T03-S01, R1: Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Relevant sections. Bounded microRTS representation and evaluation description.

Relevance: UG-T03-E09; motivates explicit mechanics without establishing open-ended adaptation.

Limits: Four causes and four effects; 20 games per candidate; ten shown units; one map. Agent competence limits interpretation.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Kynan Sorochan and Matthew Guzdial. Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search. Experimental AI in Games Workshop, 2022. arXiv:2212.03387v1.

Original attachment inspection/access statement: Kynan Sorochan and Matthew Guzdial. **Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search.** Experimental AI in Games Workshop, 2022. [arXiv:2212.03387v1](https://arxiv.org/html/2212.03387v1). Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.

#### UG-T03-S51. Evolutionary Game Design

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R2.

Direct URL:
https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf

Origin and locator: UG-T03-S01, R2: Evolutionary Game Design.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Author manuscript identity and introductory summary.

Relevance: UG-T03-E09; preference prediction and originality distinctions.

Limits: Attachment reports 1,389 games, 19 viable, 127 paired comparisons from 27 participants. Those are source-reported results, not reproduced or TD preference coefficients.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Cameron Browne and Frédéric Maire. Evolutionary Game Design. IEEE Transactions on Computational Intelligence and AI in Games 2(1), pages 1 to 16, 2010. Author PDF.

Original attachment inspection/access statement: Cameron Browne and Frédéric Maire. **Evolutionary Game Design.** IEEE Transactions on Computational Intelligence and AI in Games 2(1), pages 1 to 16, 2010. [Author PDF](https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf). Inspected: quality measurements, human comparisons, Experiment III and conclusions. The filename is not the publication year.

#### UG-T03-S52. GAVEL: Generating Games Via Evolution and Language Models

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R3.

Direct URLs:
https://arxiv.org/html/2407.09388v2
https://github.com/gdrtodd/gavel

Origin and locator: UG-T03-S01, R3: GAVEL: Generating Games Via Evolution and Language Models.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv bibliographic record and method scope.

Relevance: UG-T03-E09; external checks, unused mechanics and budget accounting.

Limits: Attachment reports three runs, about 48 hours per run on RTX8000 plus 16 CPU cores and 128 GB RAM; novel occupied-cell gains not significant. Not interactive-latency evidence.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. GAVEL: Generating Games Via Evolution and Language Models. NeurIPS 2024. arXiv:2407.09388v2.

Original attachment inspection/access statement: Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. **GAVEL: Generating Games Via Evolution and Language Models.** NeurIPS 2024. [arXiv:2407.09388v2](https://arxiv.org/html/2407.09388v2). Inspected: representation, mutation, evaluation, experiments, results, limitations and references. [Code link](https://github.com/gdrtodd/gavel), not executed in this investigation.

#### UG-T03-S53. Procedural Content Generation through Quality Diversity

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R4.

Direct URLs:
https://arxiv.org/abs/1907.04053
https://arxiv.org/html/1907.04053
https://doi.org/10.1109/CIG.2019.8848053

Origin and locator: UG-T03-S01, R4: Procedural Content Generation through Quality Diversity.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv survey abstract and metadata.

Relevance: UG-T03-E09; retain useful alternatives, not mandatory MAP-Elites infrastructure.

Limits: Quality-diversity is a search framework; chosen descriptors do not establish meaningful strategic diversity.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. Procedural Content Generation through Quality Diversity. IEEE Conference on Games, 2019. arXiv:1907.04053, HTML, DOI:10.1109/CIG.2019.8848053.

Original attachment inspection/access statement: Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. **Procedural Content Generation through Quality Diversity.** IEEE Conference on Games, 2019. [arXiv:1907.04053](https://arxiv.org/abs/1907.04053), [HTML](https://arxiv.org/html/1907.04053), [DOI:10.1109/CIG.2019.8848053](https://doi.org/10.1109/CIG.2019.8848053). Inspected: framework, surveyed applications, challenges and references.

#### UG-T03-S54. Illuminating search spaces by mapping elites

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R5.

Direct URL:
https://arxiv.org/abs/1504.04909v1

Origin and locator: UG-T03-S01, R5: Illuminating search spaces by mapping elites.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Original 2015 arXiv record, not an inconsistent HTML-generated date.

Relevance: UG-T03-E09; archive idea, not a v1 algorithm decision.

Limits: Robotics/neural-network examples; archive coverage concerns selected descriptors.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Jean-Baptiste Mouret and Jeff Clune. Illuminating search spaces by mapping elites. 2015 preprint. arXiv:1504.04909v1.

Original attachment inspection/access statement: Jean-Baptiste Mouret and Jeff Clune. **Illuminating search spaces by mapping elites.** 2015 preprint. [arXiv:1504.04909v1](https://arxiv.org/abs/1504.04909v1). Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.

#### UG-T03-S55. Abandoning Objectives: Evolution through the Search for Novelty Alone

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R6.

Direct URL:
https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf

Origin and locator: UG-T03-S01, R6: Abandoning Objectives: Evolution through the Search for Novelty Alone.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Author-hosted paper abstract and identity.

Relevance: UG-T03-E09; avoid a single narrow objective.

Limits: Maze and walking experiments do not prove that novelty alone yields useful units.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Joel Lehman and Kenneth O. Stanley. Abandoning Objectives: Evolution through the Search for Novelty Alone. Evolutionary Computation 19(2), pages 189 to 223, 2011. Primary-text PDF.

Original attachment inspection/access statement: Joel Lehman and Kenneth O. Stanley. **Abandoning Objectives: Evolution through the Search for Novelty Alone.** Evolutionary Computation 19(2), pages 189 to 223, 2011. [Primary-text PDF](https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf). Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.

#### UG-T03-S56. Monte-Carlo Tree Search for Simulation-based Strategy Analysis

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R7.

Direct URL:
https://arxiv.org/html/1908.01423v1

Origin and locator: UG-T03-S01, R7: Monte-Carlo Tree Search for Simulation-based Strategy Analysis.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record; deposit year distinguished from original conference.

Relevance: UG-T03-E09; behavior traces and evaluator competence.

Limits: MCTS budget is not a validated human skill model.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Alexander Zook, Brent Harrison and Mark O. Riedl. Monte-Carlo Tree Search for Simulation-based Strategy Analysis. The paper identifies FDG 2015 as its conference; the arXiv deposit is 2019. arXiv:1908.01423v1.

Original attachment inspection/access statement: Alexander Zook, Brent Harrison and Mark O. Riedl. **Monte-Carlo Tree Search for Simulation-based Strategy Analysis.** The paper identifies FDG 2015 as its conference; the arXiv deposit is 2019. [arXiv:1908.01423v1](https://arxiv.org/html/1908.01423v1). Inspected: methodology, skill-based metrics, case studies, limitations and references. R1 cites the deposit year.

#### UG-T03-S57. Evaluating Competitive Game Balance with Restricted Play

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R8.

Direct URL:
https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf

Origin and locator: UG-T03-S01, R8: Evaluating Competitive Game Balance with Restricted Play.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Author-hosted paper abstract and identity.

Relevance: UG-T03-E09; future controlled mechanic comparisons.

Limits: Restricting an agent and removing a mechanic are different interventions. Attachment adaptations are not original TD measurements.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Alexander Jaffe, Alex Miller, Erik Andersen, Yun-En Liu, Anna Karlin and Zoran Popović. Evaluating Competitive Game Balance with Restricted Play. AIIDE 2012. Author-hosted PDF.

Original attachment inspection/access statement: Alexander Jaffe, Alex Miller, Erik Andersen, Yun-En Liu, Anna Karlin and Zoran Popović. **Evaluating Competitive Game Balance with Restricted Play.** AIIDE 2012. [Author-hosted PDF](https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf). Inspected: framework, example application and conclusions. Reached through R1 → R7 → R8.

#### UG-T03-S58. Automated Game Design via Conceptual Expansion

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R9.

Direct URL:
https://arxiv.org/html/1809.02232v1

Origin and locator: UG-T03-S01, R9: Automated Game Design via Conceptual Expansion.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv abstract and metadata.

Relevance: UG-T03-E09; preserve changed mechanical relationships.

Limits: Held-out game reconstruction using learned approximations, not unrestricted historical originality.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Matthew Guzdial and Mark Riedl. Automated Game Design via Conceptual Expansion. AIIDE 2018, 14(1), pages 31 to 37. arXiv:1809.02232v1.

Original attachment inspection/access statement: Matthew Guzdial and Mark Riedl. **Automated Game Design via Conceptual Expansion.** AIIDE 2018, 14(1), pages 31 to 37. [arXiv:1809.02232v1](https://arxiv.org/html/1809.02232v1). Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.

#### UG-T03-S59. Constrained Novelty Search: A Study on Game Content Generation

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R10.

Direct URLs:
https://antoniosliapis.com/papers/constrained_novelty_search.pdf
https://doi.org/10.1162/EVCO_a_00123

Origin and locator: UG-T03-S01, R10: Constrained Novelty Search: A Study on Game Content Generation.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Author-hosted abstract and identity.

Relevance: UG-T03-D05; keeping unsupported ideas without claiming validity.

Limits: Feasible/infeasible populations for levels; storing proposals separately does not require implementing that algorithm.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. Constrained Novelty Search: A Study on Game Content Generation. Evolutionary Computation 23(1), pages 101 to 129, 2015. Author PDF, DOI:10.1162/EVCO_a_00123.

Original attachment inspection/access statement: Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Constrained Novelty Search: A Study on Game Content Generation.** Evolutionary Computation 23(1), pages 101 to 129, 2015. [Author PDF](https://antoniosliapis.com/papers/constrained_novelty_search.pdf), [DOI:10.1162/EVCO_a_00123](https://doi.org/10.1162/EVCO_a_00123). Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.

#### UG-T03-S60. Sentient Sketchbook: Computer-Aided Game Level Authoring

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R11.

Direct URL:
https://antoniosliapis.com/papers/sentient_sketchbook.pdf

Origin and locator: UG-T03-S01, R11: Sentient Sketchbook: Computer-Aided Game Level Authoring.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Author-hosted abstract and identity.

Relevance: UG-T03-E09; designer review and controllability.

Limits: Small map-authoring study; validity does not establish useful suggestions or obedience to intent.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. Sentient Sketchbook: Computer-Aided Game Level Authoring. Foundations of Digital Games, 2013, pages 213 to 220. Author PDF.

Original attachment inspection/access statement: Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Sentient Sketchbook: Computer-Aided Game Level Authoring.** Foundations of Digital Games, 2013, pages 213 to 220. [Author PDF](https://antoniosliapis.com/papers/sentient_sketchbook.pdf). Inspected: tool interaction, checks, alternatives and small-scale designer study. Found through R4.

#### UG-T03-S61. Generation and Evaluation in the Human Invention Process through the Lens of Game Design

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R12.

Direct URL:
https://arxiv.org/html/2508.10914v1

Origin and locator: UG-T03-S01, R12: Generation and Evaluation in the Human Invention Process through the Lens of Game Design.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Current arXiv record; original attachment reviewed v1.

Relevance: UG-T03-E09; distinguish proposing from evaluating.

Limits: Attachment describes 484 submitted games, not all formalized/simulated. Current record identifies a non-archival conference paper.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Katherine M. Collins et al. Generation and Evaluation in the Human Invention Process through the Lens of Game Design. CogSci 2025. arXiv:2508.10914v1.

Original attachment inspection/access statement: Katherine M. Collins et al. **Generation and Evaluation in the Human Invention Process through the Lens of Game Design.** CogSci 2025. [arXiv:2508.10914v1](https://arxiv.org/html/2508.10914v1). Inspected: dataset, proposal/evaluation model, results, limitations and references.

#### UG-T03-S62. People use fast, goal-directed simulation to reason about novel games

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R13.

Direct URL:
https://arxiv.org/html/2407.14095v1

Origin and locator: UG-T03-S01, R13: People use fast, goal-directed simulation to reason about novel games.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Current arXiv record; original attachment reviewed v1.

Relevance: UG-T03-E09; novice comprehension differs from expert depth.

Limits: Simple connect-n judgments do not validate real-time TD experience prediction. Current v2 author list includes Mauricio Barba, absent from original v1-style entry.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Cedegao E. Zhang, Katherine M. Collins, Lionel Wong, Adrian Weller and Joshua B. Tenenbaum. People use fast, goal-directed simulation to reason about novel games. CogSci 2024. arXiv:2407.14095v1.

Original attachment inspection/access statement: Cedegao E. Zhang, Katherine M. Collins, Lionel Wong, Adrian Weller and Joshua B. Tenenbaum. **People use fast, goal-directed simulation to reason about novel games.** CogSci 2024. [arXiv:2407.14095v1](https://arxiv.org/html/2407.14095v1). Inspected: model, human judgment task and discussion. Found through R12.

#### UG-T03-S63. Mathematical discoveries from program search with large language models

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R14.

Direct URLs:
https://www.nature.com/articles/s41586-023-06924-6
https://github.com/google-deepmind/funsearch

Origin and locator: UG-T03-S01, R14: Mathematical discoveries from program search with large language models.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Relevant sections. Primary Nature article method and evaluation scope.

Relevance: UG-T03-E09; independent evaluation and equal total budget.

Limits: Externally checkable math/algorithm objectives do not establish fun, canon fidelity or LLM self-evaluation reliability.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Bernardino Romera-Paredes et al. Mathematical discoveries from program search with large language models. Nature 625, pages 468 to 475, 2024; published online in December 2023. DOI:10.1038/s41586-023-06924-6.

Original attachment inspection/access statement: Bernardino Romera-Paredes et al. **Mathematical discoveries from program search with large language models.** Nature 625, pages 468 to 475, 2024; published online in December 2023. [DOI:10.1038/s41586-023-06924-6](https://www.nature.com/articles/s41586-023-06924-6). Inspected: primary article's method and reported applications. [Code link](https://github.com/google-deepmind/funsearch), not executed here.

#### UG-T03-S64. Ludax: A GPU-Accelerated Description Language for Board Games

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R15.

Direct URLs:
https://arxiv.org/abs/2506.22609
https://arxiv.org/html/2506.22609v2

Origin and locator: UG-T03-S01, R15: Ludax: A GPU-Accelerated Description Language for Board Games.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Current arXiv record; original v2 retained.

Relevance: UG-T03-E08; declared operators require implementation.

Limits: Original attachment title says "Description Language"; current record says "Domain Specific Language". Supported board games are not a TD runtime.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Graham Todd et al. Ludax: A GPU-Accelerated Description Language for Board Games. arXiv:2506.22609, first submitted in 2025; v2 inspected, March 25, 2026.

Original attachment inspection/access statement: Graham Todd et al. **Ludax: A GPU-Accelerated Description Language for Board Games.** [arXiv:2506.22609](https://arxiv.org/abs/2506.22609), first submitted in 2025; [v2 inspected](https://arxiv.org/html/2506.22609v2), March 25, 2026. Inspected: representation, compilation, supported game class and references. R12 cites the earlier 2025 version; this record deliberately distinguishes it from the newer text inspected.

#### UG-T03-S65. Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R16.

Direct URLs:
https://arxiv.org/abs/2409.04109
https://arxiv.org/html/2409.04109v1

Origin and locator: UG-T03-S01, R16: Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv abstract/record; original attachment inspected preprint v1.

Relevance: UG-T03-E09; independent review, diversity and self-evaluation limits.

Limits: Idea ratings concern research proposals, not finished games or demonstrated discoveries.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Chenglei Si, Diyi Yang and Tatsunori Hashimoto. Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers. ICLR 2025; preprint 2024. arXiv:2409.04109, v1 inspected.

Original attachment inspection/access statement: Chenglei Si, Diyi Yang and Tatsunori Hashimoto. **Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers.** ICLR 2025; preprint 2024. [arXiv:2409.04109](https://arxiv.org/abs/2409.04109), [v1 inspected](https://arxiv.org/html/2409.04109v1). Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.

#### UG-T03-S66. Mortar: Evolving Mechanics for Automatic Game Design

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R17.

Direct URLs:
https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf
https://doi.org/10.1145/3795095.3805100

Origin and locator: UG-T03-S01, R17: Mortar: Evolving Mechanics for Automatic Game Design.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Metadata only. Author/institutional publication record.

Relevance: UG-T03-E09; context-dependent contribution, proxy failure.

Limits: Attachment reports 14 participants, six games, three pairs; human/proxy preference directions disagree for one pair. CITS is not exhaustive Shapley evaluation.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Muhammad U. Nasir, Yuchen Li, Steven James and Julian Togelius. Mortar: Evolving Mechanics for Automatic Game Design. GECCO 2026. Primary PDF, DOI:10.1145/3795095.3805100.

Original attachment inspection/access statement: Muhammad U. Nasir, Yuchen Li, Steven James and Julian Togelius. **Mortar: Evolving Mechanics for Automatic Game Design.** GECCO 2026. [Primary PDF](https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf), [DOI:10.1145/3795095.3805100](https://doi.org/10.1145/3795095.3805100). Inspected: method, mechanic contribution measure, human comparisons and related work.

#### UG-T03-S67. Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R18.

Direct URLs:
https://www.possibilityspace.org/papers/evo13.pdf
https://doi.org/10.1007/978-3-642-37192-9_29

Origin and locator: UG-T03-S01, R18: Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Author-hosted paper identity and abstract.

Relevance: UG-T03-E09; implementation and readability differ.

Limits: Attachment flags inconsistent participant counts and poor teleport-mechanic comprehension despite solver usefulness. No precise pilot sample adopted.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design. EvoApplications 2013. Author PDF, DOI:10.1007/978-3-642-37192-9_29.

Original attachment inspection/access statement: Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. **Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design.** EvoApplications 2013. [Author PDF](https://www.possibilityspace.org/papers/evo13.pdf), [DOI:10.1007/978-3-642-37192-9_29](https://doi.org/10.1007/978-3-642-37192-9_29). Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.

#### UG-T03-S68. Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R19.

Direct URLs:
https://arxiv.org/html/2309.09476v3
https://doi.org/10.1609/aiide.v19i1.27522

Origin and locator: UG-T03-S01, R19: Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: UG-T03-E09; evaluator dependence and reward risk.

Limits: RL and A* produce different rules; attachment warns of slowness and fitness exploitation. No human-subject evidence for suggested usability improvement.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Johor Jara Gonzalez, Seth Cooper and Matthew Guzdial. Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules. AIIDE 19(1), pages 266 to 275, 2023. arXiv:2309.09476v3, DOI:10.1609/aiide.v19i1.27522.

Original attachment inspection/access statement: Johor Jara Gonzalez, Seth Cooper and Matthew Guzdial. **Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules.** AIIDE 19(1), pages 266 to 275, 2023. [arXiv:2309.09476v3](https://arxiv.org/html/2309.09476v3), [DOI:10.1609/aiide.v19i1.27522](https://doi.org/10.1609/aiide.v19i1.27522). Inspected: environment recreation, evaluator comparison and limitations.

#### UG-T03-S69. Pixie: Code-Level Mechanic Generation for Game Designers

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R20.

Direct URLs:
https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962
https://doi.org/10.1609/aiide.v21i1.36824

Origin and locator: UG-T03-S01, R20: Pixie: Code-Level Mechanic Generation for Game Designers.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Publisher record/abstract now accessible.

Relevance: UG-T03-D05; future code-extension interface lead.

Limits: Original primary-PDF access gap remains for detailed Unity annotation, insertion-point and pooling claims. Metadata recovery is not verification of those details.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Michael Cook. Pixie: Code-Level Mechanic Generation for Game Designers. AIIDE 21(1), pages 206 to 215, 2025. Publisher PDF, DOI:10.1609/aiide.v21i1.36824.

Original attachment inspection/access statement: Michael Cook. **Pixie: Code-Level Mechanic Generation for Game Designers.** AIIDE 21(1), pages 206 to 215, 2025. [Publisher PDF](https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962), [DOI:10.1609/aiide.v21i1.36824](https://doi.org/10.1609/aiide.v21i1.36824). Access gap: the publisher download failed and an attempted author copy was unavailable. Mortar's related-work discussion independently confirms Pixie as code-level mechanic evolution using a genetic-programming approach.

#### UG-T03-S70. Automatic Game Design via Mechanic Generation

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R21.

Direct URLs:
https://arxiv.org/html/1908.01420v1
https://doi.org/10.1609/aaai.v28i1.8788

Origin and locator: UG-T03-S01, R21: Automatic Game Design via Mechanic Generation.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract; 2019 deposit is not publication year.

Relevance: UG-T03-E08; formal mechanics and behavioral checks.

Limits: Reachability within a planning model does not prove readability, taste or strategic depth.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Alexander Zook and Mark O. Riedl. Automatic Game Design via Mechanic Generation. AAAI 2014. arXiv:1908.01420v1, DOI:10.1609/aaai.v28i1.8788.

Original attachment inspection/access statement: Alexander Zook and Mark O. Riedl. **Automatic Game Design via Mechanic Generation.** AAAI 2014. [arXiv:1908.01420v1](https://arxiv.org/html/1908.01420v1), [DOI:10.1609/aaai.v28i1.8788](https://doi.org/10.1609/aaai.v28i1.8788). Inspected: representation, constraint-based generation, planning and demonstration domains. The arXiv deposit is 2019; the paper is from 2014. Promoted from the original R3/R7 citation lead.

#### UG-T03-S71. TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R22.

Direct URLs:
https://arxiv.org/html/2601.05899v1
https://doi.org/10.1609/aaai.v40i31.39818
https://github.com/tb6147877/TowerMind

Origin and locator: UG-T03-S01, R22: TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: UG-T03-E09; evaluation infrastructure, not mandatory v1 dependency.

Limits: Five-level agent benchmark, not novel-rule generator; compatibility with this Consumer untested.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Dawei Wang et al. TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents. AAAI 2026. arXiv:2601.05899v1, DOI:10.1609/aaai.v40i31.39818.

Original attachment inspection/access statement: Dawei Wang et al. **TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents.** AAAI 2026. [arXiv:2601.05899v1](https://arxiv.org/html/2601.05899v1), [DOI:10.1609/aaai.v40i31.39818](https://doi.org/10.1609/aaai.v40i31.39818). Inspected: environment, observation formats, evaluation scope and reported agent failures. [Code](https://github.com/tb6147877/TowerMind), not executed.

#### UG-T03-S72. Automatic generation of tower defense levels using PCG

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R23.

Direct URLs:
https://doi.org/10.1145/3337722.3337723
https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG

Origin and locator: UG-T03-S01, R23: Automatic generation of tower defense levels using PCG.

Bibliographic record: Unverified in current pass; supplied bibliographic record retained.

Content inspected: Not inspected. Current DOI fetch failed; historical ResearchGate 403 retained.

Relevance: Context only; later map/wave evaluation, UG-T03-Q10.

Limits: Kingdom Rush-style simulator and flat Monte Carlo details remain supplied-review claims.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Simon Liu et al. Automatic generation of tower defense levels using PCG. FDG 2019. DOI:10.1145/3337722.3337723, supplied manuscript entry.

Original attachment inspection/access statement: Simon Liu et al. **Automatic generation of tower defense levels using PCG.** FDG 2019. [DOI:10.1145/3337722.3337723](https://doi.org/10.1145/3337722.3337723), [supplied manuscript entry](https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG). Access gap: ResearchGate returned HTTP 403; the primary text was not independently inspected.

#### UG-T03-S73. Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R24.

Direct URLs:
https://doi.org/10.1007/978-3-030-75275-0_54
https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms

Origin and locator: UG-T03-S01, R24: Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms.

Bibliographic record: Unverified in current pass; supplied bibliographic record retained.

Content inspected: Not inspected. Current DOI fetch failed; historical ResearchGate 403 retained.

Relevance: Context only; no complete-game or asset-generation claim for UG.

Limits: Unity/Blender implementation and human-authored premise details remain supplied claims.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Vid Kraner, Iztok Fister Jr. and Lucija Brezočnik. Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms. NT 2021, Lecture Notes in Networks and Systems 233. DOI:10.1007/978-3-030-75275-0_54, supplied manuscript entry.

Original attachment inspection/access statement: Vid Kraner, Iztok Fister Jr. and Lucija Brezočnik. **Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms.** NT 2021, Lecture Notes in Networks and Systems 233. [DOI:10.1007/978-3-030-75275-0_54](https://doi.org/10.1007/978-3-030-75275-0_54), [supplied manuscript entry](https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms). Access gap: ResearchGate returned HTTP 403.

#### UG-T03-S74. Computational Intelligence and Tower Defence Games

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R25.

Direct URLs:
https://doi.org/10.1109/CEC.2011.5949738
https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games

Origin and locator: UG-T03-S01, R25: Computational Intelligence and Tower Defence Games.

Bibliographic record: Unverified in current pass; supplied bibliographic record retained.

Content inspected: Not inspected. Current DOI fetch failed; historical access gap retained.

Relevance: Context only; TD research lead, UG-T03-Q10.

Limits: No primary implementation or taxonomy detail independently rechecked.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Phillipa Avery, Julian Togelius, Elvis Alistar and Robert Pieter van Leeuwen. Computational Intelligence and Tower Defence Games. IEEE Congress on Evolutionary Computation, 2011. DOI:10.1109/CEC.2011.5949738, supplied manuscript entry.

Original attachment inspection/access statement: Phillipa Avery, Julian Togelius, Elvis Alistar and Robert Pieter van Leeuwen. **Computational Intelligence and Tower Defence Games.** IEEE Congress on Evolutionary Computation, 2011. [DOI:10.1109/CEC.2011.5949738](https://doi.org/10.1109/CEC.2011.5949738), [supplied manuscript entry](https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games). Access gap: ResearchGate returned HTTP 403.

#### UG-T03-S75. Controllable procedural map generation via multiobjective evolution

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R26.

Direct URL:
https://link.springer.com/article/10.1007/s10710-012-9174-5

Origin and locator: UG-T03-S01, R26: Controllable procedural map generation via multiobjective evolution.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Publisher abstract/metadata.

Relevance: Context only; later map-context evaluation.

Limits: Visual judgments of maps do not equal competitive playtests; heuristic objectives need domain transfer.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Julian Togelius et al. Controllable procedural map generation via multiobjective evolution. Genetic Programming and Evolvable Machines 14, pages 245 to 277, 2013. Publisher article and DOI.

Original attachment inspection/access statement: Julian Togelius et al. **Controllable procedural map generation via multiobjective evolution.** Genetic Programming and Evolvable Machines 14, pages 245 to 277, 2013. [Publisher article and DOI](https://link.springer.com/article/10.1007/s10710-012-9174-5). Abstract and reference list inspected; the full article requires access.

#### UG-T03-S76. Learning Controllable 3D Level Generators

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R27.

Direct URLs:
https://arxiv.org/html/2206.13623v3
https://doi.org/10.1145/3555858.3563273

Origin and locator: UG-T03-S01, R27: Learning Controllable 3D Level Generators.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv abstract/record.

Relevance: Context only; v1 does not require map generation.

Limits: Simplified Minecraft-style level-editing policies, not production RTS geometry.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Zehua Jiang, Sam Earle, Michael Cerny Green and Julian Togelius. Learning Controllable 3D Level Generators. FDG 2022. arXiv:2206.13623v3, DOI:10.1145/3555858.3563273.

Original attachment inspection/access statement: Zehua Jiang, Sam Earle, Michael Cerny Green and Julian Togelius. **Learning Controllable 3D Level Generators.** FDG 2022. [arXiv:2206.13623v3](https://arxiv.org/html/2206.13623v3), [DOI:10.1145/3555858.3563273](https://doi.org/10.1145/3555858.3563273). Inspected: PCGRL formulation, 3D movement requirements and control experiments.

#### UG-T03-S77. Scaling, Control and Generalization in Reinforcement Learning Level Generators

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R28.

Direct URL:
https://arxiv.org/html/2408.12525v1

Origin and locator: UG-T03-S01, R28: Scaling, Control and Generalization in Reinforcement Learning Level Generators.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Current arXiv record; original v1 title retained.

Relevance: Context only; no RL or GPU infrastructure commitment.

Limits: Original title omitted "PCGRL+"; current record includes it. Faster RL training does not establish useful gameplay metrics.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Sam Earle, Zehua Jiang and Julian Togelius. Scaling, Control and Generalization in Reinforcement Learning Level Generators. 2024. arXiv:2408.12525v1.

Original attachment inspection/access statement: Sam Earle, Zehua Jiang and Julian Togelius. **Scaling, Control and Generalization in Reinforcement Learning Level Generators.** 2024. [arXiv:2408.12525v1](https://arxiv.org/html/2408.12525v1). Inspected: method, scaling and generalization scope. The supplied list calls this "PCGRL+"; the citation here uses the manuscript title.

#### UG-T03-S78. A Database-Driven Framework for 3D Level Generation with LLMs

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R29.

Direct URLs:
https://arxiv.org/html/2508.18533v1
https://doi.org/10.1609/aiide.v21i1.36840

Origin and locator: UG-T03-S01, R29: A Database-Driven Framework for 3D Level Generation with LLMs.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: Context only; reusable content/repair lead, not Unit generation proof.

Limits: Attachment records manual curation, prompting effort, unplayable outputs and no player evaluation.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Kaijie Xu and Clark Verbrugge. A Database-Driven Framework for 3D Level Generation with LLMs. AIIDE 21(1), pages 367 to 376, 2025. arXiv:2508.18533v1, DOI:10.1609/aiide.v21i1.36840.

Original attachment inspection/access statement: Kaijie Xu and Clark Verbrugge. **A Database-Driven Framework for 3D Level Generation with LLMs.** AIIDE 21(1), pages 367 to 376, 2025. [arXiv:2508.18533v1](https://arxiv.org/html/2508.18533v1), [DOI:10.1609/aiide.v21i1.36840](https://doi.org/10.1609/aiide.v21i1.36840). Inspected: assembly pipeline, repair and limitations.

#### UG-T03-S79. Infinite Photorealistic Worlds using Procedural Generation

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R30.

Direct URL:
https://arxiv.org/abs/2306.09310

Origin and locator: UG-T03-S01, R30: Infinite Photorealistic Worlds using Procedural Generation.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv abstract/record.

Relevance: Context only; assets excluded from automatic v1 promise.

Limits: Synthetic natural-scene geometry/materials, not validated tactical terrain or units.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Alexander Raistrick et al. Infinite Photorealistic Worlds using Procedural Generation. CVPR 2023. arXiv:2306.09310.

Original attachment inspection/access statement: Alexander Raistrick et al. **Infinite Photorealistic Worlds using Procedural Generation.** CVPR 2023. [arXiv:2306.09310](https://arxiv.org/abs/2306.09310). Primary HTML inspected for system purpose and generation scope.

#### UG-T03-S80. Holodeck: Language Guided Generation of 3D Embodied AI Environments

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R31.

Direct URL:
https://arxiv.org/abs/2312.09067

Origin and locator: UG-T03-S01, R31: Holodeck: Language Guided Generation of 3D Embodied AI Environments.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv abstract/record.

Relevance: Context only; assets/spatial realization are separate work.

Limits: Embodied navigation/scene preference evidence is not TD map balance.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Yue Yang et al. Holodeck: Language Guided Generation of 3D Embodied AI Environments. CVPR 2024. arXiv:2312.09067.

Original attachment inspection/access statement: Yue Yang et al. **Holodeck: Language Guided Generation of 3D Embodied AI Environments.** CVPR 2024. [arXiv:2312.09067](https://arxiv.org/abs/2312.09067). Primary HTML inspected for planning, retrieval, spatial constraints and evaluation scope.

#### UG-T03-S81. Automatic Content Generation in the Galactic Arms Race Video Game

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R32.

Direct URLs:
https://doi.org/10.1109/TCIAIG.2009.2038365
https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game

Origin and locator: UG-T03-S01, R32: Automatic Content Generation in the Galactic Arms Race Video Game.

Bibliographic record: Unverified in current pass; supplied bibliographic record retained.

Content inspected: Not inspected. Current DOI fetch failed; historical ResearchGate 403 retained.

Relevance: UG-T03-E09 context; use frequency is not necessarily preference.

Limits: More than 1,000 registered players and player-guided weapon evolution remain supplied-review claims here.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. Automatic Content Generation in the Galactic Arms Race Video Game. IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. DOI:10.1109/TCIAIG.2009.2038365, supplied manuscript entry.

Original attachment inspection/access statement: Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. **Automatic Content Generation in the Galactic Arms Race Video Game.** IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. [DOI:10.1109/TCIAIG.2009.2038365](https://doi.org/10.1109/TCIAIG.2009.2038365), [supplied manuscript entry](https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game). Access gap: ResearchGate returned HTTP 403.

#### UG-T03-S82. OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R33.

Direct URL:
https://arxiv.org/html/2405.15568v3

Origin and locator: UG-T03-S01, R33: OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: UG-T03-E09; preserve failed proposals, avoid overstated execution evidence.

Limits: Attachment distinguishes 200-iteration long run with no RL training from shorter actual-learning runs. Interestingness is model-judged.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Maxence Faldor, Jenny Zhang, Antoine Cully and Jeff Clune. OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code. ICLR 2025. arXiv:2405.15568v3.

Original attachment inspection/access statement: Maxence Faldor, Jenny Zhang, Antoine Cully and Jeff Clune. **OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code.** ICLR 2025. [arXiv:2405.15568v3](https://arxiv.org/html/2405.15568v3). Inspected: archive, task/environment generation, interestingness filtering and evaluation modes.

#### UG-T03-S83. DreamGarden: A Designer Assistant for Growing Games from a Single Prompt

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R34.

Direct URLs:
https://arxiv.org/html/2410.01791v2
https://doi.org/10.1145/3706598.3714233

Origin and locator: UG-T03-S01, R34: DreamGarden: A Designer Assistant for Growing Games from a Single Prompt.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: Context only; no multi-agent/orchestration requirement for v1.

Limits: Ten-participant assistant usability study, not full novel-game strategic validation.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Sam Earle, Samyak Parajuli and Andrzej Banburski-Fahey. DreamGarden: A Designer Assistant for Growing Games from a Single Prompt. CHI 2025. arXiv:2410.01791v2, DOI:10.1145/3706598.3714233.

Original attachment inspection/access statement: Sam Earle, Samyak Parajuli and Andrzej Banburski-Fahey. **DreamGarden: A Designer Assistant for Growing Games from a Single Prompt.** CHI 2025. [arXiv:2410.01791v2](https://arxiv.org/html/2410.01791v2), [DOI:10.1145/3706598.3714233](https://doi.org/10.1145/3706598.3714233). Inspected: hierarchical planning, Unreal implementation modules and usability study.

#### UG-T03-S84. AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R35.

Direct URLs:
https://arxiv.org/html/2603.07106v1
https://github.com/Pluto156/AutoUE

Origin and locator: UG-T03-S01, R35: AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: UG-T03-E09; later Consumer integration, not an approved dependency.

Limits: PlayGen-20 has twenty tasks. Generated tests/model assessments can share generator omissions.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Lei Yin et al. AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems. 2026 preprint. arXiv:2603.07106v1.

Original attachment inspection/access statement: Lei Yin et al. **AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems.** 2026 preprint. [arXiv:2603.07106v1](https://arxiv.org/html/2603.07106v1). Inspected: engine workflow, documentation retrieval, runtime tests and benchmark. [Code](https://github.com/Pluto156/AutoUE), not executed.

#### UG-T03-S85. Genie: Generative Interactive Environments

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R36.

Direct URL:
https://arxiv.org/abs/2402.15391

Origin and locator: UG-T03-S01, R36: Genie: Generative Interactive Environments.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. ArXiv record and abstract.

Relevance: Context only; distinguish output types.

Limits: Interactive visual world model is not explicit editable mechanics or conventional engine code.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Jake Bruce et al. Genie: Generative Interactive Environments. ICML 2024. arXiv:2402.15391.

Original attachment inspection/access statement: Jake Bruce et al. **Genie: Generative Interactive Environments.** ICML 2024. [arXiv:2402.15391](https://arxiv.org/abs/2402.15391). Primary HTML inspected for model and output representation.

#### UG-T03-S86. Emergence of player tactics by expert-guided machine learning: an industry tower defence case study

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R37.

Direct URLs:
https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/
https://doi.org/10.1016/j.entcom.2025.100963

Origin and locator: UG-T03-S01, R37: Emergence of player tactics by expert-guided machine learning: an industry tower defence case study.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Institutional record/abstract.

Relevance: UG-T03-E09; contextual tactics rather than pick rates.

Limits: Actual commercial TD behavior analysis, not generation; full-text details still not checked.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Stuart Anderson and Ruth Falconer. Emergence of player tactics by expert-guided machine learning: an industry tower defence case study. Entertainment Computing 54, article 100963, 2025. Institutional record, DOI:10.1016/j.entcom.2025.100963.

Original attachment inspection/access statement: Stuart Anderson and Ruth Falconer. **Emergence of player tactics by expert-guided machine learning: an industry tower defence case study.** Entertainment Computing 54, article 100963, 2025. [Institutional record](https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/), [DOI:10.1016/j.entcom.2025.100963](https://doi.org/10.1016/j.entcom.2025.100963). Institutional abstract and metadata inspected; the linked full-text download returned HTTP 403.

#### UG-T03-S87. GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development

See original bibliographic entry below. See original bibliographic entry below. Attachment catalogue paper or survey. See original bibliographic entry below.

Identifier: R38.

Direct URL:
https://arxiv.org/html/2609.21293v1

Origin and locator: UG-T03-S01, R38: GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development.

Bibliographic record: Verified source identity from primary record/text; publication details not checked there remain attributed to supplied catalogue.

Content inspected: Abstract only. Current arXiv record and abstract.

Relevance: UG-T03-E09; complete-task success differs from average passed checks.

Limits: Attachment reports 47 tasks, best mean check pass 93.2%, strict success 55.3% (26/47). Different maxima need not come from the same stack; no replication here.

Associated artifacts: Associated code/data links are retained in URLs; none installed, benchmarked or license-audited.

Record notes: Original bibliographic entry: Xiuhui Zhang et al. GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development. Preprint, September 18, 2026. arXiv:2609.21293v1.

Original attachment inspection/access statement: Xiuhui Zhang et al. **GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development.** Preprint, September 18, 2026. [arXiv:2609.21293v1](https://arxiv.org/html/2609.21293v1). Inspected: evaluation contract, task suite and aggregate results.

#### UG-T03-S88. Experience-Driven Procedural Content Generation

Georgios N. Yannakakis; Julian Togelius. 2011. Published work, supplied bibliographic record. IEEE Transactions on Affective Computing 2(3).

Identifier: DOI:10.1109/T-AFFC.2011.6.

Direct URL:
https://doi.org/10.1109/T-AFFC.2011.6

Origin and locator: UG-T03-S01, Additional references and access gaps, first table.

Bibliographic record: Unverified in current pass; publisher/DOI access failed.

Content inspected: Not inspected. Historical source/access-gap wording retained.

Relevance: Frame player/experience-dependent evaluation. Related UG-T03-E09.

Limits: No findings imported from unavailable full text; no TD transfer presumed.

#### UG-T03-S89. Generative AI enhances individual creativity but reduces the collective diversity of novel content

Anil R. Doshi; Oliver P. Hauser. 2024. Published work, supplied bibliographic record. Science Advances 10(28).

Identifier: DOI:10.1126/sciadv.adn5290.

Direct URL:
https://doi.org/10.1126/sciadv.adn5290

Origin and locator: UG-T03-S01, Additional references and access gaps, first table.

Bibliographic record: Unverified in current pass; publisher/DOI access failed.

Content inspected: Not inspected. Historical source/access-gap wording retained.

Relevance: Potential cross-output similarity warning; original domain is short fiction. Related UG-T03-E09.

Limits: No findings imported from unavailable full text; no TD transfer presumed.

#### UG-T03-S90. Large language models are homogeneously creative

Emily Wenger; Yoed N. Kenett. 2026. Published work, supplied bibliographic record. PNAS Nexus 5(3).

Identifier: DOI:10.1093/pnasnexus/pgag042.

Direct URL:
https://doi.org/10.1093/pnasnexus/pgag042

Origin and locator: UG-T03-S01, Additional references and access gaps, first table.

Bibliographic record: Unverified in current pass; publisher/DOI access failed.

Content inspected: Not inspected. Historical source/access-gap wording retained.

Relevance: Cross-model similarity lead; provider diversity need not be behavioral diversity. Related UG-T03-E09.

Limits: No findings imported from unavailable full text; no TD transfer presumed.

#### UG-T03-S91. Superhuman artificial intelligence can improve human decision-making by increasing novelty

Minkyu Shin; Jin Kim; Bas van Opheusden; Thomas L. Griffiths. 2023. Published work, supplied bibliographic record. PNAS 120(12).

Identifier: DOI:10.1073/pnas.2214840120.

Direct URL:
https://doi.org/10.1073/pnas.2214840120

Origin and locator: UG-T03-S01, Additional references and access gaps, first table.

Bibliographic record: Unverified in current pass; publisher/DOI access failed.

Content inspected: Not inspected. Historical source/access-gap wording retained.

Relevance: AI-influenced play within existing rules, not new-unit invention. Related UG-T03-E09.

Limits: No findings imported from unavailable full text; no TD transfer presumed.

#### UG-T03-S92. Creativity and artificial intelligence

Margaret A. Boden. 1998. Published work, supplied bibliographic record. Artificial Intelligence 103(1 to 2).

Identifier: DOI:10.1016/S0004-3702(98)00055-1.

Direct URL:
https://doi.org/10.1016/S0004-3702(98)00055-1

Origin and locator: UG-T03-S01, Additional references and access gaps, first table.

Bibliographic record: Unverified in current pass; publisher/DOI access failed.

Content inspected: Not inspected. Historical source/access-gap wording retained.

Relevance: Conceptual vocabulary for creativity, not empirical LLM performance. Related UG-T03-E09.

Limits: No findings imported from unavailable full text; no TD transfer presumed.

#### UG-T03-S93. Program Synthesis as a Generative Method

Eric Butler; Kristin Siu; Alexander Zook. 2017. Published conference paper. FDG 2017, article 6.

Identifier: DOI:10.1145/3102071.3102076.

Direct URLs:
https://doi.org/10.1145/3102071.3102076
https://www.ericbutler.net/publications

Origin and locator: UG-T03-S01, additional lead via R1; visible browsing trace.

Bibliographic record: Verified title/DOI via primary index and author listing; full paper unavailable.

Content inspected: Metadata only. Publisher record search and author bibliographic route; article returned 403.

Relevance: Dynamic behavior generation, UG-T03-D05.

Limits: No result claim made from reference-list presence. Original file did not supply DOI; recovered now.

#### UG-T03-S94. Talakat: Bullet Hell Generation through Constrained Map-Elites

Ahmed Khalifa; Scott Lee; Andy Nealen; Julian Togelius. 2018. Published conference paper; preprint consulted. GECCO 2018.

Identifier: arXiv:1806.04718; DOI:10.1145/3205455.3205470.

Direct URLs:
https://arxiv.org/abs/1806.04718
https://www.akhalifa.com/documents/talakat-bullet-hell.pdf

Origin and locator: UG-T03-S01, additional lead via R4 reference 18.

Bibliographic record: Verified arXiv and publisher record.

Content inspected: Abstract only. DSL/search/simulation task scope.

Relevance: Constrained combat-pattern diversity; UG-T03-E09.

Limits: Bullet-hell level generation, not open character mechanics. Recovered URL does not make it new literature for this thread.

#### UG-T03-S95. Mapping Hearthstone Deck Spaces through MAP-Elites with Sliding Boundaries

Matthew C. Fontaine; Scott Lee; L. B. Soros; Fernando De Mesentier Silva; Julian Togelius; Amy K. Hoover. 2019. Published conference paper; preprint consulted. GECCO 2019.

Identifier: arXiv:1904.10656; DOI:10.1145/3321707.3321794.

Direct URL:
https://arxiv.org/abs/1904.10656

Origin and locator: UG-T03-S01, additional lead via R4 reference 13.

Bibliographic record: Verified arXiv and publisher identity.

Content inspected: Abstract only. Sliding-boundary archive and deck-evaluation task scope.

Relevance: Roster/deck interaction and descriptors; UG-T03-E09.

Limits: Deck selection/rebalancing, not new mechanic creation.

Record notes: Original attachment says "Michael C. Fontaine" and "with MAP-Elites". Primary record says "Matthew C. Fontaine" and "through MAP-Elites". This is a newly checked bibliographic correction, not a change to the historical citation.

#### UG-T03-S96. Multiobjective Exploration of the StarCraft Map Space

Julian Togelius et al., per supplied lead. 2010. Conference paper lead. CIG 2010.

Identifier: No identifier recovered.

Direct URL: Not available or not applicable; no URL inferred.

Origin and locator: UG-T03-S01, additional lead via R1 references.

Bibliographic record: Unverified in current pass.

Content inspected: Not inspected. Supplied citation context only; targeted search did not recover a usable primary record.

Relevance: Conflicting map objectives, later evaluation context.

Limits: No guessed URL, author completion or empirical claim. Parent recovery path: UG-T03-S50.

#### UG-T03-S97. Mixed-Initiative Co-Creativity

Georgios N. Yannakakis; Antonios Liapis; Constantine Alexopoulos. 2014. Conference paper. FDG 2014.

Identifier: FDG 2014 paper 37.

Direct URL:
https://www.fdg2014.org/papers/fdg2014_paper_37.pdf

Origin and locator: UG-T03-S01, additional lead via R3 reference 68.

Bibliographic record: Verified title/source route in primary search record only.

Content inspected: Metadata only. Proceedings PDF indexed; direct fetch failed.

Relevance: Designer control and revision dialogue; UG-T03-D09.

Limits: No detailed method or effect claim based on indexed snippet.

## B. Newly identified literature leads

Every entry below is newly found, not part of the original conversation. These are leads for later work, not approved dependencies.

#### UG-T03-S101. Is This LLM Library Learning? Evaluation Must Account For Compute and Behaviour

Ian Berlot-Attwell; Tobias Sesterhenn; Frank Rudzicz; Xujie Si. 2026. Published conference paper. EACL 2026, long papers, pages 3534 to 3568.

Identifier: DOI:10.18653/v1/2026.eacl-long.163.

Direct URLs:
https://aclanthology.org/2026.eacl-long.163/
https://aclanthology.org/2026.eacl-long.163.pdf

Origin and locator: Newly found, not part of the original conversation.

Bibliographic record: Verified publisher record and first page.

Content inspected: Relevant sections. Abstract and sections on reuse and compute correction; PDF pages 1 and 5 visually inspected.

Relevance: Challenges attribution of improvement to libraries; UG-T03-E11 and UG-T03-Q11.

Limits: Three examined ICL library-learning systems do not consistently beat cost-corrected baselines; mathematical/reasoning tasks, not proof a reviewed TD mechanics library has no value.

Associated artifacts: https://github.com/ikb-a/llm_lib_learning_fails

Record notes: Distinct from original 2024 workshop paper UG-T03-S25; do not retroactively replace it.

#### UG-T03-S102. babble: Learning Better Abstractions with E-Graphs and Anti-Unification

David Cao; Rose Kunkel; Chandrakana Nandi; Max Willsey; Zachary Tatlock; Nadia Polikarpova. 2022 preprint; 2023 publication. Published conference journal article. Proceedings of the ACM on Programming Languages 7, POPL, article 14.

Identifier: arXiv:2212.04596v1; DOI:10.1145/3571207.

Direct URLs:
https://arxiv.org/abs/2212.04596
https://arxiv.org/html/2212.04596
https://doi.org/10.1145/3571207

Origin and locator: Newly found, not part of the original conversation.

Bibliographic record: Verified primary record and text.

Content inspected: Relevant sections. Introduction, overview and library-learning-modulo-theory formulation.

Relevance: Abstraction beyond superficial syntactic matches; UG-T03-E12 and UG-T03-Q11.

Limits: Requires a supplied equational theory; compression is not fidelity or design quality. Does not discover correct cross-franchise interaction semantics automatically.

Associated artifacts: Artifact record: https://doi.org/10.5281/zenodo.7120897 ; not downloaded.

#### UG-T03-S103. Large Language Models as Tool Makers

Tianle Cai; Xuezhi Wang; Tengyu Ma; Xinyun Chen; Denny Zhou. 2023 preprint; 2024 publication. Published conference paper; preprint consulted. ICLR 2024.

Identifier: arXiv:2305.17126.

Direct URLs:
https://arxiv.org/abs/2305.17126
https://arxiv.org/html/2305.17126

Origin and locator: Newly found, not part of the original conversation.

Bibliographic record: Verified primary record/text.

Content inspected: Relevant sections. Tool-making, verification/application loop and task scope.

Relevance: Generated and cached reusable tools with a separate tool user; UG-T03-E13 and UG-T03-D10.

Limits: Reasoning-task tools are not arbitrary engine extensions. Any transfer should account for generation, repair and cache-construction costs.

Associated artifacts: https://github.com/ctlllll/LLM-ToolMaker ; not executed.

## Literature gaps and stated omissions

Later review should target faithful narrative-to-mechanics translation, specification of previously unsupported real-time mechanics, formal interaction semantics, abstraction reuse under changing Definitions, actual engineering effort saved, and evaluations against strong equal-budget direct prompting. The original long catalogue also spans map/asset generation and entire-game tools; those references are retained, not promoted into v1 scope.

No exhaustive literature-review claim is made. Unavailable full texts, incomplete bibliographic records, author lists still abbreviated as "et al.", mutable URLs, opaque historical citation targets and missing repository-local documents remain explicit gaps. Generic unsuccessful search queries and duplicated retrievals are not individual sources. Papers merely appearing in newly opened papers' reference lists have not been imported into the original source set.
