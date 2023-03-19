import { FileParser } from ".";
import { Transaction, UnsupportedFileFormatError } from "../types";

describe("FileParser", () => {
  let parser: FileParser;

  const createFile = (
    text: string,
    name = "text.txt",
    format = "text/plain"
  ) => {
    return new File([text], name, { type: format });
  };

  it("should throw an Exception when a file with wrong explicity format is passed", async () => {
    parser = new FileParser(createFile("", "test.json", "application/json"));

    await expect(parser.parse()).rejects.toThrowError(
      UnsupportedFileFormatError
    );
  });

  it("should throw Exception when empty file is passed", async () => {
    parser = new FileParser(createFile(""));

    await expect(parser.parse()).rejects.toThrowError(
      UnsupportedFileFormatError
    );
  });

  it("should throw Exception when file with invalid text format is passed", async () => {
    parser = new FileParser(createFile("aeeeaseasmneaks"));

    await expect(parser.parse()).rejects.toThrowError(
      UnsupportedFileFormatError
    );

    parser = new FileParser(
      createFile(
        "1kl2022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS"
      )
    );
    await expect(parser.parse()).rejects.toThrowError(
      UnsupportedFileFormatError
    );
  });

  it("should return valid transactions when valid file is passed", async () => {
    parser = new FileParser(
      createFile(
        "12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS"
      )
    );
    let data = await parser.parse();

    expect(data.name).toBe("text.txt");
    expect(data.transactions[0]).toEqual<Transaction>({
      date: new Date("2022-01-15T19:20:30-03:00"),
      product: "CURSO DE BEM-ESTAR",
      seller: "JOSE CARLOS",
      type: "ProducerSale",
      value: 12750,
    });
  });
});
