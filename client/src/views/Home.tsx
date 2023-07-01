import { Outlet } from "react-router-dom";
import Topbar from "../components/Topbar/Topbar";
import HomeViewBar from "../components/sidebars/HomeViewBar";
import { useProjects } from "../hooks/projectHooks";
import "../layout/HasSidebar.css";
import "./Home.css";

function Home() {
  const projects = useProjects();

  return (
    <div className="layout--wrapper-sidebar">
      <Topbar />
      <HomeViewBar projects={projects} />
      <main className="layout--content-sidebar">
        <Outlet />
      </main>
    </div>
  );
}

export default Home;
