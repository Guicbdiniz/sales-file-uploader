import React, { useState } from "react";
import ErrorText from "../../components/ErrorText";
import Footer from "../../components/Footer";
import GoBackButton from "../../components/GoBackButton";
import Loading from "../../components/Loading";
import { FileParser } from "../../utils/fileParser";
import { FileData, UnsupportedFileFormatError } from "../../utils/types";
import "./style.css";

type Props = {
  goBack: () => void;
};

const UNSUPPORTED_FORMAT_ERROR =
  "File format is not supported. Check your file before trying to reupload.";
const UNKNOWN_PARSING_ERROR =
  "There was an unknown error while parsing your file. Try again.";

const Upload: React.FC<Props> = (props) => {
  const { goBack } = props;
  const [loading, setLoading] = useState(false);
  const [fileData, setFileData] = useState<FileData | null>(null);
  const [errorText, setErrorText] = useState("");

  const changeHandler: React.ChangeEventHandler<HTMLInputElement> = (e) => {
    setLoading(true);
    setErrorText("");
    const file = e.target.files![0];
    const parser = new FileParser(file);
    parser
      .parse()
      .then((parsed) => {
        setFileData(parsed);
      })
      .catch((e) => {
        if (e instanceof UnsupportedFileFormatError) {
          setErrorText(UNSUPPORTED_FORMAT_ERROR);
        } else {
          console.error("Unknown parsing error captured.\n", e);
          setErrorText(UNKNOWN_PARSING_ERROR);
        }
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const renderInput = () => {
    if (loading) {
      return <Loading />;
    }
    if (fileData) {
      return (
        <>
        <div className="file-data">
          <p className="file-name">Selected file: {fileData.name}</p>
          {fileData.transactions.map((transaction, index) => (
            <div className="transaction" key={index}>
              {transaction.type} - {transaction.value} centavos of{" "}
              {transaction.product} - {transaction.seller} - {transaction.date.toDateString()}
            </div>
          ))}
        </div>
        <div className="confirm">Upload</div>
        </>
      );
    }
    return (
      <>
        <label htmlFor="input-file">Choose a File</label>
        <input type="file" id="input-file" onChange={changeHandler} />
      </>
    );
  };

  return (
    <div className="upload">
      <GoBackButton goBack={goBack} />
      {renderInput()}
      {errorText && <ErrorText text={errorText} />}
      <Footer />
    </div>
  );
};

export default Upload;
