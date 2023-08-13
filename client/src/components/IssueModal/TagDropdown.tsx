import { useState } from "react";
import { useOutletContext } from "react-router-dom";
import { type IProjectViewOutletContext } from "../../types/project";
import "./TagDropdown.css";

export default function TagDropdown() {
  const [isShown, setIsShown] = useState(false);
  const { project } = useOutletContext<IProjectViewOutletContext>();

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
        {[
          ...project.tags,
          ...project.tags,
          ...project.tags,
          ...project.tags,
          ...project.tags,
          ...project.tags,
          ...project.tags,
          ...project.tags,
        ].map((tag, idx) => {
          return (
            <button
              key={`tagoption-${project.id}-${idx}`}
              className="tagselector--item"
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
