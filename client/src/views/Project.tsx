import { Outlet, useLoaderData, type Params } from "react-router-dom";
import Topbar from "../components/Topbar/Topbar";
import ProjectViewBar from "../components/sidebars/ProjectViewBar";
import "../layout/HasSidebar.css";
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
  if (project === null) {
    return <p>No project found!</p>;
  }

  return (
    <div className="layout--wrapper-sidebar">
      <Topbar />
      <ProjectViewBar title={project.name} />
      <main className="layout--content-sidebar">
        <Outlet context={project} />
      </main>
    </div>
  );
}
