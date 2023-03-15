import React, { useEffect, useState } from "react";
import "./style.css";

const Loading: React.FC = () => {
  const [dots, setDots] = useState(".");
  useEffect(() => {
    const timer = setInterval(() => {
      setDots((dots) => {
        switch (dots) {
          case "":
            return ".";
          case ".":
            return "..";
          case "..":
            return "...";
          default:
            return "";
        }
      });
    }, 500);
    return () => clearInterval(timer);
  }, []);

  return <div className="loading">Loading{dots}</div>;
};

export default Loading;
