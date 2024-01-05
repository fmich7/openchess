import { BrowserRouter, Route, Routes } from "react-router-dom";
import Game from "./pages/Game";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />}></Route>
          <Route path="game" element={<Game />}></Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
