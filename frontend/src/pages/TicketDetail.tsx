import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { api } from '../api';
import type { Ticket } from '../api';

const statusStyles: Record<string, string> = {
  open: 'bg-amber-50 text-amber-700',
  in_progress: 'bg-blue-50 text-blue-700',
  closed: 'bg-green-50 text-green-700',
};

const nextStatus: Partial<Record<Ticket['status'], Ticket['status']>> = {
  open: 'in_progress',
  in_progress: 'closed',
};

const nextLabel: Record<string, string> = {
  open: 'Start progress',
  in_progress: 'Mark closed',
};

export default function TicketDetail() {
  const { id } = useParams<{ id: string }>();
  const [ticket, setTicket] = useState<Ticket | null>(null);
  const [loading, setLoading] = useState(true);
  const [updating, setUpdating] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    if (!id) return;
    api.getTicket(id)
      .then(setTicket)
      .catch(() => navigate('/tickets'))
      .finally(() => setLoading(false));
  }, [id, navigate]);

  const handleStatusUpdate = async () => {
    if (!ticket) return;
    const next = nextStatus[ticket.status];
    if (!next) return;
    setUpdating(true);
    setError('');
    try {
      const updated = await api.updateStatus(ticket.id, next);
      setTicket(updated);
    } catch (err: unknown) {
      setError((err as { error?: string }).error || 'Failed to update');
    } finally {
      setUpdating(false);
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <p className="text-muted text-sm">Loading...</p>
      </div>
    );
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
            &larr; Back
          </button>
        </div>
      </header>

      <main className="max-w-2xl mx-auto px-6 py-10">
        <div className="flex items-start justify-between gap-4 mb-8">
          <h1 className="text-2xl font-semibold text-text">{ticket.title}</h1>
          <span className={`shrink-0 text-xs font-medium px-2.5 py-0.5 rounded-full ${statusStyles[ticket.status]}`}>
            {ticket.status.replace('_', ' ')}
          </span>
        </div>

        <p className="text-muted leading-relaxed mb-8">
          {ticket.description || 'No description provided.'}
        </p>

        <div className="text-xs text-muted/60 mb-8 space-y-0.5">
          <p>Created {new Date(ticket.created_at).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}</p>
          <p>Updated {new Date(ticket.updated_at).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}</p>
        </div>

        {error && <p className="text-danger text-sm mb-4">{error}</p>}

        {next ? (
          <button
            onClick={handleStatusUpdate}
            disabled={updating}
            className="bg-accent text-white px-5 py-2.5 rounded-xl font-medium hover:bg-accent-hover disabled:opacity-40 transition-colors cursor-pointer"
          >
            {updating ? 'Updating...' : nextLabel[ticket.status]}
          </button>
        ) : (
          <span className="text-sm text-green-600 font-medium">Completed</span>
        )}
      </main>
    </div>
  );
}
