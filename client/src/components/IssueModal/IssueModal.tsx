import { forwardRef, useImperativeHandle, useRef, useState } from "react";
import { ITask } from "../../types/project";
import { Icons } from "../util/icons";
import "./IssueModal.css";

export type OpenModalFunc = (issue: ITask) => void;

const IssueModal = forwardRef<OpenModalFunc, object>(function IssueModal(
  _,
  ref,
) {
  const [issue, setIssue] = useState<ITask>();
  const [isOpen, setIsOpen] = useState(false);
  const modalRef = useRef<HTMLDialogElement>(null);
  useImperativeHandle(
    ref,
    () => {
      return (issue) => {
        setIssue({ ...issue });
        modalRef.current?.showModal();
        setIsOpen(true);
      };
    },
    [],
  );

  return (
    <dialog
      ref={modalRef}
      className={isOpen ? "issue--modal-open" : "issue--modal-closed"}
      onClose={() => setIsOpen(false)}
    >
      <button
        className="close-button"
        onClick={() => modalRef.current?.close()}
      >
        {Icons.close}
      </button>
      <h1 className="heading">{issue?.title}</h1>
      <p className="description">{issue?.description}</p>
      <span className="created">{issue?.createdAt.toLocaleString()}</span>
      <span className="updated">{issue?.updatedAt.toLocaleString()}</span>
      <span className="estimate">{issue?.estimate}</span>
    </dialog>
  );
});

export default IssueModal;
