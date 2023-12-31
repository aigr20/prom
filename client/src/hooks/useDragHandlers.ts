import { useRef, type MouseEvent } from "react";
import { updateStatus } from "../services/issues";
import { type IColumn } from "../types/board";
import { type Setter } from "../types/general";
import { type ITask } from "../types/project";

type useDragHandlerReturn<E extends Element> = {
  onMouseMove: (event: MouseEvent<E>) => void;
  onMouseLeave: (event: MouseEvent<E>) => void;
  onMouseUp: (event: MouseEvent<E>) => void;
};

export function useDragHandlers<E extends HTMLElement>(
  columns: IColumn[],
  setColumns: Setter<IColumn[]>,
  setTasks: Setter<ITask[]>,
): useDragHandlerReturn<E> {
  const dragged = useRef<{ data: ITask; element: E }>();

  function move(x: number, y: number) {
    if (dragged.current) {
      const element = dragged.current.element;
      element.style.left = `${x - element.offsetWidth / 2}px`;
      element.style.top = `${y - element.offsetHeight / 2}px`;
    }
  }

  function onMouseMove(event: MouseEvent<E>) {
    if (event.buttons !== 1) return;
    if (!dragged.current) {
      const target = event.currentTarget;
      const col = Number(target.dataset["colIndex"]);
      const card = Number(target.dataset["index"]);
      const { width, height } = target.getBoundingClientRect();

      dragged.current = {
        data: columns[col].issues[card],
        element: target,
      };

      target.classList.add("board--card-dragged");
      target.style.width = `${width}px`;
      target.style.height = `${height}px`;
    }
    move(event.pageX, event.pageY);
  }

  function onMouseLeave(event: MouseEvent<E>) {
    move(event.pageX, event.pageY);
  }

  function onMouseUp(event: MouseEvent<E>) {
    if (dragged.current) {
      dragged.current.element.classList.remove("board--card-dragged");
      const column = document
        .elementFromPoint(event.clientX, event.clientY)
        ?.closest(".board--column") as HTMLElement | null;

      if (column) {
        const oldColIndex = Number(dragged.current.element.dataset["colIndex"]);
        const colIndex = Number(column.dataset["colIndex"] ?? oldColIndex);
        columns[colIndex].issues.push(dragged.current.data);
        columns[oldColIndex].issues.splice(
          Number(dragged.current.element.dataset.index),
          1,
        );

        updateStatus({
          issueId: dragged.current.data.id,
          newStatus: columns[colIndex].heading,
        });
        setColumns([...columns]);
        setTasks((tasks) => {
          const task = tasks.find(
            (task) => task.id === dragged.current?.data.id,
          );
          if (task) task.status = columns[colIndex].heading;
          return [...tasks];
        });
      }

      dragged.current.element.removeAttribute("style");
      dragged.current = undefined;
    }
  }

  return { onMouseLeave, onMouseMove, onMouseUp };
}
