import { useIssueCreation } from "../../hooks/issueHooks";
import "./CreateIssueForm.css";

type Props = {
  projectId: number;
};

export default function CreateIssueForm({ projectId }: Props) {
  const { title, setTitle, description, setDescription, onSubmit } =
    useIssueCreation({ projectId });

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
