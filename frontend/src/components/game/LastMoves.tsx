const LastMoves = ({ moves }: { moves: string }) => {
  const generateGrid = () => {
    const grid: JSX.Element[] = [];
    if (moves == "") return <div></div>;
    moves.split(" ").forEach((value, index) => {
      let i = 0;
      for (; i < value.length; i++) {
        if (value[i] >= "0" && value[i] <= "9") break;
      }

      grid.push(
        <div key={index} className="grid grid-cols-5">
          <div className="text-center bg-border">{index + 1}</div>
          <div className="col-span-2">
            <span className="ml-5">{value.substring(0, i + 1)}</span>
          </div>
          <div className="col-span-2">
            <span className="ml-5">{value.substring(i + 1, value.length)}</span>
          </div>
        </div>
      );
    });
    return grid;
  };

  return <div className="h-32 overflow-auto">{generateGrid()}</div>;
};

export default LastMoves;
