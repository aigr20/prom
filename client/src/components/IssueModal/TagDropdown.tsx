import { useTagDropdown } from "../../hooks/issueModalHooks";
import type { Setter } from "../../types/general";
import { type ITag, type ITask } from "../../types/project";
import "./TagDropdown.css";

type Props = {
  issueId?: number;
  projectTags: ITag[];
  tags?: ITag[];
  setIssue: Setter<ITask | null>;
};

export default function TagDropdown({
  issueId,
  projectTags,
  tags,
  setIssue,
}: Props) {
  const { isShown, toggleDropdown, selectedTags, selectTag } = useTagDropdown({
    issueId,
    tags: tags ?? [],
    setIssue,
  });

  return (
    <div className="tagselector--wrapper">
      <button onClick={toggleDropdown} className="tagselector--toggle">
        Change tags
      </button>
      <div
        style={{ display: isShown ? "" : "none" }}
        className="tagselector--item-wrapper"
      >
        {projectTags.map((tag, idx) => {
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
