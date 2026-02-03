# t-cli

t-cli is an educational command-line tool designed to execute, verify, and grade local code submissions against a remote learning platform. It acts as a bridge between a student's local development environment and a centralized curriculum backend.

## Overview

The CLI retrieves tasks from an API, executes the required shell commands or scripts on the user's machine, compares the output against expected results, and submits the final grade back to the server. It is built with Go and utilizes the [Cobra](https://github.com/spf13/cobra) library for command management and [Charm](https://github.com/charmbracelet) libraries for the user interface.

## Features

- **Task Retrieval**: Fetches task details (commands and expected outputs) securely using a task token.
- **Local Execution**: Runs commands directly on the host machine, supporting real-world development workflows.
- **Silent Error Handling**: Captures standard error output as valid feedback for grading, preventing application crashes on user syntax errors.
- **Automated Grading**: Compares local execution output against strict expected criteria, handling whitespace normalization.
- **Visual Feedback**: Provides real-time step-by-step execution logs and a polished success/failure interface.
- **Mock Backend**: Includes a built-in mock server for local development and testing.

## Installation

### Prerequisites

- **Go**: Version 1.18 or higher.

### Build from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/Tikkaaa3/t-cli.git
   cd t-cli
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the binary:

   ```bash
   go build -o t-cli main.go
   ```

## Usage

The CLI requires a **Task Token** to function. This token identifies the user and the specific task to be graded.

### Basic Usage

```bash
./t-cli <task_token>
```

## Development Mode (Mock Server)

To test the CLI without a real backend, use the included mock server.

### Start the Mock Server

Open a new terminal window and run:

```bash
go run mock/main.go
```

This will start a local server at <http://localhost:8080>.

### Run the CLI: In your main terminal, run the CLI with any placeholder token

```bash
go run main.go TEST-TOKEN-123
```

## Production Build

To build the CLI for a production environment (pointing to a real API), use the -ldflags option to overwrite the base URL variable during compilation.

```bash
go build -ldflags "-X 'github.com/Tikkaaa3/t-cli/internal/api.BaseURL=https://api.your-platform.com'" -o t-cli
```

## Project Structure

- cmd/: Contains the main entry point and Cobra command definitions (root.go).

- internal/api/: Handles HTTP communication with the backend (fetching tasks, submitting results).

- internal/executor/: Logic for running system commands and capturing output.

- internal/grader/: Logic for comparing actual output against expected output.

- internal/ui/: Styling definitions using [Lipgloss](https://github.com/charmbracelet/lipgloss) and interface helpers.

- mock/: A standalone HTTP server for testing the CLI in isolation.

## Architecture

1. Orchestration: The root command coordinates the workflow.

2. API Layer: Authenticates using the Bearer token pattern.

3. Execution Layer: Runs tasks sequentially. If a step fails or prints to stderr, the output is captured for grading rather than halting the process.

4. Grading Layer: Trims whitespace and performs string comparison to determine pass/fail status.

5. Reporting: Submits a boolean status back to the platform.

## Motivation & Inspiration

This tool was inspired by the excellent educational experience provided by [Boot.dev](https://boot.dev). Their platform bridges the gap between web-based coding tutorials and real local environments.

Specifically, the design and goal of this project were influenced by the [Boot.dev CLI](https://github.com/bootdotdev/bootdev), which allows students to submit locally written code for verification. This project aims to explore similar architectural patterns for building robust, user-friendly educational CLIs in Go.
