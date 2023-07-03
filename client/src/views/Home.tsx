import { useMemo, useState } from "react";
import CreateProjectForm from "../components/CreateProjectForm/CreateProjectForm";
import Topbar from "../components/Topbar/Topbar";
import HomeViewBar from "../components/sidebars/HomeViewBar";
import { ProjectContext } from "../context/ProjectContext";
import { IProject } from "../types/project";
import "./Home.css";

function Home() {
  const [projects, setProjects] = useState<IProject[]>([]);
  const [showCreateProject, setShowCreateProject] = useState(false);
  const layout = useMemo(() => {
    return showCreateProject
      ? "layout--wrapper-leftrightbar"
      : "layout--wrapper-sidebar";
  }, [showCreateProject]);

  return (
    <ProjectContext.Provider value={{ projects, setProjects }}>
      <div className={layout}>
        <Topbar />
        <HomeViewBar showFormCB={setShowCreateProject} />
        <main className="layout--content">Some content here soon</main>
        {showCreateProject && (
          <CreateProjectForm showFormSetter={setShowCreateProject} />
        )}
      </div>
    </ProjectContext.Provider>
  );
}

export default Home;
