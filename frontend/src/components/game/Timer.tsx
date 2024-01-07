interface TimerProps {
  nickname: string;
  time: number;
  isWhite: boolean;
}

const Timer = (props: TimerProps) => {
  const colorEmoji = props.isWhite ? "⚪" : "⚫";
  return (
    <div className="flex">
      <p className="flex-1 ">{colorEmoji + " " + props.nickname}</p>
      <div className="px-3 py-1 text-2xl font-bold rounded bg-border">
        01:35
      </div>
    </div>
  );
};

export default Timer;
