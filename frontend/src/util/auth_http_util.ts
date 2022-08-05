import { authStore } from "../stores/global_stores";
import { request } from "./http_util";

export async function fetchWithAuthorization(
  url: string,
  method?: string,
  body?: Object,
  overrideToken?: string
) {
  return requestWithAuthorization(
    url,
    method ?? "get",
    body,
    overrideToken
  ).then((res) => res.json());
}

export async function requestWithAuthorization(
  url: string,
  method?: string,
  body?: Object,
  overrideToken?: string
) {
  return request(url, {
    method,
    body,
    headers: {
      Authorization: `Bearer ${overrideToken ?? authStore.getToken()}`,
    },
  });
}
