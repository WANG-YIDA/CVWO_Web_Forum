# CVWO Web Forum

A full-stack web forum application built with Go (backend) and React (frontend).

## Table of Contents

- Features
- Setup Instructions
- Running the Application
- Project Structure
- AI Usage Documentation

## Features

- User registration and login
- topic, post, comment CRUD

## Setup Instructions

### Prerequisites

- Go (v1.18 or above)
- Node.js (v16 or above) and npm/yarn

### 1. Clone the repository

```sh
git clone https://github.com/WANG-YIDA/CVWO_Web_Forum.git
```

### 2. Backend Setup (Go)

1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Start the backend server (default port: 8000):
   ```sh
   go run cmd/server/main.go
   ```
   Or, for live reload (if you have [air](https://github.com/cosmtrek/air) installed):
   ```sh
   air
   ```

### 3. Frontend Setup (React)

1. Go to the client directory:
   ```sh
   cd client
   ```
2. Install dependencies:
   ```sh
   yarn install
   # or
   npm install
   ```
3. Start the frontend development server (default port: 3000):
   ```sh
   yarn start
   # or
   npm start
   ```

### 4. Access the Application

- Open your browser and go to [http://localhost:3000](http://localhost:3000)

### 5. Notes

- Ensure the backend is running before using the frontend.
- If you encounter CORS issues, make sure CORS is enabled in the Go backend.
- Default backend API URL is `http://localhost:8000`.

### 6. Environment Variables

#### Frontend (`client/.env`)

Create a file named `.env` in the `client` directory with:

```
REACT_APP_API_DOMAIN=http://localhost
```

#### Backend (`.env` or shell export)

Set the environment variable before running the backend:

```
REACT_APP_API_DOMAIN=http://localhost
FRONTEND_ORIGIN_DOMAIN=http://localhost
PORT=8000 # Backend
```

## Project Structure

```
CVWO_Web_Forum/
├── client/           # React frontend
├── internal/         # Go backend (API, models, handlers, etc.)
├── cmd/server/       # Go server entry point
├── go.mod            # Go module file
└── README.md         # This file
```

## AI Usage Documentation

This project utilized AI tools (such as GitHub Copilot and Gemini) in the following ways:

- **API and Library Usage:**
  - Guidance and code examples for using some APIs and libraries e.g. Fetch API in the frontend (React) for HTTP requests.
  - Reference and troubleshooting for using the `sqlite3` library in the Go backend for database operations.
- **Frontend CSS Styles Adjustments:**
  - Suggestions for improving and customizing CSS styles, including responsive design and UI polish.
- **Bug Investigation:**
  - Assistance in diagnosing and resolving tricky bugs, such as CORS issues.
- **README Generation:**
  - Drafting and polishing this README document to ensure clarity and completeness.

All AI-generated suggestions were reviewed, tested, and integrated by the developer. No AI-generated content is used without human oversight.
