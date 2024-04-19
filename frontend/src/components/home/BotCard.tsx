import axios from "axios";
import { useNavigate } from "react-router-dom";

interface BotCardProps {
  name: string;
  description: string;
  image: string;
}

const BotCard = (props: BotCardProps) => {
  const navigate = useNavigate();
  const hostGame = async () => {
    try {
      const response = await axios.post(`/api/game`, {
        hostID: 1,
        whitePlayerID: 1,
        blackPlayerID: 2,
        isRanked: true,
        time: 60000,
        timeAdded: 10,
        gameType: "ranked",
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
    <div className="flex flex-col w-56 shadow rounded-xl bg-foreground">
      <div className="p-4">
        <img src={props.image}></img>
        <hr className="my-3"></hr>
        <p className="font-bold text-copy">{props.name}</p>
        <p className="text-copy-lighter">{props.description}</p>

        <button
          onClick={hostGame}
          className="w-full h-10 mt-4 border-2 rounded bg-primary text-background border-primary-light"
        >
          Play now
        </button>
      </div>
    </div>
  );
};

export default BotCard;
