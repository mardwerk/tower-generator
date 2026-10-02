import type { CharacterSource } from '../types.js';
import { Disclosure } from '../ui/disclosure.js';
import { safeUrl } from '../ui/utils.js';

export function SourceEvidence({ source }: { source: CharacterSource }) {
  return (
    <div className="min-w-0 text-xs">
      {source.notes && (
        <p className="my-3 whitespace-pre-wrap text-muted-foreground">{source.notes}</p>
      )}
      {source.documents.map((document) => (
        <Disclosure key={document.id} bare title={document.title}>
          {document.url && safeUrl(document.url) && (
            <a href={document.url} target="_blank" rel="noreferrer" className="break-all">
              {document.url}
            </a>
          )}
          <p className="my-2 text-muted-foreground">
            {document.access === 'retrieved'
              ? `Retrieved ${document.retrievedAt?.slice(0, 10)}`
              : 'Supplied evidence'}
            {document.excerpt ? ' · Retained excerpt' : ''}
          </p>
          <pre className="max-h-[380px] overflow-auto font-sans text-xs leading-relaxed whitespace-pre-wrap [overflow-wrap:anywhere]">
            {document.text}
          </pre>
        </Disclosure>
      ))}
    </div>
  );
}
