import { useOutletContext } from "react-router-dom";
import { useTagDropdown } from "../../hooks/issueModalHooks";
import type { IIssueModalOutletContext } from "../../types/board";
import { type ITag } from "../../types/project";
import "./TagDropdown.css";

type Props = {
  issueId?: number;
  tags?: ITag[];
};

export default function TagDropdown({ issueId, tags }: Props) {
  const { isShown, setIsShown, selectedTags, selectTag } = useTagDropdown({
    issueId,
    tags: tags ?? [],
  });
  const { project } = useOutletContext<IIssueModalOutletContext>();

  return (
    <div className="tagselector--wrapper">
      <button
        onClick={() => setIsShown((prev) => !prev)}
        className="tagselector--toggle"
      >
        Add tag
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
              onClick={() => selectTag(tag.id)}
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
