# UG-T04 source register

Extracted from the full internal handover. Source IDs are shared with that handover.

## 7. Sources and research leads

### Register conventions and coverage

This register contains **105 records: 102 already present references or explicit source gaps, and three newly found papers**. Records are grouped by provenance, not by perceived importance. UG-T04-S101 to UG-T04-S138 map exactly to the supplied catalogue's R1 to R38. A source that appeared only in a research-tool request is not treated as an adopted method or as historically read.

Every entry separates bibliographic verification from content inspected for this handover. Verification may cover only the title/authors/identifier, with venue or other details still attributed to the supplied catalogue. "Abstract only" does not support a new detailed empirical claim. Earlier "Inspected" statements are preserved as inherited access reports, not promoted to fresh full-text checks.

Checks were made for this handover on September 24, 2026. Mutable model cards, documentation, repository branches and publisher pages are not pinned historical snapshots unless an exact version is stated. No paper experiment, external code, training run or vendor benchmark was reproduced. No code, weights, data or asset license was audited.

Direct URLs are given below as ordinary text. Local paths and unresolved titles are retained where no recoverable URL exists. Several records remain incomplete. The register does not silently replace them with a different paper.

### A. Sources already present in this thread

#### Supplied documents, direct references and sources from visible research requests


##### UG-T04-S01: Research sources for unit generation and game design

**Type and publication status:** Supplied research catalogue. Internal curated research catalogue; file says research checked September 22, 2026
**Authors/organization:** Author not identified in the supplied file
**Persistent identifier or locator:** Uploaded GAME-DESIGN-RESEARCH.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** User-supplied attachment; all 38 numbered entries and both additional-reference tables reviewed.
**Bibliographic verification:** Verified as the supplied document, not as independent verification of its bibliography.
**Content inspected for this handover:** Full supplied file.

**Relevance and limits:** Contains representation, search, evaluation, taste, maps and engine-research precedents. R1 to R38 map to S101 to S138; additional references map to S139 to S148. The document explicitly disclaims systematic-review status, reproduction and implementation commitment. Its historical 'Inspected' labels belong to its author.
**Associated artifacts:** SHA-256: 8b1611fa74f65009d3c71644220f74880a46d95457302888ce3bd02e4e7c27cc
**Linked records:** UG-T04-E04, UG-T04-E11, UG-T04-E12, UG-T04-E19


##### UG-T04-S02: Unit Generator context

**Type and publication status:** Supplied contributor/context document. Internal contributor contract; no version/date supplied
**Authors/organization:** Author not identified in the supplied file
**Persistent identifier or locator:** Uploaded CONTEXT.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** User-supplied attachment; Shared terms, Unit terms, Generation and evaluation terms, and closing implementation caveat.
**Bibliographic verification:** Verified as the supplied document; underlying repositories not verified.
**Content inspected for this handover:** Full supplied file.

**Relevance and limits:** Defines Unit, Mechanic, Interpretation, Game Definition, Profile, Review and Validation. Separates the Engine from the combat runtime. Responsibility statements are historical document content, not a demonstrated implementation or fresh owner approval. Private Foundation guidance was not available.
**Associated artifacts:** SHA-256: fee3f4ebbcaf0f7c44013db8102ca98510676e8723a399e95ee323f7a946ee50
**Linked records:** UG-T04-E11, UG-T04-E19, UG-T04-Q03, UG-T04-Q09


##### UG-T04-S03: caveman SKILL.md

**Type and publication status:** Repository guidance file, stylistic instruction. Mutable repository guidance, not research
**Authors/organization:** Repository account JuliusBrussee; individual authorship not independently established
**Persistent identifier or locator:** main branch; no commit pinned

**Direct URLs**
https://raw.githubusercontent.com/JuliusBrussee/caveman/refs/heads/main/skills/caveman/SKILL.md

**Where and why it appeared:** Initial user request appended "LITE".
**Bibliographic verification:** Repository file retrieved; no historical commit verified.
**Content inspected for this handover:** Full current text.

**Relevance and limits:** Communication-style reference only. Not an accepted Unit-Generator architecture, model skill, product requirement or evidence of design quality.
**Associated artifacts:** No code or dependency adopted.
**Linked records:** UG-T04-E05; UG-T04-E15


##### UG-T04-S04: Laya model card and checkpoint

**Type and publication status:** Model card, weights reference and vendor benchmark. Mutable model card and vendor-reported benchmarks; not an independently reproduced paper
**Authors/organization:** Convai Innovations / convaiinnovations
**Persistent identifier or locator:** convaiinnovations/laya; checkpoint revision not pinned

**Direct URLs**
https://huggingface.co/convaiinnovations/laya
https://huggingface.co/convaiinnovations/laya/raw/main/README.md

**Where and why it appeared:** Initial user model example; first assistant answer, 'Where Laya fits'.
**Bibliographic verification:** Official model-card identity verified.
**Content inspected for this handover:** Architecture, checkpoint table, typed-decisions benchmark and Honest Limits sections.

**Relevance and limits:** The card describes a 421M encoder-plus-head root model, a 512-token question budget and scoring of supplied options. Reports 0.362 root accuracy, 0.461 majority baseline and 0.766 specialized-checkpoint accuracy on typed-decisions. Vendor-reported task-specific results, not TD design or taste results. Different checkpoints must not be conflated; calibration and serving claims were not reproduced.
**Associated artifacts:** Associated repository https://github.com/NandhaKishorM/laya and benchmark/training notebook links are listed by the card. None executed or license-audited.
**Linked records:** UG-T04-E13, UG-T04-D07


##### UG-T04-S05: From LLM-Driven Trading Card Generation to Procedural Relatedness: A Pokémon Case Study

**Type and publication status:** Research paper. 2026 preprint; submitted April 30, 2026
**Authors/organization:** Johannes Pfau; Panagiotis Vrettis
**Persistent identifier or locator:** arXiv:2604.27972v1

**Direct URLs**
https://arxiv.org/abs/2604.27972
https://arxiv.org/html/2604.27972v1

**Where and why it appeared:** First assistant answer, missing-paper entry A.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Relevant system and study sections.

**Relevance and limits:** Description-conditioned card generation and perceived representativeness. Reports Qwen3-14B, retrieval over 993 cards and a study with 49 participants producing 196 cards. Natural-language effects in structured output do not establish executable rules, balance or transfer to anime TD progression.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05, UG-T04-E15


##### UG-T04-S06: Open Role-Playing with Delta-Engines

**Type and publication status:** Research paper. 2024 preprint; inspected revision v5 dated March 7, 2025
**Authors/organization:** Hongqiu Wu; Zekai Xu; Tianyang Xu; Shize Wei; Yan Wang; Jiale Hong; Weiqi Wu; Hai Zhao
**Persistent identifier or locator:** arXiv:2408.05842v5

**Direct URLs**
https://arxiv.org/abs/2408.05842
https://arxiv.org/html/2408.05842v5

**Where and why it appeared:** First assistant answer, missing-paper entry B.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Relevant methods and Section 6.1, including HTML Table 1; PDF retrieval failed.

**Relevance and limits:** Description-to-character-code adaptation with retrieval and LoRA-trained CodeGemma-7B. Co-designed training set has 175 roles and 502 examples. Code accuracy is GPT-4 judged only on executable outputs. Engine patching and battle execution are not taste or balance validation under a fixed TD Definition.
**Associated artifacts:** CodeGemma model is named in the paper. Neither training nor engine was run here.
**Linked records:** UG-T04-E06, UG-T04-D07, UG-T04-D08


##### UG-T04-S07: Game-O-Matic: Generating Videogames that Represent Ideas

**Type and publication status:** Research paper. 2012, Third Workshop on Procedural Content Generation in Games; published workshop paper
**Authors/organization:** Mike Treanor; Bryan Blackford; Michael Mateas; Ian Bogost
**Persistent identifier or locator:** DOI:10.1145/2538528.2538537

**Direct URLs**
https://mtreanor.com/publications/Game-O-Matic-PCG2012.pdf
https://dl.acm.org/doi/abs/10.1145/2538528.2538537

**Where and why it appeared:** First assistant answer, entry C; earlier visible DOI/PDF tool request resolves to this same work.
**Bibliographic verification:** Title/authors verified in author PDF; publisher search record corroborates year and DOI. Publisher article fetch failed.
**Content inspected for this handover:** Relevant representation and assembly sections; PDF first page visually checked.

**Relevance and limits:** Concept graphs map relationships to human-authored gameplay patterns. Relevant to an explicit Interpretation rather than a bag of named powers. It does not establish learned character fidelity or that players identify the intended meaning reliably.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05, UG-T04-D04


##### UG-T04-S08: Chaos Cards: Creating Novel Digital Card Games through Grammatical Content Generation and Meta-Based Card Evaluation

**Type and publication status:** Research paper. AIIDE 2020, published conference paper
**Authors/organization:** Tiannan Chen; Stephen J. Guy
**Persistent identifier or locator:** AAAI OJS record 7430

**Direct URLs**
https://cdn.aaai.org/ojs/7430/7430-52-10748-1-2-20200923.pdf

**Where and why it appeared:** First assistant answer, entry D.
**Bibliographic verification:** Title/authors/venue/year verified in primary proceedings PDF.
**Content inspected for this handover:** Relevant abstract and generation/evaluation description; first page visually checked.

**Relevance and limits:** Grammar-generated executable effects and a separate neural strength predictor illustrate different roles for generation and evaluation. A strength predictor is not a character-adaptation generator or a human-taste model. No Unit-Generator reuse was tested.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05, UG-T04-D08


##### UG-T04-S09: Latent Predictor Networks for Code Generation

**Type and publication status:** Research paper. ACL 2016, long papers, pages 599 to 609
**Authors/organization:** Wang Ling; Phil Blunsom; Edward Grefenstette; Karl Moritz Hermann; Tomáš Kočiský; Fumin Wang; Andrew Senior
**Persistent identifier or locator:** DOI:10.18653/v1/P16-1057; arXiv:1603.06744

**Direct URLs**
https://aclanthology.org/P16-1057/
https://arxiv.org/html/1603.06744

**Where and why it appeared:** First assistant answer, entry E; original visible research calls.
**Bibliographic verification:** Verified against ACL Anthology.
**Content inspected for this handover:** Publisher abstract and metadata in this handover; previous assistant reported fuller reading.

**Relevance and limits:** Maps natural-language card descriptions plus structured attributes to implementation code. The input already specifies mechanics. It does not decide which mechanics should represent an unrelated fictional character.
**Associated artifacts:** Magic/Hearthstone code datasets are associated with the task; exact data snapshot, rights and reuse not inspected.
**Linked records:** UG-T04-E05, UG-T04-D08


##### UG-T04-S10: Mystical Tutor: A Magic: The Gathering Design Assistant via Denoising Sequence-to-Sequence Learning

**Type and publication status:** Research paper. AIIDE 12(1), pages 86 to 92, 2016; published
**Authors/organization:** Adam Summerville; Michael Mateas
**Persistent identifier or locator:** DOI:10.1609/aiide.v12i1.12851

**Direct URLs**
https://ojs.aaai.org/index.php/AIIDE/article/view/12851

**Where and why it appeared:** First assistant answer, entry F.
**Bibliographic verification:** Verified against publisher record.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Denoising sequence-to-sequence completion of partially specified cards; relevant to preserving fixed designer choices. Earlier and current extraction do not establish a full-paper review. No evidence here about source-grounded TD taste.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05, UG-T04-D04


