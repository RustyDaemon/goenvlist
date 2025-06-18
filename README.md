# goenvlist

A simple cross-platform CLI tool for displaying environment variables with advanced filtering options.

## Features

- List all environment variables
- Show only common variables
- Display only PATH variable
- Filter by specific variable names
- Cross-platform support (macOS, Linux, Windows)
- Color-coded output for better readability
- Raw output mode for scripting

## Usage

```bash
# Show all environment variables
goenvlist

# Show only common environment variables
goenvlist -s
goenvlist --simple

# Show only PATH variable
goenvlist -p
goenvlist --path

# Show specific variables (comma-separated)
goenvlist -f HOME,USER
goenvlist --filter HOME,USER

# Use raw output (unformatted). Useful for scripting)
goenvlist -r
goenvlist --raw

# Help
goenvlist -h
goenvlist --help
```

## License

The project is licensed under the [MIT License](LICENSE).
