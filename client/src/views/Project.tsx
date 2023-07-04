import { useMemo, useState } from "react";
import { Outlet, useLoaderData, type Params } from "react-router-dom";
import CreateIssueForm from "../components/CreateIssueForm/CreateIssueForm";
import Topbar from "../components/Topbar/Topbar";
import ProjectViewBar from "../components/sidebars/ProjectViewBar";
import { useProjectTasks } from "../hooks/projectHooks";
import { getProject } from "../services/projects";
import { IProject } from "../types/project";

type LoaderProps = {
  params: Params<"projectId">;
};

export async function loader({
  params,
}: LoaderProps): Promise<{ project: IProject | null }> {
  const { data } = await getProject({ projectId: Number(params.projectId) });
  return { project: data };
}

export default function Project() {
  const { project } = useLoaderData() as Awaited<ReturnType<typeof loader>>;
  const tasks = useProjectTasks({ projectId: project?.id });
  const [showCreateIssue, setShowCreateIssue] = useState(false);
  const layout = useMemo(() => {
    return showCreateIssue
      ? "layout--wrapper-leftrightbar"
      : "layout--wrapper-sidebar";
  }, [showCreateIssue]);

  if (project === null) {
    return <p>No project found!</p>;
  }

  return (
    <div className={layout}>
      <Topbar />
      <ProjectViewBar title={project.name} />
      <main className="layout--content-sidebar">
        <Outlet context={{ project, setShowCreateIssue, tasks }} />
      </main>
      {showCreateIssue && (
        <CreateIssueForm
          projectId={project.id}
          setShowCreateIssue={setShowCreateIssue}
        />
      )}
    </div>
  );
}
