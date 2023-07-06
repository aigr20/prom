import { useEffect } from "react";
import { useLocation, type Location } from "react-router-dom";

type OnRouteChangeArgs = {
  callback: (location: Location) => void;
};
export function useOnRouteChange({ callback }: OnRouteChangeArgs) {
  const location = useLocation();
  useEffect(() => callback(location), [callback, location]);
}
