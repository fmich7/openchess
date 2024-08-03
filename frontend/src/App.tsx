import { BrowserRouter, Route, Routes } from "react-router-dom";
import Game from "./pages/Game";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import Leaderboard from "./pages/Leaderboard";
import Login from "./pages/Login";
import Profile from "./pages/Profile";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />}></Route>
          <Route path="live/:gameID" element={<Game />}></Route>
          <Route path="profile/:profileID" element={<Profile />}></Route>
          <Route path="login/" element={<Login />}></Route>
          <Route path="leaderboard/" element={<Leaderboard />}></Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
};

export default App;
