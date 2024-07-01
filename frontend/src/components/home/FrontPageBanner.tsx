import Board from "../chessboard/Chessboard";

const Banner = () => {
  return (
    <div className="flex flex-col items-center gap-6">
      <div className="md:flex md:gap-6">
        {/* left panel */}
        <div className="flex justify-center align-middle w-[340px]">
          <Board
            draggable={false}
            fen={"8/1r1p4/p6k/PP1BK3/1p1P1R2/6Pp/7p/2q2n2 w - - 0 1"}
          />
        </div>
        {/* right panel */}

        <div className="flex flex-col mt-6 md:w-96 md:mt-0 text-copy">
          <div>
            <p className="mb-4 text-3xl font-bold">
              Ready for an ad-free and tracker-free chess experience?
            </p>

            <p className="mb-6 text-lg leading-normal">
              Immerse yourself in the world of chess without interruptions. We
              prioritize your privacy by not storing any personal data.
            </p>

            <p className="text-lg text-copy-lighter">
              10 users ● 132 games ● 0 trackers
            </p>
          </div>

          <div className="flex flex-col gap-4 mt-6 font-medium md:flex-row">
            <button className="flex-1 py-3 rounded bg-primary text-background md:py-4">
              Join now
            </button>
            <button className="flex-1 py-3 mt-4 rounded bg-primary text-background md:py-4 md:mt-0">
              Quickplay
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Banner;
