import { useEffect, useReducer, useRef, useState } from "react";
import { useOutletContext } from "react-router-dom";
import { updateIssue } from "../services/issues";
import type { IIssueModalOutletContext } from "../types/board";
import type { Setter } from "../types/general";
import type { ITag, ITask } from "../types/project";

type IssueModalReturn = {
  issue?: ITask;
  modifyFunction: (
    field: keyof ITask,
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void;
  onModalClose: () => void;
};
export function useIssueModal(): IssueModalReturn {
  const currentCopyRef = useRef<ITask>();
  const updateFieldsRef = useRef<Partial<ITask>>({});
  const updateTimerRef = useRef<number>();
  const { task: openedIssue, setTasks } =
    useOutletContext<IIssueModalOutletContext>();
  const [issue, setIssue] = useState(openedIssue);
  useEffect(() => {
    if (!issue && openedIssue) setIssue({ ...openedIssue });
  }, [issue, openedIssue]);

  if (!currentCopyRef.current) currentCopyRef.current = openedIssue;

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
      clearTimeout(updateTimerRef.current);
      if (issueHasChanged(updateFieldsRef.current)) {
        updateTimerRef.current = setTimeout(() => {
          updateIssue({
            issueId: issue.id,
            fields: { ...updateFieldsRef.current },
          });
          updateFieldsRef.current = {};
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
      updateIssue({
        issueId: issue.id,
        fields: { ...updateFieldsRef.current },
      });
      updateFieldsRef.current = {};
    }
    setTasks?.((oldTasks) => {
      if (!issue) return oldTasks;
      const index = oldTasks.findIndex((task) => openedIssue?.id === task.id);
      oldTasks[index] = issue;
      return [...oldTasks];
    });
  }

  return { issue: issue, modifyFunction, onModalClose };
}

function issueHasChanged(issue: Partial<ITask>): boolean {
  return Object.entries(issue).some(([, val]) => val !== undefined);
}

type TagDropdownArgs = {
  issueId?: number;
  tags: ITag[];
};
type TagDropdownReturn = {
  isShown: boolean;
  setIsShown: Setter<boolean>;
  selectTag: (tag: number) => void;
  selectedTags: number[];
};
export function useTagDropdown({
  issueId,
  tags,
}: TagDropdownArgs): TagDropdownReturn {
  const [isShown, setIsShown] = useState(false);
  const [selectedTags, dispatch] = useReducer(
    (prev: number[], touched: number) => {
      if (prev.includes(touched)) {
        return [...prev.filter((id) => id !== touched)];
      } else {
        return [...prev, touched];
      }
    },
    [],
    () => tags.map((tag) => tag.id),
  );

  return { isShown, setIsShown, selectTag: dispatch, selectedTags };
}
