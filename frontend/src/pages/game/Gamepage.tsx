import Board from "../../components/chessboard/Chessboard";

const Gamepage = () => {
  return (
    <div>
      <Board
        draggable={true}
        fen={"8/1r1p4/p6k/PP1BK3/1p1P1R2/6Pp/7p/2q2n2 w - - 0 1"}
      />
    </div>
  );
};

export default Gamepage;
