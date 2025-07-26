// Package int24 provides 24-bit fixed-width integer types for Go.
// This package implements signed and unsigned 24-bit integer types (Int24, Uint24).
package int24

import (
	"errors"
	"strconv"
)

// Common errors for the int24 package
var (
	ErrInt24OutOfRange        = errors.New("value exceeds range for Int24")
	ErrUint24OutOfRange       = errors.New("value exceeds maximum for Uint24")
	ErrInt24InvalidByteLength = errors.New("invalid byte length")
	ErrInt24EmptyData         = errors.New("empty data")
)

// Int24 represents a 24-bit signed integer stored in a 32-bit field.
type Int24 struct {
	value int32 // underlying value, only lower 24 bits are used
}

// NewInt24 creates a new Int24 from an int64 value.
// Returns an error if the value is out of range (-8388608 to 8388607).
func NewInt24(val int64) (Int24, error) {
	if val < -0x800000 || val > 0x7FFFFF {
		return Int24{}, ErrInt24OutOfRange
	}
	return Int24{value: int32(val)}, nil
}

// MustInt24 creates a new Int24 from an int64 value.
// Panics if the value is out of range.
func MustInt24(val int64) Int24 {
	if val < -0x800000 || val > 0x7FFFFF {
		panic("value exceeds range for Int24")
	}
	return Int24{value: int32(val)}
}

// Int64 returns the Int24 as an int64.
func (i Int24) Int64() int64 { return int64(i.value) }

// ToBytes returns a 3-byte big-endian representation of the Int24.
func (i Int24) ToBytes() [3]byte {
	var out [3]byte
	val := uint32(i.value)
	out[0] = byte(val >> 16)
	out[1] = byte(val >> 8)
	out[2] = byte(val)
	return out
}

// ToLittleEndianBytes returns a 3-byte little-endian representation of the Int24.
func (i Int24) ToLittleEndianBytes() [3]byte {
	var out [3]byte
	val := uint32(i.value)
	out[0] = byte(val)
	out[1] = byte(val >> 8)
	out[2] = byte(val >> 16)
	return out
}

// FromInt24Bytes creates an Int24 from a 3-byte big-endian slice.
func FromInt24Bytes(b []byte) (Int24, error) {
	if len(b) != 3 {
		return Int24{}, ErrInt24InvalidByteLength
	}
	val := int32(b[0])<<16 | int32(b[1])<<8 | int32(b[2])
	if b[0]&0x80 != 0 {
		val |= ^0x7FFFFF
	}
	return Int24{value: val}, nil
}

// FromInt24LittleEndianBytes creates an Int24 from a 3-byte little-endian slice.
func FromInt24LittleEndianBytes(b []byte) (Int24, error) {
	if len(b) != 3 {
		return Int24{}, ErrInt24InvalidByteLength
	}
	val := int32(b[2])<<16 | int32(b[1])<<8 | int32(b[0])
	if b[2]&0x80 != 0 {
		val |= ^0x7FFFFF
	}
	return Int24{value: val}, nil
}

// String returns the string representation of the Int24.
func (i Int24) String() string { return string(strconv.AppendInt(nil, int64(i.value), 10)) }

// MarshalJSON implements json.Marshaler for Int24.
func (i Int24) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendInt to avoid allocations
	return strconv.AppendInt(nil, int64(i.value), 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Int24.
func (i *Int24) UnmarshalJSON(data []byte) error {
	// Optimized: direct parsing without string allocation
	if len(data) == 0 {
		return ErrInt24EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	newI, err := NewInt24(val)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Int24.
func (i Int24) MarshalBinary() ([]byte, error) {
	// Optimized: return slice without allocation
	bytes := i.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Int24.
func (i *Int24) UnmarshalBinary(data []byte) error {
	newI, err := FromInt24Bytes(data)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// Uint24 represents a 24-bit unsigned integer stored in a 32-bit field.
type Uint24 struct {
	value uint32 // underlying value, only lower 24 bits are used
}

// NewUint24 creates a new Uint24 from a uint64 value.
// Returns an error if the value is out of range (0 to 16777215).
func NewUint24(val uint64) (Uint24, error) {
	if val > 0xFFFFFF {
		return Uint24{}, ErrUint24OutOfRange
	}
	return Uint24{value: uint32(val)}, nil
}

// MustUint24 creates a new Uint24 from a uint64 value.
// Panics if the value is out of range.
func MustUint24(val uint64) Uint24 {
	if val > 0xFFFFFF {
		panic("value exceeds range for Uint24")
	}
	return Uint24{value: uint32(val)}
}

// Uint64 returns the Uint24 as a uint64.
func (u Uint24) Uint64() uint64 { return uint64(u.value) }

// ToBytes returns a 3-byte big-endian representation of the Uint24.
func (u Uint24) ToBytes() [3]byte {
	var out [3]byte
	out[0] = byte(u.value >> 16)
	out[1] = byte(u.value >> 8)
	out[2] = byte(u.value)
	return out
}

// ToLittleEndianBytes returns a 3-byte little-endian representation of the Uint24.
func (u Uint24) ToLittleEndianBytes() [3]byte {
	var out [3]byte
	out[0] = byte(u.value)
	out[1] = byte(u.value >> 8)
	out[2] = byte(u.value >> 16)
	return out
}

// FromUint24Bytes creates a Uint24 from a 3-byte big-endian slice.
func FromUint24Bytes(b []byte) (Uint24, error) {
	if len(b) != 3 {
		return Uint24{}, ErrInt24InvalidByteLength
	}
	value := uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
	return Uint24{value: value}, nil
}

// FromUint24LittleEndianBytes creates a Uint24 from a 3-byte little-endian slice.
func FromUint24LittleEndianBytes(b []byte) (Uint24, error) {
	if len(b) != 3 {
		return Uint24{}, ErrInt24InvalidByteLength
	}
	value := uint32(b[2])<<16 | uint32(b[1])<<8 | uint32(b[0])
	return Uint24{value: value}, nil
}

// String returns the string representation of the Uint24.
func (u Uint24) String() string { return string(strconv.AppendUint(nil, uint64(u.value), 10)) }

// MarshalJSON implements json.Marshaler for Uint24.
func (u Uint24) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendUint to avoid allocations
	return strconv.AppendUint(nil, uint64(u.value), 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Uint24.
func (u *Uint24) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt24EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}

	newU, err := NewUint24(val)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Uint24.
func (u Uint24) MarshalBinary() ([]byte, error) {
	// Optimized: return slice without allocation
	bytes := u.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Uint24.
func (u *Uint24) UnmarshalBinary(data []byte) error {
	newU, err := FromUint24Bytes(data)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}
