export type FileData = {
  name: string;
  transactions: Transaction[];
};

export type TransactionType =
  | "ProducerSale"
  | "AffiliatedSale"
  | "PaidCommission"
  | "ReceivedCommission";

export type Transaction = {
  type: TransactionType;
  date: Date;
  product: string;
  value: number;
  seller: string;
};

export type Balance = {
  isProducer: boolean;
  balance: number;
  name: string;
}

export class UnsupportedFileFormatError extends Error {}

export class UnknownFileError extends Error {}