import { type FormEvent } from 'react';
import { useNavigate } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { api } from '../api';
import { newTicketTitleState, newTicketDescriptionState, newTicketErrorState } from '../store';

export default function NewTicket() {
  const [title, setTitle] = useRecoilState(newTicketTitleState);
  const [description, setDescription] = useRecoilState(newTicketDescriptionState);
  const [error, setError] = useRecoilState(newTicketErrorState);
  const navigate = useNavigate();

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError('');
    try {
      await api.createTicket(title, description);
      navigate('/tickets');
    } catch (err: unknown) {
      setError((err as { error?: string }).error || 'Failed to create ticket');
    }
  };

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
        <h1 className="text-2xl font-semibold text-text mb-8">New ticket</h1>
        {error && <p className="text-danger text-sm mb-4">{error}</p>}
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-text mb-1.5">Title</label>
            <input
              type="text"
              placeholder="What's the issue?"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full px-4 py-3 bg-surface border border-border rounded-xl text-text placeholder:text-muted focus:outline-none focus:border-accent transition-colors"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-text mb-1.5">Description</label>
            <textarea
              placeholder="Add details..."
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              rows={4}
              className="w-full px-4 py-3 bg-surface border border-border rounded-xl text-text placeholder:text-muted focus:outline-none focus:border-accent transition-colors resize-none"
            />
          </div>
          <button
            type="submit"
            className="bg-accent text-white px-6 py-2.5 rounded-xl font-medium hover:bg-accent-hover transition-colors cursor-pointer"
          >
            Create
          </button>
        </form>
      </main>
    </div>
  );
}
