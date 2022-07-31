export default interface AssignmentListing {
  id: number;
  name: string;
  comment: string | null;
  status: "locked" | "open" | "finished";
}
