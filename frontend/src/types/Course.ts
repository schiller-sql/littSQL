import type CourseListing from "./CourseListing";

export default interface Course extends CourseListing {
  participant_count: number;
}