##### UG-T04-S11: PICARD: Parsing Incrementally for Constrained Auto-Regressive Decoding from Language Models

**Type and publication status:** Research paper. EMNLP 2021, pages 9895 to 9901; published
**Authors/organization:** Torsten Scholak; Nathan Schucher; Dzmitry Bahdanau
**Persistent identifier or locator:** DOI:10.18653/v1/2021.emnlp-main.779

**Direct URLs**
https://aclanthology.org/2021.emnlp-main.779/

**Where and why it appeared:** First assistant answer, constrained structure.
**Bibliographic verification:** Verified against ACL Anthology.
**Content inspected for this handover:** Abstract and bibliographic record.

**Relevance and limits:** Incremental parsing rejects inadmissible continuations in text-to-SQL generation. Grammar admissibility is not game semantics, preference or runtime correctness.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D08, UG-T04-E11


##### UG-T04-S12: XGrammar-2: Dynamic and Efficient Structured Generation Engine for Agentic LLMs

**Type and publication status:** Research paper. 2026 preprint; other publication status not established here
**Authors/organization:** Linzhang Li; Yixin Dong; Guanjie Wang; Ziyi Xu; Alexander Jiang; Tianqi Chen
**Persistent identifier or locator:** arXiv:2601.04426v1

**Direct URLs**
https://arxiv.org/abs/2601.04426
https://arxiv.org/html/2601.04426v1

**Where and why it appeared:** First assistant answer and original visible research call.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Relevant introduction and structured-generation scope.

**Relevance and limits:** Reference for dynamic structured-output constraints. Not a game-rule validator or a demonstrated source of good design decisions.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D08


##### UG-T04-S13: Qwen3-14B model card and checkpoint

**Type and publication status:** Model card and checkpoint. Official model artifact; mutable card
**Authors/organization:** Qwen team
**Persistent identifier or locator:** Qwen/Qwen3-14B

**Direct URLs**
https://huggingface.co/Qwen/Qwen3-14B

**Where and why it appeared:** First assistant answer, proposed reference-model comparison.
**Bibliographic verification:** Official card/checkpoint identity verified.
**Content inspected for this handover:** Model description and usage metadata.

**Relevance and limits:** Suggested pretrained generator candidate, not a chosen backend. No TD generation, serving cost, hardware fit or preference performance was tested. A model-card license tag does not establish overall reuse compatibility.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07, UG-T04-Q05


##### UG-T04-S14: Qwen3-8B model card and checkpoint

**Type and publication status:** Model card and checkpoint. Official model artifact; mutable card
**Authors/organization:** Qwen team
**Persistent identifier or locator:** Qwen/Qwen3-8B

**Direct URLs**
https://huggingface.co/Qwen/Qwen3-8B

**Where and why it appeared:** First assistant answer, proposed reference-model comparison.
**Bibliographic verification:** Official card/checkpoint identity verified.
**Content inspected for this handover:** Model description and usage metadata.

**Relevance and limits:** Suggested pretrained generator candidate, not a chosen backend. No TD generation, serving cost, hardware fit or preference performance was tested. A model-card license tag does not establish overall reuse compatibility.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07, UG-T04-Q05


##### UG-T04-S15: LoRA: Low-Rank Adaptation of Large Language Models

**Type and publication status:** Research paper. 2021 preprint; ICLR 2022 publication identified in the inspected Delta-Engines reference list
**Authors/organization:** Edward J. Hu; Yelong Shen; Phillip Wallis; Zeyuan Allen-Zhu; Yuanzhi Li; Shean Wang; Lu Wang; Weizhu Chen
**Persistent identifier or locator:** arXiv:2106.09685v2; OpenReview:nZeVKeeFYf9

**Direct URLs**
https://arxiv.org/abs/2106.09685
https://openreview.net/forum?id=nZeVKeeFYf9

**Where and why it appeared:** First assistant answer, proposed SFT method.
**Bibliographic verification:** arXiv bibliography verified; ICLR record corroborated by a primary-paper reference, not freshly opened.
**Content inspected for this handover:** Abstract; citation context in Delta-Engines.

**Relevance and limits:** Parameter-efficient adaptation of pretrained models. A training method does not provide the target taste dataset or establish sufficient sample size.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07, UG-T04-E14


##### UG-T04-S16: QLoRA: Efficient Finetuning of Quantized LLMs

**Type and publication status:** Research paper. 2023 preprint inspected; NeurIPS 2023 attribution not independently rechecked in publisher proceedings
**Authors/organization:** Tim Dettmers; Artidoro Pagnoni; Ari Holtzman; Luke Zettlemoyer
**Persistent identifier or locator:** arXiv:2305.14314v1

**Direct URLs**
https://arxiv.org/abs/2305.14314

**Where and why it appeared:** First assistant answer, lower-memory adapter training.
**Bibliographic verification:** arXiv record verified.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Training adapters through a frozen quantized model can reduce training memory. Published benchmark results are not hardware estimates or taste results for Unit-Generator.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07, UG-T04-E14


##### UG-T04-S17: SFT Trainer documentation

**Type and publication status:** Technical documentation. Live technical documentation; no package version pinned
**Authors/organization:** Hugging Face TRL maintainers
**Persistent identifier or locator:** TRL SFTTrainer

**Direct URLs**
https://huggingface.co/docs/trl/sft_trainer

**Where and why it appeared:** First assistant answer, implementation lead for supervised fine-tuning.
**Bibliographic verification:** Official documentation verified.
**Content inspected for this handover:** Relevant SFT and adapter-configuration guidance.

**Relevance and limits:** Possible implementation route using TRL and PEFT. Neither a dependency decision nor a run. API compatibility and exact version remain unspecified.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S18: Direct Preference Optimization: Your Language Model is Secretly a Reward Model

**Type and publication status:** Research paper. 2023 preprint v2 inspected; NeurIPS 2023 attribution from the conversation retained, proceedings not independently rechecked
**Authors/organization:** Rafael Rafailov; Archit Sharma; Eric Mitchell; Stefano Ermon; Christopher D. Manning; Chelsea Finn
**Persistent identifier or locator:** arXiv:2305.18290v2

**Direct URLs**
https://arxiv.org/abs/2305.18290
https://arxiv.org/html/2305.18290v2

**Where and why it appeared:** Both assistant answers, later preference-training option.
**Bibliographic verification:** arXiv bibliography verified.
**Content inspected for this handover:** Relevant objective and preference-data formulation.

**Relevance and limits:** Learns from preferred/rejected outputs for the same input without a separate online RL loop. Does not require Elo scores. No TD training or taste generalization demonstrated. Common-request supervision and preference assumptions matter; see new S302.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07, UG-T04-E07, UG-T04-Q10


##### UG-T04-S19: Distilling System 2 into System 1

**Type and publication status:** Research paper. 2024 preprint; final venue not established in this extraction
**Authors/organization:** Ping Yu; Jing Xu; Jason Weston; Ilia Kulikov
**Persistent identifier or locator:** arXiv:2407.06023

**Direct URLs**
https://arxiv.org/abs/2407.06023
https://arxiv.org/html/2407.06023

**Where and why it appeared:** First assistant answer, fast specialist through distillation.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Relevant method and limitations.

**Relevance and limits:** Train on filtered results of more expensive inference as a later efficiency route. Task-dependent evidence, not a guarantee that a small model will reproduce a TD design process.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S20: LILO: Learning Interpretable Libraries by Compressing and Documenting Code

**Type and publication status:** Research paper. 2023 preprint record inspected; final venue not established here
**Authors/organization:** Gabriel Grand; Lionel Wong; Maddy Bowers; Theo X. Olausson; Muxin Liu; Joshua B. Tenenbaum; Jacob Andreas
**Persistent identifier or locator:** arXiv:2310.19791

**Direct URLs**
https://arxiv.org/abs/2310.19791

**Where and why it appeared:** First assistant answer, longer-term reusable mechanic abstractions.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract and metadata.

**Relevance and limits:** A lead for finding reusable program abstractions. No automatic permission to add runtime operators; not needed for the proposed v1.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-Q14


##### UG-T04-S21: Circular 14: Copyright in Derivative Works and Compilations

**Type and publication status:** Official copyright guidance. Government guidance, displayed edition 14.0720; not research
**Authors/organization:** U.S. Copyright Office
**Persistent identifier or locator:** Circular 14

**Direct URLs**
https://www.copyright.gov/circs/circ14.pdf

**Where and why it appeared:** First assistant answer, separate rights review.
**Bibliographic verification:** Official document verified.
**Content inspected for this handover:** Relevant first two pages; first page visually checked.

**Relevance and limits:** Historical reminder that generation does not supply a rights clearance. U.S.-specific and not an assessment of this project's rights. The circular itself says registration policies/practices were superseded by the Compendium. No license choice follows.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-Q13


##### UG-T04-S22: Experience-Driven Procedural Content Generation (Extended Abstract)

**Type and publication status:** Research paper. ACII 2015 invited paper, pages 519 to 525; published
**Authors/organization:** Georgios N. Yannakakis; Julian Togelius
**Persistent identifier or locator:** DOI:10.1109/ACII.2015.7344619

**Direct URLs**
https://yannakakis.net/wp-content/uploads/2015/11/PID3821875.pdf
https://ieeexplore.ieee.org/document/7344619/

**Where and why it appeared:** Second assistant answer, experience-driven PCG framing.
**Bibliographic verification:** Author PDF title/authors and IEEE proceedings metadata verified.
**Content inspected for this handover:** Relevant player-modeling text; first page visually checked.

**Relevance and limits:** Frames content generation around a stated player or experience, including preference learning. Distinct from the 2011 article at S139. Neither is a universal taste metric or a test of this TD task.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E07, UG-T04-Q01


##### UG-T04-S23: Fine-Grained Human Feedback Gives Better Rewards for Language Model Training

**Type and publication status:** Research paper. NeurIPS 2023, published; arXiv version inspected
**Authors/organization:** Zeqiu Wu; Yushi Hu; Weijia Shi; Nouha Dziri; Alane Suhr; Prithviraj Ammanabrolu; Noah A. Smith; Mari Ostendorf; Hannaneh Hajishirzi
**Persistent identifier or locator:** arXiv:2306.01693

**Direct URLs**
https://arxiv.org/abs/2306.01693
https://arxiv.org/html/2306.01693

**Where and why it appeared:** Second assistant answer, feedback reasons and aspect-level judgments.
**Bibliographic verification:** arXiv verified; proceedings record identified.
**Content inspected for this handover:** Relevant feedback representation and method sections.

**Relevance and limits:** Local/category feedback can supply information absent from one whole-output preference. Language-task evidence, not a validated TD rubric; granular labels do not by themselves remove presentation bias.
**Associated artifacts:** Associated code https://github.com/allenai/FineGrainedRLHF identified, not executed.
**Linked records:** UG-T04-D05, UG-T04-E07


##### UG-T04-S24: Chatbot Arena: An Open Platform for Evaluating LLMs by Human Preference

