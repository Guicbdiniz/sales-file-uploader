import React, { useEffect, useState } from "react";
import ErrorText from "../../components/ErrorText";
import GoBackButton from "../../components/GoBackButton";
import Loading from "../../components/Loading";
import api from "../../services/api";
import { Transaction } from "../../utils/types";
import "./style.css";

type Props = {
  goBack: () => void;
};

const UNKNOWN_FETCH_ERROR =
  "There was an error while fetching the transactions. Try again.";

const Transactions: React.FC<Props> = (props) => {
  const { goBack } = props;
  const [loading, setLoading] = useState(true);
  const [errorText, setErrorText] = useState("");
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    (async () => {
      try {
        const transactions = await api.fetchTransactions();
        setTransactions(transactions);
      } catch (e) {
        console.error("Error captured while fetching transactions:\n", e);
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
        {transactions.map((transaction, index) => (
          <div className="transaction" key={index}>
            {transaction.type} - {transaction.value} centavos of{" "}
            {transaction.product} - {transaction.seller} -{" "}
            {transaction.date.toString()}
          </div>
        ))}
      </div>
    );
  };

  return (
    <div className="transactions">
      <GoBackButton goBack={goBack} />
      <h1>Transactions</h1>
      {renderContent()}
    </div>
  );
};
export default Transactions;
