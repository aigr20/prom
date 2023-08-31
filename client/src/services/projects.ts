import { type CreateProjectArg, type ResponseData } from "../types/general";
import {
  type IProject,
  type ITag,
  type ITagCount,
  type ITask,
} from "../types/project";

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

type CreateTagArgs = {
  projectId: number;
  tagText: string;
  tagColor: string;
};
export async function createTag({
  projectId,
  tagText,
  tagColor,
}: CreateTagArgs): Promise<ResponseData<ITag | null>> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "POST",
    headers,
    body: JSON.stringify({ projectId, text: tagText, color: tagColor }),
  };

  return fetch(`${import.meta.env.VITE_API_URL}/tags/create`, options)
    .then((res) => {
      if (res.ok && res.status === 201) return res.json();
      throw new Error("Tag creation failed");
    })
    .catch((err: Error) => alert(err.message));
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

export async function getProjectBacklog({
  projectId,
}: ProjectIDArg): Promise<ResponseData<ITask[]>> {
  return fetch(`${import.meta.env.VITE_API_URL}/projects/${projectId}/backlog`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}

export async function getProjectTagCounts({
  projectId,
}: ProjectIDArg): Promise<ResponseData<ITagCount[]>> {
  return fetch(
    `${import.meta.env.VITE_API_URL}/projects/${projectId}/tag_counts`,
  )
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}
