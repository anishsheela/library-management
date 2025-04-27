import React, { useState, useEffect } from 'react';

function BorrowedBooks() {
  const [borrowedBooks, setBorrowedBooks] = useState([]);
  const [notification, setNotification] = useState('');

  useEffect(() => {
    // Get borrowed books from local storage
    const storedBooks = JSON.parse(localStorage.getItem('borrowedBooks') || '[]');
    setBorrowedBooks(storedBooks);
  }, []);

  const handleReturn = (bookId) => {
    // Update borrowed books list
    const updatedBorrowedBooks = borrowedBooks.filter(book => book.id !== bookId);
    setBorrowedBooks(updatedBorrowedBooks);
    localStorage.setItem('borrowedBooks', JSON.stringify(updatedBorrowedBooks));
    
    // Update available copies count in the book list (would be an API call in a real app)
    const bookList = JSON.parse(localStorage.getItem('books') || '[]');
    const updatedBookList = bookList.map(book => 
      book.id === bookId ? { ...book, copies: book.copies + 1 } : book
    );
    localStorage.setItem('books', JSON.stringify(updatedBookList));
    
    setNotification('Book returned successfully!');
    setTimeout(() => {
      setNotification('');
    }, 2000);
  };

  return (
    <div className="borrowed-books">
      <h2>My Borrowed Books</h2>
      {notification && <div className="notification">{notification}</div>}
      
      {borrowedBooks.length === 0 ? (
        <p className="no-books-message">You haven't borrowed any books yet.</p>
      ) : (
        <div className="borrowed-books-list">
          {borrowedBooks.map(book => (
            <div key={book.id} className="borrowed-book-item">
              <img src={book.coverImage} alt={book.title} className="book-cover-small" />
              <div className="borrowed-book-info">
                <h3>{book.title}</h3>
                <p>{book.author}</p>
              </div>
              <button 
                className="btn-return" 
                onClick={() => handleReturn(book.id)}
              >
                Return Book
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default BorrowedBooks;