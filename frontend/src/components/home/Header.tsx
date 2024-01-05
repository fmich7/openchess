import { FaSearch, FaUser } from "react-icons/fa";
import { IoIosNotifications } from "react-icons/io";
import { Link } from "react-router-dom";

const Header = () => {
  return (
    <div className="flex items-center justify-between">
      {/* Header LOGO */}
      <Link to={"/"}>
        <div className="flex my-4 text-4xl font-bold text-copy">
          openchess
          <span id="logo_underscore" className="text-primary-light">
            __
          </span>
        </div>
      </Link>

      {/* header mid links */}
      <div className="flex gap-6 text-copy-light">
        <Link to={"/"}>Home</Link>
        <Link to={"play"}>Play</Link>
        <Link to={"tools"}>Tools</Link>
      </div>

      {/* header right menu */}
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
