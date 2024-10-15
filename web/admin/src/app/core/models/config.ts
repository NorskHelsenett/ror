export interface Config {
  auth: Auth;
  regex: Regex;
  rowsPerPage: number[];
  rows: number;
  rorApi: string;
  sse: SSE;
}

export interface Auth {
  issuer: string;
  clientId: string;
  redirectUri: string;
  scope: string;
  responseType: string;
  logoutUrl: string;
  postLogoutRedirectUri: string;
  requireHttps: boolean;
  strictDiscoveryDocumentValidation: boolean;
}

export interface Regex {
  forms: string;
}

export interface SSE {
  postfixUrl: string;
  timeout: number;
  method: string;
}
