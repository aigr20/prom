import {
  ForwardedRef,
  RefObject,
  useImperativeHandle,
  useRef,
  useState,
} from "react";
import { OpenModalFunc } from "../components/IssueModal/IssueModal";
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

type IssueModalReturn = {
  issue?: ITask;
  modalRef: RefObject<HTMLDialogElement>;
  modifyFunction: (
    field: keyof ITask,
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void;
};
export function useIssueModal(
  ref: ForwardedRef<OpenModalFunc>,
): IssueModalReturn {
  const [issue, setIssue] = useState<ITask>();
  const modalRef = useRef<HTMLDialogElement>(null);

  useImperativeHandle(
    ref,
    () => {
      return (issue) => {
        setIssue({ ...issue });
        modalRef.current?.showModal();
      };
    },
    [],
  );

  function modifyFunction(
    field: keyof ITask,
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) {
    setIssue((oldIssue) => {
      if (!oldIssue) return;
      let number = false;
      switch (field) {
        case "estimate":
          number = true;
          break;
        default:
          number = false;
      }
      const val = number ? Number(event.target.value) : event.target.value;
      return { ...oldIssue, [field]: val };
    });
  }

  return { issue, modalRef, modifyFunction };
}
