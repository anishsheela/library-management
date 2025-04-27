import React from 'react';

function BookDetails({ book, onClose, onCheckout }) {
  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button className="close-button" onClick={onClose}>×</button>
        <div className="book-details">
          <div className="book-image">
            <img src={book.coverImage} alt={book.title} />
          </div>
          <div className="book-info-details">
            <h2>{book.title}</h2>
            <h3>by {book.author}</h3>
            <p>{book.description}</p>
            <p className="copies-info">Available copies: {book.copies}</p>
            <button 
              className="btn-checkout" 
              disabled={book.copies === 0}
              onClick={() => onCheckout(book.id)}
            >
              {book.copies > 0 ? 'Checkout Book' : 'Out of Stock'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default BookDetails;