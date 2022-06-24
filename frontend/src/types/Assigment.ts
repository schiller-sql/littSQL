export default interface Assigment {
  project_id: number;
  course_id: number;
  status: boolean;
  solution_mode: "tryout" | "exam" | "no-solutions-tryout" | "voluntary";
}
