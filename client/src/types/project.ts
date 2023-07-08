import { Setter } from "./general";

export type IProject = {
  id: number;
  name: string;
  createdAt: Date;
  updatedAt: Date;
};

export type IProjectsAndSetter = {
  projects: IProject[];
  setProjects: React.Dispatch<React.SetStateAction<IProject[]>>;
};

export type ITask = {
  id: number;
  title: string;
  desciption?: string;
  createdAt: Date;
  updatedAt: Date;
  status: string;
};

export type IProjectViewOutletContext = {
  project: IProject;
  tasks: ITask[];
  setShowCreateIssue: Setter<boolean>;
};
