import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthContext } from "./context/AuthContext";
import { useAuth } from "./hooks/useAuth";
import Game from "./pages/Game";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import Login from "./pages/Login";

function App() {
  const { user, setUser } = useAuth();
  return (
    <BrowserRouter>
      <AuthContext.Provider value={{ user, setUser }}>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />}></Route>
            <Route path="game/:gameID" element={<Game />}></Route>
            <Route path="login" element={<Login />}></Route>
          </Route>
        </Routes>
      </AuthContext.Provider>
    </BrowserRouter>
  );
}

export default App;
