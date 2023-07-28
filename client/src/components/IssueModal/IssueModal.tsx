import { forwardRef, useImperativeHandle, useRef, useState } from "react";
import { ITask } from "../../types/project";
import { formatDate } from "../util/date";
import { Icons } from "../util/icons";
import "./IssueModal.css";

export type OpenModalFunc = (issue: ITask) => void;

const IssueModal = forwardRef<OpenModalFunc, object>(function IssueModal(
  _,
  ref,
) {
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

  return (
    <dialog ref={modalRef} className="issue--modal">
      <input
        className="heading"
        name="issue-heading"
        value={issue?.title ?? ""}
        onChange={(e) =>
          setIssue((oldIssue) => {
            if (oldIssue === undefined) return;
            return { ...oldIssue, title: e.target.value };
          })
        }
      />
      <button
        className="close-button"
        onClick={() => modalRef.current?.close()}
      >
        {Icons.close}
      </button>
      <textarea
        className="description"
        onChange={(e) =>
          setIssue((oldIssue) => {
            if (oldIssue === undefined) return;
            return { ...oldIssue, description: e.target.value };
          })
        }
        value={issue?.description ?? ""}
      ></textarea>
      <span className="created" title="Skapat">
        {issue?.createdAt && formatDate(issue.createdAt)}
      </span>
      <span className="updated" title="Uppdaterat">
        {issue?.updatedAt && formatDate(issue.updatedAt)}
      </span>
      <input
        className="estimate"
        name="issue-estimate"
        type="number"
        min={0}
        step={1}
        value={issue?.estimate ?? 0}
        onChange={(e) =>
          setIssue((oldIssue) => {
            if (oldIssue === undefined) return;
            return { ...oldIssue, estimate: e.target.valueAsNumber };
          })
        }
      />
    </dialog>
  );
});

export default IssueModal;
