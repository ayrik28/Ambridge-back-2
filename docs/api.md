# Ambridge API Documentation

## Authentication Endpoints

### Register User
- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "name": "John",
    "surname": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword"
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "message": "User registered successfully",
      "user_id": 1
    }
    ```
- **Error Response**:
  - **Code**: 400 Bad Request
  - **Content**:
    ```json
    {
      "error": "Email is required"
    }
    ```
  OR
  - **Code**: 409 Conflict
  - **Content**:
    ```json
    {
      "error": "Email already registered"
    }
    ```

### Login
- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "john.doe@example.com",
    "password": "securepassword"
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "6fd8d272...",
      "user": {
        "id": 1,
        "name": "John",
        "surname": "Doe",
        "email": "john.doe@example.com",
        "role": "user"
      }
    }
    ```
- **Error Response**:
  - **Code**: 401 Unauthorized
  - **Content**:
    ```json
    {
      "error": "Invalid email or password"
    }
    ```

### Logout
- **URL**: `/api/auth/logout`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer {token}`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "message": "Logged out successfully"
    }
    ```
- **Error Response**:
  - **Code**: 401 Unauthorized
  - **Content**:
    ```json
    {
      "error": "Unauthorized"
    }
    ```

### Refresh Token
- **URL**: `/api/auth/refresh-token`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "refresh_token": "6fd8d272..."
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "7fe9d383..."
    }
    ```
- **Error Response**:
  - **Code**: 401 Unauthorized
  - **Content**:
    ```json
    {
      "error": "Invalid refresh token"
    }
    ```