export type ResponseData<T> = {
  data: T;
};

export type ProjectIDArg = {
  projectId: number;
};

export type CreateProjectArg = {
  projectName: string;
};

export type Setter<PropName extends string, T> = {
  [P in PropName]: React.Dispatch<React.SetStateAction<T>>;
};
