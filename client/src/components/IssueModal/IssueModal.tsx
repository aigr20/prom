import { useLoaderData, useNavigate, type Params } from "react-router-dom";
import { useIssueModal } from "../../hooks/issueModalHooks";
import { getIssue } from "../../services/issues";
import type { ITask } from "../../types/project";
import { formatDate } from "../util/date";
import { Icons } from "../util/icons";
import "./IssueModal.css";
import TagDropdown from "./TagDropdown";

type LoaderProps = {
  params: Params<"issueId">;
};

export async function issueLoader({
  params,
}: LoaderProps): Promise<{ issue: ITask | null }> {
  const { data } = await getIssue({ issueId: Number(params.issueId) });
  return { issue: data };
}

export default function IssueModal() {
  const navigate = useNavigate();
  const { issue } = useLoaderData() as Awaited<ReturnType<typeof issueLoader>>;
  const { issueValues, modifyFunction, onModalClose, setIssue } =
    useIssueModal(issue);

  return (
    <div className="issue--modal">
      <input
        className="heading"
        name="issue-heading"
        value={issueValues?.title ?? ""}
        onChange={(e) => modifyFunction("title", e)}
      />
      <button
        className="close-button"
        onClick={() => {
          onModalClose();
          navigate("..");
        }}
      >
        {Icons.close}
      </button>
      <textarea
        className="description"
        onChange={(e) => modifyFunction("description", e)}
        value={issueValues?.description ?? ""}
      />
      <div className="tags">
        {issueValues?.tags.map((tag) => {
          return (
            <span
              key={`modal-tag-${tag.id}-${issue?.id}`}
              className="tag"
              style={{ background: tag.color }}
            >
              {tag.text}
            </span>
          );
        })}
      </div>
      <TagDropdown
        issueId={issue?.id}
        tags={issueValues?.tags}
        setIssue={setIssue}
      />
      <span className="created" title="Skapat">
        {issue?.createdAt && formatDate(issue.createdAt)}
      </span>
      <span className="updated" title="Uppdaterat">
        {issue?.updatedAt && formatDate(issue.updatedAt)}
      </span>
      <input
        className="estimate"
        name="issue-estimate"
        type="number"
        min={0}
        step={1}
        value={String(issueValues?.estimate ?? 0)}
        onChange={(e) => modifyFunction("estimate", e)}
      />
    </div>
  );
}
