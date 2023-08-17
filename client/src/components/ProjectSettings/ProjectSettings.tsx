import { useOutletContext } from "react-router-dom";
import type { IProjectViewOutletContext } from "../../types/project";
import "./ProjectSettings.css";

export default function ProjectSettings() {
  const { project } = useOutletContext<IProjectViewOutletContext>();

  return (
    <>
      <h2>Project Settings</h2>
      <section>
        <h3>Tags</h3>
        <ul className="settings--tag_list-wrapper">
          {project.tags.map((tag) => {
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
        <form className="settings-create_tag-wrapper">
          <input type="text" maxLength={15} placeholder="Tag label" />
        </form>
      </section>
    </>
  );
}
