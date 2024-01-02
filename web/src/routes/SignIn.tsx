import React from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import {userStore} from "../lib/stores/userStore.ts";
import {useNavigate} from "react-router-dom";
import APIAuth from "../lib/api/APIAuth.tsx";

export const SignIn: React.FC = () => {
    const navigate = useNavigate();
    const setUser = userStore(state => state.setUser);
    const signIn = () => {
        APIAuth.login()
            .then(resp => console.log(resp))
            .catch(err => console.error(err));

        setUser({
            id: '1',
            username: 'aloha',
            email: 'abcd',
            image: 'abcd',
            authToken: 'abcd',
        });
        navigate('/signin');
    }

    return (
        <AppLayout>
            <>
                <button onClick={signIn}>Login</button>
            </>
        </AppLayout>
    );
};
