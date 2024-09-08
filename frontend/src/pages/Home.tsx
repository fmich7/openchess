import BotCard from "../components/home/BotCard";
import Banner from "../components/home/FrontPageBanner";

import axios from "axios";
import { useEffect, useState } from "react";
import EasyBotLogo from "../assets/bot1.jpg";
import MediumBotLogo from "../assets/bot2.jpg";
import HardBotLogo from "../assets/bot3.jpg";

const fetchLeaderboard = () => {
  return axios
    .get("/api/leaderboard")
    .then((response) => {
      const data = response.data;
      console.log(data);
      return data.map((profile: any, index: number) => (
        <div
          key={index}
          className="p-5 rounded w-96 bg-foreground text-copy-light"
        >
          <span className="float-left">{index + 1}.</span>
          <span className="float-left ml-3">🇵🇱 {profile.nickname}</span>
          <p className="float-right">
            {profile.elo} elo | {profile.games_won}/{profile.games_played} WR
          </p>
        </div>
      ));
    })
    .catch((err) => {
      console.error(err);
      return [];
    });
};

const Home = () => {
  const [leaderboard, setLeaderboard] = useState<JSX.Element[]>([]);
  useEffect(() => {
    fetchLeaderboard().then((profiles) => {
      setLeaderboard(profiles);
    });
  }, []);
  return (
    <div className="flex flex-col gap-6">
      <Banner />
      <hr className="border-t border-gray-600"></hr>
      {/* bots section */}
      <div className="flex flex-col items-center h-[230px]">
        <span className="mb-6 text-2xl font-bold text-copy">Leaderboard</span>
        <div className="flex flex-col gap-6">{leaderboard}</div>
      </div>
    </div>
  );
};

export default Home;
