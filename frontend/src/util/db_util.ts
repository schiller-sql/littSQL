import { writable, type Readable } from "svelte/store";
import type { QueryExecResult } from "../sql-js/sql-wasm";
import { databaseStore } from "../stores/global_stores";

export function execStatementOnDatabase(
  databaseSql: string,
  statementSql: string
): QueryExecResult[] | string {
  const db = databaseStore.newDatabase();
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
  const db = databaseStore.newDatabase();
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
  const db = databaseStore.newDatabase();
  try {
    if (databaseSql) {
      db.run(databaseSql);
    }
    db.run(sql);
  } catch (e) {
    return e;
  }
}
