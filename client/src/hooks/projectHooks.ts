import { useContext, useEffect, useState } from "react";
import { ProjectContext } from "../context/ProjectContext";
import {
  createProject,
  getProjectTasks,
  getProjects,
} from "../services/projects";
import type { ProjectIDArg, Setter } from "../types/general";
import type {
  IProjectCreationReturn,
  IProjectsAndSetter,
  ITask,
} from "../types/project";

export function useProjects(): IProjectsAndSetter {
  const { projects, setProjects } = useContext(ProjectContext);
  useEffect(() => {
    getProjects().then(({ data }) => setProjects([...data]));
  }, [setProjects]);

  return { projects, setProjects };
}

export function useProjectTasks({ projectId }: ProjectIDArg): ITask[] | null {
  const [tasks, setTasks] = useState<ITask[] | null>(null);
  useEffect(() => {
    getProjectTasks({ projectId }).then(({ data }) => setTasks([...data]));
  }, [projectId]);

  return tasks;
}

type ProjectCreationArgs = {
  showFormSetter: Setter<boolean>;
};
export function useProjectCreation({
  showFormSetter,
}: ProjectCreationArgs): IProjectCreationReturn {
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
