import { atom, selector } from 'recoil';

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
