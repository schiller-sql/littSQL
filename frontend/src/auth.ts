import { writable, type Writable } from "svelte/store";

export enum UserType {
  teacher,
  student,
}

export interface User {
  token: string;
  type: UserType;
}

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

export const DEFAULT_URL = "api/";

const localStorageTokenKey = "token";
const localStorageUserTypeKey = "user_type";

// TODO: token should first be refreshed before being used
function getUserFromLocalStorage(): User | null {
  const token: string | null = localStorage.getItem(localStorageTokenKey);
  const rawUserType: string | null = localStorage.getItem(
    localStorageUserTypeKey
  );
  const userType: UserType | null = userTypeFromString(rawUserType);
  if (token === null || userType === null) {
    localStorage.clear();
    return null;
  }
  return {
    token,
    type: userType,
  };
}

function writeUserToLocalStorage(user: User) {
  if (user === null) {
    localStorage.clear();
  } else {
    localStorage.setItem(localStorageTokenKey, user.token);
    localStorage.setItem(localStorageUserTypeKey, userTypeToString(user.type));
  }
}

export const authStore: Writable<User | null> = writable(
  getUserFromLocalStorage()
);

authStore.subscribe((user) => {
  writeUserToLocalStorage(user);
});

export async function fetchWithToken(
  url: string,
  method: string,
  token: string,
  body?: Object
) {
  return requestWithToken(url, method, token, body).then((res) => res.json());
}

export async function requestWithToken(
  url: string,
  method: string,
  token: string,
  body?: Object
) {
  const res = await fetch(DEFAULT_URL + url, {
    method: method,
    body: JSON.stringify(body),
    headers: { Authorization: `Bearer ${token}` },
  });
  if (!res.ok) {
    throw (await res.json()).error;
  }
  return res;
}
