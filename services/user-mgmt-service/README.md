
# User Management Module

This module provides user registration, login, and user listing functionalities using Django REST Framework and JWT authentication.

## 🔧 Features

- User registration (`/register/`)
- User login and JWT token issuance (`/login/` or `/api/token/`)
- Token refresh (`/api/token/refresh/`)
- List all users (`/users/`)
- Retrieve individual user details (`/users/<id>/`)
- JWKS endpoint for JWT verification (`/.well-known/jwks.json`)

## 📦 API Endpoints

| Method | Endpoint                      | Description                     |
|--------|-------------------------------|---------------------------------|
| POST   | `/register/`                  | Register a new user             |
| POST   | `/login/`                     | Login with username & password  |
| POST   | `/api/token/`                 | Obtain access & refresh tokens  |
| POST   | `/api/token/refresh/`         | Refresh the access token        |
| GET    | `/users/`                     | Get list of all users           |
| GET    | `/users/<id>/`                | Get user by ID                  |
| GET    | `/.well-known/jwks.json`      | Get public key for JWT verify   |

## 🧪 Testing with Postman

### Register
```
POST /register/
{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
}
```

### Login
```
POST /login/
{
    "username": "testuser",
    "password": "password123"
}
```

### Get Users (Requires Authorization)
Add `Authorization: Bearer <access_token>` header.

```
GET /users/
```

## 🔐 JWT Auth Setup

Ensure you have the following in your `settings.py`:

```python
SIMPLE_JWT = {
    'SIGNING_KEY': PRIVATE_KEY,
    'VERIFYING_KEY': PUBLIC_KEY,
    'ACCESS_TOKEN_LIFETIME': timedelta(minutes=15),
    'REFRESH_TOKEN_LIFETIME': timedelta(days=7),
    ...
}
```

## 👩‍💻 How to Contribute

1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Open a Pull Request
