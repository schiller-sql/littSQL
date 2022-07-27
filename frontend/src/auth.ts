import { writable, type Readable, type Writable } from "svelte/store";
import { request } from "./util/http_util";

export enum UserType {
  teacher,
  student,
}

interface User {
  token: string;
  type: UserType;
}

export type AuthState =
  | ({ status: "logged_in" } & User)
  | { status: "autologin_loading" }
  | { status: "logged_out" };

function userTypeFromString(rawUserType: string | null): UserType | null {
  switch (rawUserType) {
    case "teacher":
      return UserType.teacher;
    case "student":
      return UserType.student;
  }
  return null;
}

export function userTypeToString(userType: UserType): string {
  switch (userType) {
    case UserType.teacher:
      return "teacher";
    case UserType.student:
      return "student";
  }
}

const localStorageTokenKey = "token";
const localStorageUserTypeKey = "user_type";

function deleteUserFromLocalStorage() {
  localStorage.removeItem(localStorageTokenKey);
  localStorage.removeItem(localStorageUserTypeKey);
}

function writeUserToLocalStorage(user: User) {
  localStorage.setItem(localStorageTokenKey, user.token);
  localStorage.setItem(localStorageUserTypeKey, userTypeToString(user.type));
}

function getUserFromLocalStorage(): User | undefined {
  const token: string | null = localStorage.getItem(localStorageTokenKey);
  const rawUserType: string | null = localStorage.getItem(
    localStorageUserTypeKey
  );
  const userType: UserType | null = userTypeFromString(rawUserType);
  if (token === null || userType === null) {
    deleteUserFromLocalStorage();
    return undefined;
  }
  return {
    token,
    type: userType,
  };
}

export interface AuthStore extends Readable<AuthState> {
  getToken(): string;
  getUserType(): UserType;
  logIn(token: string, userType: UserType): void;
  logOut(): void;
}

// TODO: remove need to always call refresh_token, by looking at the expire date of the token
// TODO: timer to autorefresh after refresh time (possibly based on exp, instead of fixed time)
export function createAuthStore(): AuthStore {
  let lastToken: string | undefined;
  let lastUserType: UserType | undefined;
  let initialAuthState: AuthState;
  const user = getUserFromLocalStorage();
  if (!user) {
    initialAuthState = { status: "logged_out" };
  } else {
    initialAuthState = { status: "autologin_loading" };
  }
  const w: Writable<AuthState> = writable(initialAuthState);
  function logIn(token: string, userType: UserType) {
    w.set({ status: "logged_in", token, type: userType });
    lastToken = token;
    lastUserType = userType;
    writeUserToLocalStorage({ token, type: userType });
  }
  function logOut() {
    w.set({ status: "logged_out" });
    deleteUserFromLocalStorage();
  }
  function getToken(): string {
    return lastToken!;
  }
  function getUserType(): UserType {
    return lastUserType!;
  }
  async function autologin() {
    try {
      const { token } = await request("auth/refresh_token", {
        returnJson: true,
        headers: {
          Authorization: "Bearer " + user.token,
        },
        errorAsRes: true,
      });
      logIn(token, user.type);
    } catch (res) {
      if (res.status === "401") {
        deleteUserFromLocalStorage();
      }
      logOut();
    }
  }
  if (user) {
    autologin();
  }
  return {
    ...w,
    logIn,
    logOut,
    getToken,
    getUserType,
  };
}
