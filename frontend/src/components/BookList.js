import React, { useState } from 'react';
import BookDetails from './BookDetails';

// Mock book data
const mockBooks = [
  {
    id: 1,
    title: 'To Kill a Mockingbird',
    author: 'Harper Lee',
    coverImage: '/api/placeholder/150/220',
    description: 'A classic novel about racial injustice in the American South.',
    copies: 5
  },
  {
    id: 2,
    title: '1984',
    author: 'George Orwell',
    coverImage: '/api/placeholder/150/220',
    description: 'A dystopian novel set in a totalitarian regime.',
    copies: 3
  },
  {
    id: 3,
    title: 'The Great Gatsby',
    author: 'F. Scott Fitzgerald',
    coverImage: '/api/placeholder/150/220',
    description: 'A novel about the American Dream in the Jazz Age.',
    copies: 4
  },
  {
    id: 4,
    title: 'Pride and Prejudice',
    author: 'Jane Austen',
    coverImage: '/api/placeholder/150/220',
    description: 'A romantic novel about manners and marriage.',
    copies: 2
  },
  {
    id: 5,
    title: 'The Catcher in the Rye',
    author: 'J.D. Salinger',
    coverImage: '/api/placeholder/150/220',
    description: 'A novel about teenage alienation and identity.',
    copies: 6
  },
  {
    id: 6,
    title: 'Brave New World',
    author: 'Aldous Huxley',
    coverImage: '/api/placeholder/150/220',
    description: 'A dystopian novel about a technologically advanced future society.',
    copies: 3
  }
];

function BookList() {
  const [books, setBooks] = useState(mockBooks);
  const [selectedBook, setSelectedBook] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [notification, setNotification] = useState('');

  const handleBookSelect = (book) => {
    setSelectedBook(book);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setSelectedBook(null);
  };

  const handleCheckout = (bookId) => {
    setBooks(books.map(book => 
      book.id === bookId && book.copies > 0 
        ? { ...book, copies: book.copies - 1 } 
        : book
    ));
    
    // Add to local storage for borrowed books
    const borrowedBooks = JSON.parse(localStorage.getItem('borrowedBooks') || '[]');
    const bookToBorrow = books.find(book => book.id === bookId);
    
    if (!borrowedBooks.some(book => book.id === bookId)) {
      borrowedBooks.push(bookToBorrow);
      localStorage.setItem('borrowedBooks', JSON.stringify(borrowedBooks));
    }
    
    setNotification('Book checkout successful!');
    setTimeout(() => {
      setNotification('');
      handleCloseModal();
    }, 2000);
  };

  return (
    <div className="books-container">
      <h2>Available Books</h2>
      {notification && <div className="notification">{notification}</div>}
      <div className="book-grid">
        {books.map(book => (
          <div key={book.id} className="book-card" onClick={() => handleBookSelect(book)}>
            <img src={book.coverImage} alt={book.title} className="book-cover" />
            <div className="book-info">
              <h3>{book.title}</h3>
              <p>{book.author}</p>
              <p className="book-copies">Available: {book.copies}</p>
            </div>
          </div>
        ))}
      </div>
      
      {showModal && selectedBook && (
        <BookDetails 
          book={selectedBook} 
          onClose={handleCloseModal} 
          onCheckout={handleCheckout} 
        />
      )}
    </div>
  );
}

export default BookList;