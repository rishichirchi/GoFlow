import React, { useState } from "react";
import { HashRouter, Routes, Route, Link } from "react-router-dom";
import HomePage from "./pages/HomePage";
import ProfilePage from "./pages/ProfilePage";
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
              <img src="/public/icons/settings.svg" alt="Settings"/>
            </Link>
          </div>
        </div>

        {/* Main Content */}
        <div className="main-content">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/email" element={<EmailPage />} />
            <Route path="/content-gen" element={<ContentGen />} />
            <Route path="/profile" element={<ProfilePage />} />
            <Route path="/help" element={<div>Help Page</div>} />
          </Routes>
        </div>
      </div>
    </HashRouter>
  );
}

export default App;
