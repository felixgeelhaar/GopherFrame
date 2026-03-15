# Security Policy

## Supported Versions

| Version | Supported          |
|---------|--------------------|
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability in GopherFrame, please report it responsibly.

### How to Report

1. **Do NOT open a public GitHub issue** for security vulnerabilities
2. Email security concerns to the maintainer via [GitHub private vulnerability reporting](https://github.com/felixgeelhaar/GopherFrame/security/advisories/new)
3. Include:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

### Response Timeline

- **Acknowledgment**: Within 48 hours
- **Initial Assessment**: Within 5 business days
- **Fix Timeline**: Depends on severity
  - **Critical**: Patch within 7 days
  - **High**: Patch within 14 days
  - **Medium**: Patch within 30 days
  - **Low**: Next scheduled release

## Security Measures

### Current Protections

- **Path Traversal Protection**: All file I/O operations validate paths against directory traversal attacks
- **Input Validation**: All user-provided data is validated before processing
- **Dependency Scanning**: Automated vulnerability scanning via `gosec` and Dependabot
- **CI Security Gates**: Security scanning runs on every PR via GitHub Actions
- **Supply Chain**: GitHub Actions pinned to commit SHAs to prevent tag manipulation attacks

### Dependency Management

- Dependencies are kept minimal (Apache Arrow Go as the only core dependency)
- Automated dependency updates via Dependabot
- All dependencies are scanned for known vulnerabilities
- Go module checksums verified via `go mod verify`

### Safe Coding Practices

- No use of `unsafe` package in production code
- No shell command execution
- No network operations (library is purely computational)
- Memory-safe operations with proper bounds checking
- Concurrent operations use proper synchronization

## Scope

The following are in scope for security reports:

- Memory safety issues (buffer overflows, use-after-free)
- Path traversal in I/O operations
- Denial of service via crafted input files
- Information disclosure through error messages
- Dependency vulnerabilities

The following are out of scope:

- Performance issues (unless they enable DoS)
- Feature requests
- Issues in example code that doesn't run in production
