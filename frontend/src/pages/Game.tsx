import { FaChessPawn } from "react-icons/fa";
import Board from "../components/chessboard/Chessboard";
import GameController from "../components/game/GameController";
const Game = () => {
  const gameController: GameController = new GameController(5);
  console.log(gameController);

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
              <p>1+0 • Casual • Bullet</p>
              <span className=" text-copy-lighter">21 minutes ago</span>
            </div>
          </div>
          {/* match players */}
          <div>
            ⚪ Anonymous <br />⚫ Anonymous
          </div>
          <hr className="border-t border-copy-lighter"></hr>
          {/* game status */}
          <div className="grid justify-center">Game aborted</div>
        </div>
      </div>
      {/* BOARD SECTION */}
      <div className="grid justify-center">
        <div className="w-96">
          <Board draggable={true} />
        </div>
      </div>
      {/* TIMERS SECTION */}
      <div className="flex-1 w-full md:max-w-96 text-copy-light">
        <div className="flex flex-col gap-4 p-4 border rounded shadow-2xl bg-foreground border-border">
          {/* opponents timer */}
          <div className="flex">
            <p className="flex-1 ">⚪ Anonymous</p>
            <div className="px-3 py-1 text-2xl font-bold rounded bg-border">
              01:35
            </div>
          </div>
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
          <div className="flex">
            <p className="flex-1 ">⚪ Anonymous</p>
            <div className="px-3 py-1 text-2xl font-bold rounded bg-secondary-dark">
              01:35
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Game;
