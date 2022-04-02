import { writable, type Writable } from "svelte/store";

export enum UserType {
  teacher,
  student,
}

export interface User {
  jwt: string;
  type: UserType;
}

export const DEFAULT_URL = 'http://localhost:8080/';


export const authStore: Writable<User | null> = writable();
