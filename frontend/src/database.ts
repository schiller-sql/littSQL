import { get, writable, type Writable } from "svelte/store";
import initSqlJs from "./sql-js/sql-wasm";
import type { SqlJsStatic, Database } from "./sql-js/sql-wasm";

let databaseInitCalled = false;

export const databaseIsReadyStore: Writable<boolean> = writable(false);

databaseIsReadyStore.subscribe((value) => {
  if (!value) {
    console.log("sqlite: not ready yet");
  } else {
    console.log("sqlite: ready!");
  }
});

let sql: SqlJsStatic | undefined;

export function newDatabase(data?: ArrayLike<number> | Buffer): Database {
  const db = new sql!.Database(data);
  db.run("PRAGMA foreign_keys = ON");
  return db;
}

export async function initSqlite() {
  // await new Promise((r) => setTimeout(r, 2000));
  if (databaseInitCalled) return;
  databaseInitCalled = true;
  console.log("sqlite: loading...");
  sql = await initSqlJs({
    locateFile: (file) => {
      console.log(`sqlite: loading '${file}'...`);
      return `static/sql.js/${file}`;
    },
  });
  databaseIsReadyStore.set(true);
}
