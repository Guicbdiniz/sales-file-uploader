import React from "react";
import "./style.css";

type Props = {
  goBack: () => void;
};

const GoBackButton: React.FC<Props> = (props) => {
  const { goBack } = props;

  return (
    <div className="goBack" onClick={goBack}>
      Go Back
    </div>
  );
};

export default GoBackButton;
