import { Outlet } from "react-router-dom";
import Sidebar from "../components/Sidebar/Sidebar";
import Topbar from "../components/Topbar/Topbar";
import { useProjects } from "../hooks/projectHooks";
import "../layout/HasSidebar.css";
import "./Home.css";

function Home() {
  const projects = useProjects();

  return (
    <div className="layout--wrapper-sidebar">
      <Topbar projects={projects} />
      <Sidebar projects={projects} />
      <main className="content--wrapper">
        <Outlet />
      </main>
    </div>
  );
}

export default Home;
