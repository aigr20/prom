import { useState } from "react";
import { Setter } from "../types/general";

type IssueCreationArgs = {
  projectId: number;
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
}: IssueCreationArgs): IssueCreationReturn {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  function onSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
  }

  return {
    title,
    setTitle,
    description,
    setDescription,
    onSubmit,
  };
}
