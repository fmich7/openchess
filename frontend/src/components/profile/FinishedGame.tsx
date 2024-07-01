import React from "react";

export type FGameProps = {
  white_player_id: number;
  black_player_id: number;
  created_at: string;
  is_ranked: boolean;
  game_ended: boolean;
  game_outcome: string;
  game_status: string;
  game_type: string;
  id: number;
  time: number;
  time_added: number;
};
const FinishedGame: React.FC<{ games: FGameProps[] }> = ({ games }) => {
  const gameList: JSX.Element[] = [];

  for (let i = 0; i < games.length; i++) {
    gameList.push(
      <div key={i} className="p-2 border-2 rounded bg-background border-border">
        <div className="flex flex-row justify-between">
          <div>
            <p>⚪ {games[i].white_player_id}</p>
            <p>⚫ {games[i].black_player_id}</p>
          </div>
          <div className="text-xs text-right">
            <p>{games[i].is_ranked ? "Ranked" : "Unranked"}</p>
            <p>
              {games[i].time} + {games[i].time_added}
            </p>
            <p>Played: {games[i].created_at.split("T")[0]}</p>
          </div>
        </div>
      </div>
    );
  }
  if (gameList.length === 0)
    gameList.push(<div className="text-center">No games found</div>);

  return gameList;
};

export default FinishedGame;
