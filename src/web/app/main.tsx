import { createRoot } from 'react-dom/client';
import { App } from './app.js';
import './styles.css';
import { TooltipProvider } from '../ui/tooltip.js';

createRoot(document.getElementById('root')!).render(
  <TooltipProvider>
    <App />
  </TooltipProvider>,
);
