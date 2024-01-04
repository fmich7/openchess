import { FaSearch, FaUser } from "react-icons/fa";
import { IoIosNotifications } from "react-icons/io";

const Header = () => {
  return (
    <div className="flex items-center justify-between">
      <div className="flex my-4 text-4xl font-bold text-copy">
        openchess
        <span id="logo_underscore" className="text-primary-light">
          __
        </span>
      </div>

      <div className="flex gap-6 text-copy-light">
        <div>Home</div>
        <div>Play</div>
        <div>Tools</div>
      </div>

      <div className="flex gap-4 shrink-0">
        <button className="grid items-center justify-center w-10 h-10 text-2xl rounded-full bg-primary text-background">
          <IoIosNotifications />
        </button>
        <button className="grid items-center justify-center w-10 h-10 text-lg rounded-full bg-primary text-background">
          <FaSearch />
        </button>
        <button className="grid items-center justify-center w-10 h-10 text-lg rounded-full bg-primary text-background">
          <FaUser />
        </button>
      </div>
    </div>
  );
};

export default Header;
