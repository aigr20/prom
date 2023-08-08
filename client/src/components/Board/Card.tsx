import { useDragHandlers } from "../../hooks/useDragHandlers";
import { type IColumn } from "../../types/board";
import { type Setter } from "../../types/general";
import { type ITask } from "../../types/project";
import { type OpenModalFunc } from "../IssueModal/IssueModal";
import { Icons } from "../util/icons";
import "./Card.css";

type Props = {
  task: ITask;
  index: number;
  colIndex: number;
  columns: IColumn[];
  setColumns: Setter<IColumn[]>;
  setTasks: Setter<ITask[]>;
  openAsModal: OpenModalFunc | null;
};

export default function Card({
  task,
  index,
  colIndex,
  columns,
  setColumns,
  setTasks,
  openAsModal,
}: Props) {
  const handlers = useDragHandlers(columns, setColumns, setTasks);

  return (
    <div
      className="board--card"
      data-index={index}
      data-col-index={colIndex}
      onMouseDown={() => openAsModal?.(task)}
      {...handlers}
    >
      <div className="card--title-row">
        <h4 className="card--title">{task.title}</h4>
        <span
          className="card--drag-handle"
          onMouseDown={(e) => e.stopPropagation()}
        >
          {Icons.drag}
        </span>
      </div>
      <span className="card--tag">Some</span>
      <span className="card--tag">Tags</span>
    </div>
  );
}
