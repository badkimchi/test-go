import React from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import {loginInfoStore} from "../lib/stores/loginInfoStore.ts";
import {useNavigate} from "react-router-dom";
import {APIAuth} from "../lib/api/APIAuth.tsx";
import {LoginInfo} from "../lib/models/loginInfo.ts";
import {APIAccount} from "../lib/api/APIAccount.tsx";
import {useGoogleLogin} from "@react-oauth/google";
import {Button} from "@chakra-ui/react";

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
    // const code = "ya29.a0AfB_byCsk3YvQFSPh6Sz71-m0SHWUrLNUev2lzkac1qw02d0S-PRdRGw6QbijM7jdeSn-LrBQXrPa20mQHCBnX9sa6pGJUUiC5zmD8cjT65J0OQVV5Wp66zl1d36McnGN8eQMTmGRF75JZBHUjc57_O9pppaP3hINFCQaCgYKAfASARISFQHGX2Mia1b39IO-5KJy-2kVfry1dg0171"
    // axios
    //     .get('https://www.googleapis.com/oauth2/v3/userinfo', {
    //         headers: {Authorization: `Bearer ${code}`},
    //     })
    //     .then(res => console.log(res.data));

    const googleLogin = useGoogleLogin({
        onSuccess: async tokenResponse => {
            signIn(tokenResponse.code);
            // console.log(tokenResponse);
            // // fetching userinfo can be done on the client or the server
            // const userInfo = await axios
            //     .get('https://www.googleapis.com/oauth2/v3/userinfo', {
            //         headers: {Authorization: `Bearer ${tokenResponse.access_token}`},
            //     })
            //     .then(res => res.data);
            //
            // console.log(userInfo);
        },
        flow: 'auth-code',
    });

    return (
        <AppLayout>
            <>
                <Button onClick={() => {
                    googleLogin()
                }}> Google Login </Button>
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
