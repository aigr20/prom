import { Link } from "react-router-dom";
import "./Topbar.css";

export default function Topbar() {
  return (
    <nav className="topbar--wrapper">
      <h1 className="topbar--logo">
        <Link to="/">Prom</Link>
      </h1>
    </nav>
  );
}
