import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Backlog from "./components/Backlog/Backlog";
import Board from "./components/Board/Board";
import IssueModal, { issueLoader } from "./components/IssueModal/IssueModal";
import "./index.css";
import "./layout/base.css";
import Home from "./views/Home";
import Project, { loader as projectLoader } from "./views/Project";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "project/:projectId",
    loader: projectLoader,
    element: <Project />,
    children: [
      {
        path: "backlog",
        element: <Backlog />,
      },
      {
        path: "board",
        element: <Board />,
        children: [
          {
            path: ":issueId",
            element: <IssueModal />,
            loader: issueLoader,
          },
        ],
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
