// import create from 'zustand';
// import { persist } from 'zustand/middleware';
// import { Account } from '../models/user';
//
// type AccountState = {
//   current: Account | null;
//   setUser: (user: Account) => void;
//   logout: () => void;
// };
//
// export const userStore = create(
//   persist<AccountState>(
//     (set, _) => ({
//       current: null,
//       setUser: (user) => set({ current: user }),
//       logout: () => set({ current: null }),
//     }),
//     {
//       name: 'user-storage',
//     }
//   )
// );
