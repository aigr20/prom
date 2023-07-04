import "./Spinner.css";

type Props = {
  size?: number;
};

export default function Spinner({ size = 80 }: Props) {
  return (
    <div
      style={{ width: `${size}px`, height: `${size}px` }}
      className="spinner--spinner"
    ></div>
  );
}
