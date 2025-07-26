package intx

import (
	"encoding/json"

	. "github.com/CVDpl/go-intx/24"
	. "github.com/CVDpl/go-intx/40"
	. "github.com/CVDpl/go-intx/48"
	. "github.com/CVDpl/go-intx/56"

	"testing"
)

// Benchmark constructors
func BenchmarkNewUint24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUint24(uint64(i % 0xFFFFFF))
	}
}

func BenchmarkNewInt24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInt24(int64(i % 0x7FFFFF))
	}
}

func BenchmarkNewUint40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUint40(uint64(i % 0xFFFFFFFFFF))
	}
}

func BenchmarkNewInt40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInt40(int64(i % 0x7FFFFFFFFF))
	}
}

func BenchmarkNewUint48(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUint48(uint64(i % 0xFFFFFFFFFFFF))
	}
}

func BenchmarkNewInt48(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInt48(int64(i % 0x7FFFFFFFFFFF))
	}
}

func BenchmarkNewUint56(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUint56(uint64(i % 0xFFFFFFFFFFFFFF))
	}
}

func BenchmarkNewInt56(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewInt56(int64(i % 0x7FFFFFFFFFFFFF))
	}
}

// Benchmark Must constructors
func BenchmarkMustUint24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustUint24(uint64(i % 0xFFFFFF))
	}
}

func BenchmarkMustInt24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustInt24(int64(i % 0x7FFFFF))
	}
}

func BenchmarkMustUint40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustUint40(uint64(i % 0xFFFFFFFFFF))
	}
}

func BenchmarkMustInt40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustInt40(int64(i % 0x7FFFFFFFFF))
	}
}

func BenchmarkMustUint48(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustUint48(uint64(i % 0xFFFFFFFFFFFF))
	}
}

func BenchmarkMustInt48(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustInt48(int64(i % 0x7FFFFFFFFFFF))
	}
}

func BenchmarkMustUint56(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustUint56(uint64(i % 0xFFFFFFFFFFFFFF))
	}
}

func BenchmarkMustInt56(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustInt56(int64(i % 0x7FFFFFFFFFFFFF))
	}
}

// Benchmark ToBytes operations
func BenchmarkUint24ToBytes(b *testing.B) {
	u := MustUint24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToBytes()
	}
}

func BenchmarkInt24ToBytes(b *testing.B) {
	intVar := MustInt24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToBytes()
	}
}

func BenchmarkUint40ToBytes(b *testing.B) {
	u := MustUint40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToBytes()
	}
}

func BenchmarkInt40ToBytes(b *testing.B) {
	intVar := MustInt40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToBytes()
	}
}

func BenchmarkUint48ToBytes(b *testing.B) {
	u := MustUint48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToBytes()
	}
}

func BenchmarkInt48ToBytes(b *testing.B) {
	intVar := MustInt48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToBytes()
	}
}

func BenchmarkUint56ToBytes(b *testing.B) {
	u := MustUint56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToBytes()
	}
}

func BenchmarkInt56ToBytes(b *testing.B) {
	intVar := MustInt56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToBytes()
	}
}

// Benchmark ToLittleEndianBytes operations
func BenchmarkUint24ToLittleEndianBytes(b *testing.B) {
	u := MustUint24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToLittleEndianBytes()
	}
}

func BenchmarkInt24ToLittleEndianBytes(b *testing.B) {
	intVar := MustInt24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToLittleEndianBytes()
	}
}

func BenchmarkUint40ToLittleEndianBytes(b *testing.B) {
	u := MustUint40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToLittleEndianBytes()
	}
}

func BenchmarkInt40ToLittleEndianBytes(b *testing.B) {
	intVar := MustInt40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToLittleEndianBytes()
	}
}

func BenchmarkUint48ToLittleEndianBytes(b *testing.B) {
	u := MustUint48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToLittleEndianBytes()
	}
}

func BenchmarkInt48ToLittleEndianBytes(b *testing.B) {
	intVar := MustInt48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToLittleEndianBytes()
	}
}

func BenchmarkUint56ToLittleEndianBytes(b *testing.B) {
	u := MustUint56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ToLittleEndianBytes()
	}
}

func BenchmarkInt56ToLittleEndianBytes(b *testing.B) {
	intVar := MustInt56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.ToLittleEndianBytes()
	}
}

// Benchmark FromBytes operations
func BenchmarkFromUint24Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint24Bytes(bytes)
	}
}

func BenchmarkFromInt24Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt24Bytes(bytes)
	}
}

func BenchmarkFromUint40Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint40Bytes(bytes)
	}
}

