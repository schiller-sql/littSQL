import type ProjectListing from "./ProjectListing";
import type Task from "./Task";

export default interface Project extends ProjectListing {
  sql: string | null;
  documentation_md: string;
  tasks: Task[];
}
