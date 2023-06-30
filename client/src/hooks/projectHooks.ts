import { useEffect, useState } from "react";
import { getProjects } from "../services/projects";
import { IProject } from "../types/project";

export function useProjects(): IProject[] {
  const [projects, setProjects] = useState<IProject[]>([]);
  useEffect(() => {
    getProjects().then(({ data }) => setProjects([...data]));
  }, []);

  return projects;
}
