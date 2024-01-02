import {Token} from "./token.ts";

export class LoginInfo {
  id: string = '';
  username: string = '';
  email: string = '';
  image: string = '';
  authToken: Token = new Token();
}
