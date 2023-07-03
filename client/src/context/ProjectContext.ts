import { createContext } from "react";
import { IProjectsAndSetter } from "../types/project";

export const ProjectContext = createContext<IProjectsAndSetter>({
  projects: [],
  setProjects: () => undefined,
});
