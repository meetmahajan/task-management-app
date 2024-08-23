# Task Management Web Application

This project is a simple task management web application built using Golang for the backend, MySQL as the database, and Angular for the frontend.

## Table of Contents

1. [Project Structure](#project-structure)
2. [Backend (Golang) Overview](#backend-golang-overview)
    - [API Endpoints](#api-endpoints)
    - [Database Configuration](#database-configuration)
3. [Frontend (Angular) Overview](#frontend-angular-overview)
    - [Components](#components)
    - [Services](#services)
4. [Running the Project](#running-the-project)
    - [Running Locally](#running-locally)
    - [Running with Docker](#running-with-docker)

## Project Structure

```
task-management-app/
├── backend/
│   ├── main.go                    # Entry point of the Go application
│   ├── go.mod                     # Module dependencies
│   ├── go.sum                     # Dependency checksums
│   ├── Dockerfile                 # Dockerfile for backend service
│   ├── controllers/
│   │   └── taskController.go      # Handles HTTP requests related to tasks
│   ├── models/
│   │   └── task.go                # Task model definition
│   ├── routes/
│   │   └── taskRoutes.go          # API route definitions
│   ├── config/
│   │   └── database.go            # Database connection setup
│   └── .env                       # Environment variables for the backend
├── frontend/
│   ├── src/
│   │   ├── app/
│   │   │   ├── components/
│   │   │   │   ├── task-list/
│   │   │   │   │   ├── task-list.component.ts     # TaskListComponent logic
│   │   │   │   │   └── task-list.component.html   # TaskListComponent template
|   |   |   |   └── task-detail/
│   │   │   │       ├── task-detail.component.ts   # TaskDetailComponent logic
│   │   │   │       └── task-detail.component.html # TaskDetailComponent template
│   │   │   ├── services/
│   │   │   |   └── task.service.ts                # Handles HTTP calls to backend
|   |   |   └── models/
│   │   │       └── task.model.ts                  # Handles schema of task object
│   │   ├── index.html                             # Main HTML file for Angular
│   │   └── main.ts                                # Angular entry point
│   ├── angular.json                               # Angular project configuration
│   ├── package.json                               # NPM dependencies and scripts
│   ├── tsconfig.json                              # TypeScript configuration
│   └── Dockerfile                                 # Dockerfile for frontend service
└── docker-compose.yml                             # Docker Compose file to orchestrate services
```

## Backend (Golang) Overview

The backend is responsible for handling all API requests, managing the MySQL database, and serving as the core of the task management application.

### API Endpoints

- **GET /api/tasks**: Fetches all tasks from the database.
  - **Function**: `GetTasks` in `taskController.go`
  - **Description**: Retrieves and returns a list of all tasks.

- **GET /api/tasks/{id}**: Fetches a single task by its ID.
  - **Function**: `GetTask` in `taskController.go`
  - **Description**: Retrieves a specific task based on the provided ID.

- **POST /api/tasks**: Creates a new task.
  - **Function**: `CreateTask` in `taskController.go`
  - **Description**: Takes task details from the request body and creates a new task in the database.

- **PUT /api/tasks/{id}**: Updates an existing task by its ID.
  - **Function**: `UpdateTask` in `taskController.go`
  - **Description**: Updates the details of an existing task using the data provided in the request body.

- **DELETE /api/tasks/{id}**: Deletes a task by its ID.
  - **Function**: `DeleteTask` in `taskController.go`
  - **Description**: Deletes the task from the database based on the provided ID.

### Database Configuration

- The application uses MySQL as the database.
- Connection details are configured in the `.env` file located in the `backend/` directory.

#### **.env Example:**
```
DB_USER=root
DB_PASSWORD=password
DB_HOST=localhost
DB_NAME=task_management
```

## Frontend (Angular) Overview

The frontend is built with Angular and provides the user interface for managing tasks.

### Components

- **TaskListComponent (`task-list.component.ts`)**
  - **Description**: Displays a list of tasks fetched from the backend.
  - **Methods**: `ngOnInit` - Fetches the list of tasks on component initialization.

- **TaskDetailComponent (`task-detail.component.ts`)**
  - **Description**: Displays and allows editing of a single task's details.
  - **Methods**: `ngOnInit` - Fetches the task details by ID when the component initializes.

### Services

- **TaskService (`task.service.ts`)**
  - **Description**: Handles HTTP requests to the backend API.
  - **Methods**:
    - `getTasks()`: Fetches all tasks.
    - `getTask(id: number)`: Fetches a specific task by ID.
    - `createTask(task: Task)`: Creates a new task.
    - `updateTask(task: Task)`: Updates an existing task.
    - `deleteTask(id: number)`: Deletes a task by ID.

### Task Model (`task.model.ts`)

The `task.model.ts` file defines the structure of a Task object in the application. It ensures that all tasks have a consistent shape, with properties such as `id`, `title`, `description`, and `completed` status. This model is used across the application to maintain type safety and data integrity when handling tasks.

## Running the Project

### Running Locally

To run the project locally, you need to have Go, Node.js, Angular CLI, and MySQL installed on your machine.

#### **Backend Setup**

1. **Navigate to the backend directory:**
   ```bash
   cd backend
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up MySQL database:**
   - Ensure MySQL is running.
   - Create a database named `task_management`.
   - Update the `.env` file with your MySQL credentials.

4. **Run the backend server:**
   ```bash
   go run main.go
   ```

   The backend server will start on `http://localhost:8080`.

#### **Frontend Setup**

1. **Navigate to the frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

3. **Run the frontend server:**
   ```bash
   ng serve
   ```

   The frontend will be available at `http://localhost:4200`.

### Running with Docker

1. **Navigate to the root of the project:**
   ```bash
   cd task-management-app
   ```

2. **Build and run the application with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

   This command will start the backend, frontend, and MySQL database in Docker containers.

3. **Access the Application:**
   - **Frontend**: `http://localhost:4200`
   - **Backend API**: `http://localhost:8080/api/tasks`