func BenchmarkFromInt40Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt40Bytes(bytes)
	}
}

func BenchmarkFromUint48Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint48Bytes(bytes)
	}
}

func BenchmarkFromInt48Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt48Bytes(bytes)
	}
}

func BenchmarkFromUint56Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint56Bytes(bytes)
	}
}

func BenchmarkFromInt56Bytes(b *testing.B) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt56Bytes(bytes)
	}
}

// Benchmark FromLittleEndianBytes operations
func BenchmarkFromUint24LittleEndianBytes(b *testing.B) {
	bytes := []byte{0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint24LittleEndianBytes(bytes)
	}
}

func BenchmarkFromInt24LittleEndianBytes(b *testing.B) {
	bytes := []byte{0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt24LittleEndianBytes(bytes)
	}
}

func BenchmarkFromUint40LittleEndianBytes(b *testing.B) {
	bytes := []byte{0x9A, 0x78, 0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint40LittleEndianBytes(bytes)
	}
}

func BenchmarkFromInt40LittleEndianBytes(b *testing.B) {
	bytes := []byte{0x9A, 0x78, 0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt40LittleEndianBytes(bytes)
	}
}

func BenchmarkFromUint48LittleEndianBytes(b *testing.B) {
	bytes := []byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint48LittleEndianBytes(bytes)
	}
}

func BenchmarkFromInt48LittleEndianBytes(b *testing.B) {
	bytes := []byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt48LittleEndianBytes(bytes)
	}
}

func BenchmarkFromUint56LittleEndianBytes(b *testing.B) {
	bytes := []byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromUint56LittleEndianBytes(bytes)
	}
}

func BenchmarkFromInt56LittleEndianBytes(b *testing.B) {
	bytes := []byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FromInt56LittleEndianBytes(bytes)
	}
}

// Benchmark JSON operations
func BenchmarkUint24JSONMarshal(b *testing.B) {
	u := MustUint24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

func BenchmarkInt24JSONMarshal(b *testing.B) {
	intVar := MustInt24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(intVar)
	}
}

func BenchmarkUint40JSONMarshal(b *testing.B) {
	u := MustUint40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

func BenchmarkInt40JSONMarshal(b *testing.B) {
	intVar := MustInt40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(intVar)
	}
}

func BenchmarkUint48JSONMarshal(b *testing.B) {
	u := MustUint48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

func BenchmarkInt48JSONMarshal(b *testing.B) {
	intVar := MustInt48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(intVar)
	}
}

func BenchmarkUint56JSONMarshal(b *testing.B) {
	u := MustUint56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

func BenchmarkInt56JSONMarshal(b *testing.B) {
	intVar := MustInt56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(intVar)
	}
}

// Benchmark Binary operations
func BenchmarkUint24BinaryMarshal(b *testing.B) {
	u := MustUint24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalBinary()
	}
}

func BenchmarkInt24BinaryMarshal(b *testing.B) {
	intVar := MustInt24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.MarshalBinary()
	}
}

func BenchmarkUint40BinaryMarshal(b *testing.B) {
	u := MustUint40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalBinary()
	}
}

func BenchmarkInt40BinaryMarshal(b *testing.B) {
	intVar := MustInt40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.MarshalBinary()
	}
}

func BenchmarkUint48BinaryMarshal(b *testing.B) {
	u := MustUint48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalBinary()
	}
}

func BenchmarkInt48BinaryMarshal(b *testing.B) {
	intVar := MustInt48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.MarshalBinary()
	}
}

func BenchmarkUint56BinaryMarshal(b *testing.B) {
	u := MustUint56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalBinary()
	}
}

func BenchmarkInt56BinaryMarshal(b *testing.B) {
	intVar := MustInt56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intVar.MarshalBinary()
	}
}

// Benchmark String operations
func BenchmarkUint24String(b *testing.B) {
	u := MustUint24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}

func BenchmarkInt24String(b *testing.B) {
	intVar := MustInt24(0x123456)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = intVar.String()
	}
}

func BenchmarkUint40String(b *testing.B) {
	u := MustUint40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}

func BenchmarkInt40String(b *testing.B) {
	intVar := MustInt40(0x123456789A)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = intVar.String()
	}
}

func BenchmarkUint48String(b *testing.B) {
	u := MustUint48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}

func BenchmarkInt48String(b *testing.B) {
	intVar := MustInt48(0x123456789ABC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = intVar.String()
	}
}

func BenchmarkUint56String(b *testing.B) {
	u := MustUint56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u.String()
	}
}

func BenchmarkInt56String(b *testing.B) {
	intVar := MustInt56(0x123456789ABCDE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = intVar.String()
	}
}
