import {create} from 'zustand';
import {persist} from 'zustand/middleware';
import {LoginInfo} from '../models/loginInfo.ts';

type AccountState = {
    current: LoginInfo | null;
    setUser: (account: LoginInfo) => void;
    logout: () => void;
};

export const userStore = create(
    persist<AccountState>(
        (set, _) => ({
            current: null,
            setUser: (account) => set({current: account}),
            logout: () => set({current: null}),
        }),
        {
            name: 'account-storage',
        }
    )
);
