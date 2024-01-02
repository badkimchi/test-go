import React, {useEffect, useState} from 'react';
import {AppLayout} from '../components/layouts/AppLayout';
import APIDevice from "../lib/api/APIDevice.tsx";

export const Home: React.FC = () => {
    const [data, setData] = useState<string>('none');
    useEffect(() => {
        console.log('here')
    }, [])

    return (
        <AppLayout>
            <>
                {data}
            </>
        </AppLayout>
    );
};
