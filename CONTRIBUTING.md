# Contributing to Landing AI Go SDK

Thank you for your interest in contributing to the Landing AI Go SDK! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Code Style](#code-style)
- [Release Process](#release-process)

## Code of Conduct

This project adheres to a code of conduct that all contributors are expected to follow:

- Be respectful and inclusive
- Welcome newcomers and help them get started
- Focus on what is best for the community
- Show empathy towards other community members

## Getting Started

### Prerequisites

- Go 1.24.0 or higher
- Git
- A GitHub account

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/landingai.git
cd landingai
```

3. Add the upstream repository:

```bash
git remote add upstream https://github.com/youssefsiam38/landingai.git
```

## Development Setup

### Install Dependencies

```bash
go mod download
go mod verify
```

### Install Development Tools

```bash
# Install golangci-lint for linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install goreleaser for releases (optional)
go install github.com/goreleaser/goreleaser@latest
```

### Verify Setup

```bash
# Run tests
go test ./...

# Run linter
golangci-lint run

# Build the project
go build ./...
```

## Making Changes

### Create a Branch

Create a feature branch for your changes:

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

Branch naming conventions:
- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `refactor/` - Code refactoring
- `test/` - Test additions or modifications

### Commit Messages

Write clear, descriptive commit messages following the conventional commits format:

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:

```
feat(parse): add support for spreadsheet parsing

Add support for parsing Excel and CSV files.
Includes automatic format detection and conversion.

Closes #123
```

```
fix(client): handle timeout correctly in Parse requests

Previously, timeouts were not properly detected in some edge cases.
Now explicitly checks for context deadline exceeded.

Fixes #456
```

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...

# Run specific test
go test -v ./tests/ -run TestNewClient
```

### Writing Tests

- Place tests in the `tests/` directory or alongside the code
- Use table-driven tests when appropriate
- Test both success and error cases
- Mock external dependencies (API calls)
- Aim for >80% code coverage

Example test:

```go
func TestParseRequest(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   "test.pdf",
            want:    "expected",
            wantErr: false,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ParseDocument(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("ParseDocument() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("ParseDocument() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Submitting Changes

### Before Submitting

1. **Run all tests**:
   ```bash
   go test ./...
   ```

2. **Run linter**:
   ```bash
   golangci-lint run
   ```

3. **Format code**:
   ```bash
   gofmt -w .
   ```

4. **Update documentation** if needed:
   - Update README.md
   - Add/update examples
   - Update CHANGELOG.md

### Pull Request Process

1. **Update your branch** with the latest upstream changes:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

2. **Push your changes**:
   ```bash
   git push origin feature/your-feature-name
   ```

3. **Create a Pull Request** on GitHub with:
   - Clear title describing the change
   - Detailed description of what changed and why
   - Link to related issues (e.g., "Closes #123")
   - Screenshots/examples if applicable

4. **Address review comments**:
   - Respond to all comments
   - Make requested changes
   - Push updates to the same branch

5. **Wait for approval** from maintainers

### Pull Request Guidelines

- Keep PRs focused on a single feature or fix
- Include tests for new functionality
- Update documentation as needed
- Ensure CI passes (all tests and lints)
- Rebase on main before merging
- Squash commits if requested

## Code Style

### Go Style Guide

Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) and [Effective Go](https://golang.org/doc/effective_go).

Key points:

- Use `gofmt` for formatting
- Use meaningful variable names
- Keep functions small and focused
- Document all exported types and functions
- Handle errors explicitly
- Use interfaces for abstraction

### Documentation

- All exported types, functions, and methods must have comments
- Comments should be complete sentences
- Use examples to illustrate usage

Example:

```go
// NewClient creates a new Landing AI client with the given API key and options.
// The client can be configured with region selection, custom HTTP clients, and timeouts.
//
// Example:
//
//    client := landingai.NewClient(
//        "your-api-key",
//        landingai.WithRegion(landingai.RegionEU),
//    )
func NewClient(apiKey string, opts ...ClientOption) *Client {
    // implementation
}
```

### Error Handling

- Return errors, don't panic
- Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Use custom error types for specific errors
- Check all errors

```go
// Good
resp, err := client.Parse(ctx).WithFile("doc.pdf").Do()
if err != nil {
    return fmt.Errorf("failed to parse document: %w", err)
}

// Bad
resp, _ := client.Parse(ctx).WithFile("doc.pdf").Do()  // Don't ignore errors
```

## Project Structure

```
landingai/
├── *.go              # Core SDK files
├── tests/            # Test files
├── examples/         # Example programs
├── .github/          # GitHub workflows
└── docs/             # Additional documentation
```

## Adding New Features

### Process

1. **Open an issue** first to discuss the feature
2. **Wait for approval** from maintainers
3. **Implement the feature** following guidelines
4. **Add tests** with good coverage
5. **Add examples** if applicable
6. **Update documentation**
7. **Submit PR** for review

### Feature Checklist

- [ ] Implementation complete
- [ ] Tests added and passing
- [ ] Documentation updated
- [ ] Examples added (if applicable)
- [ ] CHANGELOG.md updated
- [ ] Linter passes
- [ ] No breaking changes (or clearly documented)

## Release Process

Releases are automated using GoReleaser and GitHub Actions.

### Versioning

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for new functionality (backward compatible)
- **PATCH** version for bug fixes (backward compatible)

### Creating a Release

Maintainers only:

1. Update `CHANGELOG.md` with release notes
2. Create and push a tag:
   ```bash
   git tag -a v0.1.0 -m "Release v0.1.0"
   git push origin v0.1.0
   ```
3. GitHub Actions will automatically:
   - Run tests
   - Build artifacts
   - Create GitHub release
   - Update documentation

## Getting Help

- **Questions**: Open a [GitHub Discussion](https://github.com/youssefsiam38/landingai/discussions)
- **Bugs**: Open a [GitHub Issue](https://github.com/youssefsiam38/landingai/issues)
- **Security**: Email security concerns to youssef.siam38@gmail.com

## Recognition

Contributors will be recognized in:
- CHANGELOG.md for their contributions
- GitHub contributors page
- Release notes

Thank you for contributing to the Landing AI Go SDK! 🎉
