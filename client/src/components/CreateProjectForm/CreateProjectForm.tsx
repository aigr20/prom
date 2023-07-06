import { useProjectCreation } from "../../hooks/projectHooks";
import { Setter } from "../../types/general";
import "./CreateProjectForm.css";

type Props = {
  showFormSetter: Setter<boolean>;
};

export default function CreateProjectForm({ showFormSetter }: Props) {
  const { projectName, setProjectName, submitCallback } = useProjectCreation({
    showFormSetter,
  });
  return (
    <form
      className="create_project--form rightbar--wrapper"
      onSubmit={submitCallback}
    >
      <h2>Create Project</h2>
      <input
        type="text"
        placeholder="Name"
        value={projectName}
        onChange={(e) => setProjectName(e.target.value)}
      />
      <button type="submit">Create Project</button>
    </form>
  );
}
