import { apiUrl } from "../config";

export async function request(
  url: string,
  options?: {
    method?: string;
    body?: Object;
    headers?: HeadersInit;
    returnJson?: boolean;
    nonErrorStatusCodes?: number[];
    errorAsRes?: boolean;
  }
) {
  const res = await fetch(apiUrl + "/" + url, {
    method: options?.method ?? "get",
    body:
      options?.body === undefined ? undefined : JSON.stringify(options?.body),
    headers:
      options?.body === undefined
        ? {
            "Content-Type": "application/json",
            ...(options?.headers ?? {}),
          }
        : options?.headers,
  });
  if (
    !res.ok &&
    (!options?.nonErrorStatusCodes ||
      !options?.nonErrorStatusCodes.includes(res.status))
  ) {
    if (options?.errorAsRes) {
      throw res;
    }
    throw (await res.json()).error;
  }
  if (options?.returnJson) return await res.json();
  return res;
}
