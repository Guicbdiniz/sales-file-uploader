import React, { useState } from "react";
import Balances from "./screens/Balances";
import Menu from "./screens/Menu";
import Transactions from "./screens/Transactions";
import Upload from "./screens/Upload";

function App(): JSX.Element {
  const [screen, setScreen] = useState<
    "menu" | "upload" | "transactions" | "balances"
  >("menu");

  const goToBalances = () => {
    setScreen("balances");
  };

  const goToTransactions = () => {
    setScreen("transactions");
  };

  const goToUpload = () => {
    setScreen("upload");
  };

  const goToMenu = () => {
    setScreen("menu");
  };

  switch (screen) {
    case "menu": {
      return (
        <Menu
          goToBalances={goToBalances}
          goToTransactions={goToTransactions}
          goToUpload={goToUpload}
        />
      );
    }
    case "upload": {
      return <Upload goBack={goToMenu} />;
    }
    case "transactions": {
      return <Transactions goBack={goToMenu} />;
    }
    case "balances": {
      return <Balances goBack={goToMenu} />;
    }
  }
}

export default App;
