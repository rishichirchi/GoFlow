import { useState } from "react";
import "./email.css";

export default function EmailPage() {
  const [input, setInput] = useState("");
  const [currentMode, setCurrentMode] = useState("email");
  const [response, setResponse] = useState("");
  const handleSubmit = async () => {
    const response = await fetch(`http://localhost:8000/cold-email?prompt=${input}`);
    const data = await response.json();

    setResponse(data.data.response);
  };

  return (
    <div className="emailpage" style={{ padding: "1rem" }}>
      <div className="tabs">
        <button className="tab active">Write</button>
      </div>
      <div className="options">
        <button
          className={`option ${currentMode === "email" ? "active" : ""}`}
          onClick={() => setCurrentMode("email")}
        >
          Email
        </button>
        
      </div>
      <div className="input-container">
        <textarea
          className="input-box"
          placeholder="Enter the topic you want to write about"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        ></textarea>
        <button className="submit-btn" onClick={handleSubmit}>
          Submit
        </button>
      </div>
      {response && (
        <div className="input-container" style={{ marginTop: "10px" }}>
          <h2>Response: </h2>
          <div className="footer response">{response}</div>
          <div className="action-buttons"></div>
        </div>
      )}
    </div>
  );
}
