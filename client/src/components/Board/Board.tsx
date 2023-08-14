import { useMemo } from "react";
import { Outlet, useOutletContext, useParams } from "react-router-dom";
import { useColumns } from "../../hooks/useColumns";
import { type IProjectViewOutletContext } from "../../types/project";
import "./Board.css";
import Card from "./Card";

export default function Board() {
  const { tasks, setTasks, project } =
    useOutletContext<IProjectViewOutletContext>();
  const headings = useMemo(() => ["TODO", "In Progress", "Finished"], []);
  const { columns, setColumns } = useColumns({
    headings,
    tasks,
  });
  const params = useParams();
  const activeIssue = Number(params["issueId"]);

  return (
    <>
      <h2>Board</h2>
      <div className="board--wrapper">
        {columns.map(({ heading, issues }, colIdx) => (
          <div
            key={`col-${colIdx}`}
            className="board--column"
            data-col-index={colIdx}
          >
            <h3 className="board--column-title">{heading}</h3>
            {issues.map((task, idx) => (
              <Card
                key={`issue-${task.id}`}
                task={task}
                index={idx}
                colIndex={colIdx}
                columns={columns}
                setColumns={setColumns}
                setTasks={setTasks}
              />
            ))}
          </div>
        ))}
      </div>
      <Outlet
        context={{
          task: tasks.find((task) => task.id === activeIssue),
          setTasks,
          project,
        }}
      />
    </>
  );
}
