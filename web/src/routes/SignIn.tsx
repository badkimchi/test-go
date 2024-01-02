import React from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import {userStore} from "../lib/stores/userStore.ts";
import {useNavigate} from "react-router-dom";
import APIAuth from "../lib/api/APIAuth.tsx";
import {Account} from "../lib/models/account.ts";
import APIAccount from "../lib/api/APIAccount.tsx";

export const SignIn: React.FC = () => {
    const navigate = useNavigate();
    const setUser = userStore(state => state.setUser);
    const signIn = () => {
        APIAuth.login()
            .then(resp => {
                const user = new Account();
                user.authToken = resp;
                setUser(user);
                console.log(user);
                navigate('/');
            })
            .catch(err => console.error(err));
    }


    return (
        <AppLayout>
            <>
                <button onClick={signIn}>Login</button>
                <button onClick={() => {
                    APIAccount.getAccount()
                        .then((resp) => {
                            console.log(resp);
                        })
                        .catch(err => console.error(err));
                }}>Test</button>
            </>
        </AppLayout>
    );
};
