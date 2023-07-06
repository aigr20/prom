import { useOutletContext } from "react-router-dom";
import { Setter } from "../../types/general";
import { IProject, ITask } from "../../types/project";
import SpinIfNull from "../util/SpinIfNull";
import { Icons } from "../util/icons";

type BacklogContext = {
  project: IProject;
  tasks: ITask[];
  setShowCreateIssue: Setter<boolean>;
};

export default function Backlog() {
  const { project, setShowCreateIssue, tasks } =
    useOutletContext<BacklogContext>();

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
