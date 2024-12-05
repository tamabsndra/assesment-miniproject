# Simple POST Management

## Project Overview

Assessment-Miniproject is a full-stack web application built with a Next.js frontend and a Go backend. This project allows users to register, log in, and manage posts. With essential CRUD functionality, it provides a practical example of a full-stack application, connecting frontend and backend technologies seamlessly.

## Features

- **User Authentication**: Register and log in securely
- **Post Management**: Create, read, update, and delete posts through an intuitive interface

## Prerequisites

- Node.js (v16+)
- Go (v1.16+)
- Redis
- PostgreSQL (v12+)
- SQLC
- Swag

## Project Structure

```
├── miniproject-frontend/          # React frontend
├── miniproject-backend/           # Go backend
└── README.md
```

## Database Setup

1. Install PostgreSQL:
```bash
# Ubuntu
sudo apt install postgresql postgresql-contrib

# macOS
brew install postgresql
```

2. Create database:
```bash
psql postgres postgres
CREATE DATABASE miniproject;
```

3. Apply schema:
```bash
psql -d miniproject-backend -f schema.sql
```

4. Generate SQLC code:
```bash
sqlc generate
```
## Backend Setup (Go)

### Project Information
- **Module**: github.com/tamabsndra/miniproject/miniproject-backend
- **Go Version**: 1.23.0

### Dependencies
- Gin: HTTP web framework
- Redis: For caching and session management
- Swagger: API documentation
- Validator: Input validation
- Godotenv: Environment variable management

### Setup Instructions

1. Clone the backend repository:
```bash
git clone https://github.com/tamabsndra/assesment-miniproject.git
cd miniproject-backend
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
   - Create a `.env` file in the root directory and configure keys such as Redis connection strings, JWT secrets, database URLs, etc.

4. Run the application:
```bash
go run main.go
```

The backend server should now be running. By default, it will be accessible at `http://localhost:8080`.

### Redis Setup

To use Redis for caching, make sure Redis is installed and running on your system.

1. Start Redis:
   - Run `redis-server` to start a Redis instance locally
2. Configuration:
   - Ensure that your `.env` file contains the correct Redis connection string

### Swagger Setup

Swagger is used for API documentation. You can view the Swagger docs by following these steps:

1. Install Swag CLI (if not already installed):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate Swagger documentation:
```bash
swag init
```

3. Access Swagger UI:
   - After running the backend server, visit `http://localhost:8080/swagger/index.html` to view the API documentation

## Frontend Setup (Next.js)

### Project Information
- **Name**: miniproject-frontend
- **Version**: 0.1.0

### Scripts
- **Development Server**: `npm run dev` - Runs the app in development mode
- **Production Build**: `npm run build` - Builds the app for production
- **Production Server**: `npm run start` - Starts the production server
- **Lint**: `npm run lint` - Lints the code for syntax and style issues

### Dependencies
The project uses various libraries for functionality and styling, including:
- React and Next.js for frontend structure
- Radix UI and Lucide React for UI components
- React Query for data fetching and caching
- Next Auth for authentication
- Tailwind CSS for styling

### Setup Instructions

1. Clone the frontend repository:
```bash
git clone https://github.com/tamabsndra/assesment-miniproject.git
cd miniproject-frontend
```

2. Install dependencies:
```bash
npm install
```

3. Set up environment variables:
   - Create a `.env.local` file in the root directory with necessary keys (API endpoints, authentication secrets, etc.)

4. Run the development server:
```bash
npm run dev
```

Visit `http://localhost:3000` in your browser to access the frontend application.

### Build Instructions

To create a production build:
```bash
npm run build
```

Start the production server with:
```bash
npm run start
```

## Usage

1. Start Backend: Make sure the backend is running on `http://localhost:8080`
2. Start Frontend: Start the frontend development server on `http://localhost:3000`
3. You can now register and log in to manage posts from the frontend interface

## Troubleshooting

If you encounter any issues, make sure all dependencies are installed and environment variables are correctly set.

## License

This project is licensed under the MIT License.
