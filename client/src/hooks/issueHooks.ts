import { useState } from "react";
import { createIssue } from "../services/issues";
import { Setter } from "../types/general";
import { ITask } from "../types/project";

type IssueCreationArgs = {
  projectId: number;
  setShowCreateIssue: Setter<boolean>;
  issues: ITask[];
  setIssues: Setter<ITask[]>;
};
type IssueCreationReturn = {
  title: string;
  setTitle: Setter<string>;
  description: string;
  setDescription: Setter<string>;
  onSubmit: React.FormEventHandler<HTMLFormElement>;
};
export function useIssueCreation({
  projectId,
  setShowCreateIssue,
  issues,
  setIssues,
}: IssueCreationArgs): IssueCreationReturn {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  function onSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    createIssue({ title, description, projectId }).then(({ data }) => {
      if (data !== null) {
        setIssues([...issues, data]);
        setShowCreateIssue(false);
      }
    });
  }

  return {
    title,
    setTitle,
    description,
    setDescription,
    onSubmit,
  };
}
