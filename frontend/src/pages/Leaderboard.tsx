import axios from "axios";
import { useEffect, useState } from "react";
import { FaUserTie } from "react-icons/fa";
import { IoMdPersonAdd } from "react-icons/io";

type Profile = {
  nickname: string;
  firstname: string;
  lastname: string;
  id: string;
  elo: number;
};

const CreateProfileList = (profiles: Profile[]) => {
  const list: JSX.Element[] = [];
  profiles.forEach((profile) => {
    list.push(<div>{profile.firstname}</div>);
  });
  return list;
};

const Leaderboard = () => {
  const [profiles, setProfiles] = useState<JSX.Element[]>();

  useEffect(() => {
    axios
      .get(`/api/leaderboard/`)
      .then((response) => {
        const lb = response.data["leaderboard"];
        const profiles: Profile[] = [];

        for (let i = 0; i < lb.length; i++) {
          profiles.push({
            nickname: lb[0].nickname,
            firstname: lb[0].firstname,
            lastname: lb[0].lastname,
            id: lb[0].id,
            elo: lb[0].elo,
          });
        }

        const profileList: JSX.Element[] = CreateProfileList(profiles);
        setProfiles(profileList);
      })
      .catch((error) => {
        throw new Error("Error fetching data:" + error);
      });
  }, []);

  return (
    <div className="flex justify-center flex-1 w-full text-copy-light">
      <div className="flex flex-col w-3/4 gap-4 p-4 border rounded shadow-2xl sm:max-w-xl bg-foreground border-border">
        {/* PROFILE */}
        <div className="flex flex-col justify-between gap-4 sm:flex-row">
          {/* LEFT SECTION */}
          <div className="flex flex-row gap-4">
            <div className="text-9xl">
              <FaUserTie />
            </div>
          </div>

          {/* RIGHT SECTION */}
          <div>
            <button className="flex gap-3 px-3 py-1 align-middle border-2 rounded bg-primary text-background border-primary-light">
              <div className="flex flex-col justify-center">
                <IoMdPersonAdd size={"1.25rem"} />
              </div>
              <p>Follow</p>
            </button>
          </div>
        </div>
        <hr className="border-t border-copy-lighter"></hr>
        {/* GAMES */}
        <p className="text-copy-light">Recent games</p>
        <div className="flex flex-col gap-2">{profiles}</div>
      </div>
    </div>
  );
};

export default Leaderboard;
