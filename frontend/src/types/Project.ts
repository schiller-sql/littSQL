import type ProjectListing from "./ProjectListing";

export default interface Project extends ProjectListing {
  sql: string | null;
  documentation_md: string;
  tasks: Task[];
}

interface Task {
  description: string;
  isVoluntary: boolean;
  questions: Question[];
}

interface Question {
  question: string;
  type: "text" | "sql";
  solution: "";
}
