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
