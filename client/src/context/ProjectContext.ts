import { createContext } from "react";
import { type IProjectsAndSetter } from "../types/project";

export const ProjectContext = createContext<IProjectsAndSetter>({
  projects: [],
  setProjects: () => undefined,
});
