import { writable, type Readable } from "svelte/store";
import { newDatabase } from "../database";
import type { QueryExecResult } from "../sql-js/sql-wasm";

export function execStatementOnDatabase(
  databaseSql: string,
  statementSql: string
): QueryExecResult[] | string {
  const db = newDatabase();
  try {
    db.run(databaseSql);
  } catch (e) {
    return "Error in database: " + e.toString();
  }
  let result: QueryExecResult[];
  try {
    result = db.exec(statementSql);
  } catch (e) {
    return e.toString();
  }
  db.close();
  return result;
}

export function getAllTables(
  databaseSql: string
): Map<string, QueryExecResult | undefined> {
  const allTables = new Map<string, QueryExecResult | undefined>();
  const db = newDatabase();
  db.run(databaseSql);
  const tableNamesQueryResult = db.exec(
    "select distinct tbl_name from sqlite_master order by tbl_name"
  )[0];
  if (tableNamesQueryResult === undefined) {
    return allTables;
  }
  const tableNames = tableNamesQueryResult.values.map(
    (row) => row[0]
  ) as string[];
  for (const tableName of tableNames) {
    const tableValues = db.exec(`select * from ${tableName}`)[0];
    allTables.set(tableName, tableValues);
  }
  return allTables;
}

export function checkForError(
  sql: string,
  databaseSql?: string
): string | undefined {
  const db = newDatabase();
  try {
    if (databaseSql) {
      db.run(databaseSql);
    }
    db.run(sql);
  } catch (e) {
    return e;
  }
}

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
