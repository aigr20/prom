import { Link } from "react-router-dom";
import { useProjects } from "../../hooks/projectHooks";
import { type Setter } from "../../types/general";
import { Icons } from "../util/icons";
import "./HomeViewBar.css";

type Props = {
  showFormCB: Setter<boolean>;
};

export default function HomeViewBar({ showFormCB }: Props) {
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
