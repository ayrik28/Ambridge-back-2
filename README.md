# Ambridge Backend

Backend system for the Ambridge platform, a service that helps startups with development solutions.

## Features

- JWT-based authentication
- User registration with email verification
- Login/logout functionality
- Password reset capabilities
- Role-based access control (admin, regular users)
- User profile management

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login a user
- `POST /api/auth/logout` - Logout a user (requires authentication)
- `POST /api/auth/refresh-token` - Refresh JWT token
- `POST /api/auth/forgot-password` - Request password reset
- `POST /api/auth/reset-password` - Reset password with OTP
- `POST /api/auth/verify-email` - Verify email with OTP

For detailed API documentation with request/response examples, see [API Documentation](docs/api.md).

## Setup

1. Clone the repository
2. Run the initialization script:

```bash
# Make the script executable
chmod +x init.sh

# Run the initialization script
./init.sh
```

3. Update the `.env` file with your configuration
4. Run the server:

```bash
# Using the Makefile
make run

# Or directly with Go
go run cmd/server/main.go

# Or build and run the binary
make build
./bin/server
```

## Database

The project uses MySQL with GORM ORM for database operations. The database schema is automatically migrated on server startup.

## Environment Variables

- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET` - Secret key for JWT tokens
- `PORT` - Server port
- `EMAIL_FROM` - Email address for sending emails
- `EMAIL_PASSWORD` - Password for email account
- `SMTP_HOST` - SMTP host for sending emails
- `SMTP_PORT` - SMTP port for sending emails
