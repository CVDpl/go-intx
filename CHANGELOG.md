# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of go-intx library
- Support for 24-bit, 40-bit, 48-bit, and 56-bit integer types
- Both signed (Int24, Int40, Int48, Int56) and unsigned (Uint24, Uint40, Uint48, Uint56) types
- Safe constructors with error handling (`NewInt24`, `NewUint24`, etc.)
- Panic constructors for convenience (`MustInt24`, `MustUint24`, etc.)
- Big-endian and little-endian byte conversion methods
- JSON marshaling and unmarshaling support
- Binary marshaling and unmarshaling support
- String representation methods
- Comprehensive test suite with 160+ tests
- Performance benchmarks
- Modular package structure for selective imports
- Complete documentation and examples

### Features
- **Range Validation**: All constructors validate input ranges
- **Byte Conversion**: `ToBytes()` and `ToLittleEndianBytes()` methods
- **From Bytes**: `FromInt24Bytes()`, `FromUint24Bytes()`, etc. with error handling
- **JSON Support**: Implements `json.Marshaler` and `json.Unmarshaler` interfaces
- **Binary Support**: Implements `encoding.BinaryMarshaler` and `encoding.BinaryUnmarshaler` interfaces
- **String Support**: Implements `fmt.Stringer` interface
- **Performance Optimized**: Zero allocations for most operations
- **Modular Design**: Import only the types you need

### Performance
- Constructors: ~0.25 ns/op
- Byte conversion: ~0.25-0.30 ns/op
- JSON marshaling: ~95-125 ns/op
- String conversion: ~17-22 ns/op

### Documentation
- Comprehensive README.md with examples
- API reference documentation
- Performance benchmarks
- Usage examples
- Contributing guidelines
- Code of conduct
- Security policy

## [1.0.0] - 2024-12-19

### Added
- Initial release
- All core functionality as described above

---

## Version History

- **1.0.0**: Initial release with all core features 