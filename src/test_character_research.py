"""Verify name-only research, equal-budget repeat improvement and safe reuse."""

from copy import deepcopy
import json
from pathlib import Path
import shutil
import tempfile
import unittest
from unittest.mock import patch

from src import character_research as research, wiki
from src.model import Model


class FakeModel:
    name = 'test/research'

    def __init__(self, *, bad_citation=False, extra=False, ambiguous=False):
        self.calls, self.usage = [], []
        self.bad_citation, self.extra, self.ambiguous = bad_citation, extra, ambiguous

    def json(self, instructions, data, *, tokens, search_results=0):
        self.calls.append((tokens, search_results, deepcopy(data)))
        self.usage.append({'total_tokens': 100, 'cost': .001})
        if search_results:
            return {'ambiguous': self.ambiguous, 'alternatives': ['Usopp, One Piece', 'Usopp, Another Work'],
                    'identity': {'name': 'Usopp', 'work': 'One Piece', 'aliases': ['Sogeking'],
                                 'scope': 'Manga canon, documented techniques only',
                                 'evidenceUrls': ['https://example.com/usopp']},
                    'urls': ['https://example.com/usopp', 'https://example.com/techniques']}, [
                        {'url': 'https://example.com/usopp', 'title': 'Usopp', 'content': 'Usopp from One Piece uses a slingshot.'},
                        {'url': 'https://example.com/techniques', 'title': 'Techniques', 'content': 'Pop Green seeds rapidly grow into plants.'}]
        source = data['sources'][0]
        reference = {'source': source['source'], 'passage': source['passages'][0]['id']}
        if self.bad_citation:
            reference['passage'] = 'p-invented'
        records = [{'id': 'slingshot', 'kind': 'equipment', 'name': 'Slingshot',
                    'description': 'Launches physical ammunition.', 'scope': data['identity']['scope'],
                    'limitations': 'The source does not define numerical damage.', 'evidence': [reference],
                    'classification': {'delivery': 'projectile', 'functions': ['damage'], 'status': 'reviewed'}}]
        if self.extra:
            records.append({'id': 'pop-green', 'kind': 'ability', 'name': 'Pop Green', 'description': 'Uses seeds to grow plants.',
                            'scope': data['identity']['scope'], 'evidence': [reference]})
        checks = [{'id': a['id'], 'verdict': 'supported', 'note': 'Supported by the supplied passage.',
                   'evidence': [reference]} for a in data['selectedRecords']]
        return {'summary': 'Sourced capabilities. Exact game values are not established.', 'records': records, 'checks': checks}, []


def document(url):
    return {'id': 'source-' + url.rsplit('/', 1)[-1], 'title': 'English evidence', 'url': url,
            'access': 'retrieved', 'language': 'en', 'retrievedAt': '2026-10-02T10:00:00+00:00',
            'text': 'Usopp from One Piece uses a slingshot. Pop Green seeds grow into plants.'}


class CharacterResearchTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix='character-research-test-')
        self.root = Path(self.temp.name) / 'wiki'

    def tearDown(self):
        self.temp.cleanup()

    def run_research(self, model=None, query='Usopp', **kwargs):
        with patch.object(research, 'fetch_document', side_effect=document):
            return research.research(query, root=self.root, model=model or FakeModel(), **kwargs)

    def snapshot(self):
        return {p.relative_to(self.root): p.read_bytes() for p in self.root.rglob('*') if p.is_file()}

    def test_new_and_repeat_get_identical_budgets_and_expand_same_entry(self):
        first_model = FakeModel()
        first = self.run_research(first_model)
        second_model = FakeModel(extra=True)
        second = self.run_research(second_model, query='Sogeking')
        self.assertEqual(first['mode'], 'create')
        self.assertEqual(second['mode'], 'improve')
        self.assertEqual(first['budget'], second['budget'])
        self.assertEqual([(t, s) for t, s, _ in first_model.calls], [(t, s) for t, s, _ in second_model.calls])
        self.assertEqual(second['newRecords'], 1)
        self.assertEqual(second['entry']['metadata']['research']['runs'], 2)
        self.assertEqual(len(list(self.root.glob('*/*/README.md'))), 1)
        self.assertEqual(len(second_model.calls[0][2]['existing']), 1)
        self.assertEqual(second['entry']['metadata']['abilities'][0]['verification']['verdict'], 'supported')
        self.assertTrue(all(a['classification'].get('status') == 'suggested' for a in second['entry']['metadata']['abilities']))

    def test_old_folder_spelling_is_found_and_normalized_without_losing_manual_files(self):
        first = self.run_research()['entry']
        old = self.root / 'ONE PIECE' / 'Usopp'
        old.parent.mkdir()
        shutil.move(first['path'], old)
        (old / 'notes.md').write_text('Keep this manual note.')
        result = self.run_research(query='ＵＳＯＰＰ')
        self.assertEqual(result['mode'], 'improve')
        self.assertFalse(old.exists())
        self.assertEqual((self.root / 'one-piece/usopp/notes.md').read_text(), 'Keep this manual note.')
        self.assertEqual(result['entry']['metadata']['research']['runs'], 2)

    def test_human_reviewed_records_and_tags_are_preserved_with_suggested_updates(self):
        entry = self.run_research()['entry']
        meta, body = wiki.parse_markdown(entry['markdown'])
        meta['abilities'][0]['classification']['status'] = 'reviewed'
        meta['abilities'][0]['description'] = 'A manually reviewed description.'
        body += '\nKeep these manual notes.\n'
        entry = wiki.save_entry(entry['key'], wiki.markdown(meta, body), entry['revision'], self.root)
        entry = wiki.review_entry(entry['key'], entry['revision'], self.root)
        first = self.run_research(FakeModel(extra=True))['entry']
        second = self.run_research(FakeModel(extra=True))['entry']
        record = next(a for a in second['metadata']['abilities'] if a['id'] == 'slingshot')
        self.assertEqual(record['description'], 'A manually reviewed description.')
        self.assertEqual(record['classification']['status'], 'reviewed')
        self.assertEqual(record['reviewStatus'], 'reviewed')
        self.assertEqual(second['metadata']['status'], 'draft')
        self.assertEqual(second['metadata']['suggestedUpdates'][0]['classification']['status'], 'suggested')
        self.assertIn('Keep these manual notes.', first['body'])

    def test_tag_review_alone_does_not_invent_human_fact_review(self):
        entry = self.run_research()['entry']
        meta, body = wiki.parse_markdown(entry['markdown'])
        meta['abilities'][0]['classification']['status'] = 'reviewed'
        entry = wiki.save_entry(entry['key'], wiki.markdown(meta, body), entry['revision'], self.root)
        updated = self.run_research()['entry']['metadata']['abilities'][0]
        self.assertEqual(updated['classification']['status'], 'reviewed')
        self.assertEqual(updated['reviewStatus'], 'draft')

    def test_bad_model_citations_and_ambiguous_identity_preserve_existing_files(self):
        self.run_research()
        before = self.snapshot()
        for model in [FakeModel(bad_citation=True), FakeModel(ambiguous=True)]:
            with self.assertRaises(ValueError):
                self.run_research(model)
            self.assertEqual(before, self.snapshot())

    def test_blocked_page_uses_disclosed_search_excerpt_and_keeps_prior_citations(self):
        self.run_research()
        with patch.object(research, 'fetch_document', side_effect=OSError('Blocked')):
            result = research.research('Usopp', root=self.root, model=FakeModel(extra=True))
        self.assertTrue(result['warnings'])
        sources = result['entry']['documents']
        self.assertTrue(any(d.get('retrievalMethod') == 'search-excerpt' for d in sources))
        self.assertTrue(any(d.get('retainedPassages') for d in sources))
        with self.assertRaisesRegex(ValueError, 'English'):
            research.source_excerpt({'url': 'https://example.com/japanese', 'content': '日本語の情報です。' * 20}, '2026-10-02T10:00:00+00:00')

    def test_duplicate_identity_and_changed_scope_are_not_merged(self):
        entry = self.run_research()['entry']
        with self.assertRaisesRegex(ValueError, 'canon scope'):
            self.run_research(scope='Another canon boundary')
        other = self.root / 'one-piece/another-usopp'
        shutil.copytree(entry['path'], other)
        before = self.snapshot()
        with self.assertRaisesRegex(ValueError, 'Several Wiki entries'):
            self.run_research()
        self.assertEqual(before, self.snapshot())

    def test_matching_uses_series_aliases_and_normalized_metadata_before_fuzzy_scores(self):
        entry = self.run_research()['entry']
        identity = {'name': 'Usópp', 'work': 'ONE PIECE', 'aliases': []}
        self.assertEqual(research.match(identity, [entry])['key'], entry['key'])
        self.assertIsNone(research.match({**identity, 'work': 'Another series'}, [entry]))
        self.assertEqual(research.normalized('Monkey D. Luffy'), research.normalized('ｍｏｎｋｅｙ-d-luffy'))
        long = 'a' * 100
        copy = deepcopy(entry)
        copy['metadata'].update(name=long, aliases=[])
        with self.assertRaisesRegex(ValueError, 'Possible duplicate'):
            research.match({'name': long[:-1] + 'b', 'work': 'One Piece', 'aliases': []}, [copy])

    def test_stale_write_during_research_preserves_newer_manual_edit(self):
        entry = self.run_research()['entry']
        model = FakeModel(extra=True)
        original_call = model.json
        def call(*args, **kwargs):
            result = original_call(*args, **kwargs)
            if not kwargs.get('search_results'):
                path = Path(entry['path']) / 'README.md'
                path.write_text(path.read_text() + '\nA newer manual edit.\n')
            return result
        model.json = call
        with self.assertRaisesRegex(ValueError, 'changed during research'):
            self.run_research(model)
        self.assertIn('A newer manual edit.', wiki.read_entry(entry['key'], self.root)['body'])
        self.assertEqual(wiki.read_entry(entry['key'], self.root)['metadata']['research']['runs'], 1)

    def test_malformed_or_guessed_url_suggestions_are_ignored_when_identity_has_real_evidence(self):
        model = FakeModel()
        original_call = model.json
        def call(*args, **kwargs):
            result, sources = original_call(*args, **kwargs)
            if kwargs.get('search_results'):
                result['identity']['evidenceUrls'].append('https://invented.example.com/character')
                result['urls'] += ['not-a-url', 'file:///etc/passwd', 'https://invented.example.com/character']
            return result, sources
        model.json = call
        result = self.run_research(model)
        self.assertEqual(result['entry']['metadata']['research']['identityUrls'], ['https://example.com/usopp'])
        self.assertTrue(any('not returned' in w for w in result['warnings']))

    def test_no_usable_search_evidence_does_not_replace_a_saved_entry(self):
        self.run_research()
        before = self.snapshot()
        with patch.object(research, 'fetch_document', side_effect=OSError('Blocked')), patch.object(research, 'source_excerpt', side_effect=ValueError('No English excerpt')):
            with self.assertRaisesRegex(ValueError, 'No usable English'):
                research.research('Usopp', root=self.root, model=FakeModel())
        self.assertEqual(before, self.snapshot())

    def test_tracking_query_differences_do_not_consume_extra_page_attempts(self):
        annotation = {'url': 'https://example.com/usopp?utm_source=wiki', 'title': 'Usopp', 'content': 'English evidence.'}
        current = {'documents': [document('https://example.com/usopp')]}
        with patch.object(research, 'fetch_document', side_effect=document) as fetch:
            documents, _, attempts = research.fetch_sources([annotation], current, 10, '2026-10-02T10:00:00+00:00')
        self.assertEqual(attempts, 1)
        self.assertEqual(fetch.call_count, 1)
        self.assertEqual(list(documents), ['https://example.com/usopp'])

    def test_refresh_preserves_passages_cited_by_verification_and_review_proposals(self):
        entry = self.run_research()['entry']
        source = entry['documents'][0]
        source_path = Path(entry['path']) / 'sources' / (source['id'] + '.md')
        meta, body = wiki.parse_markdown(source_path.read_text())
        body += '\n\n<!-- passage:p-verification -->\nEvidence used only for verification.\n'
        body += '\n\n<!-- passage:p-proposal -->\nEvidence used only for a proposed update.\n'
        source_path.write_text(wiki.markdown(meta, body))
        entry = wiki.read_entry(entry['key'], self.root)
        meta, body = wiki.parse_markdown(entry['markdown'])
        reference = {'source': f"sources/{source['id']}.md", 'passage': 'p-verification'}
        meta['abilities'][0]['verification'] = {'verdict': 'supported', 'evidence': [reference]}
        proposal = deepcopy(meta['abilities'][0])
        proposal['evidence'] = [{**reference, 'passage': 'p-proposal'}]
        meta['suggestedUpdates'] = [proposal]
        entry = wiki.save_entry(entry['key'], wiki.markdown(meta, body), entry['revision'], self.root)
        updated = wiki.collect({k: meta[k] for k in ['name', 'work', 'scope']}, [source['url']],
                               refresh=True, root=self.root, fetcher=document)['entry']
        refreshed = next(d for d in updated['documents'] if d['id'] == source['id'])
        retained = {p['id'] for p in refreshed['retainedPassages']}
        self.assertTrue({'p-verification', 'p-proposal'} <= retained)

    def test_generated_roleplay_biographies_are_excluded_from_model_evidence(self):
        entry = self.run_research()['entry']
        source = deepcopy(entry['documents'][0])
        source['url'] = 'https://daddyjim.ai/one-piece/character/Monkey-D-Luffy'
        source['id'] = 'generated-roleplay'
        entry['documents'].append(source)
        context, allowed, _ = research.context_for(entry)
        self.assertFalse(any('daddyjim' in d['url'] for d in context['sources']))
        self.assertFalse(any(s == 'sources/generated-roleplay.md' for s, _ in allowed))

    def test_provider_budget_search_and_secret_handling(self):
        response = {'choices': [{'message': {'content': '{"ok":true}', 'annotations': []}, 'finish_reason': 'stop'}],
                    'usage': {'total_tokens': 4, 'cost': .002}}
        from unittest.mock import MagicMock
        stream = MagicMock()
        stream.__enter__.return_value.read.return_value = json.dumps(response).encode()
        with patch('src.model.configuration', return_value={'OPENROUTER_API_KEY': 'private-test-key', 'OPENROUTER_MODEL': 'test/model'}), patch('src.model.urlopen', return_value=stream) as request:
            model = Model()
            result, _ = model.json('Instructions', {'query': 'Usopp'}, tokens=100, search_results=5)
            sent = request.call_args.args[0]
            body = json.loads(sent.data)
            self.assertEqual(body['plugins'][0]['max_results'], 5)
            self.assertEqual(body['max_tokens'], 100)
            self.assertNotIn('private-test-key', sent.data.decode())
            self.assertTrue(result['ok'])
            with self.assertRaisesRegex(ValueError, 'character budget'):
                model.json('Instructions', {'query': 'x' * 100_001}, tokens=100)
            self.assertEqual(request.call_count, 1)
