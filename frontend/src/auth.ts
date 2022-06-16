import { writable, type Writable } from "svelte/store";

export enum UserType {
  teacher,
  student,
}

export interface User {
  token: string;
  type: UserType;
}

export const DEFAULT_URL = "api/";

export const authStore: Writable<User | null> = writable();

export function fetchWithToken(url: string, method: string, token: string) {
  return requestWithToken(url, method, token).then((res) => res.json());
}

export function requestWithToken(url: string, method: string, token: string) {
  return fetch(DEFAULT_URL + url, {
    method: method,
    headers: { Authorization: `Bearer ${token}` },
  });
}
