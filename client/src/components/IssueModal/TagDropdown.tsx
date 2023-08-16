import { useOutletContext } from "react-router-dom";
import { useTagDropdown } from "../../hooks/issueModalHooks";
import type { IIssueModalOutletContext } from "../../types/board";
import type { Setter } from "../../types/general";
import { type ITag, type ITask } from "../../types/project";
import "./TagDropdown.css";

type Props = {
  issueId?: number;
  tags?: ITag[];
  setIssue: Setter<ITask | null>;
};

export default function TagDropdown({ issueId, tags, setIssue }: Props) {
  const { isShown, toggleDropdown, selectedTags, selectTag } = useTagDropdown({
    issueId,
    tags: tags ?? [],
    setIssue,
  });
  const { project } = useOutletContext<IIssueModalOutletContext>();

  return (
    <div className="tagselector--wrapper">
      <button onClick={toggleDropdown} className="tagselector--toggle">
        Change tags
      </button>
      <div
        style={{ display: isShown ? "" : "none" }}
        className="tagselector--item-wrapper"
      >
        {project?.tags.map((tag, idx) => {
          const isSelected = selectedTags.includes(tag.id);
          return (
            <button
              key={`tagoption-${issueId}-${idx}`}
              className={`tagselector--item ${
                isSelected ? "tagselector--item-selected" : ""
              }`}
              onClick={() => selectTag({ action: "do", tag: tag.id })}
            >
              <svg width={10} height={10} viewBox="0 0 2.6458 2.6458">
                <rect fill={tag.color} height="100%" width="100%" />
              </svg>
              <span>{tag.text}</span>
            </button>
          );
        })}
      </div>
    </div>
  );
}
