import { Outlet } from "react-router-dom";
import Topbar from "../components/Topbar/Topbar";
import HomeViewBar from "../components/sidebars/HomeViewBar";
import { ProjectContext } from "../context/ProjectContext";
import { useProjects } from "../hooks/projectHooks";
import "../layout/HasSidebar.css";
import "./Home.css";

function Home() {
  const projectContext = useProjects();

  return (
    <ProjectContext.Provider value={projectContext}>
      <div className="layout--wrapper-sidebar">
        <Topbar />
        <HomeViewBar />
        <main className="layout--content">
          <Outlet />
        </main>
      </div>
    </ProjectContext.Provider>
  );
}

export default Home;
