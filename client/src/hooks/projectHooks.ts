import { useEffect, useState } from "react";
import { getProjects } from "../services/projects";
import { Project } from "../types/project";

export function useProjects(): Project[] {
  const [projects, setProjects] = useState<Project[]>([]);
  useEffect(() => {
    getProjects().then(({ data }) => setProjects([...data]));
  }, []);

  return projects;
}
