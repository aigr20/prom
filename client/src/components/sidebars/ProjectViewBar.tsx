import { Link } from "react-router-dom";
import "./ProjectViewBar.css";

type Props = {
  title: string;
};

function isActive(current: string, shouldBe: string): "active" | "" {
  return current === shouldBe ? "active" : "";
}

export default function ProjectViewBar({ title }: Props) {
  const pathSplit = window.location.pathname.split("/");
  const curr = pathSplit.length < 4 ? "backlog" : pathSplit.pop() ?? "backlog";

  return (
    <aside className="layout--sidebar sidebar--project">
      <h2>{title}</h2>
      <ul>
        <li>
          <Link to="backlog" className={isActive(curr, "backlog")}>
            Backlog
          </Link>
        </li>
        <li>
          <Link to="board" className={isActive(curr, "board")}>
            Board
          </Link>
        </li>
        <li>
          <Link to="settings" className={isActive(curr, "settings")}>
            Project Settings
          </Link>
        </li>
      </ul>
    </aside>
  );
}
