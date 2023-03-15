import React from "react";
import "./style.css";

type Props = {
  text: string;
};

const ErrorText: React.FC<Props> = (props) => {
  const { text } = props;
  return <div className="error-text">{text}</div>;
};

export default ErrorText;
