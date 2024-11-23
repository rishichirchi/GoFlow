import { useState } from "react";
import "./email.css";

export default function EmailPage() {
  const [input, setInput] = useState("");
  const [currentMode, setCurrentMode] = useState("email");
  const [response, setResponse] = useState("");
  const handleSubmit = () => {
    setResponse(input);
  };

  return (
    <div className="emailpage" style={{ padding: "1rem" }}>
      <div className="tabs">
        <button className="tab active">Write</button>
        <button className="tab">Reply</button>
      </div>
      <div className="options">
        <button
          className={`option ${currentMode === "email" ? "active" : ""}`}
          onClick={() => setCurrentMode("email")}
        >
          Email
        </button>
        <button
          className={`option ${currentMode === "letter" ? "active" : ""}`}
          onClick={() => setCurrentMode("letter")}
        >
          Essay
        </button>
        <button
          className={`option ${currentMode === "paragraph" ? "active" : ""}`}
          onClick={() => setCurrentMode("paragraph")}
        >
          Paragraph
        </button>
        <button
          className={`option ${currentMode === "idea" ? "active" : ""}`}
          onClick={() => setCurrentMode("idea")}
        >
          Idea
        </button>
        <button
          className={`option ${currentMode === "blog" ? "active" : ""}`}
          onClick={() => setCurrentMode("blog")}
        >
          Blog Post
        </button>
      </div>
      <div className="settings">
        Type:
        <span className="setting">Formal</span> -
        <span className="setting">Short</span> -
        <span className="setting">English</span>
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
