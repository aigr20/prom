import { useContext, useEffect, useState } from "react";
import { ProjectContext } from "../context/ProjectContext";
import {
  createProject,
  getProjectTasks,
  getProjects,
} from "../services/projects";
import type { Setter } from "../types/general";
import type { IProjectsAndSetter, ITask } from "../types/project";

export function useProjects(): IProjectsAndSetter {
  const { projects, setProjects } = useContext(ProjectContext);
  useEffect(() => {
    getProjects().then(({ data }) => setProjects([...data]));
  }, [setProjects]);

  return { projects, setProjects };
}

type ProjectTasksArgs = {
  projectId?: number;
};
export function useProjectTasks({
  projectId,
}: ProjectTasksArgs): ITask[] | null {
  const [tasks, setTasks] = useState<ITask[] | null>(null);
  useEffect(() => {
    if (projectId === undefined) return;
    getProjectTasks({ projectId }).then(({ data }) => setTasks([...data]));
  }, [projectId]);

  return tasks;
}

type ProjectCreationArgs = {
  showFormSetter: Setter<boolean>;
};
type ProjectCreationReturn = {
  projectName: string;
  setProjectName: React.Dispatch<React.SetStateAction<string>>;
  submitCallback: React.FormEventHandler<HTMLFormElement>;
};
export function useProjectCreation({
  showFormSetter,
}: ProjectCreationArgs): ProjectCreationReturn {
  const [projectName, setProjectName] = useState("");
  const { projects, setProjects } = useContext(ProjectContext);

  function onSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    createProject({ projectName }).then(({ data }) => {
      if (data !== null) {
        showFormSetter(false);
        setProjects([...projects, data]);
      } else {
        alert("error!!!");
      }
    });
  }

  return { projectName, setProjectName, submitCallback: onSubmit };
}
