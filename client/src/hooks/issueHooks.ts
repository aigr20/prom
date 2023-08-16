import { useState } from "react";
import { createIssue } from "../services/issues";
import { type Setter } from "../types/general";
import { type ITask } from "../types/project";

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
  estimate?: number;
  setEstimate: Setter<number | undefined>;
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
  const [estimate, setEstimate] = useState<number>();

  function onSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    createIssue({ title, description, estimate, projectId }).then(
      ({ data }) => {
        if (data !== null) {
          setIssues([...issues, data]);
          setShowCreateIssue(false);
        }
      },
    );
  }

  return {
    title,
    setTitle,
    description,
    setDescription,
    estimate,
    setEstimate,
    onSubmit,
  };
}
