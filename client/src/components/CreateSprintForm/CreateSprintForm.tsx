import { useState } from "react";
import { padZero } from "../util/date";

function formatDate(date: Date): string {
  return `${date.getFullYear()}-${padZero(date.getMonth() + 1)}-${padZero(
    date.getDate(),
  )}`;
}

type Props = {
  projectId: number;
};
export default function CreateSprintForm({ projectId }: Props) {
  const [isShown, setIsShown] = useState(false);
  const [sprintName, setSprintName] = useState("");
  const [start, setStart] = useState(new Date());

  return (
    <>
      <button onClick={() => setIsShown((prev) => !prev)}>
        {isShown ? "Cancel" : "Create Sprint"}
      </button>
      <form
        style={{ display: isShown ? "" : "none" }}
        onSubmit={(e) => e.preventDefault()}
      >
        <input
          type="text"
          placeholder="Sprint name"
          value={sprintName}
          onChange={(e) => setSprintName(e.currentTarget.value)}
        />
        <input
          type="date"
          value={formatDate(start)}
          min={formatDate(new Date())}
          onChange={(e) => setStart(new Date(e.currentTarget.value))}
        />
        <input type="date" />
        <button type="submit">Create Sprint</button>
      </form>
    </>
  );
}
