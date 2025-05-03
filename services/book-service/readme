# **Book Service**
The **Book Service** manages the book catalog, tracking availability and allowing operations such as adding, retrieving, updating, and deleting books. It interacts with the **Borrow Service** to update availability when books are borrowed or returned.

## **Features**
1. **Add new books** to the catalog (Admin functionality).  
2. **Retrieve all books**, optionally filtering by availability.  
3. **Get details of a specific book**, including availability status.  
4. **Update book availability**, typically triggered by Borrow Service.  
5. **Delete a book** from the catalog (Admin functionality).  

---

## **Database Schema**
Book Service stores book records in **MongoDB**, allowing flexibility for future enhancements.

```json
{
  "_id": "unique-object-id",
  "bookId": "book_456",
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "available": true
}
```
1. `_id` → MongoDB's auto-generated object ID.  
2. `bookId` → Unique identifier for books.  
3. `title` → Name of the book.  
4. `author` → Author of the book.  
5. `available` → Tracks if the book is available for borrowing.  

---

## **API Endpoints & Detailed Flow**
### **1. Add a New Book**
**POST** `/books`  
**Stores a new book in MongoDB (Admin functionality)**  

#### **Request**
```json
{
  "bookId": "book_456",
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "available": true
}
```
#### **Flow**
1. Admin sends book details.  
2. Book Service **saves the book to MongoDB**.  

#### **Response**
```json
{
  "message": "Book added successfully"
}
```

---

### **2. Retrieve All Books**
**GET** `/books?availability=true` (Optional query filter)  
**Gets all books, optionally filtering by availability**  

#### **Request**
```http
GET /books
```
Or, filter by availability:
```http
GET /books?availability=true
```
#### **Flow**
1. Fetches book records from MongoDB.  
2. If `availability=true`, it filters only available books.  

#### **Response**
```json
[
  {
    "bookId": "book_123",
    "title": "1984",
    "author": "George Orwell",
    "available": true
  },
  {
    "bookId": "book_456",
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "available": false
  }
]
```

---

### **3. Get Book Details by ID**
**GET** `/books/{bookId}`  
**Fetches book details including availability**  

#### **Request**
```http
GET /books/book_456
```
#### **Flow**
1. Looks up `bookId` in MongoDB.  

#### **Response**
```json
{
  "bookId": "book_456",
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "available": false
}
```

---

### **4. Update Book Availability**
**PUT** `/books/{bookId}`  
**Triggered by Borrow Service to mark books as borrowed or returned**  

#### **Request**
```json
{
  "available": false
}
```
#### **Flow**
1. Borrow Service sends request when book is borrowed.  
2. Book Service **updates availability in MongoDB**.  

#### **Response**
```json
{
  "message": "Book availability updated successfully"
}
```

---

### **5. Delete a Book**
**DELETE** `/books/{bookId}`  
**Removes a book from the catalog (Admin functionality)**  

#### **Request**
```http
DELETE /books/book_456
```
#### **Flow**
1. Removes book record from MongoDB.  

#### **Response**
```json
{
  "message": "Book removed successfully"
}
```

---

## **Book Service Architecture**
**Backend:** Golang  
   - Manages book catalog operations (add, retrieve, update, delete).  
   - Accepts availability updates from Borrow Service when books are borrowed or returned.  

**Database:** MongoDB  
   - Stores book details (`bookId`, `title`, `author`, `available`).  
   - Uses **document-based storage** for flexibility in data structure.  

**Inter-Service Communication:** Kubernetes DNS  
   - Used by Borrow Service to **check book availability** before borrowing.  
   - Uses `borrow-service.default.svc.cluster.local` for updates.  

**Deployment:** Kubernetes (Dockerized)  
   - Runs as a **microservice** inside Kubernetes.  
   - MongoDB database is deployed separately with **Persistent Volumes (PV)** for data retention.  

**Data Persistence:** MongoDB with PV & PVC  
   - Book records **remain intact** across pod restarts using **Persistent Volume Claim (PVC)**.  
   - Prevents data loss and ensures catalog consistency.  

---

## **Deployment in Kubernetes**
### **Book Service Deployment**
1. Creates a **pod** for Book Service using a Docker image.  
2. Runs on **port 5000**, accessible within Kubernetes.  
3. Can be **scaled dynamically** for high traffic loads.  

### **Book Service Kubernetes Service**
1. Exposes Book Service **internally** using `ClusterIP`.  
2. Accessible via Kubernetes DNS at:
   ```
   book-service.default.svc.cluster.local
   ```
3. Used by Borrow Service to **check availability** and **update book status**.  

### **MongoDB Deployment**
1. Deploys MongoDB **inside Kubernetes** for storing books.  
2. Uses environment variables to initialize the database.  
3. Runs on **port 27017** for book data transactions.  

### **MongoDB Kubernetes Service**
1. Exposes MongoDB **internally** (`ClusterIP`) for database communication.  
2. Book Service connects via:
   ```
   book-db.default.svc.cluster.local:27017
   ```
3. Ensures **data persistence for the book catalog**.  

### **Persistent Storage (PV & PVC)**
1. **Persistent Volume (PV)** provides static storage for MongoDB.  
2. **Persistent Volume Claim (PVC)** ensures data remains intact across pod restarts.  
3. Prevents **book catalog data loss when MongoDB pods restart**.  

---

## **Setup & Deployment**
### **1. Build & Push Docker Image**
```sh
docker build -t your-dockerhub-username/book-service .
docker push your-dockerhub-username/book-service:latest
```

### **2. Apply Kubernetes Deployment**
```sh
kubectl apply -f book-service.yaml
```

### **3. Verify Deployment**
```sh
kubectl get pods
kubectl get services
```

---