import React, {useEffect} from 'react';
import {AppLayout} from '../components/layouts/AppLayout';

export const Home: React.FC = () => {
    useEffect(() => {
        console.log('here')
    }, [])

    return (
        <AppLayout>
            <>
                Home Page
            </>
        </AppLayout>
    );
};