**Type and publication status:** Research paper. 2024 arXiv record inspected; final venue not independently established here
**Authors/organization:** Wei-Lin Chiang; Lianmin Zheng; Ying Sheng; Anastasios N. Angelopoulos; Tianle Li; Dacheng Li; Banghua Zhu; Hao Zhang; Michael I. Jordan; Joseph E. Gonzalez; Ion Stoica
**Persistent identifier or locator:** arXiv:2403.04132v1

**Direct URLs**
https://arxiv.org/abs/2403.04132
https://arxiv.org/html/2403.04132v1

**Where and why it appeared:** Second assistant answer, pairwise ranking and uncertainty.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Sections 3 to 5 and relevant ranking discussion.

**Relevance and limits:** Anonymous comparisons, Bradley-Terry estimation and uncertainty support the proposed review/ranking distinction. An estimated candidate/model ranking is not a model of unseen unit content. Its inference assumptions and comparison population must be stated.
**Associated artifacts:** Platform described in paper; no platform code, data or statistical implementation imported.
**Linked records:** UG-T04-D05, UG-T04-D06, UG-T04-E08


##### UG-T04-S25: Elo Uncovered: Robustness and Best Practices in Language Model Evaluation

**Type and publication status:** Research paper. GEM 2023 workshop, pages 339 to 352; later NeurIPS 2024 record also surfaced, not treated as the same inspected version
**Authors/organization:** Meriem Boubdir; Edward Kim; Beyza Ermis; Sara Hooker; Marzieh Fadaee
**Persistent identifier or locator:** arXiv:2311.17295v1; ACL:2023.gem-1.28

**Direct URLs**
https://arxiv.org/html/2311.17295v1
https://aclanthology.org/2023.gem-1.28/

**Where and why it appeared:** Second assistant answer, alternative to relying on sequential Elo.
**Bibliographic verification:** arXiv metadata and ACL proceedings record verified.
**Content inspected for this handover:** Relevant ordering, hyperparameter and comparison experiments in v1.

**Relevance and limits:** Documents ranking sensitivity to comparison order and hyperparameters. Does not establish that a particular Bradley-Terry implementation is best for this group or that ranking differences equal design-quality differences.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D06, UG-T04-E08


##### UG-T04-S26: SynthesizeMe! Inducing Persona-Guided Prompts for Personalized Reward Models in LLMs

**Type and publication status:** Research paper. ACL 2025, pages 8045 to 8078; published
**Authors/organization:** Michael J. Ryan; Omar Shaikh; Aditri Bhagirath; Daniel Frees; William Held; Diyi Yang
**Persistent identifier or locator:** DOI:10.18653/v1/2025.acl-long.397

**Direct URLs**
https://aclanthology.org/2025.acl-long.397/
https://aclanthology.org/2025.acl-long.397.pdf

**Where and why it appeared:** Second assistant answer, preference history and examples before weight training.
**Bibliographic verification:** Verified against ACL Anthology.
**Content inspected for this handover:** Relevant methods, Sections 3.1 and 3.2; corresponding PDF page visually checked.

**Relevance and limits:** Personalized prompt/example conditioning for predicting preferences, relevant to editable design-guidance experiments. Chatbot preferences are not TD design preferences. Broad demographic/persona inference was not required by the thread's proposed use.
**Associated artifacts:** PersonalRewardBench is associated with the work; dataset files and reuse terms not inspected.
**Linked records:** UG-T04-D07, UG-T04-E09


##### UG-T04-S27: Comparison-based Active Preference Learning for Multi-dimensional Personalization

**Type and publication status:** Research paper. ACL 2025, pages 33145 to 33166; published
**Authors/organization:** Minhyeon Oh; Seungjoon Lee; Jungseul Ok
**Persistent identifier or locator:** DOI:10.18653/v1/2025.acl-long.1590

**Direct URLs**
https://aclanthology.org/2025.acl-long.1590/
https://aclanthology.org/2025.acl-long.1590.pdf

**Where and why it appeared:** Second assistant answer, AMPLe and informative comparisons.
**Bibliographic verification:** Verified against ACL Anthology.
**Content inspected for this handover:** Sections 3.1 and 3.2; corresponding PDF page visually checked.

**Relevance and limits:** Active comparisons can estimate priorities over a provided multidimensional reward representation. The method assumes supplied attributes and a preference model. It does not discover the correct unit-design dimensions or remove the need for representative reviewers.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E09, UG-T04-Q06


##### UG-T04-S28: Scaling Laws for Reward Model Overoptimization

**Type and publication status:** Research paper. 2022 preprint record inspected; final publication venue not independently rechecked here
**Authors/organization:** Leo Gao; John Schulman; Jacob Hilton
**Persistent identifier or locator:** arXiv:2210.10760

**Direct URLs**
https://arxiv.org/abs/2210.10760
https://arxiv.org/html/2210.10760

**Where and why it appeared:** Second assistant answer, learned-score exploitation.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Relevant experimental setup and findings.

**Relevance and limits:** Shows proxy-reward overoptimization under RL and best-of-N selection in a synthetic gold-reward setup. Not a measurement of this project's risk or real player enjoyment.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E10, UG-T04-Q10


##### UG-T04-S29: Qwen3.5-9B model card and checkpoint

**Type and publication status:** Model card and checkpoint. Official model artifact; mutable
**Authors/organization:** Qwen team
**Persistent identifier or locator:** Qwen/Qwen3.5-9B

**Direct URLs**
https://huggingface.co/Qwen/Qwen3.5-9B

**Where and why it appeared:** Visible original research tool calls; not selected in the final model recommendation.
**Bibliographic verification:** Official card identity verified.
**Content inspected for this handover:** Metadata/model description only.

**Relevance and limits:** Alternative model considered during research. No adoption, training or task-specific comparison established.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S30: XGrammar: Flexible and Efficient Structured Generation Engine for Large Language Models

**Type and publication status:** Research paper. 2024 preprint record; later venue not rechecked
**Authors/organization:** Yixin Dong; Charlie F. Ruan; Yaxing Cai; Ruihang Lai; Ziyi Xu; Yilong Zhao; Tianqi Chen
**Persistent identifier or locator:** arXiv:2411.15100

**Direct URLs**
https://arxiv.org/abs/2411.15100

**Where and why it appeared:** Visible original structured-generation search; distinguish from XGrammar-2 at S12.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Structured-generation implementation lead. No guarantee of semantics, source fidelity or taste.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D08


##### UG-T04-S31: Consistent Game Content Creation via Function Calling for Large Language Models

**Type and publication status:** Research paper. 2024 paper; venue not independently resolved in this handover
**Authors/organization:** Roberto Gallotta; Antonios Liapis; Georgios Yannakakis
**Persistent identifier or locator:** Author-hosted PDF

**Direct URLs**
https://antoniosliapis.com/papers/consistent_game_content_creation_via_function_calling_for_large_language_models.pdf

**Where and why it appeared:** Visible original research tool requests; supporting first assistant discussion of typed calls.
**Bibliographic verification:** Title/authors/year verified in primary PDF; venue unverified.
**Content inspected for this handover:** Abstract and relevant scope; first page visually checked.

**Relevance and limits:** LLMaker uses function calls for consistency in dungeon-crawler content authoring. Not a learned TD taste system. A tool-call interface does not establish good adaptation.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D08


##### UG-T04-S32: Type-Constrained Code Generation with Language Models

**Type and publication status:** Research paper. 2025 preprint version inspected; later venue not checked
**Authors/organization:** Niels Mündler; Jingxuan He; Hao Wang; Koushik Sen; Dawn Song; Martin Vechev
**Persistent identifier or locator:** arXiv:2504.09246v1

**Direct URLs**
https://arxiv.org/abs/2504.09246v1

**Where and why it appeared:** Visible original research tool calls.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract and metadata.

**Relevance and limits:** Lead on constraints beyond syntax during code generation. No Unit-Generator type system, implementation or taste evidence follows.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D08


##### UG-T04-S33: System One

**Type and publication status:** Vendor technical documentation. Vendor technical documentation; mutable
**Authors/organization:** TypeSafe AI
**Persistent identifier or locator:** TypeSafe AI concepts/system-one

**Direct URLs**
https://docs.typesafe.ai/concepts/system-one

**Where and why it appeared:** Visible original research tool request in response to the user's System One interest.
**Bibliographic verification:** Official documentation identity verified.
**Content inspected for this handover:** Relevant product description.

**Relevance and limits:** Describes bounded decision outputs as distinct from generative text. Vendor terminology and claims, not an independently verified model class suitable for this task.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S34: Introducing System One Models & Jev

**Type and publication status:** Vendor announcement. Vendor announcement; mutable
**Authors/organization:** TypeSafe AI
**Persistent identifier or locator:** TypeSafe AI blog

**Direct URLs**
https://typesafe.ai/blog/introducing-system-one-models-and-jev

**Where and why it appeared:** Visible original research tool request.
**Bibliographic verification:** Official page identity verified.
**Content inspected for this handover:** Relevant product description.

**Relevance and limits:** Commercial decision-model reference. Not an adopted dependency; no private API accessed, benchmark reproduced or model comparison established.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S35: RPGAgent: Driving Coherent Story-to-Play [incomplete wording retained]

**Type and publication status:** Incomplete research lead. Possible CHI 2026 reference; publication status unverified
**Authors/organization:** Unresolved
**Persistent identifier or locator:** Candidate DOI:10.1145/3772318.3790326; title/DOI match unverified

**Direct URLs**
https://dl.acm.org/doi/full/10.1145/3772318.3790326
https://programs.sigchi.org/chi/2026/program/content/222237

**Where and why it appeared:** Visible original search wording and URL requests; original results are redacted.
**Bibliographic verification:** Unverified. DOI fetch failed; program URL returned only an application shell.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Potential story-to-play authoring lead. No findings should be attributed to this unresolved record.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05; UG-T04-E15


##### UG-T04-S36: Neural Networks Models for Analyzing Magic: the Gathering Cards

**Type and publication status:** Research paper. ICONIP 2018 accepted-paper status on arXiv
**Authors/organization:** Felipe Zilio; Marcelo Prates; Luis Lamb
**Persistent identifier or locator:** arXiv:1810.03744v1

**Direct URLs**
https://arxiv.org/abs/1810.03744
https://ar5iv.labs.arxiv.org/html/1810.03744

**Where and why it appeared:** Visible original ar5iv request, recovered by identifier.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Card text/image classification and text-generation matching lead. Not source-grounded TD adaptation or human preference learning.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05; UG-T04-E15


##### UG-T04-S37: Procedural Content Generation via Machine Learning (PCGML)

**Type and publication status:** Research paper. 2017 preprint, v3 May 7, 2018; IEEE Transactions on Games 10(3), 257 to 270, 2018 identified in the author-hosted PCG-QD bibliography
**Authors/organization:** Adam Summerville; Sam Snodgrass; Matthew Guzdial; Christoffer Holmgård; Amy K. Hoover; Aaron Isaksen; Andy Nealen; Julian Togelius
**Persistent identifier or locator:** arXiv:1702.00539v3

**Direct URLs**
https://arxiv.org/abs/1702.00539v3

**Where and why it appeared:** Visible original research tool request.
**Bibliographic verification:** arXiv verified; journal details corroborated in primary-paper references rather than publisher page.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Background taxonomy for learning generation, repair and critique from content. A survey, not a demonstrated complete system for this project.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05; UG-T04-E15


##### UG-T04-S38: UrzaGPT: LoRA-Tuned Large Language Models for Card Selection in Collectible Card Games

