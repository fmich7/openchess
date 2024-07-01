import axios from "axios";
import { Chess, Square } from "chess.js";
import { useEffect, useMemo, useRef, useState } from "react";
import { Chessboard } from "react-chessboard";
import { FaChessPawn } from "react-icons/fa";
import { useParams } from "react-router-dom";
import GameController from "../components/game/GameController";
import LastMoves from "../components/game/LastMoves";
import Timer from "../components/game/Timer";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

const Game = () => {
  const { gameID } = useParams();
  const whiteToMove = useRef(true);
  const [game, setGame] = useState(new Chess());
  const playerOneId = "Player";
  const playerTwoId = "Opponent";
  const gc: GameController = useMemo(() => {
    return new GameController(true, 10, 0, playerOneId, playerTwoId, "Bullet");
  }, []);

  useEffect(() => {
    // if (isMounted.current === false) {
    // isMounted.current = true;
    // Fetch game state from server
    axios
      .get(`/api/live/${gameID}`)
      .then((response) => {
        const data = response.data;
        // FEN
        if (data["details"]["game_fen"]) {
          setGame(new Chess(data["details"]["game_fen"]));
        } else {
          throw new Error("Something is wrong with server...");
        }
      })
      .catch((error) => {
        throw new Error("Error fetching data:" + error);
      });
    // }
  }, [gameID]);

  const makeMove = (move: Move | string) => {
    const gameCopy = new Chess(game.fen());
    const result = gameCopy.move(move);
    if (result === null) {
      return null;
    }

    setGame(gameCopy);
    return result;
  };

  // handle dropping piece on board
  const onDrop = (sourceSquare: Square, targetSquare: Square) => {
    const fen_before = game.fen();

    const move = makeMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });
    if (move === null) return false;

    axios
      .put(`/api/live/${gameID}`, {
        move: move?.lan,
      })
      .then((resp) => {
        const data = resp.data;
        if (data["fen"]) {
          setGame(new Chess(data["fen"]));
        }
      })
      .catch((error) => {
        setGame(new Chess(fen_before));
        throw new Error(error);
      });

    return true;
  };

  // TIMERS
  // const intervalRef = useRef<NodeJS.Timeout>();
  // const updateTimers = () => {
  //   clearInterval(intervalRef.current);

  //   if (isGameOver) return;

  //   intervalRef.current = setInterval(() => {
  //     if (whiteToMove.current) {
  //       setWhiteTime((t) => Math.max(0, t - 100));
  //     } else {
  //       setBlackTime((t) => Math.max(0, t - 100));
  //     }
  //   }, 100);
  // };

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
                {(gc.time / 60).toString() + "+" + gc.timeAdded.toString()} •{" "}
                {gc.isRanked ? "Ranked" : "Casual"} • {gc.gameType}
              </p>
              <span className=" text-copy-lighter">Just started now</span>
            </div>
          </div>
          {/* match players */}
          <div>
            {(!gc.isPlayerOneWhite ? "⚪" : "⚫") + " " + playerTwoId}
            <br />
            {(gc.isPlayerOneWhite ? "⚪" : "⚫") + " " + playerOneId}
          </div>
          <hr className="border-t border-copy-lighter"></hr>
          {/* game status */}
          <div className="grid justify-center">{gc.gameStatus}</div>
        </div>
      </div>
      {/* BOARD SECTION */}
      <div className="grid justify-center">
        <div className="w-96 text-fuchsia-50">
          <Chessboard
            boardOrientation={"white"}
            position={game.fen()}
            onPieceDrop={onDrop}
            arePiecesDraggable={true}
            autoPromoteToQueen={true}
          />
          {game.fen()}
        </div>
      </div>
      {/* TIMERS SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* opponents timer */}
          <Timer
            nickname={playerTwoId}
            time={1000}
            isActive={
              gc.isPlayerOneWhite ? !whiteToMove.current : whiteToMove.current
            }
            isWhite={!gc.isPlayerOneWhite}
          />
          <hr className="border-t border-copy-lighter"></hr>
          {/* last moves */}
          <LastMoves moves={["1", "a", "3"]} />
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
            nickname={playerOneId}
            time={1000}
            isWhite={gc.isPlayerOneWhite}
            isActive={
              gc.isPlayerOneWhite ? whiteToMove.current : !whiteToMove.current
            }
          />
        </div>
      </div>
    </div>
  );
};

export default Game;
