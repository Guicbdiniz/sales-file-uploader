import React, { useState } from "react";
import Menu from "./screens/Menu";
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
    default:
      return <div></div>;
  }
}

export default App;
