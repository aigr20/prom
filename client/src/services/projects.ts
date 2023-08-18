import { type CreateProjectArg, type ResponseData } from "../types/general";
import { type IProject, type ITag, type ITask } from "../types/project";

const endpoint = "projects";

type ProjectIDArg = {
  projectId: number;
};

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

  return fetch(`http://localhost:8080/${endpoint}/create`, options)
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

  return fetch("http://localhost:8080/tags/create", options)
    .then((res) => {
      if (res.ok && res.status === 201) return res.json();
      throw new Error("Tag creation failed");
    })
    .catch((err: Error) => alert(err.message));
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
