import { forwardRef } from "react";
import { useIssueModal } from "../../hooks/issueHooks";
import { type ITask } from "../../types/project";
import { formatDate } from "../util/date";
import { Icons } from "../util/icons";
import "./IssueModal.css";
import TagDropdown from "./TagDropdown";

export type OpenModalFunc = (issue: ITask) => void;

const IssueModal = forwardRef<OpenModalFunc, object>(function IssueModal(
  _,
  ref,
) {
  const { issue, modalRef, modifyFunction, onModalClose, availableTags } =
    useIssueModal(ref);

  return (
    <dialog ref={modalRef} className="issue--modal" onClose={onModalClose}>
      <input
        className="heading"
        name="issue-heading"
        value={issue?.title ?? ""}
        onChange={(e) => modifyFunction("title", e)}
      />
      <button
        className="close-button"
        onClick={() => modalRef.current?.close()}
      >
        {Icons.close}
      </button>
      <textarea
        className="description"
        onChange={(e) => modifyFunction("description", e)}
        value={issue?.description ?? ""}
      />
      <div className="tags">
        {issue?.tags.map((tag) => {
          return (
            <span
              key={`modal-tag-${tag.text}-${issue.id}`}
              className="tag"
              style={{ background: tag.color }}
            >
              {tag.text}
            </span>
          );
        })}
      </div>
      <TagDropdown />
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
        value={String(issue?.estimate ?? 0)}
        onChange={(e) => modifyFunction("estimate", e)}
      />
    </dialog>
  );
});

export default IssueModal;
