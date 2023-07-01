import { useOutletContext } from "react-router-dom";
import { useProjectTasks } from "../../hooks/projectHooks";
import { IProject } from "../../types/project";
import SpinIfNull from "../util/SpinIfNull";

export default function Backlog() {
  const project = useOutletContext<IProject>();
  const tasks = useProjectTasks({ projectId: project.id });

  return (
    <>
      <h2>{project.name} Backlog</h2>
      <SpinIfNull couldBeNull={tasks}>
        <ul>
          {tasks?.map((task) => {
            return <li key={task.id}>{task.title}</li>;
          })}
        </ul>
      </SpinIfNull>
    </>
  );
}
