// import { useEffect, useState } from "react";

interface TimerProps {
  nickname: string;
  time: number;
  isWhite: boolean;
  isActive: boolean;
}

function formatTime(duration: number): string {
  const hours = Math.floor(duration / (1000 * 60 * 60));
  const minutes = Math.floor((duration % (1000 * 60 * 60)) / (1000 * 60));
  const seconds = Math.floor((duration % (1000 * 60)) / 1000);

  const formattedHours = hours.toString().padStart(2, "0");
  const formattedMinutes = minutes.toString().padStart(2, "0");
  const formattedSeconds = seconds.toString().padStart(2, "0");

  if (hours > 0) {
    return `${formattedHours}:${formattedMinutes}:${formattedSeconds}`;
  } else {
    return `${formattedMinutes}:${formattedSeconds}`;
  }
}

const Timer = (props: TimerProps) => {
  const colorEmoji = props.isWhite ? "⚪" : "⚫";
  const formattedTime = formatTime(props.time);
  // const [time, setTime] = useState(props.time);

  return (
    <div className="flex ">
      <p className="flex-1 ">{colorEmoji + " " + props.nickname}</p>
      <div
        className={`px-3 py-1 text-2xl font-bold rounded ${
          props.isActive ? "bg-secondary-dark" : "bg-border"
        }`}
      >
        {formattedTime}
      </div>
    </div>
  );
};

export default Timer;
