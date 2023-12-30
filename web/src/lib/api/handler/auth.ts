// import { AxiosResponse } from 'axios';
// import { request } from '../setupAxios';
// import { ChangePasswordInput, LoginDTO, RegisterDTO, ResetPasswordInput } from '../dtos/AuthInput';
// import { Account } from '../../models/user';
//
// export const register = (body: RegisterDTO): Promise<AxiosResponse<Account>> => request.post('/user/register', body);
//
// export const login = (body: LoginDTO): Promise<AxiosResponse<Account>> => request.post('/user/login', body);
//
// export const logout = (): Promise<AxiosResponse> => request.post('/user/logout');
//
// export const forgotPassword = (email: string): Promise<AxiosResponse<boolean>> =>
//   request.post('/user/forgot-password', { email });
//
// export const changePassword = (body: ChangePasswordInput): Promise<AxiosResponse> =>
//   request.put('/user/change-password', body);
//
// export const resetPassword = (body: ResetPasswordInput): Promise<AxiosResponse<Account>> =>
//   request.post('/user/reset-password', body);
