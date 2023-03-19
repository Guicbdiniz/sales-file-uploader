import axios from "axios";
import { Balance, Transaction } from "../utils/types";

const API_URL = process.env.REACT_APP_API_URL || "localhost:3001";

/**
 * postTransactions sends a POST request to the API to
 * save new transactions.
 */
const postTransactions = async (transactions: Transaction[]): Promise<void> => {
  const response = await axios.post(API_URL + "/transactions", transactions, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
  });
  if (response.status !== 201) {
    throw new Error(response.data.errorText);
  }
};

/**
 * fetchTransactions sends a GET request to the API to
 * get all transactions.
 */
const fetchTransactions = async (): Promise<Transaction[]> => {
  const response = await axios(API_URL + "/transactions");
  if (response.status !== 200) {
    throw new Error(response.data.jsonResponse.errorText);
  }
  const transactions: Transaction[] | null = response.data.data;
  return transactions || [];
};

/**
 * fetchBalances sends a GET request to the API to
 * get all balances.
 */
const fetchBalances = async (): Promise<Balance[]> => {
  const response = await axios(API_URL + "/balances");
  if (response.status !== 200) {
    throw new Error(response.data.jsonResponse.errorText);
  }
  const balances: Balance[] | null = response.data.data;
  return balances || [];
};

const ping = async (): Promise<string> => {
  const response = await axios(API_URL + "/ping", {
  });
  if (response.status !== 200) {
    throw new Error("could not ping api");
  }
  return await response.data;
};

/**
 * Set of methods to access the API.
 */
const api = {
  ping,
  postTransactions,
  fetchTransactions,
  fetchBalances,
};

export default api;
