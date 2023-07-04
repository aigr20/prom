import { ResponseData } from "../types/general";
import { ITask } from "../types/project";

type CreateIssueArgs = {
  projectId: number;
  title: string;
  description: string;
};
export function createIssue({
  projectId,
  title,
  description,
}: CreateIssueArgs): Promise<ResponseData<ITask | null>> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "POST",
    body: JSON.stringify({ project: projectId, title, description }),
    headers,
  };
  return fetch("http://localhost:8080/issues/create", options)
    .then((res) => {
      if (res.ok && res.status === 201) return res.json();
      return { data: null };
    })
    .catch(() => {
      return { data: null };
    });
}
