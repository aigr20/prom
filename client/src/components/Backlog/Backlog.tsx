import { useOutletContext } from "react-router-dom";
import {
  type ITask,
  type IProjectViewOutletContext,
} from "../../types/project";
import { Icons } from "../util/icons";
import { useEffect, useState } from "react";
import { getProjectBacklog } from "../../services/projects";
import Sprint from "../Sprint/Sprint";

export default function Backlog() {
  const { project, setShowCreateIssue } =
    useOutletContext<IProjectViewOutletContext>();
  const [backlogTasks, setBacklogTasks] = useState<ITask[]>([]);
  useEffect(() => {
    getProjectBacklog({ projectId: project.id }).then(({ data }) =>
      setBacklogTasks([...data]),
    );
  }, [project.id]);

  return (
    <>
      <h2>{project.name} Backlog</h2>
      {project.currentSprint && <Sprint sprint={project.currentSprint} />}
      <h3>Project Backlog</h3>
      <ul>
        {backlogTasks.map((task) => {
          return <li key={task.id}>{task.title}</li>;
        })}
        <li>
          <button onClick={() => setShowCreateIssue(true)}>
            {Icons.heavyPlus} New issue
          </button>
        </li>
      </ul>
    </>
  );
}
