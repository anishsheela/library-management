import React from 'react';
import { Link } from 'react-router-dom';

function Navbar({ user, onLogout }) {
  return (
    <nav className="navbar">
      <div className="navbar-brand">Library Management System</div>
      <div className="navbar-links">
        <Link to="/books">Books</Link>
        <Link to="/borrowed">My Borrowed Books</Link>
      </div>
      <div className="navbar-user">
        <span>{user?.email}</span>
        <button onClick={onLogout} className="btn-logout">Logout</button>
      </div>
    </nav>
  );
}

export default Navbar;