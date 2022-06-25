import type Question from "./Question";

export default interface Task {
  description: string;
  is_voluntary: boolean;
  questions: Question[];
}
