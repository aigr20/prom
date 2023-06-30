import { useLoaderData, type Params } from "react-router-dom";
import Topbar from "../components/Topbar/Topbar";
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
      <div className="layout--sidebar"></div>
      <main className="layout--content-sidebar"></main>
    </div>
  );
}
