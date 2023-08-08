import {
  useImperativeHandle,
  useRef,
  useState,
  type ForwardedRef,
  type RefObject,
} from "react";
import { type OpenModalFunc } from "../components/IssueModal/IssueModal";
import { createIssue, updateIssue } from "../services/issues";
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

type IssueModalReturn = {
  issue?: ITask;
  modalRef: RefObject<HTMLDialogElement>;
  modifyFunction: (
    field: keyof ITask,
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void;
  onModalClose: () => void;
};
export function useIssueModal(
  ref: ForwardedRef<OpenModalFunc>,
): IssueModalReturn {
  const [issue, setIssue] = useState<ITask>();
  const modalRef = useRef<HTMLDialogElement>(null);
  const currentCopyRef = useRef<ITask>();
  const updateFieldsRef = useRef<Partial<ITask>>({});
  const updateTimerRef = useRef<number>();
  if (!currentCopyRef.current) currentCopyRef.current = issue;

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
    let number = false;
    switch (field) {
      case "estimate":
        number = true;
        break;
      default:
        number = false;
    }

    const val = number ? Number(event.target.value) : event.target.value;
    if (currentCopyRef.current?.[field] !== val) {
      (updateFieldsRef.current[field] as ITask[keyof ITask]) = val;
    } else if (currentCopyRef.current?.[field] === val) {
      updateFieldsRef.current[field] = undefined;
    }

    if (issue) {
      clearTimeout(updateTimerRef?.current);
      if (issueHasChanged(updateFieldsRef.current)) {
        updateTimerRef.current = setTimeout(() => {
          updateIssue({ issueId: issue.id, fields: updateFieldsRef.current });
        }, 10 * 1000);
      }
    }

    setIssue((oldIssue) => {
      if (!oldIssue) return;
      return { ...oldIssue, [field]: val };
    });
  }

  function onModalClose() {
    if (issue && issueHasChanged(updateFieldsRef.current)) {
      clearTimeout(updateTimerRef.current);
      updateIssue({ issueId: issue.id, fields: updateFieldsRef.current });
    }
  }

  return { issue, modalRef, modifyFunction, onModalClose };
}

function issueHasChanged(issue: Partial<ITask>): boolean {
  return Object.entries(issue).some(([, val]) => val !== undefined);
}
