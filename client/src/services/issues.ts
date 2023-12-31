import { type ResponseData } from "../types/general";
import { type ITag, type ITask } from "../types/project";

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
  return fetch(`${import.meta.env.VITE_API_URL}/issues/create`, options)
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
  return fetch(`${import.meta.env.VITE_API_URL}/issues/update`, options)
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

  return fetch(`${import.meta.env.VITE_API_URL}/issues/status`, options)
    .then((res) => {
      if (res.ok && res.status === 204) return;
      throw new Error("Status update failed");
    })
    .catch((err: Error) => {
      console.error(err.message);
    });
}

export async function getIssue({
  issueId,
}: {
  issueId: number;
}): Promise<ResponseData<ITask | null>> {
  return fetch(`${import.meta.env.VITE_API_URL}/issues/${issueId}`)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      throw new Error("Failed to get issue");
    })
    .catch((err: Error) => {
      alert(err.message);
      return { data: null };
    });
}

type TagManipArgs = {
  issueId: number;
  tags: number[];
};
export async function addTags({
  issueId,
  tags,
}: TagManipArgs): Promise<ResponseData<ITag[]>> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "PATCH",
    body: JSON.stringify({ issueId, tags }),
    headers,
  };

  return fetch(`${import.meta.env.VITE_API_URL}/issues/tags`, options)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      throw new Error("Failed to add tags");
    })
    .catch((err: Error) => alert(err.message));
}

export async function removeTags({
  issueId,
  tags,
}: TagManipArgs): Promise<ResponseData<ITag[]>> {
  const headers = new Headers();
  headers.set("Content-Type", "application/json");
  const options: RequestInit = {
    method: "DELETE",
    body: JSON.stringify({ issueId, tags }),
    headers,
  };

  return fetch(`${import.meta.env.VITE_API_URL}/issues/tags`, options)
    .then((res) => {
      if (res.ok && res.status === 200) return res.json();
      throw new Error("Failed to remove tags");
    })
    .catch((err: Error) => alert(err.message));
}
