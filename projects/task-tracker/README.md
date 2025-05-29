# Task Tracker

Task Tracker is a simple task management CLI application built using Go. It allows you to create, view, and manage tasks from the command line.

## Features

- Add new tasks with a description
- Delete tasks by ID
- Update task descriptions
- Update task status (todo, in-progress, done)
- List all tasks
- List tasks by status

## Commands

### Add a Task

```bash
tast-tracker add "My new task"
```

### Delete a Task

```bash
tast-tracker delete [id]
```

### Update a Task Description

```bash
tast-tracker update [id] [new description]
```

### Update Task Status

- Mark as Todo:

  ```bash
  tast-tracker mark-todo [id]
  ```

- Mark as In Progress:

  ```bash
  tast-tracker mark-in-progress [id]
  ```

- Mark as Done:

  ```bash
  tast-tracker mark-done [id]
  ```

### List Tasks

- List all tasks:

  ```bash
  tast-tracker list
  ```

- List tasks by status:

  ```bash
  tast-tracker list [status]
  ```

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/selahaddinislamoglu/golang-studies.git
    ```

2. Navigate to the project directory:

    ```bash
    cd golang-studies/projects/task-tracker
    ```

3. Build the application:

    ```bash
    go build
    ```

4. Run the application:

    ```bash
    ./task-tracker
    ```

## Configuration

The application uses a JSON file (`tasks.json`) to store task data. Ensure the file is accessible and writable by the application.

## License

This project is licensed under the MIT License.
