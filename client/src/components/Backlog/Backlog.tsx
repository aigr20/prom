import { useOutletContext } from "react-router-dom";
import { type IProjectViewOutletContext } from "../../types/project";
import SpinIfNull from "../util/SpinIfNull";
import { Icons } from "../util/icons";

export default function Backlog() {
  const { project, setShowCreateIssue, tasks } =
    useOutletContext<IProjectViewOutletContext>();

  return (
    <>
      <h2>{project.name} Backlog</h2>
      {project.currentSprint && (
        <>
          <h3>Sprint Backlog - {project.currentSprint.name}</h3>
          <ul>
            {project.currentSprint.issues.map((task) => {
              return <li key={task.id}>{task.title}</li>;
            })}
          </ul>
        </>
      )}
      <h3>Project Backlog</h3>
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
