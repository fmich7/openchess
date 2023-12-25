import Tile from "./tile";

interface Piece {
  image: string;
  x: number;
  y: number;
}

const pieces: Piece[] = [];
let activePiece: HTMLElement | null = null;

// fill board with pieces
function CreateInitialBoardState() {
  // white and black color
  for (let s = 0; s < 2; s++) {
    const Y = s === 0 ? 6 : 1;
    const piece_color = Y === 6 ? "white" : "black";

    // pawns
    for (let i = 0; i < 8; i++) {
      pieces.push({
        image: `/src/assets/chess_pieces/${piece_color}_pawn.png`,
        x: i,
        y: Y,
      });
    }

    // rooks
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_rook.png`,
      x: 0,
      y: Y === 1 ? Y - 1 : Y + 1,
    });
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_rook.png`,
      x: 7,
      y: Y === 1 ? Y - 1 : Y + 1,
    });

    // knights
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_knight.png`,
      x: 1,
      y: Y === 1 ? Y - 1 : Y + 1,
    });
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_knight.png`,
      x: 6,
      y: Y === 1 ? Y - 1 : Y + 1,
    });

    // bishops
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_bishop.png`,
      x: 2,
      y: Y === 1 ? Y - 1 : Y + 1,
    });
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_bishop.png`,
      x: 5,
      y: Y === 1 ? Y - 1 : Y + 1,
    });

    // queen
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_queen.png`,
      x: 3,
      y: Y === 1 ? Y - 1 : Y + 1,
    });
    // king
    pieces.push({
      image: `/src/assets/chess_pieces/${piece_color}_king.png`,
      x: 4,
      y: Y === 1 ? Y - 1 : Y + 1,
    });
  }
}

// handle grabbing a piece
function GrabPiece(e: React.MouseEvent) {
  const element = e.target as HTMLElement;
  if (element.classList.contains("piece")) {
    activePiece = element;
    const x = e.clientX - 35;
    const y = e.clientY - 35;

    element.style.position = "absolute";
    element.style.left = `${x}px`;
    element.style.top = `${y}px`;
  }
}

// handle moving a piece
function MovePiece(e: React.MouseEvent) {
  if (activePiece) {
    const x = e.clientX - 35;
    const y = e.clientY - 35;

    activePiece.style.position = "absolute";
    activePiece.style.left = `${x}px`;
    activePiece.style.top = `${y}px`;
  }
}

// handle dropping a piece
function DropPiece() {
  activePiece = null;
}

// Chessboard component
function Chessboard() {
  const tiles = [];
  const columns = "ABCDEFGHI";
  CreateInitialBoardState();

  // generate board
  for (let i = 7; i >= 0; i--) {
    for (let j = 0; j < 8; j++) {
      let image = undefined;

      pieces.forEach((p) => {
        if (p.x === j && p.y === i) image = p.image;
      });

      tiles.push(
        <Tile
          col={j}
          row={i}
          image={image}
          key={columns[j] + (i + 1).toString()}
        />
      );
    }
  }
  return (
    <div
      onMouseDown={(e) => GrabPiece(e)}
      onMouseMove={(e) => MovePiece(e)}
      onMouseUp={() => DropPiece()}
      className="grid grid-cols-8 h-[32rem] w-[32rem]"
    >
      {tiles}
    </div>
  );
}

export default Chessboard;
