import type { ISprint } from "../../types/sprint";

type Props = {
  sprint: ISprint;
};
export default function Sprint({ sprint }: Props) {
  return <>{sprint.name}
  <ul>
      {sprint.issues.map((issue) => {
        return <li key={issue.id}>{issue.title}</li>
      })}
  </ul></>;
}
