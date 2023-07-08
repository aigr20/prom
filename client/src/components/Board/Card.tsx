import { useDragHandlers } from "../../hooks/useDragHandlers";
import { ITask } from "../../types/project";
import "./Card.css";

type Props = {
  tasks: ITask[];
  task: ITask;
  index: number;
  colIndex: number;
};

export default function Card({ task, index, colIndex, tasks }: Props) {
  const handlers = useDragHandlers(tasks);

  return (
    <div
      className="board--card"
      data-index={index}
      data-col-index={colIndex}
      {...handlers}
    >
      <h4 className="card--title">{task.title}</h4>
      <span className="card--tag">Some</span>
      <span className="card--tag">Tags</span>
    </div>
  );
}
