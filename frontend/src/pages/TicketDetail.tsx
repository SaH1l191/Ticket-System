import { useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { api } from '../api';
import type { Ticket } from '../api';
import { ticketDetailState, ticketDetailLoadingState, ticketDetailUpdatingState, ticketDetailErrorState } from '../store';

const nextStatus: Partial<Record<Ticket['status'], Ticket['status']>> = {
  open: 'in_progress',
  in_progress: 'closed',
};

const statusColors: Record<Ticket['status'], string> = {
  open: 'bg-amber-50 text-amber-700',
  in_progress: 'bg-blue-50 text-blue-700',
  closed: 'bg-green-50 text-green-700',
};

export default function TicketDetail() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const [ticket, setTicket] = useRecoilState(ticketDetailState);
  const [loading, setLoading] = useRecoilState(ticketDetailLoadingState);
  const [updating, setUpdating] = useRecoilState(ticketDetailUpdatingState);
  const [error, setError] = useRecoilState(ticketDetailErrorState);

  useEffect(() => {
    if (!id) return;

    api.getTicket(id)
      .then(setTicket)
      .catch(() => navigate('/tickets'))
      .finally(() => setLoading(false));
  }, [id, navigate, setLoading, setTicket]);

  const updateStatus = async () => {
    if (!ticket) return;

    const next = nextStatus[ticket.status];
    if (!next) return;

    setUpdating(true);
    setError('');

    try {
      const updated = await api.updateStatus(ticket.id, next);
      setTicket(updated);
    } catch {
      setError('Failed to update ticket');
    } finally {
      setUpdating(false);
    }
  };

  if (loading) {
    return <div className="p-8 text-muted text-sm">Loading...</div>;
  }

  if (!ticket) return null;

  const next = nextStatus[ticket.status];

  return (
    <div className="min-h-screen">
      <header className="border-b border-border bg-surface">
        <div className="max-w-2xl mx-auto px-6 py-4">
          <button
            onClick={() => navigate('/tickets')}
            className="text-sm text-muted hover:text-text transition-colors cursor-pointer"
          >
            ← Back
          </button>
        </div>
      </header>

      <main className="max-w-2xl mx-auto px-6 py-10">
        <div className="bg-surface border border-border rounded-xl p-6">
          
          <div className="flex items-center justify-between gap-4 mb-6">
            <h1 className="text-2xl font-semibold text-text">
              {ticket.title}
            </h1>

            <span
              className={`shrink-0 text-xs font-medium px-2.5 py-0.5 rounded-full capitalize ${
                statusColors[ticket.status]
              }`}
            >
              {ticket.status.replace('_', ' ')}
            </span>
          </div>

          <p className="text-muted mb-6">
            {ticket.description || 'No description provided.'}
          </p>

          <div className="text-sm text-muted space-y-1 mb-6">
            <p>
              Created:{' '}
              {new Date(ticket.created_at).toLocaleDateString()}
            </p>
            <p>
              Updated:{' '}
              {new Date(ticket.updated_at).toLocaleDateString()}
            </p>
          </div>

          {error && (
            <p className="text-danger text-sm mb-4">
              {error}
            </p>
          )}

          {next ? (
            <button
              onClick={updateStatus}
              disabled={updating}
              className="px-4 py-2 bg-accent text-white rounded-xl font-medium hover:bg-accent-hover transition-colors cursor-pointer disabled:opacity-50"
            >
              {updating
                ? 'Updating...'
                : ticket.status === 'open'
                  ? 'Start progress'
                  : 'Mark closed'}
            </button>
          ) : (
            <span className="text-sm font-medium text-accent">
              Completed
            </span>
          )}
        </div>
      </main>
    </div>
  );
}