**Type and publication status:** Research paper. 2025 preprint
**Authors/organization:** Timo Bertram
**Persistent identifier or locator:** arXiv:2508.08382v1

**Direct URLs**
https://arxiv.org/abs/2508.08382v1

**Where and why it appeared:** Visible original research tool request.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Domain-adapted drafting/selection, relevant to separating choosing from generating. Selecting existing cards is not inventing character-based units.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S39: Distilling Step-by-Step! Outperforming Larger Language Models with Less Training Data and Smaller Model Sizes

**Type and publication status:** Research paper. 2023 preprint record; final venue not independently checked
**Authors/organization:** Cheng-Yu Hsieh; Chun-Liang Li; Chih-Kuan Yeh; Hootan Nakhost; Yasuhisa Fujii; Alexander Ratner; Ranjay Krishna; Chen-Yu Lee; Tomas Pfister
**Persistent identifier or locator:** arXiv:2305.02301v2

**Direct URLs**
https://arxiv.org/abs/2305.02301

**Where and why it appeared:** Visible original research tool request; distinct from S19.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Distillation using additional rationale supervision. NLP task evidence, not a demonstrated TD training approach.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-D07


##### UG-T04-S40: Progression Balancing [incomplete search lead]

**Type and publication status:** Incomplete research lead. Unresolved
**Authors/organization:** Pfau (2025), as written in original search
**Persistent identifier or locator:** No confirmed identifier

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** Visible original queries included "Progression Balancing" and "Baldur".
**Bibliographic verification:** Unverified.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Potential game-progression/balance lead. Cannot attach findings, a complete title or a DOI confidently.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05; UG-T04-E15


##### UG-T04-S41: 1001 Nights [incomplete search lead]

**Type and publication status:** Incomplete research lead. Unresolved
**Authors/organization:** Unresolved
**Persistent identifier or locator:** No confirmed identifier

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** Visible original query combined "1001 Nights" and "game mechanics".
**Bibliographic verification:** Unverified.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Potential story/mechanics reference. Search wording alone does not establish which publication was intended.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05; UG-T04-E15


##### UG-T04-S42: Author publication pages used to recover research references

**Type and publication status:** Author bibliographies. Author bibliographies; mutable
**Authors/organization:** Eric Butler; Mike Treanor
**Persistent identifier or locator:** Author sites

**Direct URLs**
https://www.ericbutler.net/publications
https://mtreanor.com

**Where and why it appeared:** Visible original research calls; Butler bibliography re-opened to resolve S144.
**Bibliographic verification:** Butler bibliography verified; Treanor paper at S07 inspected instead of rechecking the whole homepage.
**Content inspected for this handover:** Relevant bibliographic entries only.

**Relevance and limits:** Recovery route for Program Synthesis and Game-O-Matic. A bibliography is not a full-paper review.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-S07, UG-T04-S144


##### UG-T04-S43: Damage optimization in video games [incomplete search lead]

**Type and publication status:** Incomplete research lead. Unresolved
**Authors/organization:** Unresolved
**Persistent identifier or locator:** No confirmed identifier

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** Visible original research query.
**Bibliographic verification:** Unverified.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Possible numerical-balancing reference. Not merged with S40 without evidence of identity.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E05; UG-T04-E15


#### Papers and research records from GAME-DESIGN-RESEARCH.md, R1 to R38


##### UG-T04-S101: Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Kynan Sorochan and Matthew Guzdial. **Generating Real-Time Strategy Game Units Using Search-Based Procedural Content Generation and Monte Carlo Tree Search.** Experimental AI in Games Workshop, 2022. [arXiv:2212.03387v1](https://arxiv.org/html/2212.03387v1). Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.

**Direct URLs**
https://arxiv.org/html/2212.03387v1
https://arxiv.org/abs/2212.03387v1

**Where and why it appeared:** UG-T04-S01, section R1, "Generating RTS units". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified for arXiv title/authors/identifier; workshop metadata retained from the supplied record.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: Search Space, Evaluator, Search Method, generated units, evaluation and references.

**Relevance and limits:** Explicit triggers/effects and tests of new-unit contribution. Predefined search vocabulary, one-map/evaluator limitations; no character-to-mechanics learning.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-D08; UG-T04-E12


##### UG-T04-S102: Evolutionary Game Design

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Cameron Browne and Frédéric Maire. **Evolutionary Game Design.** IEEE Transactions on Computational Intelligence and AI in Games 2(1), pages 1 to 16, 2010. [Author PDF](https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf). Inspected: quality measurements, human comparisons, Experiment III and conclusions. The filename is not the publication year.

**Direct URLs**
https://cambolbro.com/cv/publications/ciaig-browne-maire-19.pdf

**Where and why it appeared:** UG-T04-S01, section R2, "Evolutionary Game Design". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in author PDF; final year/volume/pages retained from the attachment.
**Content inspected for this handover:** Abstract and opening text only.
**Inherited access report:** Inspected: quality measurements, human comparisons, Experiment III and conclusions. The filename is not the publication year.

**Relevance and limits:** Paired human preferences and simulation-derived quality in formal games. Preference prediction is game-family and sample dependent; no universal taste or historical-originality certificate.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-D05; UG-T04-E12


