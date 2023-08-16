import {
  useCallback,
  useEffect,
  useMemo,
  useReducer,
  useRef,
  useState,
} from "react";
import { useOutletContext } from "react-router-dom";
import { addTags, removeTags, updateIssue } from "../services/issues";
import type { IIssueModalOutletContext } from "../types/board";
import type { Setter } from "../types/general";
import type { ITag, ITask } from "../types/project";

type IssueModalReturn = {
  issueValues: ITask | null;
  modifyFunction: (
    field: keyof ITask,
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void;
  onModalClose: () => void;
  setIssue: Setter<ITask | null>;
};
export function useIssueModal(openedIssue: ITask | null): IssueModalReturn {
  const currentCopyRef = useRef<ITask | null>(null);
  const updateFieldsRef = useRef<Partial<ITask>>({});
  const updateTimerRef = useRef<number>();
  const { setTasks } = useOutletContext<IIssueModalOutletContext>();
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
      if (!oldIssue) return null;
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

  return { issueValues: issue, modifyFunction, onModalClose, setIssue };
}

function issueHasChanged(issue: Partial<ITask>): boolean {
  return Object.entries(issue).some(([, val]) => val !== undefined);
}

type TagDropdownArgs = {
  issueId?: number;
  tags: ITag[];
  setIssue: Setter<ITask | null>;
};
type TagDropdownReturn = {
  isShown: boolean;
  toggleDropdown: () => void;
  selectTag: (arg: { action: "reset" | "do"; tag: number }) => void;
  selectedTags: number[];
};
export function useTagDropdown({
  issueId,
  tags,
  setIssue,
}: TagDropdownArgs): TagDropdownReturn {
  const [isShown, setIsShown] = useState(false);
  const originalTags = useMemo(() => tags.map((tag) => tag.id), [tags]);
  const tagReducer = useCallback(
    (prev: number[], arg: { action: "do" | "reset"; tag?: number }) => {
      if (arg.action === "reset" || arg.tag === undefined) {
        return [...originalTags];
      }
      if (prev.includes(arg.tag)) {
        return [...prev.filter((id) => id !== arg.tag)];
      } else {
        return [...prev, arg.tag];
      }
    },
    [originalTags],
  );
  const [selectedTags, dispatch] = useReducer(tagReducer, originalTags);

  function toggleDropdown() {
    setIsShown((wasShown) => {
      if (wasShown && issueId !== undefined) {
        if (selectedTags.some((tag) => !originalTags.includes(tag))) {
          addTags({ issueId, tags: selectedTags })
            .then(({ data }) => {
              setIssue((prev) => {
                if (!prev) return null;
                return { ...prev, tags: data };
              });
            })
            .then(() => dispatch({ action: "reset" }));
        }

        if (originalTags.some((tag) => !selectedTags.includes(tag))) {
          removeTags({
            issueId,
            tags: originalTags.filter((tag) => !selectedTags.includes(tag)),
          })
            .then(({ data }) => {
              setIssue((prev) => {
                if (!prev) return null;
                return { ...prev, tags: data };
              });
            })
            .then(() => dispatch({ action: "reset" }));
        }
      }
      return !wasShown;
    });
  }

  return { isShown, toggleDropdown, selectTag: dispatch, selectedTags };
}
