import React, { useState } from 'react';
import GunList from './components/GunList';
import '../src/App.css'

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';
import ProfilePage from './components/Profile';
import Navbar from './components/Navbar';
import CreateGun from './components/CreateGun';

function App() {
  const [token, setToken] = useState(localStorage.getItem('token'));

  return (
    <>
      <Router>
        <Navbar />
        <Routes>
          <Route exact path="/" element={<Login />} />
          <Route exact path="/register" element={<Register />} />
          <Route exact path="/store" element={<GunList />} />
          {token!=null && (
            <>
              <Route path="/profile" element={<ProfilePage />} />
              <Route path="/creategun" element={<CreateGun />} />
            </>
          )}
        </Routes>
      </Router>
    </>
  );
}

export default App;
