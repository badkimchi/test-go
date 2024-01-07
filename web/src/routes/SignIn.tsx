import React from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import {loginInfoStore} from "../lib/stores/loginInfoStore.ts";
import {useNavigate} from "react-router-dom";
import {APIAuth} from "../lib/api/APIAuth.tsx";
import {LoginInfo} from "../lib/models/loginInfo.ts";
import {APIAccount} from "../lib/api/APIAccount.tsx";
import {useGoogleLogin} from "@react-oauth/google";

export const SignIn: React.FC = () => {
    const navigate = useNavigate();
    const setUser = loginInfoStore(state => state.setUser);
    const signIn = (accessToken: string) => {
        console.log(accessToken)
        APIAuth.login(accessToken)
            .then(resp => {
                const user = new LoginInfo();
                user.authToken = resp;
                setUser(user);
                navigate('/');
            })
            .catch(err => console.error(err));
    }

    const googleLogin = useGoogleLogin({
        onSuccess: async tokenResponse => {
            signIn(tokenResponse.code);
        },
        // if we have implicit flow,
        // then we could skip exchanging auth code on the server side with access token.
        flow: 'auth-code',
    });

    return (
        <AppLayout>
            <>
                <button onClick={() => {
                    googleLogin()
                }}> Google Login </button>
                <button onClick={() => {
                    APIAccount.getAccount()
                        .then((resp) => {
                            console.log(resp);
                        })
                        .catch(err => {
                            console.error(err)
                        });
                }}>Test
                </button>
            </>
        </AppLayout>
    );
};
