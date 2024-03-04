import axios from "axios";
import { Chess, Square } from "chess.js";
import { useEffect, useMemo, useRef, useState } from "react";
import { Chessboard } from "react-chessboard";
import { FaChessPawn } from "react-icons/fa";
import GameController from "../components/game/GameController";
import LastMoves from "../components/game/LastMoves";
import Timer from "../components/game/Timer";
import { config } from "../config";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

const Game = () => {
  // game variables
  const isMounted = useRef(false);
  const game = useRef(new Chess());
  const whiteToMove = useRef(true);
  const [isGameOver, setIsGameOver] = useState(false);

  const playerOneId = "Player";
  const playerTwoId = "Opponent";
  const gc: GameController = useMemo(() => {
    return new GameController(true, 10, 0, playerOneId, playerTwoId, "Bullet");
  }, []);

  const [movesList, setMovesList] = useState<string[]>([]);
  const [whiteTurn, setWhiteTurn] = useState(whiteToMove.current);

  // players
  const [whiteTime, setWhiteTime] = useState(gc.time * 1000);
  const [blackTime, setBlackTime] = useState(gc.time * 1000);

  const boardOrientation = gc.isPlayerOneWhite ? "white" : "black";

  useEffect(() => {
    // init
    if (isMounted.current === false) {
      isMounted.current = true;

      (async () => {
        const response = await axios.get(`${config.apiURL}/game`);
        console.log(response.data);
      })();

      // ai start game
      if (gc.isPlayerOneWhite === false && whiteToMove.current) makeAiMove();
    }

    // Cleanup
    return () => {
      clearInterval(intervalRef.current);
    };
  }, []);

  useEffect(() => {
    if (blackTime === 0 || whiteTime == 0) {
      console.log("Ran out of time!");
      gameOver();
    }
  }, [whiteTime, blackTime]);

  // make a ai move
  const makeAiMove = () => {
    setTimeout(() => {
      const possibleMoves = game.current.moves();
      const randomIndex = Math.floor(Math.random() * possibleMoves.length);

      makeMove(possibleMoves[randomIndex]);
    }, 400);
  };

  // make move
  const makeMove = (move: Move | string) => {
    const currMove = game.current.move(move);
    if (currMove === null) {
      return null;
    }

    if (game.current.isGameOver()) gameOver();
    else updateTimers();
    setMovesList((moves) => [
      ...moves,
      game.current.history()[game.current.history().length - 1],
    ]);
    whiteToMove.current = !whiteToMove.current;
    setWhiteTurn(whiteToMove.current);

    return currMove;
  };

  // handle dropping piece on board
  const onDrop = (sourceSquare: Square, targetSquare: Square) => {
    if (
      (gc.isPlayerOneWhite && !whiteTurn) ||
      (!gc.isPlayerOneWhite && whiteTurn)
    )
      return false;

    const move = makeMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });
    if (move === null) return false;

    makeAiMove();
    return true;
  };

  const gameOver = () => {
    clearInterval(intervalRef.current);
    setIsGameOver(true);
    game.current.isGameOver = () => true;
  };

  // TIMERS
  const intervalRef = useRef<NodeJS.Timeout>();
  const updateTimers = () => {
    clearInterval(intervalRef.current);

    if (isGameOver) return;

    intervalRef.current = setInterval(() => {
      if (whiteToMove.current) {
        setWhiteTime((t) => Math.max(0, t - 100));
      } else {
        setBlackTime((t) => Math.max(0, t - 100));
      }
    }, 100);
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
            boardOrientation={boardOrientation}
            position={game.current.fen()}
            onPieceDrop={onDrop}
            arePiecesDraggable={!isGameOver}
            autoPromoteToQueen={true}
          />
          {game.current.fen()}
        </div>
      </div>
      {/* TIMERS SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* opponents timer */}
          <Timer
            nickname={playerTwoId}
            time={!gc.isPlayerOneWhite ? whiteTime : blackTime}
            isActive={
              gc.isPlayerOneWhite ? !whiteToMove.current : whiteToMove.current
            }
            isWhite={!gc.isPlayerOneWhite}
          />
          <hr className="border-t border-copy-lighter"></hr>
          {/* last moves */}
          <LastMoves moves={movesList} />
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
            time={gc.isPlayerOneWhite ? whiteTime : blackTime}
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
