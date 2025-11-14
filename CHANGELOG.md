# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2025-11-14

### Added
- Initial release of Landing AI Go SDK
- Support for ADE Parse API (`/v1/ade/parse`)
- Parse documents from file uploads
- Parse documents from URLs
- Parse in-memory file data
- Multi-region support (US and EU)
- Model version selection (DPT-1, DPT-2, DPT-2 mini)
- Page-level document splitting
- Comprehensive error handling with custom error types
- Full type definitions for all API responses
- Support for all chunk types (text, table, figure, logo, card, attestation, scan_code, marginalia)
- Grounding information with bounding box coordinates
- Client configuration options (timeout, custom HTTP client, base URL)
- Context support for request cancellation
- Complete documentation with examples
- Example code demonstrating all features

### Features
- **Document Parsing**: Parse PDFs, images, and spreadsheets
- **URL Support**: Parse documents directly from URLs
- **Flexible Configuration**: Custom timeouts, HTTP clients, and base URLs
- **Type Safety**: Full type definitions matching the Landing AI API
- **Error Handling**: Comprehensive error types for all API responses
- **Regional Support**: US (default) and EU regions
- **Model Selection**: Choose from multiple DPT model versions
- **Splitting**: Optional page-level document splitting
- **Grounding**: Bounding box coordinates for each chunk

### Documentation
- Comprehensive README with installation and usage instructions
- Complete API reference
- Working examples for all features
- Contributing guidelines
- Changelog

[Unreleased]: https://github.com/youssefsiam38/landingai/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/youssefsiam38/landingai/releases/tag/v0.1.0
