import { Chess, Square } from "chess.js";
import { useEffect, useState } from "react";
import { Chessboard } from "react-chessboard";
import { BoardOrientation } from "react-chessboard/dist/chessboard/types";
import GameController from "../game/GameController";

interface Move {
  from: string;
  to: string;
  promotion: string;
}

// fen -> setState
interface BoardProps {
  fen?: string | undefined;
  gameController?: GameController;
  boardOrientation?: BoardOrientation | undefined;
  draggable: boolean;
}

const Board = (props: BoardProps) => {
  const [game, setGame] = useState(new Chess());
  const gameController = props.gameController; // bind for shortcut

  useEffect(() => {
    if (gameController && !gameController.playerOneToMove) {
      setTimeout(makeRandomMove, 400);
      gameController.playerOneToMove = true;
    }
  }, []);

  function checkIfGameIsOver() {
    if (game.isGameOver() || game.isDraw()) {
      console.log("GAME OVEWWR!");
      if (gameController) gameController.gameStatus = "Game over!";
      if (gameController) console.log(gameController.gameStatus);

      return true;
    }
  }

  const makeMove = (move: Move | string) => {
    const result = game.move(move);
    setGame(new Chess(game.fen()));
    return result;
  };

  function makeRandomMove() {
    const possibleMoves = game.moves();
    if (checkIfGameIsOver() || possibleMoves.length === 0) return; // exit if the game is over
    const randomIndex = Math.floor(Math.random() * possibleMoves.length);
    makeMove(possibleMoves[randomIndex]);
    return true;
  }

  function onDrop(sourceSquare: Square, targetSquare: Square) {
    checkIfGameIsOver();

    if (gameController && gameController.playerOneToMove) {
      const move = makeMove({
        from: sourceSquare,
        to: targetSquare,
        promotion: "q", // always promote to a queen for example simplicity
      });
      if (move === null) return false;
      gameController.playerOneToMove = false;
      setTimeout(makeRandomMove, 200);
      gameController.playerOneToMove = true;
    }

    return true;
  }

  return (
    <Chessboard
      boardOrientation={props.boardOrientation}
      position={props.fen ? props.fen : game.fen()}
      onPieceDrop={onDrop}
      arePiecesDraggable={props.draggable && !game.isGameOver()}
      autoPromoteToQueen={true}
    />
  );
};

export default Board;
