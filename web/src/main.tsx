import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {QueryClient, QueryClientProvider} from '@tanstack/react-query'
import {GoogleOAuthProvider} from "@react-oauth/google";

const qClient = new QueryClient();

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <GoogleOAuthProvider clientId="773325553700-oluqkagk36js85vlqh55dselui6dvpar.apps.googleusercontent.com">
            <QueryClientProvider client={qClient}>
                <App/>
            </QueryClientProvider>
        </GoogleOAuthProvider>
    </React.StrictMode>,
)
