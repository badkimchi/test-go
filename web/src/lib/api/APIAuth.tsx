import {APIBase} from './base/APIBase.tsx';
import {APIConfig} from './base/conf.tsx';
import {AxiosRequestConfig} from "axios";
import {Token} from "../models/token.ts";

const publicApiPath = {...APIConfig}
publicApiPath.baseURL = publicApiPath.baseURL.substring(0, publicApiPath.baseURL.length - "api/".length)

export class API extends APIBase {

    constructor(base: AxiosRequestConfig) {
        super(base);
    }

    public login(token: string): Promise<Token> {
        return this.post<Token>('/auth/login', {
            token: token
        })
            .then((response) => {
                const {data} = response;
                return data?.data;
            })
            .catch((error) => {
                throw error;
            });
    }
}

const api = new API(publicApiPath);
export {api as APIAuth};
