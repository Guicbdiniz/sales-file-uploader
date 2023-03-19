import React, { useEffect, useState } from "react";
import ErrorText from "../../components/ErrorText";
import GoBackButton from "../../components/GoBackButton";
import Loading from "../../components/Loading";
import api from "../../services/api";
import { Balance } from "../../utils/types";
import "./style.css";

type Props = {
  goBack: () => void;
};

const UNKNOWN_FETCH_ERROR =
  "There was an error while fetching the balances. Try again.";

const Balances: React.FC<Props> = (props) => {
  const { goBack } = props;
  const [loading, setLoading] = useState(true);
  const [errorText, setErrorText] = useState("");
  const [balances, setBalances] = useState<Balance[]>([]);

  useEffect(() => {
    (async () => {
      try {
        const balances = await api.fetchBalances();
        setBalances(balances);
      } catch (e) {
        console.error("Error captured while fetching balances:\n", e);
        setErrorText(UNKNOWN_FETCH_ERROR);
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  const renderContent = () => {
    if (loading) {
      return <Loading />;
    }

    if (errorText) {
      return <ErrorText text={errorText} />;
    }

    return (
      <div className="data">
        {balances.length > 0 ? (
          balances.map((balance, index) => (
            <div className="balance" key={index}>
              {balance.isProducer ? "PRODUCER " : "NOT PRODUCER "}
              {balance.name} - {balance.balance} centavos.
            </div>
          ))
        ) : (
          <div className="transaction">No balances were saved yet.</div>
        )}
      </div>
    );
  };

  return (
    <div className="balances">
      <GoBackButton goBack={goBack} />
      <h1>Balances</h1>
      {renderContent()}
    </div>
  );
};
export default Balances;
