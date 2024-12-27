## Project organization


```yml
my-api-project/
│
├── cmd/                # Contains the main application (entry point of the program)
│   └── main.go         # The main file that starts the API
│
├── internal/           # Business logic and internal components
│   └── handler/        # Functions that handle HTTP requests
│   └── model/          # Data models (structures) of the application
│   └── service/        # Business logic, such as create, read, etc.
│
├── pkg/                # Reusable code that can be used in other projects
│   └── util/           # Utility functions (validator, logger, etc.)
│
├── go.mod              # Go dependency management file
├── go.sum              # Dependency checksums
└── README.md           # Project documentation

```