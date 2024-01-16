const LastMoves = ({ moves }: { moves: string[] }) => {
  const rows = Math.ceil(moves.length / 2);

  const generateGrid = () => {
    const grid = [];

    for (let i = 0; i < rows; i++) {
      const move1 = moves[i * 2] || ""; // Use empty string if undefined
      let move2 = moves[i * 2 + 1] || ""; // Use empty string if undefined

      // If the current iteration is the last one and moves.length is odd, store the odd move
      if (i === rows - 1 && moves.length % 2 !== 0) {
        move2 = "";
      }

      grid.push(
        <div key={i} className="grid grid-cols-5">
          <div className="text-center bg-border">{i + 1}</div>
          <div className="col-span-2">
            <span className="ml-5">{move1}</span>
          </div>
          <div className="col-span-2">
            <span className="ml-5">{move2}</span>
          </div>
        </div>
      );
    }

    return grid;
  };

  return <div className="h-32 overflow-auto">{generateGrid()}</div>;
};

export default LastMoves;
