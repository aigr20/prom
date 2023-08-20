import { useEffect, useState } from "react";
import { useOutletContext, useRevalidator } from "react-router-dom";
import { useTagCreation } from "../../hooks/projectSettingsHooks";
import { createTag, getProjectTagCounts } from "../../services/projects";
import type {
  IProjectViewOutletContext,
  ITag,
  ITagCount,
} from "../../types/project";
import ColorPicker from "../util/ColorPicker/ColorPicker";
import "./ProjectSettings.css";

export default function ProjectSettings() {
  const { project } = useOutletContext<IProjectViewOutletContext>();
  const [tags, setTags] = useState<ITag[]>(project.tags);
  const [tagCounts, setTagCounts] = useState<ITagCount[]>([]);
  const tagCreator = useTagCreation();
  const { revalidate } = useRevalidator();
  useEffect(() => {
    getProjectTagCounts({ projectId: project.id }).then(({ data }) => {
      setTagCounts([...data]);
    });
  }, [project.id]);

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
                {tag.text} (
                {tagCounts.find((t) => t.tag === tag.text)?.count ?? 0})
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
              revalidate();
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
