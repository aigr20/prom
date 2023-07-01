import { Link } from "react-router-dom";
import { IProject } from "../../types/project";
import "./HomeViewBar.css";

type Props = {
  projects: IProject[];
};

export default function HomeViewBar({ projects }: Props) {
  const projectElements = projects.map((project, idx) => {
    return (
      <li className="sidebar--list-item" key={idx}>
        <Link to={`project/${project.id}`}>{project.name}</Link>
      </li>
    );
  });

  return (
    <aside className="layout--sidebar sidebar--home">
      <ul className="sidebar--list">{...projectElements}</ul>
    </aside>
  );
}
