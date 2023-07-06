export type ResponseData<T> = {
  data: T;
};

export type CreateProjectArg = {
  projectName: string;
};

export type Setter<T> = React.Dispatch<React.SetStateAction<T>>;
