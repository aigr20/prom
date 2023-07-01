import { useOutletContext } from "react-router-dom";
import { IProject } from "../../types/project";

export default function Backlog() {
  const project = useOutletContext<IProject>();

  return (
    <>
      <h2>Backlog</h2>
      {project.name}
    </>
  );
}
