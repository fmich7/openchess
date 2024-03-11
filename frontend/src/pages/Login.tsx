import axios from "axios";
import { useState } from "react";
import { useAuth } from "../hooks/useAuth";
const Login = () => {
  return (
    <div className="flex justify-center items-center w-[100vh_+25px] h-[80%] bg-background">
      {LoginForm()}
    </div>
  );
};

const LoginForm = () => {
  const { login, logout, userToken, getUserID } = useAuth();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const test = async () => {
    const response = await axios.get(`/api/account`);
    console.log(response.data);
  };
  const handleSubmit = () => {
    axios
      .post(`/api/login`, {
        nickname: username,
        password: password,
      })
      .then((response) => {
        console.log(userToken());
        login(response.data["token"]);
        console.log(userToken());
      })
      .catch(() => {
        console.error("Error fetching data");
      });
  };

  return (
    <div className="w-72 h-72 bg-foreground">
      <div className="input-container">
        <label>Username </label>
        <input
          type="text"
          name="uname"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
      </div>
      <div className="input-container">
        <label>Password </label>
        <input
          type="password"
          name="pass"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
      </div>
      <div className="button-container">
        <button onClick={handleSubmit}>SUBMIT</button>
      </div>
      <div className="button-container">
        <button onClick={logout}>Logout</button>
      </div>
      <div className="button-container">
        <button onClick={test}>TEST</button>
      </div>
      <div className="button-container">
        <button onClick={getUserID}>CheckID</button>
      </div>
    </div>
  );
};

export default Login;
