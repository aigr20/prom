import type { Setter } from "./general";
import { type IProject, type ITask } from "./project";

export type IColumn = {
  issues: ITask[];
  heading: string;
};

export type IIssueModalOutletContext = {
  project?: IProject;
  task?: ITask;
  setTasks?: Setter<ITask[]>;
};
