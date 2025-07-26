# Contributing to go-intx

Thank you for your interest in contributing to go-intx! This document provides guidelines for contributing to this project.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

- Use the GitHub issue tracker
- Include a clear and descriptive title
- Provide detailed steps to reproduce the bug
- Include Go version and operating system information
- Add code examples if applicable

### Suggesting Enhancements

- Use the GitHub issue tracker
- Describe the enhancement clearly
- Explain why this enhancement would be useful
- Include examples of how it would be used

### Pull Requests

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass (`go test ./...`)
6. Run benchmarks to ensure no performance regression (`go test -bench=. intx_bench_test.go`)
7. Update documentation if needed
8. Commit your changes (`git commit -m 'Add amazing feature'`)
9. Push to the branch (`git push origin feature/amazing-feature`)
10. Open a Pull Request

## Development Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/CVDpl/go-intx.git
   cd go-intx
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run tests:
   ```bash
   go test -v intx_test.go
   ```

4. Run benchmarks:
   ```bash
   go test -bench=. intx_bench_test.go
   ```

## Code Style Guidelines

- Follow Go's official formatting guidelines (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused
- Write comprehensive tests for new functionality

## Testing Guidelines

- All new code must include tests
- Tests should cover both success and error cases
- Use table-driven tests for multiple test cases
- Benchmark new functionality to ensure performance

## Performance Guidelines

- Avoid unnecessary allocations
- Use `strconv.AppendInt/AppendUint` instead of `fmt.Sprintf` when possible
- Pre-define error variables to avoid runtime allocations
- Run benchmarks before and after changes

## Commit Message Guidelines

Use conventional commit format:

- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test additions or changes
- `refactor:` for code refactoring
- `perf:` for performance improvements

Example:
```
feat: add new Int32 type for 32-bit integers

- Add Int32 and Uint32 types
- Implement all standard interfaces
- Add comprehensive tests
- Update documentation
```

## Review Process

1. All pull requests require review
2. At least one maintainer must approve
3. All tests must pass
4. Code must follow style guidelines
5. Documentation must be updated if needed

## Questions?

If you have questions about contributing, please open an issue or contact the maintainers.

Thank you for contributing to go-intx! 