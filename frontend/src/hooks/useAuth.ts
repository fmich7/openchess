import axios from "axios";
import { useCookies } from "react-cookie";

export const useAuth = () => {
  const [cookies, setCookie, removeCookie] = useCookies();

  const login = (token: string) => {
    setCookie("x-jwt-token", token, {
      secure: true,
    });
  };

  const logout = () => {
    removeCookie("x-jwt-token");
  };

  const userToken = () => {
    return cookies["x-jwt-token"];
  };

  const getUserID = async () => {
    try {
      const response = await axios.get(`/api/whoami`, {
        withCredentials: true,
      });
      console.log("User id: " + response.data["id"]);
      return response.data["id"];
    } catch (error) {
      console.error("Error fetching data");
      throw error;
    }
  };

  return { login, logout, userToken, getUserID };
};
