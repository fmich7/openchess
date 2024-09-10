import axios from "axios";
import { Chess, Square } from "chess.js";
import { useCallback, useEffect, useState } from "react";
import { Chessboard } from "react-chessboard";
import { FaChessPawn } from "react-icons/fa";
import { useParams } from "react-router-dom";
import LastMoves from "../components/game/LastMoves";
import Timer from "../components/game/Timer";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

type GameDetails = {
  time: number;
  timeAdded: number;
  whitePlayerID: number;
  blackPlayerID: number;
  isRanked: boolean;
  gameType: string;
  moveHistory: string;
  gameFen: string;
  whiteTime: number;
  blackTime: number;
};

type OpponentData = {
  nickname: string;
  whiteColor: boolean;
};

const defaultGameDetails: GameDetails = {
  time: 0,
  timeAdded: 0,
  whitePlayerID: 0,
  blackPlayerID: 0,
  isRanked: true,
  gameType: "",
  moveHistory: "",
  gameFen: "",
  blackTime: 0,
  whiteTime: 0,
};

const defaultOpponentData: OpponentData = {
  nickname: "Opponent",
  whiteColor: false,
};

const Game = () => {
  const { gameID } = useParams();
  const [game, setGame] = useState(new Chess());
  const [fen, setFen] = useState(game.fen());
  const [gameDetails, setGameDetails] =
    useState<GameDetails>(defaultGameDetails);
  const [opponentData, setOpponentData] =
    useState<OpponentData>(defaultOpponentData);
  const [myTurn, setMyTurn] = useState<boolean>(true);
  const myID = 1;

  useEffect(() => {
    // Fetch game state from server
    axios
      .get(`/api/live/${gameID}`)
      .then((response) => {
        const data = response.data;
        console.log(data);
        // SET INITIAL INFORMATION
        setGameDetails({
          time: data["details"]["time"] / 1000,
          timeAdded: data["details"]["time_added"] / 1000,
          whitePlayerID: data["details"]["white_player_id"],
          blackPlayerID: data["details"]["black_player_id"],
          isRanked: data["details"]["is_ranked"],
          gameType: data["details"]["game_type"],
          moveHistory: data["details"]["move_history"],
          gameFen: data["details"]["game_fen"],
          whiteTime: data["details"]["time"],
          blackTime: data["details"]["time"],
        });
        setGame(() => new Chess(data["details"]["game_fen"]));
        setFen(data["details"]["game_fen"]);
        setMyTurn(myID == data["details"]["white_player_id"]);
        setOpponentData({
          whiteColor: myID != data["details"]["white_player_id"],
          nickname: "opp",
        });
      })
      .catch((error) => {
        throw new Error("Error fetching data:" + error);
      });
  }, []);

  const makeMove = useCallback(
    (move: Move | string) => {
      try {
        const result = game.move(move);
        if (result === null) {
          return null;
        }

        setGameDetails({ ...gameDetails });
        setFen(game.fen());
        return result;
      } catch (e) {
        return null;
      }
    },
    [game]
  );

  // handle dropping piece on board
  const onDrop = (sourceSquare: Square, targetSquare: Square) => {
    const move = makeMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });
    if (move === null) return false;
    setMyTurn(!myTurn);
    axios
      .put(`/api/live/${gameID}`, {
        move: move?.lan,
        user_id: myID,
      })
      .then((resp) => {
        const data = resp.data;
        console.log(data);
        setGame(new Chess(data["fen"]));
        setFen(data["fen"]);
        setMyTurn(true);
        setGameDetails({
          ...gameDetails,
          moveHistory: data["move_history"],
          gameFen: data["fen"],
          whiteTime: data["white_time"],
          blackTime: data["black_time"],
        });
      })
      .catch((e) => {
        throw new Error(e);
      });

    return true;
  };

  return (
    <div className="flex flex-col items-center justify-center gap-2 md:items-start md:flex-row">
      {/* INFO SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* game info */}
          <div className="flex gap-4">
            <div className="grid items-center text-3xl">
              <FaChessPawn />
            </div>
            <div>
              <p>
                {(gameDetails.time / 60).toString() +
                  "+" +
                  gameDetails.timeAdded.toString()}{" "}
                • {gameDetails.isRanked ? "Ranked" : "Casual"} •{" "}
                {gameDetails.gameType}
              </p>
              <span className=" text-copy-lighter">Just started now</span>
            </div>
          </div>
          {/* match players */}
          <div>
            {/* {(!gc.isPlayerOneWhite ? "⚪" : "⚫") + " " + playerTwoId}
            <br />
            {(gc.isPlayerOneWhite ? "⚪" : "⚫") + " " + playerOneId} */}
          </div>
          <hr className="border-t border-copy-lighter"></hr>
          {/* game status */}
          {/* <div className="grid justify-center">{gc.gameStatus}</div> */}
        </div>
      </div>
      {/* BOARD SECTION */}
      <div className="grid justify-center">
        <div className="w-96 text-fuchsia-50">
          <Chessboard
            boardOrientation={
              myID == gameDetails.whitePlayerID ? "white" : "black"
            }
            position={fen}
            onPieceDrop={onDrop}
            arePiecesDraggable={true}
            autoPromoteToQueen={true}
          />
          <p className="text-sm">FEN: {fen}</p>
        </div>
      </div>
      {/* TIMERS SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* opponents timer */}
          <Timer
            nickname={opponentData.nickname}
            time={
              opponentData.whiteColor
                ? gameDetails.whiteTime
                : gameDetails.blackTime
            }
            isActive={!myTurn}
            isWhite={opponentData.whiteColor}
          />
          <hr className="border-t border-copy-lighter"></hr>
          {/* last moves */}
          <LastMoves moves={gameDetails.moveHistory} />
          {/* offer and resign buttons */}
          <div className="grid gap-2">
            <button className="h-12 rounded-lg bg-border hover:bg-background">
              Offer draw
            </button>
            <button className="h-12 rounded-lg bg-border hover:bg-background">
              Resign
            </button>
          </div>
          {/* players timer */}
          <hr className="border-t border-copy-lighter"></hr>
          <Timer
            nickname={"sad"}
            time={
              opponentData.whiteColor
                ? gameDetails.blackTime
                : gameDetails.whitePlayerID
            }
            isActive={myTurn}
            isWhite={!opponentData.whiteColor}
          />
        </div>
      </div>
    </div>
  );
};

export default Game;
