import { writable, type Readable } from "svelte/store";
import { checkForError } from "../util/db_util";

export type SqlStatus =
  | {
      status: "loading";
    }
  | { status: "ok" }
  | { status: "error"; error: string };

interface SqlStatusStore extends Readable<SqlStatus> {
  sqlUpdate(sql: string, databaseSql?: string): void;
}

export function createSqlStatusStore(
  millisecondsTillCheck: number
): SqlStatusStore {
  let lastSql: string | undefined;
  const w = writable<SqlStatus>({ status: "loading" });
  let lastTimeoutId;
  function checkSqlStatusStore(sql: string, databaseSql?: string) {
    const error = checkForError(sql, databaseSql);
    if (error !== undefined) {
      w.set({ status: "error", error });
    } else {
      w.set({ status: "ok" });
    }
  }
  function sqlUpdate(sql: string, databaseSql?: string) {
    if (sql === lastSql) return;
    lastSql = sql;
    w.set({ status: "loading" });
    clearTimeout(lastTimeoutId);
    lastTimeoutId = setTimeout(
      checkSqlStatusStore,
      millisecondsTillCheck,
      sql,
      databaseSql
    );
  }
  return {
    sqlUpdate,
    subscribe: w.subscribe,
  };
}
