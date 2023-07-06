import { useIssueCreation } from "../../hooks/issueHooks";
import { Setter } from "../../types/general";
import { ITask } from "../../types/project";
import "./CreateIssueForm.css";

type Props = {
  projectId: number;
  issues: ITask[];
  setIssues: Setter<ITask[]>;
  setShowCreateIssue: Setter<boolean>;
};

export default function CreateIssueForm({
  projectId,
  issues,
  setIssues,
  setShowCreateIssue,
}: Props) {
  const { title, setTitle, description, setDescription, onSubmit } =
    useIssueCreation({
      projectId,
      issues,
      setIssues,
      setShowCreateIssue,
    });

  return (
    <form className="rightbar--wrapper create_issue--form" onSubmit={onSubmit}>
      <h2>Create issue</h2>
      <input
        type="text"
        placeholder="Title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />
      <textarea
        rows={10}
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      ></textarea>
      <button type="submit">Create Issue</button>
    </form>
  );
}