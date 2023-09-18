
# Chorus

## Table of Contents
- [Overview](#overview)
- [Directory Structure](#directory-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Development](#development)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [To-Do](#to-do)

## Overview
Chorus is a tool designed to replatform legacy and proprietary code to modern languages and frameworks, such as Golang, Python, TypeScript et cetera. It offers a configurable pipeline that lets you define the automated steps that you would like to run to achieve replatforming. Typically, the pipeline will include steps such as code generation, unit testing, linting produced code to a predefined threshold, running custom unit tests. 

In the example included in the [sample](./sample/) directory,  Chorus takes SAS code as input, hydrates an optimized prompt to ask Gen AI to parse the SAS code and then convert it to logically and functionally equivalent Python code. Chorus orchestates a dynamic dialog with the LLM to vet the output generated by the model by running unit tests, linting etc.

You can configure Chorus using the Chorus manifest.

## Directory Structure
```
chorus-main/
├── .gitignore      # Specifies files and folders to be ignored by Git
├── README.md       # This README file
├── chorus.yaml     # YAML configuration file for the project
├── go.mod          # Go module file for dependency management
├── go.sum          # Go sum file for dependency management
├── main.go         # Main Go source file (entry point)
├── pkg/            # Utility packages used in the project
└── sample/         # Directory for sample SAS and Python code
```

## Installation

### Prerequisites
- Go 1.18 or higher
- Git

### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/<username>/chorus.git
    ```
2. Navigate to the project directory:
    ```bash
    cd chorus
    ```
3. Install the dependencies:
    ```bash
    go mod download
    ```

## Usage
To run Chorus, pass in the top-level directory location that contains input code and a directory for 
generated output files:
```bash
go run main.go ./sample
```

## Configuration
Chorus uses a YAML configuration file (`chorus.yaml`) that contains the base prompt, model details, and pipeline definition. Edit this file to customize the prompt or add new instructions.

## Development
The project is structured as follows:
- `main.go`: The entry point of the application.
- `pkg`: Contains various utility packages used in the project.
- `sample`: Sample SAS code and corresponding Python code can be placed here for testing and demonstration.

### Dependencies
The project uses Go modules for managing dependencies. The `go.mod` and `go.sum` files contain the required modules.

## Testing
Testing features are yet to be implemented. Please refer to the [To-Do](#to-do) section for more details.

## Contributing
Contributions are welcome! Feel free to submit a pull request.

## To-Do
- Optimize the code
- Use Go routines for parallel processing
- Implement linting and unit tests
- Add ability to run custom unit tests
- Print conversations only when the debug flag is set to TRUE
- Store conversational output for logging in files
- Optimize tabular output

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
