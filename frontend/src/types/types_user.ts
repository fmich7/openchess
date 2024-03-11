export interface User {
  id: string;
  authToken?: string;
}

export interface AuthContextInterface {
  user: User | null;
  setUser: (user: User | null) => void;
}
