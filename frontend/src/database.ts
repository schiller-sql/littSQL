import { writable, type Writable } from "svelte/store";
import initSqlJs from "./sql-js/sql-wasm";
import type { SqlJsStatic, Database } from "./sql-js/sql-wasm";

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
  return new sql!.Database(data);
}

export async function initSqlite() {
  console.log("sqlite: loading...");
  sql = await initSqlJs({
    locateFile: (file) => {
      console.log(`sqlite: loading '${file}'...`);
      return `static/wasm/${file}`;
    },
  });
  databaseIsReadyStore.set(true);
}
