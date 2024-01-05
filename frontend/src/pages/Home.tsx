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
        <span className="mb-6 text-2xl font-bold text-copy">
          Challenge our bots
        </span>
        <div className="grid grid-cols-1 gap-4 md:grid-cols-3 ">
          <BotCard
            name="Random moves"
            description="You should not have any troubble beating it!"
            image={EasyBotLogo}
            redirect="game"
          />
          <BotCard
            name="Easy bot"
            description="Duel a bot that is a little challenging."
            image={MediumBotLogo}
            redirect="game"
          />
          <BotCard
            name="Stockfish"
            description="Well, I guess you won't win against that."
            image={HardBotLogo}
            redirect="game"
          />
        </div>
      </div>
    </div>
  );
};

export default Home;
