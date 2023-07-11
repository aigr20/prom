import { useProjectCreation } from "../../hooks/projectHooks";
import { Setter } from "../../types/general";
import { Icons } from "../util/icons";
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
      <div className="form_title--wrapper">
        <h2 className="form--title">Create Project</h2>
        <button className="form--close" onClick={() => showFormSetter(false)}>
          {Icons.close}
        </button>
      </div>
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
