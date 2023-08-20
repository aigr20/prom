import { type Setter } from "./general";

export type IProject = {
  id: number;
  name: string;
  tags: ITag[];
  createdAt: Date;
  updatedAt: Date;
};

export type IProjectsAndSetter = {
  projects: IProject[];
  setProjects: React.Dispatch<React.SetStateAction<IProject[]>>;
};

export type ITag = {
  id: number;
  text: string;
  color: string;
};

export type ITagCount = {
  tag: string;
  count: number;
};

export type ITask = {
  id: number;
  title: string;
  description: string;
  estimate: number;
  createdAt: Date;
  updatedAt: Date;
  status: string;
  tags: ITag[];
};

export type IProjectViewOutletContext = {
  project: IProject;
  tasks: ITask[];
  setTasks: Setter<ITask[]>;
  setShowCreateIssue: Setter<boolean>;
};
