export interface Token {
  token: string;
  expiration: Date;
  refreshToken: RefreshToken;
}

export interface RefreshToken {
  token: string;
  expiration: string;
}