// Package int48 provides 48-bit fixed-width integer types for Go.
// This package implements signed and unsigned 48-bit integer types (Int48, Uint48).
package int48

import (
	"errors"
	"strconv"
)

// Common errors for the int48 package
var (
	ErrInt48OutOfRange        = errors.New("value exceeds range for Int48")
	ErrUint48OutOfRange       = errors.New("value exceeds maximum for Uint48")
	ErrInt48InvalidByteLength = errors.New("invalid byte length")
	ErrInt48EmptyData         = errors.New("empty data")
)

// Int48 represents a 48-bit signed integer stored in a 64-bit field.
type Int48 struct {
	value int64 // underlying value, only lower 48 bits are used
}

// NewInt48 creates a new Int48 from an int64 value.
// Returns an error if the value is out of range (-140737488355328 to 140737488355327).
func NewInt48(val int64) (Int48, error) {
	if val < -0x800000000000 || val > 0x7FFFFFFFFFFF {
		return Int48{}, ErrInt48OutOfRange
	}
	return Int48{value: val}, nil
}

// MustInt48 creates a new Int48 from an int64 value.
// Panics if the value is out of range.
func MustInt48(val int64) Int48 {
	if val < -0x800000000000 || val > 0x7FFFFFFFFFFF {
		panic("value exceeds range for Int48")
	}
	return Int48{value: val}
}

// Int64 returns the Int48 as an int64.
func (i Int48) Int64() int64 { return i.value }

// ToBytes returns a 6-byte big-endian representation of the Int48.
func (i Int48) ToBytes() [6]byte {
	var out [6]byte
	val := uint64(i.value)
	out[0] = byte(val >> 40)
	out[1] = byte(val >> 32)
	out[2] = byte(val >> 24)
	out[3] = byte(val >> 16)
	out[4] = byte(val >> 8)
	out[5] = byte(val)
	return out
}

// ToLittleEndianBytes returns a 6-byte little-endian representation of the Int48.
func (i Int48) ToLittleEndianBytes() [6]byte {
	var out [6]byte
	val := uint64(i.value)
	out[0] = byte(val)
	out[1] = byte(val >> 8)
	out[2] = byte(val >> 16)
	out[3] = byte(val >> 24)
	out[4] = byte(val >> 32)
	out[5] = byte(val >> 40)
	return out
}

// FromInt48Bytes creates an Int48 from a 6-byte big-endian slice.
func FromInt48Bytes(b []byte) (Int48, error) {
	if len(b) != 6 {
		return Int48{}, ErrInt48InvalidByteLength
	}
	val := int64(b[0])<<40 | int64(b[1])<<32 | int64(b[2])<<24 | int64(b[3])<<16 | int64(b[4])<<8 | int64(b[5])
	if b[0]&0x80 != 0 {
		val |= ^0x7FFFFFFFFFFF
	}
	return Int48{value: val}, nil
}

// FromInt48LittleEndianBytes creates an Int48 from a 6-byte little-endian slice.
func FromInt48LittleEndianBytes(b []byte) (Int48, error) {
	if len(b) != 6 {
		return Int48{}, ErrInt48InvalidByteLength
	}
	val := int64(b[5])<<40 | int64(b[4])<<32 | int64(b[3])<<24 | int64(b[2])<<16 | int64(b[1])<<8 | int64(b[0])
	if b[5]&0x80 != 0 {
		val |= ^0x7FFFFFFFFFFF
	}
	return Int48{value: val}, nil
}

// String returns the string representation of the Int48.
func (i Int48) String() string {
	// Optimized: use AppendInt to avoid intermediate allocations
	return string(strconv.AppendInt(nil, i.value, 10))
}

// MarshalJSON implements json.Marshaler for Int48.
func (i Int48) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendInt to avoid allocations
	return strconv.AppendInt(nil, i.value, 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Int48.
func (i *Int48) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt48EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	newI, err := NewInt48(val)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Int48.
func (i Int48) MarshalBinary() ([]byte, error) {
	bytes := i.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Int48.
func (i *Int48) UnmarshalBinary(data []byte) error {
	newI, err := FromInt48Bytes(data)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// Uint48 represents a 48-bit unsigned integer stored in a 64-bit field.
type Uint48 struct {
	value uint64 // underlying value, only lower 48 bits are used
}

// NewUint48 creates a new Uint48 from a uint64 value.
// Returns an error if the value is out of range (0 to 281474976710655).
func NewUint48(val uint64) (Uint48, error) {
	if val > 0xFFFFFFFFFFFF {
		return Uint48{}, ErrUint48OutOfRange
	}
	return Uint48{value: val}, nil
}

// MustUint48 creates a new Uint48 from a uint64 value.
// Panics if the value is out of range.
func MustUint48(val uint64) Uint48 {
	if val > 0xFFFFFFFFFFFF {
		panic("value exceeds maximum for Uint48")
	}
	return Uint48{value: val}
}

// Uint64 returns the Uint48 as a uint64.
func (u Uint48) Uint64() uint64 { return u.value }

// ToBytes returns a 6-byte big-endian representation of the Uint48.
func (u Uint48) ToBytes() [6]byte {
	var out [6]byte
	out[0] = byte(u.value >> 40)
	out[1] = byte(u.value >> 32)
	out[2] = byte(u.value >> 24)
	out[3] = byte(u.value >> 16)
	out[4] = byte(u.value >> 8)
	out[5] = byte(u.value)
	return out
}

// ToLittleEndianBytes returns a 6-byte little-endian representation of the Uint48.
func (u Uint48) ToLittleEndianBytes() [6]byte {
	var out [6]byte
	out[0] = byte(u.value)
	out[1] = byte(u.value >> 8)
	out[2] = byte(u.value >> 16)
	out[3] = byte(u.value >> 24)
	out[4] = byte(u.value >> 32)
	out[5] = byte(u.value >> 40)
	return out
}

// FromUint48Bytes creates a Uint48 from a 6-byte big-endian slice.
func FromUint48Bytes(b []byte) (Uint48, error) {
	if len(b) != 6 {
		return Uint48{}, ErrInt48InvalidByteLength
	}
	value := uint64(b[0])<<40 | uint64(b[1])<<32 | uint64(b[2])<<24 | uint64(b[3])<<16 | uint64(b[4])<<8 | uint64(b[5])
	return Uint48{value: value}, nil
}

// FromUint48LittleEndianBytes creates a Uint48 from a 6-byte little-endian slice.
func FromUint48LittleEndianBytes(b []byte) (Uint48, error) {
	if len(b) != 6 {
		return Uint48{}, ErrInt48InvalidByteLength
	}
	value := uint64(b[5])<<40 | uint64(b[4])<<32 | uint64(b[3])<<24 | uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
	return Uint48{value: value}, nil
}

// String returns the string representation of the Uint48.
func (u Uint48) String() string {
	// Optimized: use AppendUint to avoid intermediate allocations
	return string(strconv.AppendUint(nil, u.value, 10))
}

// MarshalJSON implements json.Marshaler for Uint48.
func (u Uint48) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendUint to avoid allocations
	return strconv.AppendUint(nil, u.value, 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Uint48.
func (u *Uint48) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt48EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}

	newU, err := NewUint48(val)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Uint48.
func (u Uint48) MarshalBinary() ([]byte, error) {
	bytes := u.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Uint48.
func (u *Uint48) UnmarshalBinary(data []byte) error {
	newU, err := FromUint48Bytes(data)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}
