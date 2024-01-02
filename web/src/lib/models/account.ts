import {Token} from "./token.ts";

export interface Account {
  id: string;
  username: string;
  email: string;
  image: string;
  authToken: Token;
}
