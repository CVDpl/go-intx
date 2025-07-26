// Package int56 provides 56-bit fixed-width integer types for Go.
// This package implements signed and unsigned 56-bit integer types (Int56, Uint56).
package int56

import (
	"errors"
	"strconv"
)

// Common errors for the int56 package
var (
	ErrInt56OutOfRange        = errors.New("value exceeds range for Int56")
	ErrUint56OutOfRange       = errors.New("value exceeds maximum for Uint56")
	ErrInt56InvalidByteLength = errors.New("invalid byte length")
	ErrInt56EmptyData         = errors.New("empty data")
)

// Int56 represents a 56-bit signed integer stored in a 64-bit field.
type Int56 struct {
	value int64 // underlying value, only lower 56 bits are used
}

// NewInt56 creates a new Int56 from an int64 value.
// Returns an error if the value is out of range (-36028797018963968 to 36028797018963967).
func NewInt56(val int64) (Int56, error) {
	if val < -0x80000000000000 || val > 0x7FFFFFFFFFFFFF {
		return Int56{}, ErrInt56OutOfRange
	}
	return Int56{value: val}, nil
}

// MustInt56 creates a new Int56 from an int64 value.
// Panics if the value is out of range.
func MustInt56(val int64) Int56 {
	if val < -0x80000000000000 || val > 0x7FFFFFFFFFFFFF {
		panic("value exceeds range for Int56")
	}
	return Int56{value: val}
}

// Int64 returns the Int56 as an int64.
func (i Int56) Int64() int64 { return i.value }

// ToBytes returns a 7-byte big-endian representation of the Int56.
func (i Int56) ToBytes() [7]byte {
	var out [7]byte
	val := uint64(i.value)
	out[0] = byte(val >> 48)
	out[1] = byte(val >> 40)
	out[2] = byte(val >> 32)
	out[3] = byte(val >> 24)
	out[4] = byte(val >> 16)
	out[5] = byte(val >> 8)
	out[6] = byte(val)
	return out
}

// ToLittleEndianBytes returns a 7-byte little-endian representation of the Int56.
func (i Int56) ToLittleEndianBytes() [7]byte {
	var out [7]byte
	val := uint64(i.value)
	out[0] = byte(val)
	out[1] = byte(val >> 8)
	out[2] = byte(val >> 16)
	out[3] = byte(val >> 24)
	out[4] = byte(val >> 32)
	out[5] = byte(val >> 40)
	out[6] = byte(val >> 48)
	return out
}

// FromInt56Bytes creates an Int56 from a 7-byte big-endian slice.
func FromInt56Bytes(b []byte) (Int56, error) {
	if len(b) != 7 {
		return Int56{}, ErrInt56InvalidByteLength
	}
	val := int64(b[0])<<48 | int64(b[1])<<40 | int64(b[2])<<32 | int64(b[3])<<24 | int64(b[4])<<16 | int64(b[5])<<8 | int64(b[6])
	if b[0]&0x80 != 0 {
		val |= ^0x7FFFFFFFFFFFFF
	}
	return Int56{value: val}, nil
}

// FromInt56LittleEndianBytes creates an Int56 from a 7-byte little-endian slice.
func FromInt56LittleEndianBytes(b []byte) (Int56, error) {
	if len(b) != 7 {
		return Int56{}, ErrInt56InvalidByteLength
	}
	val := int64(b[6])<<48 | int64(b[5])<<40 | int64(b[4])<<32 | int64(b[3])<<24 | int64(b[2])<<16 | int64(b[1])<<8 | int64(b[0])
	if b[6]&0x80 != 0 {
		val |= ^0x7FFFFFFFFFFFFF
	}
	return Int56{value: val}, nil
}

// String returns the string representation of the Int56.
func (i Int56) String() string {
	// Optimized: use AppendInt to avoid intermediate allocations
	return string(strconv.AppendInt(nil, i.value, 10))
}

