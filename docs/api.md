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

## Project Endpoints

### Get All Projects
- **URL**: `/api/projects`
- **Method**: `GET`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "status": "success",
      "projects": [
        {
          "id": 1,
          "created_at": "2023-07-15T12:34:56Z",
          "projlink": "https://example.com/project",
          "title": "Sample Project",
          "type": "web",
          "cover": "https://example.com/cover.jpg",
          "logo": "https://example.com/logo.png",
          "profilename": "John Doe",
          "profilepic": "https://example.com/profile.jpg",
          "aboutproject": "This is a sample project description",
          "technologies": "React, Node.js, Go",
          "linkedin_link": "https://linkedin.com/in/johndoe",
          "telegram_link": "https://t.me/johndoe",
          "x_link": "https://x.com/johndoe",
          "youtube_link": "https://youtube.com/johndoe",
          "github_link": "https://github.com/johndoe",
          "insta_link": "https://instagram.com/johndoe"
        }
      ]
    }
    ```

### Get Project by ID
- **URL**: `/api/projects/:id`
- **Method**: `GET`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "status": "success",
      "project": {
        "id": 1,
        "created_at": "2023-07-15T12:34:56Z",
        "projlink": "https://example.com/project",
        "title": "Sample Project",
        "type": "web",
        "cover": "https://example.com/cover.jpg",
        "logo": "https://example.com/logo.png",
        "profilename": "John Doe",
        "profilepic": "https://example.com/profile.jpg",
        "aboutproject": "This is a sample project description",
        "technologies": "React, Node.js, Go",
        "linkedin_link": "https://linkedin.com/in/johndoe",
        "telegram_link": "https://t.me/johndoe",
        "x_link": "https://x.com/johndoe",
        "youtube_link": "https://youtube.com/johndoe",
        "github_link": "https://github.com/johndoe",
        "insta_link": "https://instagram.com/johndoe"
      }
    }
    ```
- **Error Response**:
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "error": "Project not found"
    }
    ```

### Create Project
- **URL**: `/api/projects`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer {token}`
- **Request Body**:
  ```json
  {
    "title": "New Project",
    "projlink": "https://example.com/new-project",
    "type": "mobile",
    "cover": "https://example.com/new-cover.jpg",
    "logo": "https://example.com/new-logo.png",
    "profilename": "John Doe",
    "profilepic": "https://example.com/profile.jpg",
    "aboutproject": "This is a new project description",
    "technologies": "Flutter, Firebase",
    "linkedin_link": "https://linkedin.com/in/johndoe",
    "telegram_link": "https://t.me/johndoe",
    "x_link": "https://x.com/johndoe",
    "youtube_link": "https://youtube.com/johndoe",
    "github_link": "https://github.com/johndoe",
    "insta_link": "https://instagram.com/johndoe"
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "status": "success",
      "message": "Project created successfully",
      "project": {
        "id": 2,
        "created_at": "2023-07-16T09:45:32Z",
        "title": "New Project",
        "projlink": "https://example.com/new-project",
        "type": "mobile"
        // Other project fields...
      }
    }
    ```
- **Error Response**:
  - **Code**: 400 Bad Request
  - **Content**:
    ```json
    {
      "error": "Key: 'ProjectRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag"
    }
    ```

### Update Project
- **URL**: `/api/projects/:id`
- **Method**: `PUT`
- **Headers**: `Authorization: Bearer {token}`
- **Request Body**:
  ```json
  {
    "title": "Updated Project",
    "type": "web",
    "aboutproject": "This is an updated description"
    // Other fields to update
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "status": "success",
      "message": "Project updated successfully",
      "project": {
        "id": 1,
        "title": "Updated Project",
        "type": "web",
        "aboutproject": "This is an updated description"
        // Other project fields...
      }
    }
    ```
- **Error Response**:
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "error": "Project not found"
    }
    ```

### Delete Project
- **URL**: `/api/projects/:id`
- **Method**: `DELETE`
- **Headers**: `Authorization: Bearer {token}`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "status": "success",
      "message": "Project deleted successfully"
    }
    ```
- **Error Response**:
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "error": "Project not found"
    }
    ```