##### UG-T04-S103: GAVEL: Generating Games Via Evolution and Language Models

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Graham Todd, Alexander G. Padula, Matthew Stephenson, Éric Piette, Dennis J. N. J. Soemers and Julian Togelius. **GAVEL: Generating Games Via Evolution and Language Models.** NeurIPS 2024. [arXiv:2407.09388v2](https://arxiv.org/html/2407.09388v2). Inspected: representation, mutation, evaluation, experiments, results, limitations and references. [Code link](https://github.com/gdrtodd/gavel), not executed in this investigation.

**Direct URLs**
https://arxiv.org/html/2407.09388v2
https://github.com/gdrtodd/gavel
https://arxiv.org/abs/2407.09388v2

**Where and why it appeared:** UG-T04-S01, section R3, "GAVEL". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified for arXiv record; supplied NeurIPS attribution retained.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: representation, mutation, evaluation, experiments, results, limitations and references. [Code link](https://github.com/gdrtodd/gavel), not executed in this investigation.

**Relevance and limits:** Executable rule mutation, staged evaluation and retention of alternatives. Board-game metrics, expensive search, unused components and inconclusive novelty-cell improvement do not establish TD taste.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-D08; UG-T04-E12


##### UG-T04-S104: Procedural Content Generation through Quality Diversity

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Daniele Gravina, Ahmed Khalifa, Antonios Liapis, Julian Togelius and Georgios N. Yannakakis. **Procedural Content Generation through Quality Diversity.** IEEE Conference on Games, 2019. [arXiv:1907.04053](https://arxiv.org/abs/1907.04053), [HTML](https://arxiv.org/html/1907.04053), [DOI:10.1109/CIG.2019.8848053](https://doi.org/10.1109/CIG.2019.8848053). Inspected: framework, surveyed applications, challenges and references.

**Direct URLs**
https://arxiv.org/abs/1907.04053
https://arxiv.org/html/1907.04053
https://doi.org/10.1109/CIG.2019.8848053

**Where and why it appeared:** UG-T04-S01, section R4, "Quality diversity for PCG". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv, including CoG 2019 record.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: framework, surveyed applications, challenges and references.

**Relevance and limits:** Retain diverse useful designs instead of one scalar winner. Descriptor coverage is not proved strategic diversity; this is a survey.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S105: Illuminating search spaces by mapping elites

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Jean-Baptiste Mouret and Jeff Clune. **Illuminating search spaces by mapping elites.** 2015 preprint. [arXiv:1504.04909v1](https://arxiv.org/abs/1504.04909v1). Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.

**Direct URLs**
https://arxiv.org/abs/1504.04909v1

**Where and why it appeared:** UG-T04-S01, section R5, "MAP-Elites". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v1.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: algorithm, experimental domains, limitations and references via arXiv HTML. The HTML contains an inconsistent generated date; use the arXiv version record for the 2015 date.

**Relevance and limits:** An archive can preserve alternatives. Robotics/search evidence, not TD preference; a grid algorithm is not required just to save candidates.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S106: Abandoning Objectives: Evolution through the Search for Novelty Alone

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Joel Lehman and Kenneth O. Stanley. **Abandoning Objectives: Evolution through the Search for Novelty Alone.** Evolutionary Computation 19(2), pages 189 to 223, 2011. [Primary-text PDF](https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf). Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.

**Direct URLs**
https://www.cs.swarthmore.edu/~meeden/DevelopmentalRobotics/lehman_ecj11.pdf

**Where and why it appeared:** UG-T04-S01, section R6, "Novelty search". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in primary PDF; publication details retained from attachment.
**Content inspected for this handover:** Abstract and opening text only.
**Inherited access report:** Inspected: motivation, behavioral novelty, experimental tasks and discussion. Found in both R4 and R5.

**Relevance and limits:** Avoid prematurely removing unusual behavior solely for one objective. Behavioral novelty alone does not establish usefulness, fidelity or suitability.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S107: Monte-Carlo Tree Search for Simulation-based Strategy Analysis

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Alexander Zook, Brent Harrison and Mark O. Riedl. **Monte-Carlo Tree Search for Simulation-based Strategy Analysis.** The paper identifies FDG 2015 as its conference; the arXiv deposit is 2019. [arXiv:1908.01423v1](https://arxiv.org/html/1908.01423v1). Inspected: methodology, skill-based metrics, case studies, limitations and references. R1 cites the deposit year.

**Direct URLs**
https://arxiv.org/html/1908.01423v1
https://arxiv.org/abs/1908.01423v1

**Where and why it appeared:** UG-T04-S01, section R7, "Simulation-based strategy analysis". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv; original publication year and deposit year kept separate.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: methodology, skill-based metrics, case studies, limitations and references. R1 cites the deposit year.

**Relevance and limits:** Analyze traces and choices, not only outcomes. Search budget is not a validated human skill model.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S108: Evaluating Competitive Game Balance with Restricted Play

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Alexander Jaffe, Alex Miller, Erik Andersen, Yun-En Liu, Anna Karlin and Zoran Popović. **Evaluating Competitive Game Balance with Restricted Play.** AIIDE 2012. [Author-hosted PDF](https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf). Inspected: framework, example application and conclusions. Reached through R1 → R7 → R8.

**Direct URLs**
https://homes.cs.washington.edu/~zoran/jaffe2012ecg.pdf

**Where and why it appeared:** UG-T04-S01, section R8, "Restricted-play evaluation". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in author PDF; venue retained from attachment.
**Content inspected for this handover:** Abstract and opening framework text.
**Inherited access report:** Inspected: framework, example application and conclusions. Reached through R1 → R7 → R8.

**Relevance and limits:** Compare access to a mechanic and rule removal under controlled conditions. Weak agents, scenario choice and policy adaptation can confound usefulness.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S109: Automated Game Design via Conceptual Expansion

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Matthew Guzdial and Mark Riedl. **Automated Game Design via Conceptual Expansion.** AIIDE 2018, 14(1), pages 31 to 37. [arXiv:1809.02232v1](https://arxiv.org/html/1809.02232v1). Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.

**Direct URLs**
https://arxiv.org/html/1809.02232v1
https://arxiv.org/abs/1809.02232v1

**Where and why it appeared:** UG-T04-S01, section R9, "Conceptual expansion". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: representation, expansion, evaluation, discussion and references. Reached through R1.

**Relevance and limits:** Make changed mechanical relationships explicit when adapting sources. Reconstruction and learned approximation are not evidence of unrestricted invention or canon fidelity.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S110: Constrained Novelty Search: A Study on Game Content Generation

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Constrained Novelty Search: A Study on Game Content Generation.** Evolutionary Computation 23(1), pages 101 to 129, 2015. [Author PDF](https://antoniosliapis.com/papers/constrained_novelty_search.pdf), [DOI:10.1162/EVCO_a_00123](https://doi.org/10.1162/EVCO_a_00123). Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.

**Direct URLs**
https://antoniosliapis.com/papers/constrained_novelty_search.pdf
https://doi.org/10.1162/EVCO_a_00123

**Where and why it appeared:** UG-T04-S01, section R10, "Constrained novelty search". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in author PDF; publication DOI/details retained from attachment.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: two-population approach, game-content application and reported qualifications. Found through R4; bibliographic metadata also checked through Crossref.

**Relevance and limits:** Keep unsupported but useful proposals separate from valid candidates. The project's proposal organization is an adaptation, not an implemented dual-population algorithm.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S111: Sentient Sketchbook: Computer-Aided Game Level Authoring

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Antonios Liapis, Georgios N. Yannakakis and Julian Togelius. **Sentient Sketchbook: Computer-Aided Game Level Authoring.** Foundations of Digital Games, 2013, pages 213 to 220. [Author PDF](https://antoniosliapis.com/papers/sentient_sketchbook.pdf). Inspected: tool interaction, checks, alternatives and small-scale designer study. Found through R4.

**Direct URLs**
https://antoniosliapis.com/papers/sentient_sketchbook.pdf

**Where and why it appeared:** UG-T04-S01, section R11, "Sentient Sketchbook". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in author PDF; proceedings details retained from attachment.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: tool interaction, checks, alternatives and small-scale designer study. Found through R4.

**Relevance and limits:** Preserve designer intent and inspect changes between revisions. Small map-specific study; valid alternatives can still ignore intent.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S112: Generation and Evaluation in the Human Invention Process through the Lens of Game Design

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Katherine M. Collins et al. **Generation and Evaluation in the Human Invention Process through the Lens of Game Design.** CogSci 2025. [arXiv:2508.10914v1](https://arxiv.org/html/2508.10914v1). Inspected: dataset, proposal/evaluation model, results, limitations and references.

**Direct URLs**
https://arxiv.org/html/2508.10914v1
https://arxiv.org/abs/2508.10914v1

**Where and why it appeared:** UG-T04-S01, section R12, "Human game invention". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: dataset, proposal/evaluation model, results, limitations and references.

**Relevance and limits:** Distinguish proposal plausibility from simulated consequences. Cognitive modeling and formalizable subsets, not a general model of AI design taste.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S113: People use fast, goal-directed simulation to reason about novel games

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Cedegao E. Zhang, Katherine M. Collins, Lionel Wong, Adrian Weller and Joshua B. Tenenbaum. **People use fast, goal-directed simulation to reason about novel games.** CogSci 2024. [arXiv:2407.14095v1](https://arxiv.org/html/2407.14095v1). Inspected: model, human judgment task and discussion. Found through R12.

**Direct URLs**
https://arxiv.org/html/2407.14095v1
https://arxiv.org/abs/2407.14095v1

**Where and why it appeared:** UG-T04-S01, section R13, "Reasoning about unfamiliar games". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: model, human judgment task and discussion. Found through R12.

**Relevance and limits:** Separate novice comprehension from expert strategic evaluation. Simple turn-based game judgments may not transfer to real-time TD.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S114: Mathematical discoveries from program search with large language models

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Bernardino Romera-Paredes et al. **Mathematical discoveries from program search with large language models.** Nature 625, pages 468 to 475, 2024; published online in December 2023. [DOI:10.1038/s41586-023-06924-6](https://www.nature.com/articles/s41586-023-06924-6). Inspected: primary article's method and reported applications. [Code link](https://github.com/google-deepmind/funsearch), not executed here.

**Direct URLs**
https://www.nature.com/articles/s41586-023-06924-6
https://github.com/google-deepmind/funsearch

**Where and why it appeared:** UG-T04-S01, section R14, "FunSearch". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against publisher article.
**Content inspected for this handover:** Relevant abstract/method overview.
**Inherited access report:** Inspected: primary article's method and reported applications. [Code link](https://github.com/google-deepmind/funsearch), not executed here.

**Relevance and limits:** Independent execution/evaluation rather than self-certifying generated explanations. Mathematical/algorithmic objectives are easier to verify than enjoyment or fidelity.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S115: Ludax: A GPU-Accelerated Description Language for Board Games

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Graham Todd et al. **Ludax: A GPU-Accelerated Description Language for Board Games.** [arXiv:2506.22609](https://arxiv.org/abs/2506.22609), first submitted in 2025; [v2 inspected](https://arxiv.org/html/2506.22609v2), March 25, 2026. Inspected: representation, compilation, supported game class and references. R12 cites the earlier 2025 version; this record deliberately distinguishes it from the newer text inspected.

**Direct URLs**
https://arxiv.org/abs/2506.22609
https://arxiv.org/html/2506.22609v2
https://arxiv.org/abs/2506.22609v2

**Where and why it appeared:** UG-T04-S01, section R15, "Ludax". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v2, with a title correction recorded below.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: representation, compilation, supported game class and references. R12 cites the earlier 2025 version; this record deliberately distinguishes it from the newer text inspected.

**Relevance and limits:** Defined operators give a specification executable meaning. Turn-based board-game DSL, not an existing real-time TD Consumer.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12

**Bibliographic correction:** Original title says "Description Language"; checked arXiv v2 says "Ludax: A GPU-Accelerated Domain Specific Language for Board Games". Version is March 25, 2026. Do not silently substitute the newer version for the one used by earlier citations.


##### UG-T04-S116: Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Chenglei Si, Diyi Yang and Tatsunori Hashimoto. **Can LLMs Generate Novel Research Ideas? A Large-Scale Human Study with 100+ NLP Researchers.** ICLR 2025; preprint 2024. [arXiv:2409.04109](https://arxiv.org/abs/2409.04109), [v1 inspected](https://arxiv.org/html/2409.04109v1). Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.

**Direct URLs**
https://arxiv.org/abs/2409.04109
https://arxiv.org/html/2409.04109v1

**Where and why it appeared:** UG-T04-S01, section R16, "LLM idea generation and self-evaluation". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv; the attachment's inspected v1 is distinct from later versions.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: study description and discussion of novelty, diversity and self-evaluation. The inspected evidence is the preprint, not an audit of changes in the final proceedings version.

**Relevance and limits:** Blind review and independent checks of idea quality. Research proposals, not completed games; diversity/self-evaluation limits remain.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S117: Mortar: Evolving Mechanics for Automatic Game Design

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Muhammad U. Nasir, Yuchen Li, Steven James and Julian Togelius. **Mortar: Evolving Mechanics for Automatic Game Design.** GECCO 2026. [Primary PDF](https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf), [DOI:10.1145/3795095.3805100](https://doi.org/10.1145/3795095.3805100). Inspected: method, mechanic contribution measure, human comparisons and related work.

**Direct URLs**
https://www.raillab.org/publication/nasir-2026-mortar/nasir-2026-mortar.pdf
https://doi.org/10.1145/3795095.3805100

**Where and why it appeared:** UG-T04-S01, section R17, "Mortar: mechanics evaluated inside games". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in primary author-hosted PDF; publication details retained from attachment.
**Content inspected for this handover:** Abstract and relevant description only; historical study counts not independently reproduced.
**Inherited access report:** Inspected: method, mechanic contribution measure, human comparisons and related work.

**Relevance and limits:** Measure mechanics in combinations and test preference separately. Small study includes disagreement between skill-order proxy and human preference.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q11


##### UG-T04-S118: Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Michael Cook, Simon Colton, Azalea Raad and Jeremy Gow. **Mechanic Miner: Reflection-Driven Game Mechanic Discovery and Level Design.** EvoApplications 2013. [Author PDF](https://www.possibilityspace.org/papers/evo13.pdf), [DOI:10.1007/978-3-642-37192-9_29](https://doi.org/10.1007/978-3-642-37192-9_29). Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.

**Direct URLs**
https://www.possibilityspace.org/papers/evo13.pdf
https://doi.org/10.1007/978-3-642-37192-9_29

**Where and why it appeared:** UG-T04-S01, section R18, "Mechanic Miner". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Title/authors verified in author PDF; publication details retained from attachment.
**Content inspected for this handover:** Abstract and opening description only.
**Inherited access report:** Inspected: representation, search, generated mechanics and pilot study. Promoted from the original R9 citation lead.

**Relevance and limits:** Provide scenarios showing what a new mechanic enables. Solver utility does not ensure player comprehension; participant-count inconsistency retained.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q11


##### UG-T04-S119: Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Johor Jara Gonzalez, Seth Cooper and Matthew Guzdial. **Mechanic Maker 2.0: Reinforcement Learning for Evaluating Generated Rules.** AIIDE 19(1), pages 266 to 275, 2023. [arXiv:2309.09476v3](https://arxiv.org/html/2309.09476v3), [DOI:10.1609/aiide.v19i1.27522](https://doi.org/10.1609/aiide.v19i1.27522). Inspected: environment recreation, evaluator comparison and limitations.

**Direct URLs**
https://arxiv.org/html/2309.09476v3
https://doi.org/10.1609/aiide.v19i1.27522
https://arxiv.org/abs/2309.09476v3

**Where and why it appeared:** UG-T04-S01, section R19, "Mechanic Maker 2.0". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v3.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: environment recreation, evaluator comparison and limitations.

**Relevance and limits:** Evaluator choice changes discovered mechanics; learning ability and final performance differ. RL fitness exploitation and slower evaluation; no human usability validation.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q11


##### UG-T04-S120: Pixie: Code-Level Mechanic Generation for Game Designers

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Michael Cook. **Pixie: Code-Level Mechanic Generation for Game Designers.** AIIDE 21(1), pages 206 to 215, 2025. [Publisher PDF](https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962), [DOI:10.1609/aiide.v21i1.36824](https://doi.org/10.1609/aiide.v21i1.36824). Access gap: the publisher download failed and an attempted author copy was unavailable. Mortar's related-work discussion independently confirms Pixie as code-level mechanic evolution using a genetic-programming approach.

**Direct URLs**
https://ojs.aaai.org/index.php/AIIDE/article/download/36824/38962
https://doi.org/10.1609/aiide.v21i1.36824
https://ojs.aaai.org/index.php/AIIDE/article/view/36824

**Where and why it appeared:** UG-T04-S01, section R20, "Pixie". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Publisher bibliographic record verified; original full-text access gap remains.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Access gap: the publisher download failed and an attempted author copy was unavailable. Mortar's related-work discussion independently confirms Pixie as code-level mechanic evolution using a genetic-programming approach.

**Relevance and limits:** Potential explicit code-level mechanic interface. Annotated-Unity and engineering claims remain unverified without full text.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S121: Automatic Game Design via Mechanic Generation

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Alexander Zook and Mark O. Riedl. **Automatic Game Design via Mechanic Generation.** AAAI 2014. [arXiv:1908.01420v1](https://arxiv.org/html/1908.01420v1), [DOI:10.1609/aaai.v28i1.8788](https://doi.org/10.1609/aaai.v28i1.8788). Inspected: representation, constraint-based generation, planning and demonstration domains. The arXiv deposit is 2019; the paper is from 2014. Promoted from the original R3/R7 citation lead.

**Direct URLs**
https://arxiv.org/html/1908.01420v1
https://doi.org/10.1609/aaai.v28i1.8788
https://arxiv.org/abs/1908.01420v1

**Where and why it appeared:** UG-T04-S01, section R21, "Formal mechanic generation". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv; 2014 publication and 2019 deposit remain distinct.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: representation, constraint-based generation, planning and demonstration domains. The arXiv deposit is 2019; the paper is from 2014. Promoted from the original R3/R7 citation lead.

**Relevance and limits:** Constraint-based actions and reachability for formal mechanic checks. Encoded reachability does not establish fun or strategic depth.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S122: TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Dawei Wang et al. **TowerMind: A Tower Defence Game Learning Environment and Benchmark for LLM as Agents.** AAAI 2026. [arXiv:2601.05899v1](https://arxiv.org/html/2601.05899v1), [DOI:10.1609/aaai.v40i31.39818](https://doi.org/10.1609/aaai.v40i31.39818). Inspected: environment, observation formats, evaluation scope and reported agent failures. [Code](https://github.com/tb6147877/TowerMind), not executed.

**Direct URLs**
https://arxiv.org/html/2601.05899v1
https://doi.org/10.1609/aaai.v40i31.39818
https://github.com/tb6147877/TowerMind
https://arxiv.org/abs/2601.05899v1

**Where and why it appeared:** UG-T04-S01, section R22, "TowerMind". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v1; later versions exist but were not substituted.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: environment, observation formats, evaluation scope and reported agent failures. [Code](https://github.com/tb6147877/TowerMind), not executed.

**Relevance and limits:** Possible TD action/observation interface reference. Agent benchmark, not a unit generator; compatibility untested.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S123: Automatic generation of tower defense levels using PCG

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Simon Liu et al. **Automatic generation of tower defense levels using PCG.** FDG 2019. [DOI:10.1145/3337722.3337723](https://doi.org/10.1145/3337722.3337723), [supplied manuscript entry](https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG). Access gap: ResearchGate returned HTTP 403; the primary text was not independently inspected.

**Direct URLs**
https://doi.org/10.1145/3337722.3337723
https://www.researchgate.net/publication/335590750_Automatic_generation_of_tower_defense_levels_using_PCG

**Where and why it appeared:** UG-T04-S01, section R23, "Tower defense level generation". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Unverified independently; DOI retrieval failed. Supplied record retained.
**Content inspected for this handover:** Not inspected.
**Inherited access report:** Access gap: ResearchGate returned HTTP 403; the primary text was not independently inspected.

**Relevance and limits:** Joint maps/placement/waves as a future evaluation context. Supplied flat Monte Carlo and implementation claims remain unverified.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S124: Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Vid Kraner, Iztok Fister Jr. and Lucija Brezočnik. **Procedural Content Generation of Custom Tower Defense Game Using Genetic Algorithms.** NT 2021, Lecture Notes in Networks and Systems 233. [DOI:10.1007/978-3-030-75275-0_54](https://doi.org/10.1007/978-3-030-75275-0_54), [supplied manuscript entry](https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms). Access gap: ResearchGate returned HTTP 403.

**Direct URLs**
https://doi.org/10.1007/978-3-030-75275-0_54
https://www.researchgate.net/publication/351488584_Procedural_Content_Generation_of_Custom_Tower_Defense_Game_Using_Genetic_Algorithms

**Where and why it appeared:** UG-T04-S01, section R24, "Save the Sheep". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Unverified independently; DOI retrieval failed. Supplied record retained.
**Content inspected for this handover:** Not inspected.
**Inherited access report:** Access gap: ResearchGate returned HTTP 403.

**Relevance and limits:** Separate generated arrangements from human-authored premise/assets. Full text unavailable; 3D game implementation does not mean autonomous design.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S125: Computational Intelligence and Tower Defence Games

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Phillipa Avery, Julian Togelius, Elvis Alistar and Robert Pieter van Leeuwen. **Computational Intelligence and Tower Defence Games.** IEEE Congress on Evolutionary Computation, 2011. [DOI:10.1109/CEC.2011.5949738](https://doi.org/10.1109/CEC.2011.5949738), [supplied manuscript entry](https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games). Access gap: ResearchGate returned HTTP 403.

**Direct URLs**
https://doi.org/10.1109/CEC.2011.5949738
https://www.researchgate.net/publication/224246975_Computational_Intelligence_and_Tower_Defence_Games

**Where and why it appeared:** UG-T04-S01, section R25, "Early tower defense research". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Unverified independently; DOI retrieval failed. Supplied record retained.
**Content inspected for this handover:** Not inspected.
**Inherited access report:** Access gap: ResearchGate returned HTTP 403.

**Relevance and limits:** Separate TD generation components and experience responsibilities. Access gap; no downstream implementation reuse established.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S126: Controllable procedural map generation via multiobjective evolution

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Julian Togelius et al. **Controllable procedural map generation via multiobjective evolution.** Genetic Programming and Evolvable Machines 14, pages 245 to 277, 2013. [Publisher article and DOI](https://link.springer.com/article/10.1007/s10710-012-9174-5). Abstract and reference list inspected; the full article requires access.

**Direct URLs**
https://link.springer.com/article/10.1007/s10710-012-9174-5

**Where and why it appeared:** UG-T04-S01, section R26, "Multiobjective RTS maps". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against publisher metadata.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Abstract and reference list inspected; the full article requires access.

**Relevance and limits:** Retain conflicting map objectives. Visual map assessment is not competitive or TD gameplay evidence.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S127: Learning Controllable 3D Level Generators

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Zehua Jiang, Sam Earle, Michael Cerny Green and Julian Togelius. **Learning Controllable 3D Level Generators.** FDG 2022. [arXiv:2206.13623v3](https://arxiv.org/html/2206.13623v3), [DOI:10.1145/3555858.3563273](https://doi.org/10.1145/3555858.3563273). Inspected: PCGRL formulation, 3D movement requirements and control experiments.

**Direct URLs**
https://arxiv.org/html/2206.13623v3
https://doi.org/10.1145/3555858.3563273
https://arxiv.org/abs/2206.13623v3

**Where and why it appeared:** UG-T04-S01, section R27, "Controllable 3D level generators". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v3.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: PCGRL formulation, 3D movement requirements and control experiments.

**Relevance and limits:** Check geometry under actual movement rules. Minecraft-style controllable maps are not production TD terrain or unit generation.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S128: Scaling, Control and Generalization in Reinforcement Learning Level Generators

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Sam Earle, Zehua Jiang and Julian Togelius. **Scaling, Control and Generalization in Reinforcement Learning Level Generators.** 2024. [arXiv:2408.12525v1](https://arxiv.org/html/2408.12525v1). Inspected: method, scaling and generalization scope. The supplied list calls this "PCGRL+"; the citation here uses the manuscript title.

**Direct URLs**
https://arxiv.org/html/2408.12525v1
https://arxiv.org/abs/2408.12525v1

**Where and why it appeared:** UG-T04-S01, section R28, "Scaling PCGRL". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v1, with a title discrepancy recorded below.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: method, scaling and generalization scope. The supplied list calls this "PCGRL+"; the citation here uses the manuscript title.

**Relevance and limits:** Control, fixed constraints and out-of-distribution size tests. Speed or generalization does not show quality of chosen objectives; RL is not required for v1.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12

**Bibliographic correction:** Original entry omits "PCGRL+" and calls the remainder the manuscript title. The checked arXiv v1 title is "PCGRL+: Scaling, Control and Generalization in Reinforcement Learning Level Generators".


##### UG-T04-S129: A Database-Driven Framework for 3D Level Generation with LLMs

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Kaijie Xu and Clark Verbrugge. **A Database-Driven Framework for 3D Level Generation with LLMs.** AIIDE 21(1), pages 367 to 376, 2025. [arXiv:2508.18533v1](https://arxiv.org/html/2508.18533v1), [DOI:10.1609/aiide.v21i1.36840](https://doi.org/10.1609/aiide.v21i1.36840). Inspected: assembly pipeline, repair and limitations.

**Direct URLs**
https://arxiv.org/html/2508.18533v1
https://doi.org/10.1609/aiide.v21i1.36840
https://arxiv.org/abs/2508.18533v1

**Where and why it appeared:** UG-T04-S01, section R29, "Database-driven 3D levels". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v1.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: assembly pipeline, repair and limitations.

**Relevance and limits:** Logical progression before scene realization, then revalidation after repair. Some unplayable outputs, substantial curation and no human-player study in the supplied account.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S130: Infinite Photorealistic Worlds using Procedural Generation

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Alexander Raistrick et al. **Infinite Photorealistic Worlds using Procedural Generation.** CVPR 2023. [arXiv:2306.09310](https://arxiv.org/abs/2306.09310). Primary HTML inspected for system purpose and generation scope.

**Direct URLs**
https://arxiv.org/abs/2306.09310

**Where and why it appeared:** UG-T04-S01, section R30, "Infinigen". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Primary HTML inspected for system purpose and generation scope.

**Relevance and limits:** Environmental asset/scene realization. Photorealism does not establish routes, tactical utility or TD balance.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S131: Holodeck: Language Guided Generation of 3D Embodied AI Environments

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Yue Yang et al. **Holodeck: Language Guided Generation of 3D Embodied AI Environments.** CVPR 2024. [arXiv:2312.09067](https://arxiv.org/abs/2312.09067). Primary HTML inspected for planning, retrieval, spatial constraints and evaluation scope.

**Direct URLs**
https://arxiv.org/abs/2312.09067

**Where and why it appeared:** UG-T04-S01, section R31, "Holodeck". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Primary HTML inspected for planning, retrieval, spatial constraints and evaluation scope.

**Relevance and limits:** Spatial constraints when assembling retrieved assets. Embodied environments, not validated tactical maps.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S132: Automatic Content Generation in the Galactic Arms Race Video Game

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Erin J. Hastings, Ratan K. Guha and Kenneth O. Stanley. **Automatic Content Generation in the Galactic Arms Race Video Game.** IEEE Transactions on Computational Intelligence and AI in Games 1(4), pages 245 to 263, 2009. [DOI:10.1109/TCIAIG.2009.2038365](https://doi.org/10.1109/TCIAIG.2009.2038365), [supplied manuscript entry](https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game). Access gap: ResearchGate returned HTTP 403.

**Direct URLs**
https://doi.org/10.1109/TCIAIG.2009.2038365
https://www.researchgate.net/publication/224089471_Automatic_Content_Generation_in_the_Galactic_Arms_Race_Video_Game

**Where and why it appeared:** UG-T04-S01, section R32, "Galactic Arms Race". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Unverified independently; DOI retrieval failed. Supplied record retained.
**Content inspected for this handover:** Not inspected.
**Inherited access report:** Access gap: ResearchGate returned HTTP 403.

**Relevance and limits:** Player usage as one possible feedback source. Usage may reflect power or access rather than preference; deployment counts unverified here.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S133: OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Maxence Faldor, Jenny Zhang, Antoine Cully and Jeff Clune. **OMNI-EPIC: Open-endedness via Models of human Notions of Interestingness with Environments Programmed in Code.** ICLR 2025. [arXiv:2405.15568v3](https://arxiv.org/html/2405.15568v3). Inspected: archive, task/environment generation, interestingness filtering and evaluation modes.

**Direct URLs**
https://arxiv.org/html/2405.15568v3
https://arxiv.org/abs/2405.15568v3

**Where and why it appeared:** UG-T04-S01, section R33, "OMNI-EPIC". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v3.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: archive, task/environment generation, interestingness filtering and evaluation modes.

**Relevance and limits:** Retain success/failure histories for proposal generation. The supplied long 200-iteration run did not train agents; interestingness is model judged.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S134: DreamGarden: A Designer Assistant for Growing Games from a Single Prompt

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Sam Earle, Samyak Parajuli and Andrzej Banburski-Fahey. **DreamGarden: A Designer Assistant for Growing Games from a Single Prompt.** CHI 2025. [arXiv:2410.01791v2](https://arxiv.org/html/2410.01791v2), [DOI:10.1145/3706598.3714233](https://doi.org/10.1145/3706598.3714233). Inspected: hierarchical planning, Unreal implementation modules and usability study.

**Direct URLs**
https://arxiv.org/html/2410.01791v2
https://doi.org/10.1145/3706598.3714233
https://arxiv.org/abs/2410.01791v2

**Where and why it appeared:** UG-T04-S01, section R34, "DreamGarden". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v2.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: hierarchical planning, Unreal implementation modules and usability study.

**Relevance and limits:** Editable plans and targeted repair preserving designer choices. Ten-person interaction study, not evidence of complete strategically good games.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S135: AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Lei Yin et al. **AutoUE: Automated Generation of 3D Games in Unreal Engine via Multi-Agent Systems.** 2026 preprint. [arXiv:2603.07106v1](https://arxiv.org/html/2603.07106v1). Inspected: engine workflow, documentation retrieval, runtime tests and benchmark. [Code](https://github.com/Pluto156/AutoUE), not executed.

**Direct URLs**
https://arxiv.org/html/2603.07106v1
https://github.com/Pluto156/AutoUE
https://arxiv.org/abs/2603.07106v1

**Where and why it appeared:** UG-T04-S01, section R35, "AutoUE". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v1.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: engine workflow, documentation retrieval, runtime tests and benchmark. [Code](https://github.com/Pluto156/AutoUE), not executed.

**Relevance and limits:** Engine documentation and actual runtime state changes for implementation. Multi-agent game prototyping/PlayGen-20, not validated original TD design or player taste.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S136: Genie: Generative Interactive Environments

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Jake Bruce et al. **Genie: Generative Interactive Environments.** ICML 2024. [arXiv:2402.15391](https://arxiv.org/abs/2402.15391). Primary HTML inspected for model and output representation.

**Direct URLs**
https://arxiv.org/abs/2402.15391

**Where and why it appeared:** UG-T04-S01, section R36, "Genie". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Primary HTML inspected for model and output representation.

**Relevance and limits:** Interactive visual exploration as a separate research path. A learned interactive world is not an explicit editable rules/runtime implementation.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S137: Emergence of player tactics by expert-guided machine learning: an industry tower defence case study

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Stuart Anderson and Ruth Falconer. **Emergence of player tactics by expert-guided machine learning: an industry tower defence case study.** Entertainment Computing 54, article 100963, 2025. [Institutional record](https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/), [DOI:10.1016/j.entcom.2025.100963](https://doi.org/10.1016/j.entcom.2025.100963). Institutional abstract and metadata inspected; the linked full-text download returned HTTP 403.

**Direct URLs**
https://rke.abertay.ac.uk/en/publications/emergence-of-player-tactics-by-expert-guided-machine-learning-an-/
https://doi.org/10.1016/j.entcom.2025.100963

**Where and why it appeared:** UG-T04-S01, section R37, "Player tactics in a commercial TD game". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against the institutional record.
**Content inspected for this handover:** Abstract only; full-text access gap retained.
**Inherited access report:** Institutional abstract and metadata inspected; the linked full-text download returned HTTP 403.

**Relevance and limits:** Contextual player action sequences rather than pick rates alone. Behavior analysis, not generation; only abstract independently inspected.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S138: GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development

**Type:** Research paper or survey; publication status and identifiers are preserved in the original bibliographic record below.
**Original bibliographic record and inherited access wording:** Xiuhui Zhang et al. **GameASG-Bench: Benchmarking Autonomous Software Generation for Game Development.** Preprint, September 18, 2026. [arXiv:2609.21293v1](https://arxiv.org/html/2609.21293v1). Inspected: evaluation contract, task suite and aggregate results.

**Direct URLs**
https://arxiv.org/html/2609.21293v1
https://arxiv.org/abs/2609.21293v1

**Where and why it appeared:** UG-T04-S01, section R38, "GameASG-Bench". This source was already in the supplied attachment, not discovered for the first time in this handover.
**Bibliographic verification:** Verified against arXiv v1; supplied September 2026 preprint remains unreproduced.
**Content inspected for this handover:** Abstract only.
**Inherited access report:** Inspected: evaluation contract, task suite and aggregate results.

**Relevance and limits:** Separate individual-check pass rate from complete-task success. Software compliance benchmark, not novelty, balance or human preference.
**Associated artifacts:** Code/data/model links, if present in the original record below, are entry points only. None were run or license-audited.
**Linked records:** UG-T04-E12; UG-T04-Q12


#### Additional papers retained by the supplied catalogue


##### UG-T04-S139: Experience-Driven Procedural Content Generation

**Type and publication status:** Research paper. IEEE Transactions on Affective Computing 2(3), 2011, as recorded in the supplied catalogue.
**Authors/organization:** Georgios N. Yannakakis; Julian Togelius
**Persistent identifier or locator:** DOI:10.1109/T-AFFC.2011.6

**Direct URLs**
https://doi.org/10.1109/T-AFFC.2011.6

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, Additional references and access gaps.
**Bibliographic verification:** Unverified independently in this handover; DOI retrieval failed.
**Content inspected for this handover:** Not inspected in this handover. Attachment also recorded an access gap.

**Relevance and limits:** Experience-based generation and preference modeling; related to UG-T04-E07. Do not merge this 2011 article with the separately cited 2015 extended abstract, UG-T04-S22.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S140: Generative AI enhances individual creativity but reduces the collective diversity of novel content

**Type and publication status:** Research paper. Science Advances 10(28), 2024, as recorded in the supplied catalogue.
**Authors/organization:** Anil R. Doshi; Oliver P. Hauser
**Persistent identifier or locator:** DOI:10.1126/sciadv.adn5290

**Direct URLs**
https://doi.org/10.1126/sciadv.adn5290

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, Additional references and access gaps.
**Bibliographic verification:** Unverified independently in this handover; DOI retrieval failed.
**Content inspected for this handover:** Not inspected; inherited publisher access gap.

**Relevance and limits:** Lead for loss of diversity across generated candidates. The supplied summary concerns short fiction, not mechanically distinct TD units.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S141: Large language models are homogeneously creative

**Type and publication status:** Research paper. PNAS Nexus 5(3), 2026, as recorded in the supplied catalogue.
**Authors/organization:** Emily Wenger; Yoed N. Kenett
**Persistent identifier or locator:** DOI:10.1093/pnasnexus/pgag042

**Direct URLs**
https://doi.org/10.1093/pnasnexus/pgag042

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, Additional references and access gaps.
**Bibliographic verification:** Unverified independently in this handover; DOI retrieval failed.
**Content inspected for this handover:** Not inspected; inherited publisher access gap.

**Relevance and limits:** Lead for cross-model similarity and candidate diversity. No transfer to this project's character adaptations established.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S142: Superhuman artificial intelligence can improve human decision-making by increasing novelty

**Type and publication status:** Research paper. PNAS 120(12), 2023, as recorded in the supplied catalogue.
**Authors/organization:** Minkyu Shin; Jin Kim; Bas van Opheusden; Thomas L. Griffiths
**Persistent identifier or locator:** DOI:10.1073/pnas.2214840120

**Direct URLs**
https://doi.org/10.1073/pnas.2214840120

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, Additional references and access gaps.
**Bibliographic verification:** Unverified independently in this handover; DOI retrieval failed.
**Content inspected for this handover:** Not inspected; inherited publisher access gap.

**Relevance and limits:** Background lead on human decision changes following AI play. Not direct evidence of inventing new rules or character-based units.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S143: Creativity and artificial intelligence

**Type and publication status:** Research paper. Artificial Intelligence 103(1 to 2), 1998, as recorded in the supplied catalogue.
**Authors/organization:** Margaret A. Boden
**Persistent identifier or locator:** DOI:10.1016/S0004-3702(98)00055-1

**Direct URLs**
https://doi.org/10.1016/S0004-3702(98)00055-1

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, Additional references and access gaps.
**Bibliographic verification:** Unverified independently in this handover; DOI retrieval failed.
**Content inspected for this handover:** Not inspected; inherited publisher access gap.

**Relevance and limits:** Conceptual distinctions between combining, exploring and transforming ideas. Conceptual reference, not empirical validation of LLM design quality.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S144: Program Synthesis as a Generative Method

**Type and publication status:** Research paper. Foundations of Digital Games, 2017.
**Authors/organization:** Eric Butler; Kristin Siu; Alexander Zook
**Persistent identifier or locator:** DOI:10.1145/3102071.3102076

**Direct URLs**
https://www.ericbutler.net/publications
https://dl.acm.org/doi/10.1145/3102071.3102076

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, citation trail from R1; also visible earlier research search.
**Bibliographic verification:** Title/authors/year verified in author bibliography; DOI corroborated by publisher search record.
**Content inspected for this handover:** Metadata only.

**Relevance and limits:** Lead for generating behavior rather than numerical configurations. Full method not inspected. DOI 10.1145/2538528.2538537 belongs to Game-O-Matic, not this paper.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S145: Talakat: Bullet Hell Generation through Constrained Map-Elites

**Type and publication status:** Research paper. GECCO, 2018.
**Authors/organization:** Ahmed Khalifa; Scott Lee; Andy Nealen; Julian Togelius
**Persistent identifier or locator:** arXiv:1806.04718v1; DOI:10.1145/3205455.3205470

**Direct URLs**
https://arxiv.org/abs/1806.04718v1
https://doi.org/10.1145/3205455.3205470

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, citation lead from R4.
**Bibliographic verification:** arXiv title/authors/year verified; publisher search corroborates DOI.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Constrained diversity of combat patterns. Does not establish character adaptation, player taste or a need for MAP-Elites in v1.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S146: Mapping Hearthstone Deck Spaces through MAP-Elites with Sliding Boundaries

**Type and publication status:** Research paper. GECCO, 2019; acceptance noted in arXiv record.
**Authors/organization:** Matthew C. Fontaine; Scott Lee; L. B. Soros; Fernando De Mesentier Silva; Julian Togelius; Amy K. Hoover
**Persistent identifier or locator:** arXiv:1904.10656v1

**Direct URLs**
https://arxiv.org/abs/1904.10656v1

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, citation lead from R4.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Abstract only.

**Relevance and limits:** Interacting content collections and selection of diversity descriptors. Original attachment says 'Michael C. Fontaine' and 'with MAP-Elites'. The primary record gives Matthew and 'through'. This correction is not a new source.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S147: Multiobjective Exploration of the StarCraft Map Space

**Type and publication status:** Research paper. IEEE Conference on Computational Intelligence and Games, 2010, pages 265 to 272.
**Authors/organization:** Julian Togelius; Mike Preuss; Nicola Beume; Simon Wessing; Johan Hagelback; Georgios N. Yannakakis
**Persistent identifier or locator:** DOI:10.1109/ITW.2010.5593346

**Direct URLs**
https://doi.org/10.1109/ITW.2010.5593346
https://pure.itu.dk/en/publications/multiobjective-exploration-of-the-starcraft-map-space
https://www.wi.uni-muenster.de/research/publications/87283

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, citation lead from R1.
**Bibliographic verification:** Institutional metadata corroborates title, authors, year and identifier.
**Content inspected for this handover:** Metadata only; attempted author PDF retrieval failed.

**Relevance and limits:** Lead on conflicting map objectives. Map design is outside the immediate taste-focused unit prototype.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


##### UG-T04-S148: Mixed-Initiative Co-Creativity

**Type and publication status:** Research paper. Foundations of Digital Games, 2014, pages 1 to 8.
**Authors/organization:** Georgios N. Yannakakis; Antonios Liapis; Constantine Alexopoulos
**Persistent identifier or locator:** University of Malta repository handle:123456789/29459

**Direct URLs**
https://www.um.edu.mt/library/oar/handle/123456789/29459
http://www.fdg2014.org/papers/fdg2014_paper_37.pdf

**Where and why it appeared:** GAME-DESIGN-RESEARCH.md, citation lead from R3.
**Bibliographic verification:** Verified against institutional record.
**Content inspected for this handover:** Institutional abstract and metadata; original PDF not retrieved.

**Relevance and limits:** Designer control, alternative suggestions and revision dialogue. Framework relevance does not establish this proposed interface or a measured benefit for Unit-Generator.
**Associated artifacts:** No associated artifact was inspected or executed.
**Linked records:** UG-T04-E12; UG-T04-Q12


#### Internal references and missing source material


These are reference locations, not newly accessed project documents. Public-looking repository URLs do not establish that their code or licenses were inspected. Foundation remains explicitly private.


##### UG-T04-S201: Unit Generator product requirements

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** unit-generator/docs/PRODUCT.md; also linked as docs/PRODUCT.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S01 Project fit and S02 final scope note
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Product requirements and implemented scope referenced but not provided. No unavailable content inferred. Product requirements and implemented scope referenced but not provided.
**Associated artifacts:** unit-generator/docs/PRODUCT.md; also linked as docs/PRODUCT.md
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S202: Towerright context

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** towerright/CONTEXT.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S01 Project fit and reading order
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Platform responsibility source referenced but not provided. No unavailable content inferred. Platform responsibility source referenced but not provided.
**Associated artifacts:** towerright/CONTEXT.md
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S203: Archived September 21 generator survey

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** unit-generator/research/archive/GENERATOR-SURVEY-2026-09-21.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S01 Project fit and reading order
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Older card/monster examples and repository/license observations cannot be reconstructed from the archive reference. No unavailable content inferred. Older card/monster examples and repository/license observations cannot be reconstructed from the archive reference.
**Associated artifacts:** unit-generator/research/archive/GENERATOR-SURVEY-2026-09-21.md
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S204: Generator references

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** GENERATOR-REFERENCES.md; fragment #framework-and-character-sources

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S01 Project fit and final paragraph
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Technical, framework and character-source index not provided. No unavailable content inferred. Technical, framework and character-source index not provided.
**Associated artifacts:** GENERATOR-REFERENCES.md; fragment #framework-and-character-sources
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S205: Unit Generator framework sources

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** unit-generator/research/FRAMEWORK-SOURCES.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S01 final paragraph
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Repository-local application and limitation notes not provided. No unavailable content inferred. Repository-local application and limitation notes not provided.
**Associated artifacts:** unit-generator/research/FRAMEWORK-SOURCES.md
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S206: API compatibility documentation

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** docs/API.md

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S02 final implemented-scope note
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Compatibility details not provided. No unavailable content inferred. Compatibility details not provided.
**Associated artifacts:** docs/API.md
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S207: Foundation

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** No path supplied

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S02 opening paragraph
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Explicitly private internal guidance. Neither accessed nor needed for public contributors according to S02. No unavailable content inferred. Explicitly private internal guidance. Neither accessed nor needed for public contributors according to S02.
**Associated artifacts:** No path supplied
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S208: src/core and experimental RulePack implementation

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** src/core; experimental RulePack API

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** S02 final implemented-scope note
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Bounded numerical implementation and layout-descriptor status reported in S02, not independently inspected code. No unavailable content inferred. Bounded numerical implementation and layout-descriptor status reported in S02, not independently inspected code.
**Associated artifacts:** src/core; experimental RulePack API
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S209: Mardwerk Map Generator repository

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** https://github.com/mardwerk/map-generator

**Direct URLs**
https://github.com/mardwerk/map-generator

**Where and why it appeared:** S02 Related repositories
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Referenced repository only. No private/project code retrieved or required here. No unavailable content inferred. Referenced repository only. No private/project code retrieved or required here.
**Associated artifacts:** https://github.com/mardwerk/map-generator
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S210: Mardwerk organization metadata repository

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** https://github.com/mardwerk/.github

**Direct URLs**
https://github.com/mardwerk/.github

**Where and why it appeared:** S02 Related repositories
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** Referenced repository only. No code, license or organization configuration inspected. No unavailable content inferred. Referenced repository only. No code, license or organization configuration inspected.
**Associated artifacts:** https://github.com/mardwerk/.github
**Linked records:** UG-T04-E19; UG-T04-Q08; UG-T04-Q09


##### UG-T04-S211: Character wiki/source packet for the Luffy example

**Type and publication status:** Unavailable project document/repository or source gap. Internal reference or source gap; no publication record.
**Authors/organization:** Project material; authorship not established.
**Persistent identifier or locator:** No URL, revision or source text supplied

**Direct URLs**
No direct URL supplied or recovered.

**Where and why it appeared:** User's initial wiki-input request and later Luffy example
**Bibliographic verification:** Unverified beyond its mention in the supplied material.
**Content inspected for this handover:** Not inspected.

**Relevance and limits:** An intended input/source gap, not a located wiki or verified canon source. No unavailable content inferred. An intended input/source gap, not a located wiki or verified canon source.
**Associated artifacts:** No URL, revision or source text supplied
**Linked records:** UG-T04-Q04


### B. Newly identified literature leads

All entries below are **newly found, not part of the original conversation**. They fill a specific gap or challenge the proposed approach. They are not retroactive support for an accepted project decision.


##### UG-T04-S301: Human Feedback is not Gold Standard

**Type and publication status:** Research paper. ICLR, 2024, published conference paper.
**Authors/organization:** Tom Hosking; Phil Blunsom; Max Bartolo
**Persistent identifier or locator:** arXiv:2309.16349v2; OpenReview:7W3GLNImfS

**Direct URLs**
https://arxiv.org/html/2309.16349v2
https://openreview.net/forum?id=7W3GLNImfS
https://github.com/cohere-ai/human-feedback-paper

**Where and why it appeared:** Newly found, not part of the original conversation.
**Bibliographic verification:** Verified against arXiv, conference record and PDF conference heading.
**Content inspected for this handover:** Relevant Sections 2 to 4; a corresponding PDF results page visually checked.

**Relevance and limits:** Shows that human feedback can underweight correctness-related defects and respond to presentation. Supports separating stated preference from factual/mechanical checks. Not TD research. This does not invalidate subjective taste or establish that the same biases occurred here.
**Associated artifacts:** Associated code/data repository identified, not executed or license-audited.
**Linked records:** UG-T04-E16; UG-T04-Q06; UG-T04-Q12


##### UG-T04-S302: COMAL: A Convergent Meta-Algorithm for Aligning LLMs with General Preferences

**Type and publication status:** Research paper. 2024 preprint; v2 dated October 13, 2025. Final venue not independently established.
**Authors/organization:** Yixin Liu; Argyris Oikonomou; Weiqiang Zheng; Yang Cai; Arman Cohan
**Persistent identifier or locator:** arXiv:2410.23223v2

**Direct URLs**
https://arxiv.org/abs/2410.23223v2
https://arxiv.org/html/2410.23223v2

**Where and why it appeared:** Newly found, not part of the original conversation.
**Bibliographic verification:** Verified against arXiv.
**Content inspected for this handover:** Relevant introduction, preference formulation and discussion of Bradley-Terry assumptions; proofs not audited.

**Relevance and limits:** Challenges assuming that all preferences can be faithfully represented by a single scalar reward. Distinct reviewers or contexts can produce non-transitive aggregate preferences. No preference cycles measured in Unit-Generator. Does not require adopting COMAL or game-theoretic training for v1. Chatbot Arena itself discusses nonparametric and misspecified-model inference.
**Associated artifacts:** No code or trained artifact executed.
**Linked records:** UG-T04-E17; UG-T04-D06; UG-T04-Q06


##### UG-T04-S303: Learning the Designer's Preferences to Drive Evolution

**Type and publication status:** Research paper. EvoApplications, 2020.
**Authors/organization:** Alberto Alvarez; Jose Font
**Persistent identifier or locator:** arXiv:2003.03268v1; DOI:10.1007/978-3-030-43722-0_28

**Direct URLs**
https://arxiv.org/html/2003.03268v1
https://doi.org/10.1007/978-3-030-43722-0_28

**Where and why it appeared:** Newly found, not part of the original conversation.
**Bibliographic verification:** Title/authors/year verified against arXiv; publication identifier corroborated by publisher metadata.
**Content inspected for this handover:** Relevant Sections 3 to 5.

**Relevance and limits:** A mixed-initiative dungeon editor learns designer preferences to guide evolutionary suggestions. It reports limitations from noisy/sparse feedback and less preference-model use than expected. Study involved 15 game-design students, not TD character reviewers. Labels were derived from suggestion selections and nearby alternatives rather than the proposed explicit pairwise unit judgments. No guaranteed sample sufficiency follows.
**Associated artifacts:** Evolutionary Dungeon Designer is the studied system; no implementation run here.
**Linked records:** UG-T04-E18; UG-T04-D07; UG-T04-Q10


### Literature gaps requiring a later, broader review

The thread still lacks direct evidence for character-source-grounded TD progression taste, transfer from concept votes to play feel, and preference generalization across characters and game constraints. Further work should examine disagreement and non-transitive preferences, the amount and type of data needed for editable preference guidance versus trained rankers, and whether outline-first review improves full-unit judgments.

Primary-text access remains incomplete for several TD, creativity and implementation leads. Unresolved titles such as RPGAgent, Progression Balancing, 1001 Nights and Damage optimization in video games are retained rather than guessed. Exact source wiki material and the linked project implementation files remain unavailable.

The 2011 and 2015 experience-driven PCG works are distinct records. The original Ludax title, PCGRL+ prefix and Hearthstone deck-paper attribution discrepancies are preserved with corrections where primary metadata was checked. A recoverable title does not verify every earlier claim about the work.

No claim of exhaustive literature coverage, original invention, license compatibility or complete historical reconstruction follows from this register.

