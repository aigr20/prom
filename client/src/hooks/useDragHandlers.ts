import { useRef, type MouseEvent } from "react";
import { ITask } from "../types/project";

type useDragHandlerReturn<E extends Element> = {
  onMouseDown: (event: MouseEvent<E>) => void;
  onMouseMove: (event: MouseEvent<E>) => void;
  onMouseLeave: (event: MouseEvent<E>) => void;
  onMouseUp: (event: MouseEvent<E>) => void;
};

export function useDragHandlers<E extends HTMLElement>(
  tasks: ITask[],
): useDragHandlerReturn<E> {
  const dragged = useRef<{ data: ITask; element: E }>();

  function move(x: number, y: number) {
    if (dragged.current) {
      const element = dragged.current.element;
      element.style.left = `${x - element.offsetWidth / 2}px`;
      element.style.top = `${y - element.offsetHeight / 2}px`;
    }
  }

  function onMouseDown(event: MouseEvent<E>) {
    if (event.button !== 0) {
      return;
    }
    const target = event.currentTarget;
    const col = Number(target.dataset["colIndex"]);
    const card = Number(target.dataset["index"]);
    const { width, height } = target.getBoundingClientRect();

    dragged.current = {
      data: tasks[card],
      element: target,
    };

    target.classList.add("board--card-dragged");
    target.style.width = `${width}px`;
    target.style.height = `${height}px`;
    move(event.pageX, event.pageY);
  }

  function onMouseMove(event: MouseEvent<E>) {
    move(event.pageX, event.pageY);
  }

  function onMouseLeave(event: MouseEvent<E>) {
    move(event.pageX, event.pageY);
  }

  function onMouseUp(event: MouseEvent<E>) {
    if (dragged.current) {
      dragged.current.element.classList.remove("board--card-dragged");
      dragged.current.element.removeAttribute("style");
      dragged.current = undefined;
    }
  }

  return { onMouseDown, onMouseLeave, onMouseMove, onMouseUp };
}
