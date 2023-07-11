import { useDragHandlers } from "../../hooks/useDragHandlers";
import { IColumn } from "../../types/board";
import { Setter } from "../../types/general";
import { ITask } from "../../types/project";
import "./Card.css";

type Props = {
  task: ITask;
  index: number;
  colIndex: number;
  columns: IColumn[];
  setColumns: Setter<IColumn[]>;
  setTasks: Setter<ITask[]>;
};

export default function Card({
  task,
  index,
  colIndex,
  columns,
  setColumns,
  setTasks,
}: Props) {
  const handlers = useDragHandlers(columns, setColumns, setTasks);

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
