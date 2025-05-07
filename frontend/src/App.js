import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';
import BookList from './components/BookList';
import BorrowedBooks from './components/BorrowedBooks';
import Navbar from './components/Navbar';
import './App.css';

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState(null);
  
  const handleLogin = (email) => {
    setIsAuthenticated(true);
    setUser({ email });
  };
  
  const handleLogout = () => {
    setIsAuthenticated(false);
    setUser(null);
  };

  return (
    <Router>
      <div className="app">
        {isAuthenticated && <Navbar user={user} onLogout={handleLogout} />}
        <div className="container">
          <Routes>
            <Route 
              path="/" 
              element={isAuthenticated ? <Navigate to="/books" /> : <Login onLogin={handleLogin} />} 
            />
            <Route path="/register" element={<Register />} />
            <Route 
              path="/books" 
              element={isAuthenticated ? <BookList /> : <Navigate to="/" />} 
            />
            <Route 
              path="/borrowed" 
              element={isAuthenticated ? <BorrowedBooks /> : <Navigate to="/" />} 
            />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;