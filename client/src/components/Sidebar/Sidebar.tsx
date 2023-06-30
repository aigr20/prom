import { Link } from "react-router-dom";
import { IProject } from "../../types/project";
import "./Sidebar.css";

type Props = {
  projects: IProject[];
};

export default function Sidebar({ projects }: Props) {
  const projectElements = projects.map((project, idx) => {
    return (
      <li className="sidebar--list-item" key={idx}>
        <Link to={`project/${project.id}`}>{project.name}</Link>
      </li>
    );
  });

  return (
    <aside className="sidebar--wrapper">
      <ul className="sidebar--list">{...projectElements}</ul>
    </aside>
  );
}
