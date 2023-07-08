import { useOutletContext } from "react-router-dom";
import { IProjectViewOutletContext } from "../../types/project";
import "./Board.css";
import Card from "./Card";

export default function Board() {
  const columns = ["TODO", "In Progress", "Finished"];
  const { tasks } = useOutletContext<IProjectViewOutletContext>();

  return (
    <>
      <h2>Board</h2>
      <div className="board--wrapper">
        {columns.map((title, colIdx) => (
          <div key={`col-${colIdx}`} className="board--column">
            <h3 className="board--column-title">{title}</h3>
            {tasks
              .filter((task) => task.status === title)
              .map((task, idx) => (
                <Card
                  key={`issue-${task.id}`}
                  tasks={tasks}
                  task={task}
                  index={idx}
                  colIndex={colIdx}
                />
              ))}
          </div>
        ))}
      </div>
    </>
  );
}
