import Board from "../chessboard/Chessboard";

const Banner = () => {
  return (
    <div>
      <div className="justify-center md:flex md:gap-6">
        {/* left panel */}
        <div className="flex justify-center align-middle w-96">
          <Board
            draggable={true}
            fen={"8/1r1p4/p6k/PP1BK3/1p1P1R2/6Pp/7p/2q2n2 w - - 0 1"}
          />
        </div>
        {/* right panel */}
        <div className="flex flex-col mt-6 md:w-96 md:mt-0 text-copy">
          <div>
            <p className="text-3xl font-bold ">
              Ready for an ad-free and tracker-free chess experience?
            </p>

            <p className="mt-6">
              Immerse yourself in the world of chess without interruptions. We
              prioritize your privacy by not storing any personal data.
            </p>
          </div>

          <div className="flex gap-4 mt-6">
            <button className="flex-1 h-10 rounded bg-primary-dark text-[#fff]">
              Play now
            </button>
            <button className="flex-1 h-10 rounded bg-primary text-background">
              Get Started
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Banner;
