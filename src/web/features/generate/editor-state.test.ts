// The Request editor's required concepts: kept when listed, left out when empty.
import type { LabRequest } from '../../api/contract.js';
import { emptyRequest } from '../../api/artifacts.js';
import { editRequest, readEditor } from './editor-state.js';

function equal(actual: string, expected: string, what: string) {
  if (actual !== expected) throw new Error(`${what}: got ${actual}, expected ${expected}`);
}

export const tests: Record<string, () => Promise<void>> = {
  'required concepts round-trip and an empty list is left out': async () => {
    const request: LabRequest = {
      ...emptyRequest(),
      requiredConcepts: [{ name: 'Gear 4', reason: 'A signature form.' }],
    };
    const input = editRequest(request);
    equal(
      JSON.stringify(readEditor(input).requiredConcepts),
      '[{"name":"Gear 4","reason":"A signature form."}]',
      'listed',
    );
    const cleared = readEditor({ ...input, requiredConcepts: '[]' });
    equal(String('requiredConcepts' in cleared), 'false', 'an empty list');
    const blank = readEditor({ ...input, requiredConcepts: '' });
    equal(String('requiredConcepts' in blank), 'false', 'a blank field');
    equal(String('requiredConcepts' in readEditor(editRequest(emptyRequest()))), 'false', 'none');
  },
};
