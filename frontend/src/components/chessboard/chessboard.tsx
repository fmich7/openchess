import { Chess, Square } from "chess.js";
import { useState } from "react";
import { Chessboard } from "react-chessboard";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

// fen -> setState
interface BoardProps {
  fen?: string | undefined;
  draggable: boolean;
}

const Board = (props: BoardProps) => {
  const [game, setGame] = useState(new Chess());

  const makeMove = (move: Move | string) => {
    const result = game.move(move);
    setGame(new Chess(game.fen()));
    return result;
  };

  function makeRandomMove() {
    const possibleMoves = game.moves();
    if (game.isGameOver() || game.isDraw() || possibleMoves.length === 0)
      return; // exit if the game is over
    const randomIndex = Math.floor(Math.random() * possibleMoves.length);
    makeMove(possibleMoves[randomIndex]);
    return true;
  }

  function onDrop(sourceSquare: Square, targetSquare: Square) {
    if (game.isGameOver() || game.isDraw()) {
      console.log("GAME OVEWWR!");
      return true;
    }

    const move = makeMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });

    // illegal move
    if (move === null) return false;
    setTimeout(makeRandomMove, 200);
    return true;
  }

  return (
    <div className="w-96">
      <Chessboard
        position={props.fen ? props.fen : game.fen()}
        onPieceDrop={onDrop}
        arePiecesDraggable={props.draggable && !game.isGameOver()}
        autoPromoteToQueen={true}
      />
    </div>
  );
};

export default Board;
