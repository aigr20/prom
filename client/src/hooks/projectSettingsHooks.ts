import React, { useCallback, useEffect, useState, type FormEvent } from "react";
import { useRevalidator } from "react-router-dom";
import { createTag, getProjectTagCounts } from "../services/projects";
import type { Setter } from "../types/general";
import type { IProject, ITag, ITagCount } from "../types/project";

type TagCreationReturn = {
  tagText: string;
  setTagText: React.ChangeEventHandler<HTMLInputElement>;
  tagColor: string;
  setTagColor: Setter<string>;
  tags: ITag[];
  tagCounts: ITagCount[];
  createTagCB: (event: FormEvent<HTMLFormElement>) => void;
};
export function useTagCreation(project: IProject): TagCreationReturn {
  const [tags, setTags] = useState<ITag[]>(project.tags);
  const [tagCounts, setTagCounts] = useState<ITagCount[]>([]);

  useEffect(() => {
    getProjectTagCounts({ projectId: project.id }).then(({ data }) => {
      setTagCounts([...data]);
    });
  }, [project.id]);

  const [tagText, setTagText] = useState("");
  const [tagColor, setTagColor] = useState("");
  const { revalidate } = useRevalidator();

  const createTagCB = useCallback(
    (event: FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      createTag({
        projectId: project.id,
        tagText: tagText,
        tagColor: tagColor,
      }).then(({ data }) => {
        if (!data) return;
        setTags((prev) => [...prev, data]);
        revalidate();
      });
    },
    [project.id, revalidate, tagColor, tagText],
  );

  return {
    tagText,
    setTagText: (event: React.ChangeEvent<HTMLInputElement>) =>
      setTagText(event.currentTarget.value),
    tagColor,
    setTagColor,
    tags,
    tagCounts,
    createTagCB,
  };
}
