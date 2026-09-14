const API_BASE = '';

interface ApiError {
  status: number;
  error: string;
  message?: string;
}

interface AuthResponse {
  token?: string;
  message?: string;
  error?: string;
}

interface Ticket {
  id: string;
  title: string;
  description: string;
  status: 'open' | 'in_progress' | 'closed';
  created_at: string;
  updated_at: string;
}

async function request<T>(path: string, options: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('token');
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...(options.headers as Record<string, string>),
  };
  if (token) headers['Authorization'] = `Bearer ${token}`;

  const res = await fetch(`${API_BASE}${path}`, { ...options, headers });
  const data = await res.json();
  if (!res.ok) throw { status: res.status, ...data } as ApiError;
  return data as T;
}

export const api = {
  register: (email: string, password: string) =>
    request<AuthResponse>('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    }),

  login: (email: string, password: string) =>
    request<AuthResponse>('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    }),

  createTicket: (title: string, description: string) =>
    request<Ticket>('/tickets', {
      method: 'POST',
      body: JSON.stringify({ title, description }),
    }),

  listTickets: () => request<Ticket[]>('/tickets'),

  getTicket: (id: string) => request<Ticket>(`/tickets/${id}`),

  updateStatus: (id: string, status: string) =>
    request<Ticket>(`/tickets/${id}/status`, {
      method: 'PATCH',
      body: JSON.stringify({ status }),
    }),
};

export type { ApiError, AuthResponse, Ticket };
