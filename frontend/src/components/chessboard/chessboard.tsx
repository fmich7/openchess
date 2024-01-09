import { Chess, Square } from "chess.js";
import { useState } from "react";
import { Chessboard } from "react-chessboard";
import { BoardOrientation } from "react-chessboard/dist/chessboard/types";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

// fen -> setState
interface BoardProps {
  fen?: string | undefined;
  boardOrientation?: BoardOrientation | undefined;
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
    const move = makeMove({
      from: sourceSquare,
      to: targetSquare,
      promotion: "q", // always promote to a queen for example simplicity
    });
    if (move === null) return false;
    setTimeout(makeRandomMove, 200);

    return true;
  }

  return (
    <Chessboard
      boardOrientation={props.boardOrientation}
      position={props.fen}
      onPieceDrop={onDrop}
      arePiecesDraggable={props.draggable}
      autoPromoteToQueen={true}
    />
  );
};

export default Board;
