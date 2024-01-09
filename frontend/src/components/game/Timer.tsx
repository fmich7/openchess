import { useEffect, useState } from "react";

interface TimerProps {
  nickname: string;
  time: number;
  isWhite: boolean;
  isActive: boolean;
}

function formatTime(seconds: number) {
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const remainingSeconds = seconds % 60;

  const parts = [];

  if (hours > 0) {
    parts.push(hours < 10 ? `0${hours}` : `${hours}`);
  }

  if (minutes > 0 || hours > 0) {
    parts.push(minutes < 10 ? `0${minutes}` : `${minutes}`);
  }

  parts.push(
    remainingSeconds < 10 ? `0${remainingSeconds}` : `${remainingSeconds}`
  );

  return parts.join(":");
}

const Timer = (props: TimerProps) => {
  const colorEmoji = props.isWhite ? "⚪" : "⚫";
  const [time, setTime] = useState(props.time);
  const [formattedTime, setFormattedTime] = useState(formatTime(time));

  useEffect(() => {
    if (!props.isActive || props.time <= 0) return;

    const timer = setInterval(() => {
      setTime((time) => time - 1);
      setFormattedTime(formatTime(time));
    }, 1000);

    return () => clearInterval(timer);
  }, [formattedTime, props, time]);

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
