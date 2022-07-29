import type AssignmentListing from "./AssignmentListing";

export type CorrectionBehaviour =
  | "show_no_correction"
  | "show_correction"
  | "show_correction_and_solution";

export default interface Assigment {
  id: number;
  name: string;
  comment: string | null;
  project_id: number | null;
  finished_date: string | null;
  locked: boolean;
  answer_config: {
    enable_auto_correction_on_sql_questions: boolean;
    show_query_solution: boolean;
    submit_only_once: boolean;
    active_correction_behaviour: CorrectionBehaviour;
    finished_correction_behavior: CorrectionBehaviour;
    finished_hide_answers: boolean;
  };
}
