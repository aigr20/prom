import { useEffect, useState } from "react";
import { IColumn } from "../types/board";
import { Setter } from "../types/general";
import { ITask } from "../types/project";

type UseColumnsArgs = {
  headings: string[];
  tasks: ITask[];
};
type UseColumnsReturn = {
  columns: IColumn[];
  setColumns: Setter<IColumn[]>;
};
export function useColumns({
  headings,
  tasks,
}: UseColumnsArgs): UseColumnsReturn {
  const [columns, setColumns] = useState<IColumn[]>(new Array(headings.length));
  useEffect(() => {
    setColumns((cols) => {
      headings.forEach((heading, idx) => {
        cols[idx] = {
          heading,
          issues: tasks.filter((task) => task.status === heading),
        };
      });
      return [...cols];
    });
  }, [headings, tasks]);

  return { columns, setColumns };
}
