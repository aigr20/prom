import { useMemo, useRef } from "react";
import { useOutletContext } from "react-router-dom";
import { useColumns } from "../../hooks/useColumns";
import { IProjectViewOutletContext } from "../../types/project";
import IssueModal, { OpenModalFunc } from "../IssueModal/IssueModal";
import "./Board.css";
import Card from "./Card";

export default function Board() {
  const { tasks, setTasks } = useOutletContext<IProjectViewOutletContext>();
  const headings = useMemo(() => ["TODO", "In Progress", "Finished"], []);
  const { columns, setColumns } = useColumns({
    headings,
    tasks,
  });
  const issueModal = useRef<OpenModalFunc>(null);

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
                openAsModal={issueModal.current}
              />
            ))}
          </div>
        ))}
      </div>
      <IssueModal
        ref={issueModal}
        // issue={tasks[0] ?? { createdAt: Date.now() }}
      />
    </>
  );
}
