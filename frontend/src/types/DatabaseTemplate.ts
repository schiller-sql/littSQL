import type DatabaseTemplateListing from "./DatabaseTemplateListing";

export default interface DatabaseTemplate extends DatabaseTemplateListing {
  sql: string;
}
