import { PropsWithChildren } from "react";
import Spinner from "./Spinner/Spinner";

type Props<T> = {
  couldBeNull: T | null;
  spinnerSize?: number;
} & PropsWithChildren;

export default function SpinIfNull<T>({
  couldBeNull,
  spinnerSize,
  children,
}: Props<T>) {
  if (couldBeNull === null) {
    return <Spinner size={spinnerSize} />;
  }
  return <>{children}</>;
}
