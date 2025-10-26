# Contributing to KGP

Thank you for your interest in contributing to KGP! This document provides guidelines for contributing to the project.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue on GitHub with:
- A clear description of the bug
- Steps to reproduce the issue
- Expected vs actual behavior
- Your Go version and OS
- Relevant terminal information (if applicable)

### Suggesting Features

Feature requests are welcome! Please open an issue describing:
- The use case for the feature
- How it would work
- Why it would be valuable

### Submitting Pull Requests

1. **Fork the repository** and create a new branch from `main`
2. **Write tests** for your changes
3. **Ensure all tests pass**: `go test -v ./...`
4. **Run go fmt**: `gofmt -w .`
5. **Run go vet**: `go vet ./...`
6. **Update documentation** if needed
7. **Submit a pull request** with a clear description

## Development Guidelines

### Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Write clear, descriptive function and variable names
- Add comments for exported functions and types
- Keep functions focused and concise

### Testing

- Maintain or improve test coverage (currently 97.4%)
- Write tests for new features
- Ensure tests are deterministic and isolated
- Use table-driven tests where appropriate

### Commit Messages

- Use clear, descriptive commit messages
- Start with a verb in present tense (e.g., "Add", "Fix", "Update")
- Reference issue numbers when applicable

Example:
```
Add support for Unicode placeholders

Implements virtual placement feature for Unicode diacritic placeholders.
Fixes #123
```

### Documentation

- Update README.md for new features
- Add examples for new functionality
- Keep API documentation up to date
- Document any breaking changes

## Testing Locally

```bash
# Run all tests
go test -v ./...

# Check coverage
go test -cover ./...

# Run with race detector
go test -race ./...

# Format code
gofmt -w .

# Vet code
go vet ./...

# Build the module
go build ./...
```

## Code Review Process

All contributions go through code review. Reviewers will check for:
- Code quality and style
- Test coverage
- Documentation
- Adherence to project guidelines

## Questions?

Feel free to open an issue for questions or discussion about contributions.

## License

By contributing to KGP, you agree that your contributions will be licensed under the MIT License.
