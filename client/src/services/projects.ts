import { type CreateProjectArg, type ResponseData } from "../types/general";
import { type IProject, type ITask } from "../types/project";

type ProjectIDArg = {
  projectId: number;
};

export async function getProjects(): Promise<ResponseData<IProject[]>> {
  return fetch(`${import.meta.env.VITE_API_URL}/projects/all`)
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
  return fetch(`${import.meta.env.VITE_API_URL}/projects/${projectId}`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: null };
    })
    .catch(() => {
      return { data: null };
    });
}

export async function createProject({
  projectName,
}: CreateProjectArg): Promise<ResponseData<IProject | null>> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "POST",
    headers,
    body: JSON.stringify({ name: projectName }),
  };

  return fetch(`${import.meta.env.VITE_API_URL}/projects/create`, options)
    .then((res) => {
      if (res.ok && res.status === 201) return res.json();
      return { data: null };
    })
    .catch(() => {
      return { data: null };
    });
}

export async function getProjectTasks({
  projectId,
}: ProjectIDArg): Promise<ResponseData<ITask[]>> {
  return fetch(`${import.meta.env.VITE_API_URL}/projects/${projectId}/issues`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}
