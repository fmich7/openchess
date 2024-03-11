import { createContext } from "react";
import { AuthContextInterface } from "../types/types_user";

export const AuthContext = createContext<AuthContextInterface>({
  user: null,
  setUser: () => {},
});
