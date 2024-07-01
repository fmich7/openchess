import BotCard from "../components/home/BotCard";
import Banner from "../components/home/FrontPageBanner";

import EasyBotLogo from "../assets/bot1.jpg";
import MediumBotLogo from "../assets/bot2.jpg";
import HardBotLogo from "../assets/bot3.jpg";

const Home = () => {
  return (
    <div className="flex flex-col gap-6">
      <Banner />
      <hr className="border-t border-gray-600"></hr>
      {/* bots section */}
      <div className="flex flex-col items-center h-[430px]">
        <span className="mb-6 text-2xl font-bold text-copy">Leaderboard</span>
        <div className="flex flex-col gap-6">
          <div className="p-5 rounded w-96 bg-foreground text-copy-light">
            <p className="float-left">1. nickname</p>
            <p className="float-right">1237 elo</p>
          </div>
          <div className="p-5 rounded w-96 bg-foreground text-copy-light">
            <p className="float-left">2. nickname</p>
            <p className="float-right">1237 elo</p>
          </div>
          <div className="p-5 rounded w-96 bg-foreground text-copy-light">
            <p className="float-left">3. nickname</p>
            <p className="float-right">1237 elo</p>
          </div>
          <div className="p-5 rounded w-96 bg-foreground text-copy-light">
            <p className="float-left">4. nickname</p>
            <p className="float-right">1237 elo</p>
          </div>

          <BotCard
            name="Random moves"
            description="You should not have any troubble beating it!"
            image={EasyBotLogo}
          />
          <BotCard
            name="Easy bot"
            description="Duel a bot that is a little challenging."
            image={MediumBotLogo}
          />
          <BotCard
            name="Stockfish"
            description="Well, I guess you won't win against that."
            image={HardBotLogo}
          />
        </div>
      </div>
    </div>
  );
};

export default Home;
