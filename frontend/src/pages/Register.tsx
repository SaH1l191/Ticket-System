import { type FormEvent } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useRecoilState, useSetRecoilState } from 'recoil';
import { api } from '../api';
import { tokenState, registerEmailState, registerPasswordState, registerErrorState } from '../store';

export default function Register() {
  const [email, setEmail] = useRecoilState(registerEmailState);
  const [password, setPassword] = useRecoilState(registerPasswordState);
  const [error, setError] = useRecoilState(registerErrorState);
  const setToken = useSetRecoilState(tokenState);
  const navigate = useNavigate();

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError('');
    try {
      await api.register(email, password);
      const data = await api.login(email, password);
      setToken(data.token!);
      navigate('/tickets');
    } catch (err: unknown) {
      setError((err as { error?: string }).error || 'Registration failed');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center px-4">
      <div className="w-full max-w-sm">
        <h1 className="text-2xl font-semibold text-text mb-8 text-center">Create account</h1>
        {error && (
          <p className="text-danger text-sm mb-4 text-center">{error}</p>
        )}
        <form onSubmit={handleSubmit} className="space-y-3">
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-4 py-3 bg-surface border border-border rounded-xl text-text placeholder:text-muted focus:outline-none focus:border-accent transition-colors"
            required
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-4 py-3 bg-surface border border-border rounded-xl text-text placeholder:text-muted focus:outline-none focus:border-accent transition-colors"
            required
          />
          <button
            type="submit"
            className="w-full bg-accent text-white py-3 rounded-xl font-medium hover:bg-accent-hover transition-colors cursor-pointer"
          >
            Create account
          </button>
        </form>
        <p className="text-center mt-6 text-sm text-muted">
          Have an account?{' '}
          <Link to="/login" className="text-accent hover:text-accent-hover transition-colors">Sign in</Link>
        </p>
      </div>
    </div>
  );
}
