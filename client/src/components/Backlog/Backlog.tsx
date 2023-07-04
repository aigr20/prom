import { useOutletContext } from "react-router-dom";
import { useProjectTasks } from "../../hooks/projectHooks";
import { Setter } from "../../types/general";
import { IProject } from "../../types/project";
import SpinIfNull from "../util/SpinIfNull";
import { Icons } from "../util/icons";

type BacklogContext = {
  project: IProject;
  setShowCreateIssue: Setter<boolean>;
};

export default function Backlog() {
  const { project, setShowCreateIssue } = useOutletContext<BacklogContext>();
  const tasks = useProjectTasks({ projectId: project.id });

  return (
    <>
      <h2>{project.name} Backlog</h2>
      <SpinIfNull couldBeNull={tasks}>
        <ul>
          {tasks?.map((task) => {
            return <li key={task.id}>{task.title}</li>;
          })}
          <li>
            <button onClick={() => setShowCreateIssue(true)}>
              {Icons.heavyPlus} New issue
            </button>
          </li>
        </ul>
      </SpinIfNull>
    </>
  );
}
