import React, { useState } from "react";
import { HashRouter, Routes, Route, Link } from "react-router-dom";
import HomePage from "./pages/HomePage";
import FileUploader from "./pages/CodeReview";
import EmailPage from "./pages/EmailPage";
import "./App.css";
import ContentGen from "./pages/ContentGen";

function App() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true); // Sidebar visible by default

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };

  return (
    <HashRouter>
      <div className="mainCon">
        {/* Sidebar */}
        <div className={`sidebar ${isSidebarOpen ? "" : "sidebar-hidden"}`}>
          <div className="navbar">
            {/* Nav Items */}
            <Link to="/" className="navbar-item">
              <img src="/public/icons/home.svg" alt="Home"/>
            </Link>
            <Link to="/email" className="navbar-item">
              <img src="/public/icons/email.svg" alt="Email"/>
            </Link>
            <Link to="/content-gen" className="navbar-item" >
              <img src="/public/icons/content.svg" alt="content"/>
            </Link>
            <Link to="/review" className="navbar-item" >
              <img src="/public/icons/check.svg" alt="check"/>
            </Link>
          </div>
        </div>

        {/* Main Content */}
        <div className="main-content">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/email" element={<EmailPage />} />
            <Route path="/content-gen" element={<ContentGen />} />
            <Route path="/review" element={<FileUploader />} />
            <Route path="/help" element={<div>Help Page</div>} />
          </Routes>
        </div>
      </div>
    </HashRouter>
  );
}

export default App;
