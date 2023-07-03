import { useContext } from "react";
import { Link } from "react-router-dom";
import { ProjectContext } from "../../context/ProjectContext";
import "./HomeViewBar.css";

export default function HomeViewBar() {
  const { projects } = useContext(ProjectContext);
  const projectElements = projects.map((project, idx) => {
    return (
      <li className="sidebar--list-item" key={idx}>
        <Link to={`project/${project.id}/backlog`}>{project.name}</Link>
      </li>
    );
  });

  return (
    <aside className="layout--sidebar sidebar--home">
      <ul className="sidebar--list">{...projectElements}</ul>
    </aside>
  );
}
