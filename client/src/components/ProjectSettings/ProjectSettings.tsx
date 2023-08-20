import { useOutletContext } from "react-router-dom";
import { useTagCreation } from "../../hooks/projectSettingsHooks";
import type { IProjectViewOutletContext } from "../../types/project";
import ColorPicker from "../util/ColorPicker/ColorPicker";
import "./ProjectSettings.css";

export default function ProjectSettings() {
  const { project } = useOutletContext<IProjectViewOutletContext>();
  const tagCreator = useTagCreation(project);

  return (
    <>
      <h2>Project Settings</h2>
      <section>
        <h3>Tags</h3>
        <ul className="settings--tag_list-wrapper">
          {tagCreator.tags.map((tag) => {
            return (
              <li
                key={`${tag.id}-${tag.text}`}
                className="tag"
                style={{ background: tag.color }}
              >
                {tag.text} (
                {tagCreator.tagCounts.find((t) => t.tag === tag.text)?.count ??
                  0}
                )
              </li>
            );
          })}
        </ul>
        <form
          className="settings-create_tag-wrapper"
          onSubmit={tagCreator.createTagCB}
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
