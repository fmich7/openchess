import axios from "axios";
import { useEffect, useState } from "react";
import { FaUserTie } from "react-icons/fa";
import { IoMdPersonAdd } from "react-icons/io";
import { useParams } from "react-router-dom";
import FinishedGame, { FGameProps } from "../components/profile/FinishedGame";
type UserDataType = {
  nickname: string;
  firstname: string;
  lastname: string;
  id: string;
  elo: number;
  createdAt: string;
};

const Profile = () => {
  const { profileID } = useParams();
  const [gameProps, setGameProps] = useState<FGameProps[]>([]);
  const [userData, setUserData] = useState<UserDataType | null>(null);
  useEffect(() => {
    axios
      .get(`/api/profile/${profileID}`)
      .then((response) => {
        const user = response.data["account"];
        const games = response.data["games"];

        const newGameProps: FGameProps[] = [];
        for (let i = 0; i < games.length; i++) {
          newGameProps.push({
            white_player_id: games[i].white_player_id,
            black_player_id: games[i].black_player_id,
            created_at: games[i].created_at,
            is_ranked: games[i].is_ranked,
            game_ended: games[i].game_ended,
            game_outcome: games[i].game_outcome,
            game_status: games[i].game_status,
            game_type: games[i].game_type,
            id: games[i].white_player_id,
            time: games[i].white_player_id,
            time_added: games[i].white_player_id,
          });
        }
        setGameProps(newGameProps);

        setUserData({
          nickname: user["nickname"],
          firstname: user["firstName"],
          lastname: user["lastName"],
          id: user["id"],
          elo: user["elo"],
          createdAt: user["createdAt"],
        });
      })
      .catch((error) => {
        throw new Error("Error fetching data:" + error);
      });
  }, []);

  if (userData === null)
    return (
      <div>
        <div className="flex justify-center flex-1 w-full text-center text-copy-light">
          <div className="flex flex-col w-3/4 gap-4 p-4 border rounded shadow-2xl sm:max-w-xl bg-foreground border-border">
            USER NOT FOUND
          </div>
        </div>
      </div>
    );
  else
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
              <div>
                <p>{userData.nickname}</p>
                <p className="text-copy-lighter">
                  🇵🇱 {userData.firstname} {userData.lastname}
                </p>
                <div className="text-[12px]">
                  <p className="mt-2">{userData.elo} elo</p>
                  <p>W/L: 43/123 | 53.22% wr</p>
                  <p>Joined in: {userData.createdAt.split("T")[0]}</p>
                </div>
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
          <div className="flex flex-col gap-2">
            {<FinishedGame games={gameProps} />}
          </div>
        </div>
      </div>
    );
};

export default Profile;
