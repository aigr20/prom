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
      <p>{issue?.title}</p>
      <button onClick={() => modalRef.current?.close()}>{Icons.close}</button>
    </dialog>
  );
});

export default IssueModal;
