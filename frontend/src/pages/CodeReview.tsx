import React, { useState } from "react";
import "./codeReview.css";
import Markdown from "react-markdown";

const CodeReview: React.FC = () => {
  const [fileName, setFileName] = useState<string | null>(null); // State for uploaded file name
  const [fileContent, setFileContent] = useState<File | null>(null); // State for file content
  const [response, setResponse] = useState<string>(""); // State for the response from the API

  // Handle file upload
  const handleFileUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setFileName(file.name); // Set the file name
      setFileContent(file); // Set the file content (File object)
    }
  };

  // Confirm and make API call
  const confirmFile = async () => {
    if (!fileContent) {
      return;
    }

    try {
      const formData = new FormData();
      formData.append("name", fileName || ""); // Append the file name to FormData
      formData.append("file_upload", fileContent); // Append the file content to FormData

      const response = await fetch(`http://localhost:8000/code-review`, {
        method: "POST",
        body: formData, // Send the formData without "Content-Type" header, browser handles it
      });

      const data = await response.json();
      console.log("API Response:", data);
      const x = JSON.parse(data.data);

      setResponse(x.Candidates[0].Content.Parts[0]); // Set the response from the API
    } catch (error) {
      console.error("Error processing the file:", error);
    }
  };

  return (
    <div className="Concontainer">
      <header className="Conheader">File Upload</header>
      <p className="Condescription">Upload your file and confirm to process it.</p>

      {/* Custom file upload button */}
      <label className="file-upload-label">
        <span>Choose a file</span>
        <input
          type="file"
          onChange={handleFileUpload}
          accept=".go"
          className="file-upload-input"
        />
      </label>

      {fileName && (
        <div>
          <p>Uploaded file: {fileName}</p>
          <button className="Conbutton" onClick={confirmFile}>
            Confirm and Process
          </button>
        </div>
      )}

      {/* {response && (
        <div>
          <p>Code review response:</p>
          <pre className="response">{response}</pre>
        </div>
      )} */}
      {response && (
        <div className="input-container" style={{ marginTop: "10px" }}>
          <h2>Response: </h2>
          <div className="footer response">

            <Markdown>
              {response}
            </Markdown>
            </div>
          <div className="action-buttons"></div>
        </div>
      )}
    </div>
  );
};

export default CodeReview;
