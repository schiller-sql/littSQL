export default interface AssignmentListing {
  id: number;
  name: string;
  status: "locked" | "open" | "finished";
}
