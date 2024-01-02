import React from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import {loginInfoStore} from "../lib/stores/loginInfoStore.ts";
import {useNavigate} from "react-router-dom";
import APIAuth from "../lib/api/APIAuth.tsx";
import {LoginInfo} from "../lib/models/loginInfo.ts";
import APIAccount from "../lib/api/APIAccount.tsx";

export const SignIn: React.FC = () => {
    const navigate = useNavigate();
    const setUser = loginInfoStore(state => state.setUser);
    const signIn = () => {
        APIAuth.login()
            .then(resp => {
                const user = new LoginInfo();
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
