import { useState } from "react";
import { useOutletContext } from "react-router-dom";
import { useTagCreation } from "../../hooks/projectSettingsHooks";
import { createTag } from "../../services/projects";
import type { IProjectViewOutletContext, ITag } from "../../types/project";
import ColorPicker from "../util/ColorPicker/ColorPicker";
import "./ProjectSettings.css";

export default function ProjectSettings() {
  const { project } = useOutletContext<IProjectViewOutletContext>();
  const [tags, setTags] = useState<ITag[]>(project.tags);
  const tagCreator = useTagCreation();

  return (
    <>
      <h2>Project Settings</h2>
      <section>
        <h3>Tags</h3>
        <ul className="settings--tag_list-wrapper">
          {tags.map((tag) => {
            return (
              <li
                key={`${tag.id}-${tag.text}`}
                className="tag"
                style={{ background: tag.color }}
              >
                {tag.text}
              </li>
            );
          })}
        </ul>
        <form
          className="settings-create_tag-wrapper"
          onSubmit={(e) => {
            e.preventDefault();
            createTag({
              projectId: project.id,
              tagText: tagCreator.tagText,
              tagColor: tagCreator.tagColor,
            }).then(({ data }) => {
              if (!data) return;
              setTags((prev) => [...prev, data]);
            });
          }}
        >
          <label htmlFor="tag-label">Tag name</label>
          <input
            id="tag-label"
            type="text"
            maxLength={15}
            placeholder="Tag label"
            value={tagCreator.tagText}
            onChange={tagCreator.setTagText}
          />
          <ColorPicker onValueChange={tagCreator.setTagColor} />
          <button type="submit">Create tag</button>
        </form>
        <span style={{ background: tagCreator.tagColor }}>
          {tagCreator.tagText}
        </span>
      </section>
    </>
  );
}
