import { BookOpen, Library, Plus, Settings, UserRound } from 'lucide-react';
import { IconButton } from '../ui/icon-button.js';
import { cn } from '../ui/utils.js';

// Retain the pre-cleanup Lab's centered navigation and responsive layout.
export function Topbar({
  view,
  busy,
  onWiki,
  onCreate,
}: {
  view: 'wiki' | 'create';
  busy: boolean;
  onWiki: () => void;
  onCreate: () => void;
}) {
  return (
    <header className="sticky top-0 z-40 grid min-h-[70px] grid-cols-[minmax(0,1fr)_auto] items-center gap-1.5 border-b border-border bg-canvas px-3.5 py-2.5 sm:grid-cols-[minmax(0,1fr)_auto_minmax(0,1fr)] sm:gap-3 sm:px-6 sm:py-0">
      <a
        href="#"
        className="col-start-1 row-start-1 flex items-center gap-2.5 text-foreground hover:no-underline"
        onClick={(event) => {
          event.preventDefault();
          if (!busy) onCreate();
        }}
      >
        <img src="/mardwerk.png" alt="Mardwerk" width={30} height={36} className="object-contain" />
        <span className="font-mono text-[13px] font-semibold tracking-tight">tower-generator</span>
      </a>
      <nav
        aria-label="Workspace"
        className="col-span-full row-start-2 flex items-center justify-center gap-3.5 sm:col-span-1 sm:col-start-2 sm:row-start-1 sm:gap-2.5"
      >
        <IconButton
          label="Create"
          disabled={busy}
          aria-current={view === 'create' ? 'page' : undefined}
          className={cn(view === 'create' && 'bg-accent text-foreground')}
          onClick={onCreate}
        >
          <Plus className="size-[19px]" />
        </IconButton>
        <IconButton label="Library" disabled>
          <Library className="size-[19px]" />
        </IconButton>
        <IconButton
          label="Wiki"
          disabled={busy}
          aria-current={view === 'wiki' ? 'page' : undefined}
          className={cn(view === 'wiki' && 'bg-accent text-foreground')}
          onClick={onWiki}
        >
          <BookOpen className="size-[19px]" />
        </IconButton>
        <IconButton label="Generations" disabled>
          <UserRound className="size-[19px]" />
        </IconButton>
      </nav>
      <div className="col-start-2 row-start-1 justify-self-end sm:col-start-3">
        <IconButton label="Settings" disabled>
          <Settings className="size-[19px]" />
        </IconButton>
      </div>
    </header>
  );
}
