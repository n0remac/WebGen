import {
    createConnectTransport,
  } from "@bufbuild/connect-web";

export const baseURL = process.env.BASE_URL;

export const transport = createConnectTransport({
  baseUrl: `${baseURL}` || 'error',
  // credentials: "include",
});
