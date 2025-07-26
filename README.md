# go-intx

A Go library providing fixed-width integer types (24, 40, 48, 56 bits) with API and behavior similar to built-in `int`/`uint` types.

## Overview

This library extends Go's built-in integer types by providing custom types for 24, 40, 48, and 56-bit signed and unsigned integers. These types are useful when working with:

- Network protocols requiring specific bit-width integers
- Database schemas with custom integer sizes
- Binary file formats with non-standard integer representations
- Memory-constrained environments where every bit counts

## Features

- **8 Integer Types**: `Int24`, `Uint24`, `Int40`, `Uint40`, `Int48`, `Uint48`, `Int56`, `Uint56`
- **Range Validation**: Safe constructors with error handling
- **Byte Conversion**: Big-endian and little-endian byte representations
- **Standard Interfaces**: Implements `fmt.Stringer`, `json.Marshaler`, `json.Unmarshaler`, `encoding.BinaryMarshaler`, `encoding.BinaryUnmarshaler`
- **Modular Design**: Import only the types you need using separate packages
- **Comprehensive Testing**: Full test coverage with benchmarks

## Installation

```bash
go get github.com/CVDpl/go-intx
```

## Quick Start

### Import all types

```go
package main

import (
    "fmt"
    "log"
    
    . "intx/24"  // Import 24-bit types
    . "intx/40"  // Import 40-bit types
    . "intx/48"  // Import 48-bit types
    . "intx/56"  // Import 56-bit types
)

func main() {
    // Create a 24-bit unsigned integer
    u24, err := NewUint24(123456)
    if err != nil {
        log.Fatal(err)
    }
    
    // Convert to bytes (big-endian)
    bytes := u24.ToBytes()
    fmt.Printf("Bytes: %x\n", bytes) // Output: Bytes: 01e240
    
    // Create from bytes
    recovered, err := FromUint24Bytes(bytes)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Value: %d\n", recovered.Uint64()) // Output: Value: 123456
}
```

### Import only specific types

```go
package main

import (
    "fmt"
    
    . "intx/24"  // Import only 24-bit types
)

func main() {
    // Use 24-bit types directly without package prefix
    u24 := MustUint24(123456)
    i24 := MustInt24(-123456)
    
    fmt.Printf("Uint24: %d\n", u24.Uint64())
    fmt.Printf("Int24: %d\n", i24.Int64())
}
```

## API Reference

### Types

+----------+---------+---------------------------------------------------+----------+
|   Type   |  Size   |                       Range                       | Storage  |
|----------+---------+---------------------------------------------------+----------+
| `Int24`  | 24 bits | -8,388,608 to 8,388,607                           | `int32`  |
| `Uint24` | 24 bits | 0 to 16,777,215                                   | `uint32` |
| `Int40`  | 40 bits | -549,755,813,888 to 549,755,813,887               | `int64`  |
| `Uint40` | 40 bits | 0 to 1,099,511,627,775                            | `uint64` |
| `Int48`  | 48 bits | -140,737,488,355,328 to 140,737,488,355,327       | `int64`  |
| `Uint48` | 48 bits | 0 to 281,474,976,710,655                          | `uint64` |
| `Int56`  | 56 bits | -36,028,797,018,963,968 to 36,028,797,018,963,967 | `int64`  |
| `Uint56` | 56 bits | 0 to 72,057,594,037,927,935                       | `uint64` |
+----------+---------+---------------------------------------------------+----------+

### Constructors

Each type provides two constructors:

```go
// Safe constructor - returns error if value is out of range
value, err := NewInt24(123456)
if err != nil {
    // Handle error
}

// Panic constructor - panics if value is out of range
value := MustInt24(123456)
```

### Methods

All types implement the following methods:

#### Value Conversion
```go
// Convert to standard Go types
int64Value := int24.Int64()
uint64Value := uint24.Uint64()
```

#### Byte Conversion
```go
// Convert to bytes (big-endian)
bytes := value.ToBytes()

// Convert to bytes (little-endian)
bytes := value.ToLittleEndianBytes()

// Create from bytes (big-endian)
value, err := FromInt24Bytes(bytes)

// Create from bytes (little-endian)
value, err := FromInt24LittleEndianBytes(bytes)
```

#### String Representation
```go
// Convert to string
str := value.String()
```

#### JSON Support
```go
// Marshal to JSON
jsonData, err := json.Marshal(value)

// Unmarshal from JSON
err := json.Unmarshal(jsonData, &value)
```

#### Binary Marshaling
```go
// Marshal to binary
binaryData, err := value.MarshalBinary()

// Unmarshal from binary
err := value.UnmarshalBinary(binaryData)
```

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    
    . "intx/24"  // Import 24-bit types
    . "intx/40"  // Import 40-bit types
)

