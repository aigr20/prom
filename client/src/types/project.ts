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

export type IProjectCreationReturn = {
  projectName: string;
  setProjectName: React.Dispatch<React.SetStateAction<string>>;
  submitCallback: React.FormEventHandler<HTMLFormElement>;
};

export type ITask = {
  id: number;
  title: string;
  desciption: string;
  createdAt: Date;
  updatedAt: Date;
};
