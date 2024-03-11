import { Outlet } from "react-router-dom";
import Sidebar from "../components/Sidebar";
import Header from "../components/home/Header";

const Layout = () => {
  return (
    <div className="flex flex-row w-[100vh_+25px] h-screen bg-background">
      {/* left siderbar */}
      <Sidebar />

      {/* main */}
      <div className="bg-background md:ml-[-25px] w-full overflow-y-scroll md:rounded-l-3xl shadow-xl">
        <div className="flex flex-col h-screen gap-6 p-8">
          <Header />

          <Outlet />
        </div>
      </div>
    </div>
  );
};

export default Layout;
