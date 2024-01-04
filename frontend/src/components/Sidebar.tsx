import { FaBookOpen } from "react-icons/fa";
import { IoHome } from "react-icons/io5";
import { MdLeaderboard } from "react-icons/md";
import { RiSettings5Fill } from "react-icons/ri";
import logo from "../assets/logo.svg";
const Sidebar = () => {
  return (
    <div className="hidden h-screen pr-8 overflow-y-auto md:flex w-28 bg-primary no-scrollbar">
      <div className="flex h-screen w-[75px] flex-col items-center gap-4 px-2 pb-8 pt-6 ">
        {/* logo */}
        <img src={logo}></img>

        {/* leaderboard,  */}
        <div className="flex flex-col gap-4 mt-4">
          <button className="grid w-12 h-12 text-2xl rounded-lg place-content-center bg-primary-light">
            <IoHome />
          </button>
          <button className="grid w-12 h-12 text-2xl rounded-lg place-content-center bg-primary-light">
            <MdLeaderboard />
          </button>
          <button className="grid w-12 h-12 text-2xl rounded-lg place-content-center bg-primary-light">
            <FaBookOpen />
          </button>
        </div>

        {/* settings */}
        <button className="mt-auto text-2xl">
          <RiSettings5Fill />
        </button>
      </div>
    </div>
  );
};

export default Sidebar;
