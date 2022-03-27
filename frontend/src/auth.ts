import { writable, type Writable } from "svelte/store";

export enum UserType {
  teacher,
  student,
}

export interface User {
  jwt: string;
  type: UserType;
}

export const authStore: Writable<User | null> = writable();
