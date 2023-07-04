import { Link } from "react-router-dom";
import { useProjects } from "../../hooks/projectHooks";
import { Setter } from "../../types/general";
import { Icons } from "../util/icons";
import "./HomeViewBar.css";

export default function HomeViewBar({
  showFormCB,
}: Setter<"showFormCB", boolean>) {
  const { projects } = useProjects();
  const projectElements = projects.map((project, idx) => {
    return (
      <li className="sidebar--list-item" key={idx}>
        <Link to={`project/${project.id}/backlog`}>{project.name}</Link>
      </li>
    );
  });

  return (
    <aside className="layout--sidebar sidebar--home">
      <button onClick={() => showFormCB(true)}>{Icons.heavyPlus}</button>
      <ul className="sidebar--list">{...projectElements}</ul>
    </aside>
  );
}
