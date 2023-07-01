import { useEffect, useState } from "react";
import { getProjectTasks, getProjects } from "../services/projects";
import { ProjectIDArg } from "../types/general";
import { IProject, ITask } from "../types/project";

export function useProjects(): IProject[] {
  const [projects, setProjects] = useState<IProject[]>([]);
  useEffect(() => {
    getProjects().then(({ data }) => setProjects([...data]));
  }, []);

  return projects;
}

export function useProjectTasks({ projectId }: ProjectIDArg): ITask[] | null {
  const [tasks, setTasks] = useState<ITask[] | null>(null);
  useEffect(() => {
    getProjectTasks({ projectId }).then(({ data }) => setTasks([...data]));
  }, [projectId]);

  return tasks;
}
