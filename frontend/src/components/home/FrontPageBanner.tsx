import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Board from "../chessboard/Chessboard";

type WebsiteStats = {
  usersCount: number;
  gamesCount: number;
};

const Banner = () => {
  const [stats, setStats] = useState<WebsiteStats>({
    usersCount: 0,
    gamesCount: 0,
  });

  useEffect(() => {
    axios
      .get("/api/stats")
      .then((response) => {
        const data = response.data;
        setStats({
          usersCount: data["users_count"],
          gamesCount: data["games_count"],
        });
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  const navigate = useNavigate();
  const hostGame = async () => {
    try {
      const response = await axios.post(`/api/game`, {
        hostID: 1,
        opponentID: 0,
        isRanked: true,
        time: 60,
        timeAdded: 10,
      });

      const gameID = response.data["id"];
      if (gameID) {
        navigate(`/live/${gameID}`);
      } else {
        console.error("Response does not contain game id");
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex flex-col items-center gap-6">
      <div className="md:flex md:gap-6">
        {/* left panel */}
        <div className="flex justify-center align-middle w-[340px]">
          <Board
            draggable={false}
            fen={"8/1r1p4/p6k/PP1BK3/1p1P1R2/6Pp/7p/2q2n2 w - - 0 1"}
          />
        </div>
        {/* right panel */}

        <div className="flex flex-col mt-6 md:w-96 md:mt-0 text-copy">
          <div>
            <p className="mb-4 text-3xl font-bold">
              Ready for an ad-free and tracker-free chess experience?
            </p>

            <p className="mb-6 text-lg leading-normal">
              Immerse yourself in the world of chess without interruptions. We
              prioritize your privacy by not storing any personal data.
            </p>

            <p className="text-lg text-copy-lighter">
              {stats.usersCount} users ● {stats.gamesCount} games ● 0 trackers
            </p>
          </div>

          <div className="flex flex-col gap-4 mt-6 font-medium md:flex-row">
            <button className="flex-1 py-3 rounded bg-primary text-background md:py-4">
              Join now
            </button>
            <button
              className="flex-1 py-3 mt-4 rounded bg-primary text-background md:py-4 md:mt-0"
              onClick={hostGame}
            >
              Quickplay
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Banner;
