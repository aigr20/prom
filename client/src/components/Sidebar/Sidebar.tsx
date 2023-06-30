import { useProjects } from "../../hooks/projectHooks";
import "./Sidebar.css";

export default function Sidebar() {
  const projects = useProjects();
  const projectElements = projects.map((project, idx) => {
    return (
      <li className="sidebar--list-item" key={idx}>
        <a href={`/project/${project.id}`}>{project.name}</a>
      </li>
    );
  });

  return (
    <aside className="sidebar--wrapper">
      <ul className="sidebar--list">{...projectElements}</ul>
    </aside>
  );
}
