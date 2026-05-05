# t-cli

<p align="center">

![License](https://img.shields.io/github/license/Tikkaaa3/t-cli?style=for-the-badge)
![Last Commit](https://img.shields.io/github/last-commit/Tikkaaa3/t-cli?style=for-the-badge)
![Repo Size](https://img.shields.io/github/repo-size/Tikkaaa3/t-cli?style=for-the-badge)

![Go](https://img.shields.io/badge/Go-1.25.5+-00ADD8?style=for-the-badge&logo=go)
![Cobra](https://img.shields.io/badge/Cobra-CLI-3B7EBB?style=for-the-badge)
![Charmbracelet](https://img.shields.io/badge/Charmbracelet-TUI-FF69B4?style=for-the-badge)
![Linux](https://img.shields.io/badge/Linux-Supported-FCC624?style=for-the-badge&logo=linux)

![Cross Platform](https://img.shields.io/badge/Cross--Platform-Linux%20%7C%20Windows%20%7C%20macOS-4CAF50?style=for-the-badge)
![CLI Tool](https://img.shields.io/badge/Type-CLI%20Tool-black?style=for-the-badge)

</p>

**t-cli** is an educational command-line tool designed to execute, verify, and grade local code submissions against a remote learning platform. It acts as a bridge between a student's local development environment and a centralized curriculum backend.

## 🎬 Demo Video

👉 Click the image below to watch the demo on YouTube:

[![▶ Watch Demo](https://img.youtube.com/vi/1L5MmrqYokI/maxresdefault.jpg)](https://www.youtube.com/watch?v=1L5MmrqYokI)

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Contract](#api-contract)
- [Development](#development)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [Motivation & Inspiration](#motivation--inspiration)
- [License](#license)

## Overview

The CLI retrieves tasks from an API, executes the required shell commands or scripts on the user's machine, compares the output against expected results, and submits the final grade back to the server. It is built with Go and utilizes the [Cobra](https://github.com/spf13/cobra) library for command management and [Charm](https://github.com/charmbracelet) libraries for the user interface.

### How It Works

1. **Authenticate**: Store your API key locally using the `login` command
2. **Fetch Tasks**: Request a lesson by ID from the remote platform
3. **Execute Locally**: Run the required commands in your local shell environment
4. **Grade Automatically**: Compare output against expected results with whitespace normalization
5. **Submit Results**: Send completion status back to the platform

## Features

- **Task Retrieval**: Fetches task details (commands and expected outputs) securely using Bearer token authentication
- **Local Execution**: Runs commands directly on the host machine via shell (`sh -c` on Unix, `cmd /C` on Windows)
- **Silent Error Handling**: Captures standard error output as valid feedback for grading, preventing application crashes on user syntax errors
- **Automated Grading**: Compares local execution output against strict expected criteria with whitespace normalization
- **Visual Feedback**: Provides real-time step-by-step execution logs with spinner animations and styled success/failure messages
- **Mock Backend**: Includes a built-in HTTP mock server for local development and testing
- **Secure Storage**: API tokens stored in `~/.t-cli` with `0600` permissions (user-read/write only)

## Installation

### Prerequisites

- **Go**: Version 1.25.5 or higher ([Download Go](https://go.dev/dl/))
- **Git**: To clone the repository
- **Make**: (Optional) For automated build workflow

### Method 1: Install Using Make (Recommended)

1. Clone the repository:

   ```bash
   git clone https://github.com/Tikkaaa3/t-cli.git
   cd t-cli
   ```

2. Build and install:

   ```bash
   # Build the binary
   make build

   # Install it globally to /usr/local/bin (requires sudo)
   sudo make install
   ```

3. Verify installation:

   ```bash
   t-cli --help
   ```

You can now run `t-cli` from any terminal window!

### Method 2: Manual Build

If you don't have Make or prefer manual installation:

```bash
# Clone the repository
git clone https://github.com/Tikkaaa3/t-cli.git
cd t-cli

# Build the binary
go build -o t-cli main.go

# (Optional) Move to your PATH
sudo mv t-cli /usr/local/bin/
```

### Method 3: Development Mode

For quick testing without installation:

```bash
go run main.go login <your-api-key>
go run main.go <lesson-id>
```

## Usage

The CLI requires an API key to verify your identity. You can generate this key from your learning platform's dashboard or API.

### Login

Authenticate your CLI with the platform. This saves your credentials locally in `~/.t-cli`.

```bash
t-cli login <YOUR_API_KEY>
```

**Example:**

```bash
t-cli login sk_live_abc123xyz789
```

**Output:**

```
Success! Login credentials saved.
```

### Run a Lesson

Execute a lesson by providing its unique lesson ID:

```bash
t-cli <LESSON_ID>
```

**Example:**

```bash
t-cli lesson-golang-basics-001
```

**Sample Workflow:**

```
  > Fetching task for lesson 'lesson-golang-basics-001'...
  > Running: echo Hello World ...
  > Running: echo Testing 123 ...

╭────────────────────────────────────────╮
│ CONGRATULATIONS!                       │
│ All steps executed correctly.          │
╰────────────────────────────────────────╯

  > Saving progress...
✓ Progress saved.
```

### Available Commands

```bash
t-cli --help              # Show help message
t-cli login <api-key>     # Authenticate with the platform
t-cli <lesson-id>         # Run a specific lesson
```

## Configuration

### Config File Location

Your API token is stored in:

```
~/.t-cli
```

### Security

- The config file is created with `0600` permissions (user read/write only)
- Never commit your `~/.t-cli` file to version control
- Your API token is sent via the `Authorization: Bearer <token>` header
- All API communication uses the base URL defined in `internal/api/client.go`

### Changing Base URL

By default, the CLI connects to `http://localhost:8080`. To point to a production API, rebuild with custom `ldflags`:

```bash
go build -ldflags "-X 'github.com/Tikkaaa3/t-cli/internal/api.BaseURL=https://api.your-platform.com'" -o t-cli
```

## API Contract

The CLI expects the backend API to implement the following endpoints:

### GET `/lessons/{lesson_id}/task`

Fetch task details for a given lesson.

**Request Headers:**

```
Authorization: Bearer <api_token>
```

**Response (200 OK):**

```json
{
  "task_id": "task-uuid-12345",
  "steps": [
    {
      "command": "echo Hello World",
      "expected_output": "Hello World"
    },
    {
      "command": "echo Testing 123",
      "expected_output": "Testing 123"
    }
  ]
}
```

**Error Responses:**

- `401 Unauthorized`: Missing or invalid API token
- `404 Not Found`: Lesson ID does not exist
- `500 Internal Server Error`: Server-side error

### POST `/tasks/{task_id}/complete`

Submit successful task completion.

**Request Headers:**

```
Authorization: Bearer <api_token>
```

**Response (200 OK):**

```json
{
  "status": "success"
}
```

**Error Responses:**

- `401 Unauthorized`: Missing or invalid API token
- `404 Not Found`: Task ID does not exist
- `500 Internal Server Error`: Server-side error

**Note:** The CLI only submits results when all steps pass. Failed tasks are not submitted.

## Development

### Local Testing with Mock Server

To test the CLI without a full backend, use the included mock HTTP server. This server returns static task data for testing.

#### Step 1: Start the Mock Server

**Important:** Ensure no other service is using port `8080`.

Open a new terminal window and run:

```bash
go run mock/main.go
```

**Expected Output:**

```
Mock Backend running on http://localhost:8080
   - GET  /lessons/{id}/task
   - POST /tasks/{id}/complete
```

#### Step 2: Login with a Dummy API Key

The mock server checks for an `Authorization` header but doesn't validate it. Use any dummy key:

```bash
go run main.go login mock-api-key
```

**Output:**

```
Success! Login credentials saved.
```

#### Step 3: Run a Mock Lesson

The mock server returns the same task regardless of the lesson ID:

```bash
go run main.go any-lesson-id
```

**Mock Task Details:**

The mock server always returns a task with two steps:

1. `echo Hello World` → expects `"Hello World"`
2. `echo Testing 123` → expects `"Testing 123"`

**Sample Execution:**

```
--- CLI Request: Fetch Task ---
Requested Lesson ID: any-lesson-id
Auth Token:          Bearer mock-api-key
Action: Sending mock task JSON...

  > Running: echo Hello World ...
  > Running: echo Testing 123 ...

╭────────────────────────────────────────╮
│ CONGRATULATIONS!                       │
│ All steps executed correctly.          │
╰────────────────────────────────────────╯

--- CLI Submission: Task Complete ---
Task ID:    task-mock-uuid-123
Auth Token: Bearer mock-api-key
RESULT: PASSED! Marking task as complete...

✓ Progress saved.
```

### Building for Production

To build the CLI for a production environment, use `-ldflags` to set the base URL:

```bash
go build -ldflags "-X 'github.com/Tikkaaa3/t-cli/internal/api.BaseURL=https://api.your-platform.com'" -o t-cli
```

### Development Workflow

```bash
# Clean build artifacts
make clean

# Build locally
make build

# Run quickly in dev mode
make run
```

### Testing Checklist

- [ ] Mock server starts on port 8080
- [ ] Login command saves token to `~/.t-cli`
- [ ] Task fetch includes Authorization header
- [ ] Commands execute in correct shell (`sh -c` or `cmd /C`)
- [ ] Output comparison handles whitespace correctly
- [ ] Successful tasks submit to `/tasks/{task_id}/complete`
- [ ] Failed tasks do not submit

## Project Structure

```
t-cli/
├── cmd/
│   ├── root.go          # Main command orchestration and workflow
│   └── login.go         # Login command and token management
├── internal/
│   ├── api/
│   │   └── client.go    # HTTP client for backend communication
│   ├── executor/
│   │   └── runner.go    # Command execution engine (cross-platform)
│   ├── grader/
│   │   └── match.go     # Output comparison and validation logic
│   └── ui/
│       └── styles.go    # Lipgloss styling and visual components
├── mock/
│   └── main.go          # Standalone HTTP mock server for testing
├── main.go              # Application entry point
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
├── Makefile             # Build automation
├── .gitignore           # Git ignore patterns
└── README.md            # This file
```

### Package Descriptions

**`cmd/`** - Command-line interface layer

- Defines all Cobra commands (`root`, `login`)
- Orchestrates the complete workflow (fetch → execute → grade → submit)
- Handles user authentication and config file management

**`internal/api/`** - Backend communication

- HTTP client with Bearer token authentication
- Endpoint definitions for task retrieval and result submission
- Timeout handling (10-second default)

**`internal/executor/`** - Command execution

- Cross-platform command runner (supports Windows `cmd /C` and Unix `sh -c`)
- Captures combined stdout and stderr
- Handles execution errors gracefully without crashing

**`internal/grader/`** - Automated grading

- Compares actual vs. expected output
- Normalizes whitespace using `strings.TrimSpace()`
- Returns pass/fail boolean

**`internal/ui/`** - User interface

- [Lipgloss](https://github.com/charmbracelet/lipgloss) styling definitions
- Success/failure message formatting
- Colored borders and text rendering

**`mock/`** - Development server

- Standalone HTTP server returning static mock data
- Simulates backend API for local testing
- Accepts any lesson ID and returns the same test task

## Architecture

### Execution Flow

```
┌─────────────┐
│  User runs  │
│  t-cli <id> │
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────────────────┐
│  1. Authentication Check                    │
│     - Read ~/.t-cli for API token           │
│     - Exit if not logged in                 │
└──────┬──────────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────────┐
│  2. Fetch Task (API Layer)                  │
│     - GET /lessons/{id}/task                │
│     - Authorization: Bearer <token>         │
│     - Parse JSON response                   │
└──────┬──────────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────────┐
│  3. Execute Commands (Executor Layer)       │
│     - For each step:                        │
│       • Run command via shell               │
│       • Capture stdout + stderr             │
│       • Store result                        │
└──────┬──────────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────────┐
│  4. Grade Output (Grader Layer)             │
│     - Compare each actual vs. expected      │
│     - Trim whitespace on both sides         │
│     - Return pass/fail                      │
└──────┬──────────────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────────────┐
│  5. Display Results (UI Layer)              │
│     - Show styled success/failure box       │
│     - If passed: submit to backend          │
└──────┬──────────────────────────────────────┘
       │
       ▼ (only if passed)
┌─────────────────────────────────────────────┐
│  6. Submit Result (API Layer)               │
│     - POST /tasks/{task_id}/complete        │
│     - Authorization: Bearer <token>         │
└─────────────────────────────────────────────┘
```

### Key Design Decisions

1. **Orchestration**: The root command (`cmd/root.go`) coordinates the entire workflow
2. **Authentication**: Uses Bearer token pattern with secure local storage (`0600` permissions)
3. **Execution**: Runs tasks sequentially; stderr is captured for grading, not treated as fatal
4. **Grading**: Whitespace-normalized string comparison (`strings.TrimSpace()`)
5. **Reporting**: Only successful completions are submitted to the backend
6. **Error Handling**: Network/API errors are displayed but don't crash the CLI
7. **Cross-Platform**: Shell command execution adapts to OS (`sh -c` vs `cmd /C`)

### Data Flow

**Task Structure:**

```go
type Task struct {
    ID    string        `json:"task_id"`
    Steps []CommandStep `json:"steps"`
}

type CommandStep struct {
    Command        string `json:"command"`
    ExpectedOutput string `json:"expected_output"`
}
```

**Execution Result:**

- Array of output strings (one per step)
- Grader compares index-by-index against expected outputs
- All steps must pass for overall success

## Troubleshooting

### Common Issues

#### "you are not logged in. Run 't-cli login <token>' first"

**Cause:** No API token found in `~/.t-cli`

**Solution:**

```bash
t-cli login <your-api-key>
```

---

#### "Could not fetch task: API request failed with status: 401"

**Cause:** Invalid or expired API token

**Solution:**

1. Generate a new API key from your platform dashboard
2. Re-login with the new key:

```bash
t-cli login <new-api-key>
```

---

#### "Could not fetch task: API request failed with status: 404"

**Cause:** Lesson ID does not exist on the backend

**Solution:**

- Verify the lesson ID is correct
- Check that the lesson is published and accessible

---

#### "TEST FAILED - Your output did not match expectation"

**Cause:** Command output doesn't match expected output (even with whitespace normalization)

**Debugging Steps:**

1. Check for extra spaces, tabs, or newlines
2. Verify the command runs correctly in your terminal
3. Check for OS-specific output differences (line endings, paths)
4. Review the task requirements carefully

---

#### "Could not save to server (but you passed locally!)"

**Cause:** Network error or backend unavailable during submission

**Solution:**

- Check your internet connection
- Verify the backend server is running
- Re-run the lesson to retry submission

---

#### Port 8080 already in use (Mock Server)

**Cause:** Another process is using port 8080

**Solution:**

```bash
# Find and kill the process using port 8080
lsof -ti:8080 | xargs kill -9

# Or change the mock server port in mock/main.go
```

---

#### Permission denied when installing

**Cause:** Installing to `/usr/local/bin` requires elevated privileges

**Solution:**

```bash
sudo make install
```

Or install to a user directory:

```bash
# Build
make build

# Move to user bin
mkdir -p ~/bin
mv t-cli ~/bin/

# Add to PATH in ~/.bashrc or ~/.zshrc
export PATH="$HOME/bin:$PATH"
```

### Debug Mode

To see detailed HTTP request/response information, modify `internal/api/client.go` to add debug logging.

### Getting Help

If you encounter issues not covered here:

1. Check the [GitHub Issues](https://github.com/Tikkaaa3/t-cli/issues)
2. Open a new issue with:
   - Your OS and Go version
   - Complete error message
   - Steps to reproduce
   - Expected vs. actual behavior

## Contributing

Contributions are welcome! This project is a learning tool and benefits from community improvements.

### How to Contribute

1. **Fork the repository**

   Click the "Fork" button on GitHub, then clone your fork:

   ```bash
   git clone https://github.com/YOUR_USERNAME/t-cli.git
   cd t-cli
   ```

2. **Create a feature branch**

   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Write clear, documented code
   - Follow existing code style and patterns
   - Test your changes locally

4. **Test thoroughly**

   ```bash
   # Run with mock server
   go run mock/main.go  # Terminal 1
   go run main.go login test-key  # Terminal 2
   go run main.go test-lesson
   ```

5. **Commit with clear messages**

   ```bash
   git commit -m "Add feature: description of what changed"
   ```

6. **Push and create a Pull Request**

   ```bash
   git push origin feature/your-feature-name
   ```

### Areas for Improvement

Potential contribution ideas:

- **Testing**: Add unit tests for executor, grader, and API client
- **CI/CD**: GitHub Actions workflow for automated builds
- **Features**:
  - Progress tracking (list completed lessons)
  - Retry failed steps
  - Verbose mode for debugging
  - Timeout configuration for long-running commands
- **Documentation**: Additional examples, video tutorials
- **Backend**: Reference implementation of the API server

### Code Style

- Follow standard Go conventions (`gofmt`, `golint`)
- Keep functions small and focused
- Add comments for complex logic
- Use descriptive variable names
- Handle errors explicitly

## Motivation & Inspiration

This tool was inspired by the excellent educational experience provided by [Boot.dev](https://boot.dev). Their platform bridges the gap between web-based coding tutorials and real local environments.

Specifically, the design and goal of this project were influenced by the [Boot.dev CLI](https://github.com/bootdotdev/bootdev), which allows students to submit locally written code for verification. This project aims to explore similar architectural patterns for building robust, user-friendly educational CLIs in Go.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
