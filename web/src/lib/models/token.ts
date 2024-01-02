export class Token {
  token: string = '';
  expiration: Date = new Date();
  refreshToken: RefreshToken = new RefreshToken();
}

export class RefreshToken {
  token: string = '';
  expiration: Date = new Date();
}