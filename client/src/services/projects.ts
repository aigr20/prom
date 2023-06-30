import { ResponseData } from "../types/general";
import { IProject } from "../types/project";

const endpoint = "projects";

export async function getProjects(): Promise<ResponseData<IProject[]>> {
  return fetch(`http://localhost:8080/${endpoint}/all`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      else return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}

export async function getProject({
  projectId,
}: {
  projectId: number;
}): Promise<ResponseData<IProject | null>> {
  return fetch(`http://localhost:8080/${endpoint}/${projectId}`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      else return { data: null };
    })
    .catch(() => {
      return { data: null };
    });
}
