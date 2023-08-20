import { useEffect, useState } from "react";
import type { Setter } from "../../../types/general";
import "./ColorPicker.css";

const colors: { color: string; text: string }[] = [
  { color: "#ff0000", text: "red" },
  { color: "#00ff00", text: "green" },
  { color: "#0000ff", text: "blue" },
  { color: "#ff00ff", text: "pink" },
  { color: "#ffff00", text: "yellow" },
  { color: "#00ffff", text: "light blue" },
];

type Props = {
  onValueChange: Setter<string>;
};

export default function ColorPicker({ onValueChange }: Props) {
  const [selected, setSelected] = useState(0);
  const [isShown, setIsShown] = useState(false);

  useEffect(() => {
    onValueChange(colors[selected].color);
  }, [onValueChange, selected]);

  return (
    <div className="color-picker--wrapper">
      <button
        className="color-picker--toggle"
        type="button"
        onClick={() => setIsShown((wasShown) => !wasShown)}
      >
        Color
        <svg width={10} height={10} viewBox="0 0 2.6458 2.6458">
          <rect fill={colors[selected].color} height="100%" width="100%" />
        </svg>
      </button>
      <ul
        className="color-picker--menu"
        style={{ display: isShown ? "" : "none" }}
      >
        {colors.map(({ color, text }, idx) => {
          return (
            <li key={color} className="color-picker--item">
              <button
                className="color-picker--button"
                onClick={() => setSelected(idx)}
                type="button"
                style={{ background: color }}
                aria-label={text}
              ></button>
            </li>
          );
        })}
      </ul>
    </div>
  );
}
