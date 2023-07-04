import { useOutletContext } from "react-router-dom";
import { IProject } from "../../types/project";

export default function Board() {
  const project = useOutletContext<IProject>();

  return (
    <>
      <h2>Board</h2>
      {project.name}
    </>
  );
}
