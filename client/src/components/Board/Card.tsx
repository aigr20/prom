import { useNavigate } from "react-router-dom";
import { useDragHandlers } from "../../hooks/useDragHandlers";
import { type IColumn } from "../../types/board";
import { type Setter } from "../../types/general";
import { type ITask } from "../../types/project";
import { Icons } from "../util/icons";
import "./Card.css";

type Props = {
  task: ITask;
  index: number;
  colIndex: number;
  columns: IColumn[];
  setColumns: Setter<IColumn[]>;
  setTasks: Setter<ITask[]>;
  // openAsModal: OpenModalFunc | null;
};

export default function Card({
  task,
  index,
  colIndex,
  columns,
  setColumns,
  setTasks,
}: // openAsModal,
Props) {
  const handlers = useDragHandlers(columns, setColumns, setTasks);
  const navigate = useNavigate();
  const tagElements = task.tags.map((tag, idx) => {
    return (
      <span
        key={`tag-${tag.text}-${idx}`}
        style={{
          background: tag.color,
          color: "white",
        }}
        className="card--tag"
      >
        {tag.text}
      </span>
    );
  });

  return (
    <div
      className="board--card"
      data-index={index}
      data-col-index={colIndex}
      onMouseDown={() => {
        console.log(task.id);
        navigate(`./${task.id}`);
      }}
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
      {tagElements}
      <span className="card--estimate">{task.estimate}</span>
    </div>
  );
}
