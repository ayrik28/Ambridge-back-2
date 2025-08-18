# API Request and Response Models

This document outlines all the request and response models used in the Ambridge API.

## Authentication Routes

All authentication routes are under the `/api/auth` prefix.

### 1. Register

**Endpoint:** `POST /api/auth/register`

**Request Model:**
```json
{
  "name": "String (required)",
  "surname": "String (required)",
  "email": "String (required, valid email)",
  "password": "String (required, min 6 characters)",
  "profileImage": "String (optional)",
  "referral": "String (optional)",
  "company": "String (optional)",
  "currentPosition": "String (optional)"
}
```

**Response Model (Success - 201):**
```json
{
  "message": "User registered successfully",
  "user_id": 1
}
```

**Response Model (Error - 400, 409, 500):**
```json
{
  "error": "Error message"
}
```

### 2. Login

**Endpoint:** `POST /api/auth/login`

**Request Model:**
```json
{
  "email": "String (required, valid email)",
  "password": "String (required)"
}
```

**Response Model (Success - 200):**
```json
{
  "token": "JWT token string",
  "refresh_token": "Refresh token string",
  "user": {
    "id": 1,
    "name": "User's name",
    "surname": "User's surname",
    "email": "user@example.com",
    "role": "user",
    "profileImage": "/path/to/image.jpg",
    "referral": "Source of referral",
    "company": "Company name",
    "currentPosition": "Current job position"
  }
}
```

**Response Model (Error - 400, 401, 500):**
```json
{
  "error": "Error message"
}
```

### 3. Refresh Token

**Endpoint:** `POST /api/auth/refresh-token`

**Request Model:**
```json
{
  "refresh_token": "String (required)"
}
```

**Response Model (Success - 200):**
```json
{
  "token": "New JWT token string",
  "refresh_token": "New refresh token string"
}
```

**Response Model (Error - 400, 401, 500):**
```json
{
  "error": "Error message"
}
```

### 4. Logout

**Endpoint:** `POST /api/auth/logout`

**Headers:**
- Authorization: Bearer {token}

**Request Model:**
No request body needed.

**Response Model (Success - 200):**
```json
{
  "message": "Logged out successfully"
}
```

**Response Model (Error - 401, 500):**
```json
{
  "error": "Error message"
}
```

### 5. Get Profile

**Endpoint:** `GET /api/auth/profile`

**Headers:**
- Authorization: Bearer {token}

**Request Model:**
No request body needed.

**Response Model (Success - 200):**
```json
{
  "user": {
    "name": "User's name",
    "surname": "User's surname",
    "email": "user@example.com",
    "role": "user",
    "profileImage": "/path/to/image.jpg",
    "referral": "Source of referral",
    "company": "Company name",
    "companyEmail": "company@example.com",
    "companyAddress": "Company address",
    "companyPhone": "Company phone number",
    "currentPosition": "Current job position",
    "resumeFile": "/path/to/resume.pdf"
  }
}
```

**Response Model (Error - 401, 404, 500):**
```json
{
  "error": "Error message"
}
```

### 6. Update Profile

**Endpoint:** `PATCH /api/auth/profile`

**Headers:**
- Authorization: Bearer {token}
- Content-Type: application/json

**Request Model:**
```json
{
  "name": "New name",                     // Optional
  "surname": "New surname",               // Optional
  "profileImage": "/path/to/image.jpg",   // Optional
  "referral": "New referral source",      // Optional
  "company": "New company name",          // Optional
  "companyEmail": "new@company.com",      // Optional
  "companyAddress": "New address",        // Optional
  "companyPhone": "New phone number",     // Optional
  "currentPosition": "New position",      // Optional
  "resumeFile": "/path/to/resume.pdf"     // Optional
}
```
Note: You can include only the fields you want to update. Email and password cannot be updated through this endpoint.

**Response Model (Success - 200):**
```json
{
  "message": "Profile updated successfully",
  "user": {
    "name": "Updated name",
    "surname": "Updated surname",
    "email": "user@example.com",
    "role": "user",
    "profileImage": "/path/to/image.jpg",
    "referral": "Updated referral source",
    "company": "Updated company name",
    "companyEmail": "updated@company.com",
    "companyAddress": "Updated address",
    "companyPhone": "Updated phone number",
    "currentPosition": "Updated position",
    "resumeFile": "/path/to/resume.pdf"
  }
}
```

**Response Model (Error - 400, 401, 404, 500):**
```json
{
  "error": "Error message"
}
```

### 7. Check Admin Status

**Endpoint:** `POST /api/auth/check-admin`

**Headers:**
- Authorization: Bearer {token}
- Content-Type: application/json

**Request Model:**
```json
{
  "username": "user@example.com"
}
```

**Response Model (Success - 200):**
```json
{
  "isAdmin": true,
  "user": {
    "email": "user@example.com",
    "role": "admin"
  }
}
```
Or if the user is not an admin:
```json
{
  "isAdmin": false,
  "user": {
    "email": "user@example.com",
    "role": "user"
  }
}
```

**Response Model (Error - 400, 401, 404, 500):**
```json
{
  "error": "Error message"
}
```

## User Model

The User model in the database contains the following fields:

```go
type User struct {
    gorm.Model
    Name           string
    Surname        string
    Email          string
    Password       string // Not exposed in JSON responses
    ProfileImage   string
    CompanyName    string
    CompanyEmail   string
    CompanyAddress string
    CompanyPhone   string
    Position       string
    ReferralSource string
    Role           string // 'admin' or 'user'
    RefreshToken   string // Not exposed in JSON responses
    ResumeFile     string
}
```

## Error Codes

- **400** - Bad Request (invalid input)
- **401** - Unauthorized (invalid credentials or token)
- **403** - Forbidden (insufficient permissions)
- **404** - Not Found
- **409** - Conflict (e.g., email already exists)
- **500** - Internal Server Error
