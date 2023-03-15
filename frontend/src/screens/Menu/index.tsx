import React from "react";
import Footer from "../../components/Footer";
import "./style.css";

type Props = {
  goToUpload: () => void;
  goToTransactions: () => void;
  goToBalances: () => void;
};

const Menu: React.FC<Props> = (props) => {
  const { goToBalances, goToTransactions, goToUpload } = props;

  return (
    <div className="menu">
      <h1 className="title">Select an action:</h1>
      <div className="options">
        <p className="option" onClick={goToUpload}>Upload</p>
        <p className="option" onClick={goToTransactions}>Transactions</p>
        <p className="option" onClick={goToBalances}>Balances</p>
      </div>
      <Footer />
    </div>
  );
};

export default Menu;
