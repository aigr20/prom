import { ProjectIDArg, ResponseData } from "../types/general";
import { IProject, ITask } from "../types/project";

const endpoint = "projects";

export async function getProjects(): Promise<ResponseData<IProject[]>> {
  return fetch(`http://localhost:8080/${endpoint}/all`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}

export async function getProject({
  projectId,
}: ProjectIDArg): Promise<ResponseData<IProject | null>> {
  return fetch(`http://localhost:8080/${endpoint}/${projectId}`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: null };
    })
    .catch(() => {
      return { data: null };
    });
}

export async function getProjectTasks({
  projectId,
}: ProjectIDArg): Promise<ResponseData<ITask[]>> {
  return fetch(`http://localhost:8080/${endpoint}/${projectId}/issues`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}