// MarshalJSON implements json.Marshaler for Int56.
func (i Int56) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendInt to avoid allocations
	return strconv.AppendInt(nil, i.value, 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Int56.
func (i *Int56) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt56EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	newI, err := NewInt56(val)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Int56.
func (i Int56) MarshalBinary() ([]byte, error) {
	bytes := i.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Int56.
func (i *Int56) UnmarshalBinary(data []byte) error {
	newI, err := FromInt56Bytes(data)
	if err != nil {
		return err
	}
	*i = newI
	return nil
}

// Uint56 represents a 56-bit unsigned integer stored in a 64-bit field.
type Uint56 struct {
	value uint64 // underlying value, only lower 56 bits are used
}

// NewUint56 creates a new Uint56 from a uint64 value.
// Returns an error if the value is out of range (0 to 72057594037927935).
func NewUint56(val uint64) (Uint56, error) {
	if val > 0xFFFFFFFFFFFFFF {
		return Uint56{}, ErrUint56OutOfRange
	}
	return Uint56{value: val}, nil
}

// MustUint56 creates a new Uint56 from a uint64 value.
// Panics if the value is out of range.
func MustUint56(val uint64) Uint56 {
	if val > 0xFFFFFFFFFFFFFF {
		panic("value exceeds maximum for Uint56")
	}
	return Uint56{value: val}
}

// Uint64 returns the Uint56 as a uint64.
func (u Uint56) Uint64() uint64 { return u.value }

// ToBytes returns a 7-byte big-endian representation of the Uint56.
func (u Uint56) ToBytes() [7]byte {
	var out [7]byte
	out[0] = byte(u.value >> 48)
	out[1] = byte(u.value >> 40)
	out[2] = byte(u.value >> 32)
	out[3] = byte(u.value >> 24)
	out[4] = byte(u.value >> 16)
	out[5] = byte(u.value >> 8)
	out[6] = byte(u.value)
	return out
}

// ToLittleEndianBytes returns a 7-byte little-endian representation of the Uint56.
func (u Uint56) ToLittleEndianBytes() [7]byte {
	var out [7]byte
	out[0] = byte(u.value)
	out[1] = byte(u.value >> 8)
	out[2] = byte(u.value >> 16)
	out[3] = byte(u.value >> 24)
	out[4] = byte(u.value >> 32)
	out[5] = byte(u.value >> 40)
	out[6] = byte(u.value >> 48)
	return out
}

// FromUint56Bytes creates a Uint56 from a 7-byte big-endian slice.
func FromUint56Bytes(b []byte) (Uint56, error) {
	if len(b) != 7 {
		return Uint56{}, ErrInt56InvalidByteLength
	}
	value := uint64(b[0])<<48 | uint64(b[1])<<40 | uint64(b[2])<<32 | uint64(b[3])<<24 | uint64(b[4])<<16 | uint64(b[5])<<8 | uint64(b[6])
	return Uint56{value: value}, nil
}

// FromUint56LittleEndianBytes creates a Uint56 from a 7-byte little-endian slice.
func FromUint56LittleEndianBytes(b []byte) (Uint56, error) {
	if len(b) != 7 {
		return Uint56{}, ErrInt56InvalidByteLength
	}
	value := uint64(b[6])<<48 | uint64(b[5])<<40 | uint64(b[4])<<32 | uint64(b[3])<<24 | uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
	return Uint56{value: value}, nil
}

// String returns the string representation of the Uint56.
func (u Uint56) String() string {
	// Optimized: use AppendUint to avoid intermediate allocations
	return string(strconv.AppendUint(nil, u.value, 10))
}

// MarshalJSON implements json.Marshaler for Uint56.
func (u Uint56) MarshalJSON() ([]byte, error) {
	// Optimized: use AppendUint to avoid allocations
	return strconv.AppendUint(nil, u.value, 10), nil
}

// UnmarshalJSON implements json.Unmarshaler for Uint56.
func (u *Uint56) UnmarshalJSON(data []byte) error {
	// Optimized: direct string parsing instead of json.Unmarshal
	if len(data) == 0 {
		return ErrInt56EmptyData
	}

	// Remove quotes if present
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	val, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}

	newU, err := NewUint56(val)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler for Uint56.
func (u Uint56) MarshalBinary() ([]byte, error) {
	bytes := u.ToBytes()
	return bytes[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Uint56.
func (u *Uint56) UnmarshalBinary(data []byte) error {
	newU, err := FromUint56Bytes(data)
	if err != nil {
		return err
	}
	*u = newU
	return nil
}
