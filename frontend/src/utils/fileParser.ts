import {
  FileData,
  Transaction,
  TransactionType,
  UnknownFileError,
  UnsupportedFileFormatError,
} from "./types";

const TEXT_FILE_FORMAT = "text/plain";

export class FileParser {
  constructor(public file: File) {}

  async parse(): Promise<FileData> {
    this.checkFileFormat();

    const fileText = await this.getTextContent();

    if (fileText.length === 0) {
      throw new UnsupportedFileFormatError();
    }

    const transactions: Transaction[] = [];

    for (const line of fileText.split("\n")) {
      if (line.length === 0) {
        continue;
      }
      const transaction = this.getTransactionFromLine(line);
      transactions.push(transaction);
    }

    return {
      name: this.file.name,
      transactions: transactions,
    };
  }

  private checkFileFormat() {
    if (this.file.type !== TEXT_FILE_FORMAT) {
      console.error("File is not text.");
      throw new UnsupportedFileFormatError();
    }
  }

  private async getTextContent() {
    return new Promise<string>((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => {
        if (reader.result && typeof reader.result === "string") {
          resolve(reader.result);
        } else {
          console.error("File is empty.");
          reject(new UnsupportedFileFormatError());
        }
      };
      reader.onerror = () => {
        reject(new UnknownFileError());
      };
      reader.readAsText(this.file);
    });
  }

  private getTransactionFromLine(line: string): Transaction {
    const type = this.getTransactionTypeFromChar(line.charAt(0));
    const date = this.getDateFromLine(line);
    const product = this.getProductFromLine(line);
    const value = this.getValueFromLine(line);
    const seller = this.getSellerFromLine(line);

    return {
      type: type,
      date: date,
      product: product,
      value: value,
      seller: seller,
    };
  }

  private getTransactionTypeFromChar(char: string): TransactionType {
    switch (char) {
      case "1":
        return "ProducerSale";
      case "2":
        return "AffiliatedSale";
      case "3":
        return "PaidCommission";
      case "4":
        return "ReceivedCommission";
      default:
        console.error("Transaction type does not have correct format.");
        throw new UnsupportedFileFormatError();
    }
  }

  private getDateFromLine(line: string): Date {
    if (line.length < 28) {
      console.error("File date could not be read.");
      throw new UnsupportedFileFormatError();
    }
    const dateString = line.substring(1, 26);
    const parsedDate = Date.parse(dateString);
    if (Number.isNaN(parsedDate)) {
      console.error("File date does not have correct format.");
      throw new UnsupportedFileFormatError();
    }
    return new Date(parsedDate);
  }

  private getProductFromLine(line: string): string {
    if (line.length < 56) {
      console.error("File product could not be read.");
      throw new UnsupportedFileFormatError();
    }
    return line.substring(26, 56).trimEnd();
  }

  private getValueFromLine(line: string): number {
    if (line.length < 66) {
      console.error("File value could not be read.");
      throw new UnsupportedFileFormatError();
    }
    const value = parseInt(line.substring(56, 66));
    if (Number.isNaN(value)) {
      console.error("File value is not a number.");
      throw new UnsupportedFileFormatError();
    }
    return value;
  }

  private getSellerFromLine(line: string): string {
    if (line.length < 68) {
      console.error("File seller could not be read.");
      throw new UnsupportedFileFormatError();
    }
    return line.substring(66);
  }
}
