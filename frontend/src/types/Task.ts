import type Question from "./Question";

export default interface Task {
  description: string;
  isVoluntary: boolean;
  questions: Question[];
}
