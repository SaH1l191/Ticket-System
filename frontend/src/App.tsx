import { RecoilRoot } from 'recoil';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { useRecoilValue } from 'recoil';
import { isAuthenticated } from './store';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import NewTicket from './pages/NewTicket';
import TicketDetail from './pages/TicketDetail';
import type { ReactNode } from 'react';

function ProtectedRoute({ children }: { children: ReactNode }) {
  const auth = useRecoilValue(isAuthenticated);
  return auth ? <>{children}</> : <Navigate to="/login" />;
}

function App() {
  return (
    <RecoilRoot>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/tickets" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
          <Route path="/tickets/new" element={<ProtectedRoute><NewTicket /></ProtectedRoute>} />
          <Route path="/tickets/:id" element={<ProtectedRoute><TicketDetail /></ProtectedRoute>} />
          <Route path="*" element={<Navigate to="/tickets" />} />
        </Routes>
      </BrowserRouter>
    </RecoilRoot>
  );
}

export default App;
