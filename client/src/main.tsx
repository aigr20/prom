import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import "./index.css";
import App from "./views/App";
import Project, { loader as projectLoader } from "./views/Project";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "project/:projectId",
    loader: projectLoader,
    element: <Project />,
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
