import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useSetRecoilState } from 'recoil';
import { api } from '../api';
import type { Ticket } from '../api';
import { tokenState } from '../store';

const statusStyles: Record<string, string> = {
  open: 'bg-amber-50 text-amber-700',
  in_progress: 'bg-blue-50 text-blue-700',
  closed: 'bg-green-50 text-green-700',
};

export default function Dashboard() {
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [loading, setLoading] = useState(true);
  const setToken = useSetRecoilState(tokenState);
  const navigate = useNavigate();

  useEffect(() => {
    api.listTickets()
      .then(setTickets)
      .catch(() => navigate('/login'))
      .finally(() => setLoading(false));
  }, [navigate]);

  return (
    <div className="min-h-screen">
      <header className="border-b border-border bg-surface">
        <div className="max-w-2xl mx-auto px-6 py-4 flex justify-between items-center">
          <h1 className="text-lg font-semibold text-text">Tickets</h1>
          <div className="flex items-center gap-4">
            <Link
              to="/tickets/new"
              className="text-sm font-medium text-accent hover:text-accent-hover transition-colors"
            >
              + New
            </Link>
            <button
              onClick={() => { setToken(''); navigate('/login'); }}
              className="text-sm text-muted hover:text-text transition-colors cursor-pointer"
            >
              Sign out
            </button>
          </div>
        </div>
      </header>

      <main className="max-w-2xl mx-auto px-6 py-10">
        {loading ? (
          <p className="text-muted text-sm">Loading...</p>
        ) : tickets.length === 0 ? (
          <div className="text-center py-20">
            <p className="text-muted mb-2">No tickets yet</p>
            <Link to="/tickets/new" className="text-accent text-sm hover:text-accent-hover transition-colors">
              Create your first ticket
            </Link>
          </div>
        ) : (
          <div className="space-y-2">
            {tickets.map((t) => (
              <Link
                key={t.id}
                to={`/tickets/${t.id}`}
                className="block bg-surface border border-border rounded-xl p-5 hover:border-accent/40 transition-colors"
              >
                <div className="flex justify-between items-start gap-4">
                  <div className="min-w-0">
                    <h3 className="font-medium text-text truncate">{t.title}</h3>
                    {t.description && (
                      <p className="text-sm text-muted mt-1 line-clamp-1">{t.description}</p>
                    )}
                  </div>
                  <span className={`shrink-0 text-xs font-medium px-2.5 py-0.5 rounded-full ${statusStyles[t.status]}`}>
                    {t.status.replace('_', ' ')}
                  </span>
                </div>
              </Link>
            ))}
          </div>
        )}
      </main>
    </div>
  );
}
