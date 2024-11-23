import React, { useState } from "react";
import { HashRouter, Routes, Route, Link } from "react-router-dom";
import HomePage from "./pages/HomePage";
import SettingsPage from "./pages/SettingsPage";
import ProfilePage from "./pages/ProfilePage";
import "./App.css";

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
              <img src="/icons/home.svg" alt="Home" />
            </Link>
            <Link to="/settings" className="navbar-item">
              <img src="/icons/settings.svg" alt="Settings" />
            </Link>
            <Link to="/profile" className="navbar-item">
              <img src="/icons/profile.svg" alt="Profile" />
            </Link>
            <Link to="/help" className="navbar-item">
              <img src="/icons/help.svg" alt="Help" />
            </Link>
          </div>
        </div>

        {/* Main Content */}
        <div className="main-content">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/settings" element={<SettingsPage />} />
            <Route path="/profile" element={<ProfilePage />} />
            <Route path="/help" element={<div>Help Page</div>} />
          </Routes>
        </div>
      </div>
    </HashRouter>
  );
}

export default App;
