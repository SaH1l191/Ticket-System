import { atom, selector } from 'recoil';
import type { Ticket } from './api';

export const tokenState = atom<string>({
  key: 'token',
  default: localStorage.getItem('token') || '',
  effects: [
    ({ onSet }) => {
      onSet((value: string) => {
        if (value) localStorage.setItem('token', value);
        else localStorage.removeItem('token');
      });
    },
  ],
});

export const isAuthenticated = selector<boolean>({
  key: 'isAuthenticated',
  get: ({ get }) => !!get(tokenState),
});

export const loginEmailState = atom<string>({ key: 'loginEmail', default: '' });
export const loginPasswordState = atom<string>({ key: 'loginPassword', default: '' });
export const loginErrorState = atom<string>({ key: 'loginError', default: '' });

export const registerEmailState = atom<string>({ key: 'registerEmail', default: '' });
export const registerPasswordState = atom<string>({ key: 'registerPassword', default: '' });
export const registerErrorState = atom<string>({ key: 'registerError', default: '' });

export const newTicketTitleState = atom<string>({ key: 'newTicketTitle', default: '' });
export const newTicketDescriptionState = atom<string>({ key: 'newTicketDescription', default: '' });
export const newTicketErrorState = atom<string>({ key: 'newTicketError', default: '' });

export const ticketListState = atom<Ticket[]>({ key: 'ticketList', default: [] });
export const ticketListLoadingState = atom<boolean>({ key: 'ticketListLoading', default: true });

export const ticketDetailState = atom<Ticket | null>({ key: 'ticketDetail', default: null });
export const ticketDetailLoadingState = atom<boolean>({ key: 'ticketDetailLoading', default: true });
export const ticketDetailUpdatingState = atom<boolean>({ key: 'ticketDetailUpdating', default: false });
export const ticketDetailErrorState = atom<string>({ key: 'ticketDetailError', default: '' });