func main() {
    // Create different types
    i24 := MustInt24(-123456)
    u40 := MustUint40(123456789012)
    
    // Convert to bytes
    i24Bytes := i24.ToBytes()
    u40Bytes := u40.ToBytes()
    
    fmt.Printf("Int24: %d -> %x\n", i24.Int64(), i24Bytes)
    fmt.Printf("Uint40: %d -> %x\n", u40.Uint64(), u40Bytes)
}
```

### JSON Serialization

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    
    . "intx/24"  // Import 24-bit types
    . "intx/48"  // Import 48-bit types
)

type Config struct {
    ID      Uint48 `json:"id"`
    Version Int24  `json:"version"`
}

func main() {
    config := Config{
        ID:      MustUint48(123456789012345),
        Version: MustInt24(-100),
    }
    
    jsonData, err := json.Marshal(config)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("JSON: %s\n", jsonData)
    
    // Unmarshal
    var newConfig Config
    err = json.Unmarshal(jsonData, &newConfig)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("ID: %d, Version: %d\n", newConfig.ID.Uint64(), newConfig.Version.Int64())
}
```

### Network Protocol Example

```go
package main

import (
    "bytes"
    "fmt"
    
    . "intx/24"  // Import 24-bit types
    . "intx/40"  // Import 40-bit types
    . "intx/48"  // Import 48-bit types
)

type Packet struct {
    Header   Uint24
    Length   Uint40
    Checksum Uint48
}

func (p Packet) ToBytes() []byte {
    var buf bytes.Buffer
    
    // Write header (3 bytes)
    headerBytes := p.Header.ToBytes()
    buf.Write(headerBytes[:])
    
    // Write length (5 bytes)
    lengthBytes := p.Length.ToBytes()
    buf.Write(lengthBytes[:])
    
    // Write checksum (6 bytes)
    checksumBytes := p.Checksum.ToBytes()
    buf.Write(checksumBytes[:])
    
    return buf.Bytes()
}

func ParsePacket(data []byte) (Packet, error) {
    if len(data) < 14 { // 3 + 5 + 6
        return Packet{}, fmt.Errorf("insufficient data")
    }
    
    header, err := FromUint24Bytes(data[0:3])
    if err != nil {
        return Packet{}, err
    }
    
    length, err := FromUint40Bytes(data[3:8])
    if err != nil {
        return Packet{}, err
    }
    
    checksum, err := FromUint48Bytes(data[8:14])
    if err != nil {
        return Packet{}, err
    }
    
    return Packet{Header: header, Length: length, Checksum: checksum}, nil
}
```

### Error Handling

```go
package main

import (
    "fmt"
    
    . "intx/24"  // Import 24-bit types
    . "intx/48"  // Import 48-bit types
)

func main() {
    // Handle overflow errors
    _, err := NewUint24(1 << 24) // 16777216
    if err != nil {
        fmt.Printf("Error: %v\n", err) // Error: value exceeds maximum for Uint24
    }
    
    // Handle invalid byte length
    _, err = FromUint48Bytes([]byte{1, 2, 3})
    if err != nil {
        fmt.Printf("Error: %v\n", err) // Error: invalid byte length
    }
}
```

## Testing

Run all tests:

```bash
go test -v intx_test.go
```

Run benchmarks:

```bash
go test -bench=. intx_bench_test.go
```

## Performance

The library is designed for high performance with zero allocations for most operations:

- **Constructors**: ~0.25 ns/op
- **Byte conversion**: ~0.25-0.30 ns/op
- **JSON marshaling**: ~95-125 ns/op
- **String conversion**: ~17-22 ns/op

## Modular Import

You can import only the types you need by importing specific packages:

```go
// Import only 24-bit types
import . "intx/24"
u24 := MustUint24(123456)

// Import only 40-bit types  
import . "intx/40"
u40 := MustUint40(123456789012)

// Import only 48-bit types
import . "intx/48"
u48 := MustUint48(123456789012345)

// Import only 56-bit types
import . "intx/56"
u56 := MustUint56(12345678901234567)

// Import multiple types
import (
    . "intx/24"
    . "intx/40"
)
```

### Package Structure

```
go-intx/
├── 24/main.go          # Int24, Uint24 types
├── 40/main.go          # Int40, Uint40 types
├── 48/main.go          # Int48, Uint48 types
├── 56/main.go          # Int56, Uint56 types
├── intx_test.go      # Comprehensive tests
├── intx_bench_test.go # Performance benchmarks
├── example/example.go # Usage examples
└── go.mod            # Module definition
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Submit a pull request

## License

BSD 3-Clause License

Copyright (c) 2024, go-intx contributors

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE. 