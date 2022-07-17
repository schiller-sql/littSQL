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

