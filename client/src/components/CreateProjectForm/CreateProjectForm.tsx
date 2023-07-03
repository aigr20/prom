import { useProjectCreation } from "../../hooks/projectHooks";
import { Setter } from "../../types/general";
import "./CreateProjectForm.css";

export default function CreateProjectForm({
  showFormSetter,
}: Setter<"showFormSetter", boolean>) {
  const { projectName, setProjectName, submitCallback } = useProjectCreation({
    showFormSetter,
  });
  return (
    <form className="rightbar--wrapper" onSubmit={submitCallback}>
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
