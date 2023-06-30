import { Outlet } from "react-router-dom";
import Sidebar from "../components/Sidebar/Sidebar";
import Topbar from "../components/Topbar/Topbar";
import { useProjects } from "../hooks/projectHooks";
import "./App.css";

function App() {
  const projects = useProjects();

  return (
    <>
      <Topbar projects={projects} />
      <Sidebar projects={projects} />
      <main className="content--wrapper">
        <Outlet />
      </main>
    </>
  );
}

export default App;
