// Package int40 provides 40-bit fixed-width integer types for Go.
// This package implements signed and unsigned 40-bit integer types (Int40, Uint40).
package int40

import (
	"errors"
	"strconv"
)

// Common errors for the int40 package
var (
	ErrInt40OutOfRange        = errors.New("value exceeds range for Int40")
	ErrUint40OutOfRange       = errors.New("value exceeds maximum for Uint40")
	ErrInt40InvalidByteLength = errors.New("invalid byte length")
	ErrInt40EmptyData         = errors.New("empty data")
)

// Int40 represents a 40-bit signed integer stored in a 64-bit field.
type Int40 struct {
	value int64 // underlying value, only lower 40 bits are used
}

// NewInt40 creates a new Int40 from an int64 value.
// Returns an error if the value is out of range (-549755813888 to 549755813887).
func NewInt40(val int64) (Int40, error) {
	if val < -0x8000000000 || val > 0x7FFFFFFFFF {
		return Int40{}, ErrInt40OutOfRange
	}
	return Int40{value: val}, nil
}

// MustInt40 creates a new Int40 from an int64 value.
// Panics if the value is out of range.
func MustInt40(val int64) Int40 {
	if val < -0x8000000000 || val > 0x7FFFFFFFFF {
		panic("value exceeds range for Int40")
	}
	return Int40{value: val}
}

// Int64 returns the Int40 as an int64.
func (i Int40) Int64() int64 { return i.value }

// ToBytes returns a 5-byte big-endian representation of the Int40.
func (i Int40) ToBytes() [5]byte {
	var out [5]byte
	val := uint64(i.value)
	out[0] = byte(val >> 32)
	out[1] = byte(val >> 24)
	out[2] = byte(val >> 16)
	out[3] = byte(val >> 8)
	out[4] = byte(val)
	return out
}

// ToLittleEndianBytes returns a 5-byte little-endian representation of the Int40.
func (i Int40) ToLittleEndianBytes() [5]byte {
	var out [5]byte
	val := uint64(i.value)
	out[0] = byte(val)
	out[1] = byte(val >> 8)
	out[2] = byte(val >> 16)
	out[3] = byte(val >> 24)
	out[4] = byte(val >> 32)
	return out
}

// FromInt40Bytes creates an Int40 from a 5-byte big-endian slice.
func FromInt40Bytes(b []byte) (Int40, error) {
	if len(b) != 5 {
		return Int40{}, ErrInt40InvalidByteLength
	}
	val := int64(b[0])<<32 | int64(b[1])<<24 | int64(b[2])<<16 | int64(b[3])<<8 | int64(b[4])
	if b[0]&0x80 != 0 {
		val |= ^0x7FFFFFFFFF
	}
	return Int40{value: val}, nil
}

// FromInt40LittleEndianBytes creates an Int40 from a 5-byte little-endian slice.
func FromInt40LittleEndianBytes(b []byte) (Int40, error) {
	if len(b) != 5 {
		return Int40{}, ErrInt40InvalidByteLength
	}
	val := int64(b[4])<<32 | int64(b[3])<<24 | int64(b[2])<<16 | int64(b[1])<<8 | int64(b[0])
	if b[4]&0x80 != 0 {
		val |= ^0x7FFFFFFFFF
	}
	return Int40{value: val}, nil
}

// String returns the string representation of the Int40.
func (i Int40) String() string {
	// Optimized: use AppendInt to avoid intermediate allocations
	return string(strconv.AppendInt(nil, i.value, 10))
}

// MarshalJSON implements json.Marshaler for Int40.
func (i Int40) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendInt to avoid allocations
	return strconv.AppendInt(nil, i.value, 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Int40.
func (i *Int40) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt40EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	newI, err := NewInt40(val)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Int40.
func (i Int40) MarshalBinary() ([]byte, error) {
	bytes := i.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Int40.
func (i *Int40) UnmarshalBinary(data []byte) error {
	newI, err := FromInt40Bytes(data)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// Uint40 represents a 40-bit unsigned integer stored in a 64-bit field.
type Uint40 struct {
	value uint64 // underlying value, only lower 40 bits are used
}

// NewUint40 creates a new Uint40 from a uint64 value.
// Returns an error if the value is out of range (0 to 1099511627775).
func NewUint40(val uint64) (Uint40, error) {
	if val > 0xFFFFFFFFFF {
		return Uint40{}, ErrUint40OutOfRange
	}
	return Uint40{value: val}, nil
}

// MustUint40 creates a new Uint40 from a uint64 value.
// Panics if the value is out of range.
func MustUint40(val uint64) Uint40 {
	if val > 0xFFFFFFFFFF {
		panic("value exceeds maximum for Uint40")
	}
	return Uint40{value: val}
}

// Uint64 returns the Uint40 as a uint64.
func (u Uint40) Uint64() uint64 { return u.value }

// ToBytes returns a 5-byte big-endian representation of the Uint40.
func (u Uint40) ToBytes() [5]byte {
	var out [5]byte
	out[0] = byte(u.value >> 32)
	out[1] = byte(u.value >> 24)
	out[2] = byte(u.value >> 16)
	out[3] = byte(u.value >> 8)
	out[4] = byte(u.value)
	return out
}

// ToLittleEndianBytes returns a 5-byte little-endian representation of the Uint40.
func (u Uint40) ToLittleEndianBytes() [5]byte {
	var out [5]byte
	out[0] = byte(u.value)
	out[1] = byte(u.value >> 8)
	out[2] = byte(u.value >> 16)
	out[3] = byte(u.value >> 24)
	out[4] = byte(u.value >> 32)
	return out
}

// FromUint40Bytes creates a Uint40 from a 5-byte big-endian slice.
func FromUint40Bytes(b []byte) (Uint40, error) {
	if len(b) != 5 {
		return Uint40{}, ErrInt40InvalidByteLength
	}
	value := uint64(b[0])<<32 | uint64(b[1])<<24 | uint64(b[2])<<16 | uint64(b[3])<<8 | uint64(b[4])
	return Uint40{value: value}, nil
}

// FromUint40LittleEndianBytes creates a Uint40 from a 5-byte little-endian slice.
func FromUint40LittleEndianBytes(b []byte) (Uint40, error) {
	if len(b) != 5 {
		return Uint40{}, ErrInt40InvalidByteLength
	}
	value := uint64(b[4])<<32 | uint64(b[3])<<24 | uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
	return Uint40{value: value}, nil
}

// String returns the string representation of the Uint40.
func (u Uint40) String() string {
	// Optimized: use AppendUint to avoid intermediate allocations
	return string(strconv.AppendUint(nil, u.value, 10))
}

// MarshalJSON implements json.Marshaler for Uint40.
func (u Uint40) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendUint to avoid allocations
	return strconv.AppendUint(nil, u.value, 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Uint40.
func (u *Uint40) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt40EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}

	newU, err := NewUint40(val)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Uint40.
func (u Uint40) MarshalBinary() ([]byte, error) {
	bytes := u.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Uint40.
func (u *Uint40) UnmarshalBinary(data []byte) error {
	newU, err := FromUint40Bytes(data)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}
