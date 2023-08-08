import { type ITask } from "./project";

export type IColumn = {
  issues: ITask[];
  heading: string;
};
