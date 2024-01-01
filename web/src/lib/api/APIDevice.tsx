import {APIBase} from './base/APIBase.tsx';
import {APIConfig} from './base/conf.tsx';
import {AxiosRequestConfig} from "axios";

const publicApiPath = {...APIConfig}
publicApiPath.baseURL = publicApiPath.baseURL.substring(0, publicApiPath.baseURL.length - "api/".length)

export class APIDevice extends APIBase {

    constructor(base: AxiosRequestConfig) {
        super(base);
    }

    public test(): Promise<string> {
        return this.get<string>('/public/test')
            .then((response) => {
                const {data} = response;
                return data?.data;
            })
            .catch((error) => {
                throw error;
            });
    }
}

export default new APIDevice(publicApiPath);

