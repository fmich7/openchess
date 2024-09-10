import { useEffect, useState } from "react";

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

  return hours > 0
    ? `${formattedHours}:${formattedMinutes}:${formattedSeconds}`
    : `${formattedMinutes}:${formattedSeconds}`;
}

const Timer = ({ nickname, time, isWhite, isActive }: TimerProps) => {
  const [currentTime, setCurrentTime] = useState(time);

  useEffect(() => {
    setCurrentTime(time);
  }, [time]);

  useEffect(() => {
    let timer: NodeJS.Timeout | null = null;

    if (isActive && currentTime > 0) {
      timer = setInterval(() => {
        setCurrentTime((prevTime) => prevTime - 1000);
      }, 1000);
    }

    return () => {
      if (timer) {
        clearInterval(timer);
      }
    };
  }, [isActive, currentTime]);

  const colorEmoji = isWhite ? "⚪" : "⚫";
  const formattedTime = formatTime(currentTime);

  return (
    <div className="flex">
      <p className="flex-1">{colorEmoji + " " + nickname}</p>
      <div
        className={`px-3 py-1 text-2xl font-bold rounded ${
          isActive ? "bg-secondary-dark" : "bg-border"
        }`}
      >
        {formattedTime}
      </div>
    </div>
  );
};

export default Timer;
