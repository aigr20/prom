import { ResponseData } from "../types/general";
import { Project } from "../types/project";

const endpoint = "projects";

export function getProjects(): Promise<ResponseData<Project[]>> {
  return fetch(`http://localhost:8080/${endpoint}/all`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      else return { data: [] };
    })
    .catch(() => {
      return { data: [] };
    });
}
