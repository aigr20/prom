import { useOutletContext } from "react-router-dom";
import { IProjectViewOutletContext } from "../../types/project";
import "./Board.css";

export default function Board() {
  const columns = ["TODO", "In Progress", "Finished"];
  const { tasks } = useOutletContext<IProjectViewOutletContext>();

  return (
    <>
      <h2>Board</h2>
      <div className="board--wrapper">
        {columns.map((title) => (
          <div key={`col-${title}`} className="board--column">
            <h3 className="board--column-title">{title}</h3>
            {tasks
              .filter((task) => task.status === title)
              .map((task) => (
                <div key={task.id}>{task.title}</div>
              ))}
          </div>
        ))}
      </div>
    </>
  );
}
