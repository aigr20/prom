import { type ResponseData } from "../types/general";
import { type ITask } from "../types/project";

type CreateIssueArgs = {
  projectId: number;
  title: string;
  description: string;
  estimate?: number;
};
export async function createIssue({
  projectId,
  title,
  description,
  estimate,
}: CreateIssueArgs): Promise<ResponseData<ITask | null>> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "POST",
    body: JSON.stringify({ project: projectId, title, description, estimate }),
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

type UpdateIssueArgs = {
  issueId: number;
  fields: Partial<ITask>;
};
export async function updateIssue({
  issueId,
  fields,
}: UpdateIssueArgs): Promise<void> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "PATCH",
    body: JSON.stringify({ issueId, updates: fields }),
    headers,
  };
  return fetch("http://localhost:8080/issues/update", options)
    .then((res) => {
      if (res.ok && res.status === 204) return;
      throw new Error("Update failed");
    })
    .catch((err: Error) => {
      console.log(err.message);
    });
}

type UpdateStatusArgs = {
  issueId: number;
  newStatus: string;
};
export async function updateStatus({
  issueId,
  newStatus,
}: UpdateStatusArgs): Promise<void> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "PATCH",
    body: JSON.stringify({ issueId, newStatus }),
    headers,
  };

  return fetch("http://localhost:8080/issues/status", options)
    .then((res) => {
      if (res.ok && res.status === 204) return;
      throw new Error("Status update failed");
    })
    .catch((err: Error) => {
      console.error(err.message);
    });
}
