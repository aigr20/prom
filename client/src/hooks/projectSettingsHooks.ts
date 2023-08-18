import { useState } from "react";
import type { Setter } from "../types/general";

type TagCreationReturn = {
  tagText: string;
  setTagText: React.ChangeEventHandler<HTMLInputElement>;
  tagColor: string;
  setTagColor: Setter<string>;
};
export function useTagCreation(): TagCreationReturn {
  const [tagText, setTagText] = useState("");
  const [tagColor, setTagColor] = useState("");

  return {
    tagText,
    setTagText: (event: React.ChangeEvent<HTMLInputElement>) =>
      setTagText(event.currentTarget.value),
    tagColor,
    setTagColor,
  };
}
