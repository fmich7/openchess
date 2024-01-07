import { FaChessPawn } from "react-icons/fa";
import Board from "../components/chessboard/Chessboard";
import GameController from "../components/game/GameController";
import Timer from "../components/game/Timer";

const Game = () => {
  const playerOneId = "Player";
  const playerTwoId = "Opponent";
  const gameController: GameController = new GameController(
    true,
    5,
    0,
    playerOneId,
    playerTwoId,
    "Bullet"
  );

  const boardOrientation = gameController.isPlayerOneWhite ? "white" : "black";

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
                {gameController.time.toString() +
                  "+" +
                  gameController.timeAdded.toString()}{" "}
                • {gameController.isRanked ? "Ranked" : "Casual"} •{" "}
                {gameController.gameType}
              </p>
              <span className=" text-copy-lighter">Just started now</span>
            </div>
          </div>
          {/* match players */}
          <div>
            {(!gameController.isPlayerOneWhite ? "⚪" : "⚫") +
              " " +
              playerTwoId}
            <br />
            {(gameController.isPlayerOneWhite ? "⚪" : "⚫") +
              " " +
              playerOneId}
          </div>
          <hr className="border-t border-copy-lighter"></hr>
          {/* game status */}
          <div className="grid justify-center">{gameController.gameStatus}</div>
        </div>
      </div>
      {/* BOARD SECTION */}
      <div className="grid justify-center">
        <div className="w-96">
          <Board
            gameController={gameController}
            boardOrientation={boardOrientation}
            draggable={true}
          />
        </div>
      </div>
      {/* TIMERS SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* opponents timer */}
          <Timer
            nickname={playerTwoId}
            time={gameController.time}
            isWhite={!gameController.isPlayerOneWhite}
          />
          <hr className="border-t border-copy-lighter"></hr>
          {/* last moves */}
          <div className="grid grid-cols-5">
            <div className="text-center bg-border">1</div>
            <div className="col-span-2">
              <span className="ml-5">e1</span>
            </div>
            <div className="col-span-2">
              <span className="ml-5">e1</span>
            </div>
            <div className="text-center bg-border">1</div>
            <div className="col-span-2">
              <span className="ml-5">e1</span>
            </div>
            <div className="col-span-2">
              <span className="ml-5">e1</span>
            </div>
            <div className="text-center bg-border">1</div>
            <div className="col-span-2">
              <span className="ml-5">e1</span>
            </div>
            <div className="col-span-2">
              <span className="ml-5">e1</span>
            </div>
          </div>
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
            time={gameController.time}
            isWhite={gameController.isPlayerOneWhite}
          />
        </div>
      </div>
    </div>
  );
};

export default Game;
