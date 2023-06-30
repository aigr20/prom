import { IProject } from "../../types/project";
import "./Topbar.css";

type Props = {
  projects: IProject[];
};

export default function Topbar({ projects }: Props) {
  return (
    <nav className="topbar--wrapper">
      <h1 className="topbar--logo">Prom</h1>
    </nav>
  );
}
