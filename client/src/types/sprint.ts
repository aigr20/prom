import type { ITask } from "./project";

export type ISprint = {
  id: number;
  name: string;
  start_date: Date;
  end: Date;
  issues: ITask[];
};
