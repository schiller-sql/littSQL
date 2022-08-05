import { writable, type Readable } from "svelte/store";
import initSqlJs from "../sql-js/sql-wasm";
import type { SqlJsStatic, Database } from "../sql-js/sql-wasm";

export enum DatabaseState {
  unready,
  ready,
}

interface DatabaseStore extends Readable<DatabaseState> {
  initSqlite(): void;
  newDatabase(data?: ArrayLike<number> | Buffer): Database;
}

export function createDatabaseStore(): DatabaseStore {
  const w = writable<DatabaseState>(DatabaseState.unready);
  let databaseInitCalled = false;
  let sql: SqlJsStatic | undefined;
  return {
    subscribe: w.subscribe,
    newDatabase(data?: ArrayLike<number> | Buffer): Database {
      const db = new sql!.Database(data);
      db.run("PRAGMA foreign_keys = ON");
      return db;
    },
    async initSqlite() {
      // await new Promise((r) => setTimeout(r, 2000));
      console.log("sqlite: not ready yet");
      if (databaseInitCalled) return;
      databaseInitCalled = true;
      console.log("sqlite: loading...");
      sql = await initSqlJs({
        locateFile: (file) => {
          console.log(`sqlite: loading '${file}'...`);
          return `static/sql.js/${file}`;
        },
      });
      console.log("sqlite: ready!");
      w.set(DatabaseState.ready);
    },
  };
}
