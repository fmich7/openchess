import { Chess, Square } from "chess.js";
import { useEffect, useMemo, useRef, useState } from "react";
import { Chessboard } from "react-chessboard";
import { FaChessPawn } from "react-icons/fa";
import GameController from "../components/game/GameController";
import LastMoves from "../components/game/LastMoves";
import Timer from "../components/game/Timer";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

const Game = () => {
  const [game, setGame] = useState(new Chess());
  const playerOneId = "Player";
  const playerTwoId = "Opponent";

  const gc: GameController = useMemo(() => {
    return new GameController(
      true,
      5 * 60,
      0,
      playerOneId,
      playerTwoId,
      "Bullet"
    );
  }, []);
  const boardOrientation = gc.isPlayerOneWhite ? "white" : "black";
  const whiteToMove = useRef(gc.isPlayerOneWhite);

  // make a ai move
  const makeAiMove = () => {
    setTimeout(() => {
      const possibleMoves = game.moves();
      const randomIndex = Math.floor(Math.random() * possibleMoves.length);
      game.move(possibleMoves[randomIndex]);
      setGame(() => new Chess(game.fen()));
      whiteToMove.current = !whiteToMove.current;
    }, 4000);
  };

  // If ai is white -> make first move
  useEffect(() => {
    if (!whiteToMove.current) {
      makeAiMove();
      whiteToMove.current = true;
    }
  }, []);

  // setgame to board after move
  const makeMove = (move: Move | string) => {
    const result = game.move(move);
    setGame(new Chess(game.fen()));
    return result;
  };

  // handle dropping piece on board
  const onDrop = (sourceSquare: Square, targetSquare: Square) => {
    const move = makeMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });
    if (move === null) return false;

    whiteToMove.current = !whiteToMove.current;
    makeAiMove();

    return true;
  };

  return (
    <div className="flex flex-col items-center justify-center gap-2 md:items-start md:flex-row">
      {/* test */}
      <div className={`${whiteToMove.current ? "bg-green-400" : "bg-red-700"}`}>
        {boardOrientation}
      </div>
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
        <div className="w-96">
          <Chessboard
            boardOrientation={boardOrientation}
            position={game.fen()}
            onPieceDrop={onDrop}
            arePiecesDraggable={true}
            autoPromoteToQueen={true}
          />
        </div>
      </div>
      {/* TIMERS SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* opponents timer */}
          <Timer
            nickname={playerTwoId}
            time={gc.time}
            isActive={!whiteToMove.current}
            isWhite={!gc.isPlayerOneWhite}
          />
          <hr className="border-t border-copy-lighter"></hr>
          {/* last moves */}
          <LastMoves />
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
            time={gc.time}
            isWhite={gc.isPlayerOneWhite}
            isActive={whiteToMove.current}
          />
        </div>
      </div>
    </div>
  );
};

export default Game;
