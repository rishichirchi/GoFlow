import { useState } from "react";
import "./email.css";

export default function EmailPage() {
  const [input, setInput] = useState("");
  const handleSubmit = () => {
    alert("Submitted" + input);
  };

  return (
    <div className="emailpage">
      <div className="tabs">
        <button className="tab active">Write</button>
        <button className="tab">Reply</button>
      </div>
      <div className="options">
        <button className="option active">Email</button>
        <button className="option">Essay</button>
        <button className="option">Paragraph</button>
        <button className="option">Idea</button>
        <button className="option">Blog Post</button>
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
    </div>
  );
}
