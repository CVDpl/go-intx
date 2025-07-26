package intx

import (
	"encoding/json"

	. "github.com/CVDpl/go-intx/24"
	. "github.com/CVDpl/go-intx/40"
	. "github.com/CVDpl/go-intx/48"
	. "github.com/CVDpl/go-intx/56"

	"testing"
)

func TestUint24(t *testing.T) {
	tests := []struct {
		name    string
		value   uint64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"max", 0xFFFFFF, false},
		{"overflow", 0x1000000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUint24(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUint24() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.value {
				t.Errorf("NewUint24() = %v, want %v", u.Uint64(), tt.value)
			}
		})
	}
}

func TestUint24Must(t *testing.T) {
	// Test valid value
	u := MustUint24(12345)
	if u.Uint64() != 12345 {
		t.Errorf("MustUint24() = %v, want %v", u.Uint64(), 12345)
	}

	// Test panic on overflow
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustUint24() should panic on overflow")
		}
	}()
	MustUint24(0x1000000)
}

func TestUint24ToBytes(t *testing.T) {
	u := MustUint24(0x123456)
	bytes := u.ToBytes()
	expected := [3]byte{0x12, 0x34, 0x56}
	if bytes != expected {
		t.Errorf("ToBytes() = %v, want %v", bytes, expected)
	}
}

func TestUint24ToLittleEndianBytes(t *testing.T) {
	u := MustUint24(0x123456)
	bytes := u.ToLittleEndianBytes()
	expected := [3]byte{0x56, 0x34, 0x12}
	if bytes != expected {
		t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, expected)
	}
}

func TestFromUint24Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0x12, 0x34, 0x56}, 0x123456, false},
		{"zero", []byte{0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF}, 0xFFFFFF, false},
		{"short", []byte{0x12, 0x34}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint24Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint24Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint24Bytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestFromUint24LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0x56, 0x34, 0x12}, 0x123456, false},
		{"zero", []byte{0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF}, 0xFFFFFF, false},
		{"short", []byte{0x56, 0x34}, 0, true},
		{"long", []byte{0x56, 0x34, 0x12, 0x78}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint24LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint24LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint24LittleEndianBytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestUint24String(t *testing.T) {
	u := MustUint24(12345)
	if u.String() != "12345" {
		t.Errorf("String() = %v, want %v", u.String(), "12345")
	}
}

func TestUint24JSON(t *testing.T) {
	u := MustUint24(12345)

	// Test MarshalJSON
	data, err := json.Marshal(u)
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
	}
	if string(data) != "12345" {
		t.Errorf("MarshalJSON() = %v, want %v", string(data), "12345")
	}

	// Test UnmarshalJSON
	var u2 Uint24
	err = json.Unmarshal(data, &u2)
	if err != nil {
		t.Errorf("UnmarshalJSON() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalJSON() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestUint24Binary(t *testing.T) {
	u := MustUint24(0x123456)

	// Test MarshalBinary
	data, err := u.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary() error = %v", err)
	}
	expected := []byte{0x12, 0x34, 0x56}
	if len(data) != len(expected) {
		t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(expected))
	}
	for i, b := range expected {
		if data[i] != b {
			t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
		}
	}

	// Test UnmarshalBinary
	var u2 Uint24
	err = u2.UnmarshalBinary(data)
	if err != nil {
		t.Errorf("UnmarshalBinary() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalBinary() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestInt24(t *testing.T) {
	tests := []struct {
		name    string
		value   int64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"negative", -1, false},
		{"max", 0x7FFFFF, false},
		{"min", -0x800000, false},
		{"overflow", 0x800000, true},
		{"underflow", -0x800001, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := NewInt24(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInt24() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.value {
				t.Errorf("NewInt24() = %v, want %v", i.Int64(), tt.value)
			}
		})
	}
}

func TestInt24Must(t *testing.T) {
	// Test valid value
	i := MustInt24(-12345)
	if i.Int64() != -12345 {
		t.Errorf("MustInt24() = %v, want %v", i.Int64(), -12345)
	}

	// Test panic on overflow
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustInt24() should panic on overflow")
		}
	}()
	MustInt24(0x800000)
}

func TestInt24ToBytes(t *testing.T) {
	i := MustInt24(0x123456)
	bytes := i.ToBytes()
	expected := [3]byte{0x12, 0x34, 0x56}
	if bytes != expected {
		t.Errorf("ToBytes() = %v, want %v", bytes, expected)
	}

	// Test negative value
	i = MustInt24(-0x123456)
	bytes = i.ToBytes()
	expected = [3]byte{0xED, 0xCB, 0xAA}
	if bytes != expected {
		t.Errorf("ToBytes() negative = %v, want %v", bytes, expected)
	}
}

func TestInt24ToLittleEndianBytes(t *testing.T) {
	i := MustInt24(0x123456)
	bytes := i.ToLittleEndianBytes()
	expected := [3]byte{0x56, 0x34, 0x12}
	if bytes != expected {
		t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, expected)
	}
}

func TestFromInt24Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"valid", []byte{0x12, 0x34, 0x56}, 0x123456, false},
		{"zero", []byte{0x00, 0x00, 0x00}, 0, false},
		{"negative", []byte{0xED, 0xCB, 0xAA}, -0x123456, false},
		{"short", []byte{0x12, 0x34}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt24Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt24Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt24Bytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestFromInt24LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"valid", []byte{0x56, 0x34, 0x12}, 0x123456, false},
		{"zero", []byte{0x00, 0x00, 0x00}, 0, false},
		{"negative", []byte{0xAA, 0xCB, 0xED}, -0x123456, false},
		{"short", []byte{0x56, 0x34}, 0, true},
		{"long", []byte{0x56, 0x34, 0x12, 0x78}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt24LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt24LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt24LittleEndianBytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestInt24String(t *testing.T) {
	i := MustInt24(12345)
	if i.String() != "12345" {
		t.Errorf("String() = %v, want %v", i.String(), "12345")
	}

	i = MustInt24(-12345)
	if i.String() != "-12345" {
		t.Errorf("String() negative = %v, want %v", i.String(), "-12345")
	}
}

func TestInt24JSON(t *testing.T) {
	i := MustInt24(-12345)

	// Test MarshalJSON
	data, err := json.Marshal(i)
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
	}
	if string(data) != "-12345" {
		t.Errorf("MarshalJSON() = %v, want %v", string(data), "-12345")
	}

	// Test UnmarshalJSON
	var i2 Int24
	err = json.Unmarshal(data, &i2)
	if err != nil {
		t.Errorf("UnmarshalJSON() error = %v", err)
	}
	if i2.Int64() != i.Int64() {
		t.Errorf("UnmarshalJSON() = %v, want %v", i2.Int64(), i.Int64())
	}
}

func TestInt24Binary(t *testing.T) {
	i := MustInt24(-0x123456)

	// Test MarshalBinary
	data, err := i.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary() error = %v", err)
	}
	expected := []byte{0xED, 0xCB, 0xAA}
	if len(data) != len(expected) {
		t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(expected))
	}
	for i, b := range expected {
		if data[i] != b {
			t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
		}
	}

	// Test UnmarshalBinary
	var i2 Int24
	err = i2.UnmarshalBinary(data)
	if err != nil {
		t.Errorf("UnmarshalBinary() error = %v", err)
	}
	if i2.Int64() != i.Int64() {
		t.Errorf("UnmarshalBinary() = %v, want %v", i2.Int64(), i.Int64())
	}
}

func TestUint40(t *testing.T) {
	tests := []struct {
		name    string
		value   uint64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"max", 0xFFFFFFFFFF, false},
		{"overflow", 0x10000000000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUint40(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUint40() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.value {
				t.Errorf("NewUint40() = %v, want %v", u.Uint64(), tt.value)
			}
		})
	}
}

func TestUint40Must(t *testing.T) {
	// Test valid value
	u := MustUint40(123456789012)
	if u.Uint64() != 123456789012 {
		t.Errorf("MustUint40() = %v, want %v", u.Uint64(), 123456789012)
	}

	// Test panic on overflow
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustUint40() should panic on overflow")
		}
	}()
	MustUint40(0x10000000000)
}

func TestUint40ToBytes(t *testing.T) {
	u := MustUint40(0x123456789A)
	bytes := u.ToBytes()
	expected := [5]byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	if bytes != expected {
		t.Errorf("ToBytes() = %v, want %v", bytes, expected)
	}
}

func TestUint40ToLittleEndianBytes(t *testing.T) {
	u := MustUint40(0x123456789A)
	bytes := u.ToLittleEndianBytes()
	expected := [5]byte{0x9A, 0x78, 0x56, 0x34, 0x12}
	if bytes != expected {
		t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, expected)
	}
}

func TestFromUint40Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0x12, 0x34, 0x56, 0x78, 0x9A}, 0x123456789A, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFFFF, false},
		{"short", []byte{0x12, 0x34, 0x56}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint40Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint40Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint40Bytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestFromUint40LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789A, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFFFF, false},
		{"short", []byte{0x9A, 0x78, 0x56}, 0, true},
		{"long", []byte{0x9A, 0x78, 0x56, 0x34, 0x12, 0xBC}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint40LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint40LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint40LittleEndianBytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestUint40String(t *testing.T) {
	u := MustUint40(123456789012)
	if u.String() != "123456789012" {
		t.Errorf("String() = %v, want %v", u.String(), "123456789012")
	}
}

func TestUint40JSON(t *testing.T) {
	u := MustUint40(123456789012)

	// Test MarshalJSON
	data, err := json.Marshal(u)
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
	}
	if string(data) != "123456789012" {
		t.Errorf("MarshalJSON() = %v, want %v", string(data), "123456789012")
	}

	// Test UnmarshalJSON
	var u2 Uint40
	err = json.Unmarshal(data, &u2)
	if err != nil {
		t.Errorf("UnmarshalJSON() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalJSON() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestUint40Binary(t *testing.T) {
	u := MustUint40(0x123456789A)

	// Test MarshalBinary
	data, err := u.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary() error = %v", err)
	}
	expected := []byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	if len(data) != len(expected) {
		t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(expected))
	}
	for i, b := range expected {
		if data[i] != b {
			t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
		}
	}

	// Test UnmarshalBinary
	var u2 Uint40
	err = u2.UnmarshalBinary(data)
	if err != nil {
		t.Errorf("UnmarshalBinary() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalBinary() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestInt40(t *testing.T) {
	tests := []struct {
		name    string
		value   int64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"negative_one", -1, false},
		{"max", 0x7FFFFFFFFF, false},
		{"min", -0x8000000000, false},
		{"overflow_positive", 0x8000000000, true},
		{"overflow_negative", -0x8000000001, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := NewInt40(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInt40() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.value {
				t.Errorf("NewInt40() = %v, want %v", i.Int64(), tt.value)
			}
		})
	}
}

func TestInt40Must(t *testing.T) {
	// Test valid value
	i := MustInt40(123456789012)
	if i.Int64() != 123456789012 {
		t.Errorf("MustInt40() = %v, want %v", i.Int64(), 123456789012)
	}

	// Test negative value
	i = MustInt40(-123456789012)
	if i.Int64() != -123456789012 {
		t.Errorf("MustInt40() = %v, want %v", i.Int64(), -123456789012)
	}

	// Test panic on overflow
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustInt40() should panic on overflow")
		}
	}()
	MustInt40(0x8000000000)
}

func TestInt40ToBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected [5]byte
	}{
		{"positive", 0x123456789A, [5]byte{0x12, 0x34, 0x56, 0x78, 0x9A}},
		{"negative", -0x123456789A, [5]byte{0xED, 0xCB, 0xA9, 0x87, 0x66}},
		{"zero", 0, [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}},
		{"max", 0x7FFFFFFFFF, [5]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF}},
		{"min", -0x8000000000, [5]byte{0x80, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt40(tt.value)
			bytes := i.ToBytes()
			if bytes != tt.expected {
				t.Errorf("ToBytes() = %v, want %v", bytes, tt.expected)
			}
		})
	}
}

func TestInt40ToLittleEndianBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected [5]byte
	}{
		{"positive", 0x123456789A, [5]byte{0x9A, 0x78, 0x56, 0x34, 0x12}},
		{"negative", -0x123456789A, [5]byte{0x66, 0x87, 0xA9, 0xCB, 0xED}},
		{"zero", 0, [5]byte{0x00, 0x00, 0x00, 0x00, 0x00}},
		{"max", 0x7FFFFFFFFF, [5]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x7F}},
		{"min", -0x8000000000, [5]byte{0x00, 0x00, 0x00, 0x00, 0x80}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt40(tt.value)
			bytes := i.ToLittleEndianBytes()
			if bytes != tt.expected {
				t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, tt.expected)
			}
		})
	}
}

func TestFromInt40Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"positive", []byte{0x12, 0x34, 0x56, 0x78, 0x9A}, 0x123456789A, false},
		{"negative", []byte{0xED, 0xCB, 0xA9, 0x87, 0x66}, -0x123456789A, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF}, 0x7FFFFFFFFF, false},
		{"min", []byte{0x80, 0x00, 0x00, 0x00, 0x00}, -0x8000000000, false},
		{"short", []byte{0x12, 0x34, 0x56}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt40Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt40Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt40Bytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestFromInt40LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"positive", []byte{0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789A, false},
		{"negative", []byte{0x66, 0x87, 0xA9, 0xCB, 0xED}, -0x123456789A, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, 0x7FFFFFFFFF, false},
		{"min", []byte{0x00, 0x00, 0x00, 0x00, 0x80}, -0x8000000000, false},
		{"short", []byte{0x9A, 0x78, 0x56}, 0, true},
		{"long", []byte{0x9A, 0x78, 0x56, 0x34, 0x12, 0xBC}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt40LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt40LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt40LittleEndianBytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestInt40String(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{"positive", 123456789012, "123456789012"},
		{"negative", -123456789012, "-123456789012"},
		{"zero", 0, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt40(tt.value)
			if i.String() != tt.expected {
				t.Errorf("String() = %v, want %v", i.String(), tt.expected)
			}
		})
	}
}

func TestInt40JSON(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{"positive", 123456789012, "123456789012"},
		{"negative", -123456789012, "-123456789012"},
		{"zero", 0, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt40(tt.value)
			data, err := json.Marshal(i)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(data), tt.expected)
			}
			var i2 Int40
			err = json.Unmarshal(data, &i2)
			if err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
			}
			if i2.Int64() != i.Int64() {
				t.Errorf("UnmarshalJSON() = %v, want %v", i2.Int64(), i.Int64())
			}
		})
	}
}

func TestInt40Binary(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected []byte
	}{
		{"positive", 0x123456789A, []byte{0x12, 0x34, 0x56, 0x78, 0x9A}},
		{"negative", -0x123456789A, []byte{0xED, 0xCB, 0xA9, 0x87, 0x66}},
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt40(tt.value)
			data, err := i.MarshalBinary()
			if err != nil {
				t.Errorf("MarshalBinary() error = %v", err)
			}
			if len(data) != len(tt.expected) {
				t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(tt.expected))
			}
			for i, b := range tt.expected {
				if data[i] != b {
					t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
				}
			}
			var i2 Int40
			err = i2.UnmarshalBinary(data)
			if err != nil {
				t.Errorf("UnmarshalBinary() error = %v", err)
			}
			if i2.Int64() != i.Int64() {
				t.Errorf("UnmarshalBinary() = %v, want %v", i2.Int64(), i.Int64())
			}
		})
	}
}

func TestUint48(t *testing.T) {
	tests := []struct {
		name    string
		value   uint64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"max", 0xFFFFFFFFFFFF, false},
		{"overflow", 0x1000000000000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUint48(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUint48() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.value {
				t.Errorf("NewUint48() = %v, want %v", u.Uint64(), tt.value)
			}
		})
	}
}

func TestUint48Must(t *testing.T) {
	u := MustUint48(123456789012345)
	if u.Uint64() != 123456789012345 {
		t.Errorf("MustUint48() = %v, want %v", u.Uint64(), 123456789012345)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustUint48() should panic on overflow")
		}
	}()
	MustUint48(0x1000000000000)
}

func TestUint48ToBytes(t *testing.T) {
	u := MustUint48(0x123456789ABC)
	bytes := u.ToBytes()
	expected := [6]byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	if bytes != expected {
		t.Errorf("ToBytes() = %v, want %v", bytes, expected)
	}
}

func TestUint48ToLittleEndianBytes(t *testing.T) {
	u := MustUint48(0x123456789ABC)
	bytes := u.ToLittleEndianBytes()
	expected := [6]byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}
	if bytes != expected {
		t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, expected)
	}
}

func TestFromUint48Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}, 0x123456789ABC, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFFFFFF, false},
		{"short", []byte{0x12, 0x34, 0x56}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint48Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint48Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint48Bytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestFromUint48LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789ABC, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFFFFFF, false},
		{"short", []byte{0xBC, 0x9A, 0x78}, 0, true},
		{"long", []byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12, 0xDE}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint48LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint48LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint48LittleEndianBytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestUint48String(t *testing.T) {
	u := MustUint48(123456789012345)
	if u.String() != "123456789012345" {
		t.Errorf("String() = %v, want %v", u.String(), "123456789012345")
	}
}

func TestUint48JSON(t *testing.T) {
	u := MustUint48(123456789012345)
	data, err := json.Marshal(u)
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
	}
	if string(data) != "123456789012345" {
		t.Errorf("MarshalJSON() = %v, want %v", string(data), "123456789012345")
	}
	var u2 Uint48
	err = json.Unmarshal(data, &u2)
	if err != nil {
		t.Errorf("UnmarshalJSON() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalJSON() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestUint48Binary(t *testing.T) {
	u := MustUint48(0x123456789ABC)
	data, err := u.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary() error = %v", err)
	}
	expected := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}
	if len(data) != len(expected) {
		t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(expected))
	}
	for i, b := range expected {
		if data[i] != b {
			t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
		}
	}
	var u2 Uint48
	err = u2.UnmarshalBinary(data)
	if err != nil {
		t.Errorf("UnmarshalBinary() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalBinary() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestInt48(t *testing.T) {
	tests := []struct {
		name    string
		value   int64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"negative_one", -1, false},
		{"max", 0x7FFFFFFFFFFF, false},
		{"min", -0x800000000000, false},
		{"overflow_positive", 0x800000000000, true},
		{"overflow_negative", -0x800000000001, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := NewInt48(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInt48() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.value {
				t.Errorf("NewInt48() = %v, want %v", i.Int64(), tt.value)
			}
		})
	}
}

func TestInt48Must(t *testing.T) {
	i := MustInt48(123456789012345)
	if i.Int64() != 123456789012345 {
		t.Errorf("MustInt48() = %v, want %v", i.Int64(), 123456789012345)
	}
	i = MustInt48(-123456789012345)
	if i.Int64() != -123456789012345 {
		t.Errorf("MustInt48() = %v, want %v", i.Int64(), -123456789012345)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustInt48() should panic on overflow")
		}
	}()
	MustInt48(0x800000000000)
}

func TestInt48ToBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected [6]byte
	}{
		{"positive", 0x123456789ABC, [6]byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}},
		{"negative", -0x123456789ABC, [6]byte{0xED, 0xCB, 0xA9, 0x87, 0x65, 0x44}},
		{"zero", 0, [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"max", 0x7FFFFFFFFFFF, [6]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		{"min", -0x800000000000, [6]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt48(tt.value)
			bytes := i.ToBytes()
			if bytes != tt.expected {
				t.Errorf("ToBytes() = %v, want %v", bytes, tt.expected)
			}
		})
	}
}

func TestInt48ToLittleEndianBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected [6]byte
	}{
		{"positive", 0x123456789ABC, [6]byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}},
		{"negative", -0x123456789ABC, [6]byte{0x44, 0x65, 0x87, 0xA9, 0xCB, 0xED}},
		{"zero", 0, [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"max", 0x7FFFFFFFFFFF, [6]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}},
		{"min", -0x800000000000, [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x80}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt48(tt.value)
			bytes := i.ToLittleEndianBytes()
			if bytes != tt.expected {
				t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, tt.expected)
			}
		})
	}
}

func TestFromInt48Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"positive", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}, 0x123456789ABC, false},
		{"negative", []byte{0xED, 0xCB, 0xA9, 0x87, 0x65, 0x44}, -0x123456789ABC, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0x7FFFFFFFFFFF, false},
		{"min", []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00}, -0x800000000000, false},
		{"short", []byte{0x12, 0x34, 0x56}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt48Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt48Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt48Bytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestFromInt48LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"positive", []byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789ABC, false},
		{"negative", []byte{0x44, 0x65, 0x87, 0xA9, 0xCB, 0xED}, -0x123456789ABC, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, 0x7FFFFFFFFFFF, false},
		{"min", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x80}, -0x800000000000, false},
		{"short", []byte{0xBC, 0x9A, 0x78}, 0, true},
		{"long", []byte{0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12, 0xDE}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt48LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt48LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt48LittleEndianBytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestInt48String(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{"positive", 123456789012345, "123456789012345"},
		{"negative", -123456789012345, "-123456789012345"},
		{"zero", 0, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt48(tt.value)
			if i.String() != tt.expected {
				t.Errorf("String() = %v, want %v", i.String(), tt.expected)
			}
		})
	}
}

func TestInt48JSON(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{"positive", 123456789012345, "123456789012345"},
		{"negative", -123456789012345, "-123456789012345"},
		{"zero", 0, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt48(tt.value)
			data, err := json.Marshal(i)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(data), tt.expected)
			}
			var i2 Int48
			err = json.Unmarshal(data, &i2)
			if err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
			}
			if i2.Int64() != i.Int64() {
				t.Errorf("UnmarshalJSON() = %v, want %v", i2.Int64(), i.Int64())
			}
		})
	}
}

func TestInt48Binary(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected []byte
	}{
		{"positive", 0x123456789ABC, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}},
		{"negative", -0x123456789ABC, []byte{0xED, 0xCB, 0xA9, 0x87, 0x65, 0x44}},
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt48(tt.value)
			data, err := i.MarshalBinary()
			if err != nil {
				t.Errorf("MarshalBinary() error = %v", err)
			}
			if len(data) != len(tt.expected) {
				t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(tt.expected))
			}
			for i, b := range tt.expected {
				if data[i] != b {
					t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
				}
			}
			var i2 Int48
			err = i2.UnmarshalBinary(data)
			if err != nil {
				t.Errorf("UnmarshalBinary() error = %v", err)
			}
			if i2.Int64() != i.Int64() {
				t.Errorf("UnmarshalBinary() = %v, want %v", i2.Int64(), i.Int64())
			}
		})
	}
}

func TestUint56(t *testing.T) {
	tests := []struct {
		name    string
		value   uint64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"max", 0xFFFFFFFFFFFFFF, false},
		{"overflow", 0x100000000000000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUint56(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUint56() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.value {
				t.Errorf("NewUint56() = %v, want %v", u.Uint64(), tt.value)
			}
		})
	}
}

func TestUint56Must(t *testing.T) {
	u := MustUint56(12345678901234567)
	if u.Uint64() != 12345678901234567 {
		t.Errorf("MustUint56() = %v, want %v", u.Uint64(), 12345678901234567)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustUint56() should panic on overflow")
		}
	}()
	MustUint56(0x100000000000000)
}

func TestUint56ToBytes(t *testing.T) {
	u := MustUint56(0x123456789ABCDE)
	bytes := u.ToBytes()
	expected := [7]byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}
	if bytes != expected {
		t.Errorf("ToBytes() = %v, want %v", bytes, expected)
	}
}

func TestUint56ToLittleEndianBytes(t *testing.T) {
	u := MustUint56(0x123456789ABCDE)
	bytes := u.ToLittleEndianBytes()
	expected := [7]byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}
	if bytes != expected {
		t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, expected)
	}
}

func TestFromUint56Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}, 0x123456789ABCDE, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFFFFFFFF, false},
		{"short", []byte{0x12, 0x34, 0x56}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint56Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint56Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint56Bytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestFromUint56LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    uint64
		wantErr bool
	}{
		{"valid", []byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789ABCDE, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFFFFFFFF, false},
		{"short", []byte{0xDE, 0xBC, 0x9A}, 0, true},
		{"long", []byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12, 0xF0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := FromUint56LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUint56LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && u.Uint64() != tt.want {
				t.Errorf("FromUint56LittleEndianBytes() = %v, want %v", u.Uint64(), tt.want)
			}
		})
	}
}

func TestUint56String(t *testing.T) {
	u := MustUint56(12345678901234567)
	if u.String() != "12345678901234567" {
		t.Errorf("String() = %v, want %v", u.String(), "12345678901234567")
	}
}

func TestUint56JSON(t *testing.T) {
	u := MustUint56(12345678901234567)
	data, err := json.Marshal(u)
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
	}
	if string(data) != "12345678901234567" {
		t.Errorf("MarshalJSON() = %v, want %v", string(data), "12345678901234567")
	}
	var u2 Uint56
	err = json.Unmarshal(data, &u2)
	if err != nil {
		t.Errorf("UnmarshalJSON() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalJSON() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestUint56Binary(t *testing.T) {
	u := MustUint56(0x123456789ABCDE)
	data, err := u.MarshalBinary()
	if err != nil {
		t.Errorf("MarshalBinary() error = %v", err)
	}
	expected := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}
	if len(data) != len(expected) {
		t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(expected))
	}
	for i, b := range expected {
		if data[i] != b {
			t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
		}
	}
	var u2 Uint56
	err = u2.UnmarshalBinary(data)
	if err != nil {
		t.Errorf("UnmarshalBinary() error = %v", err)
	}
	if u2.Uint64() != u.Uint64() {
		t.Errorf("UnmarshalBinary() = %v, want %v", u2.Uint64(), u.Uint64())
	}
}

func TestInt56(t *testing.T) {
	tests := []struct {
		name    string
		value   int64
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"negative_one", -1, false},
		{"max", 0x7FFFFFFFFFFFFF, false},
		{"min", -0x80000000000000, false},
		{"overflow_positive", 0x80000000000000, true},
		{"overflow_negative", -0x80000000000001, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := NewInt56(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInt56() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.value {
				t.Errorf("NewInt56() = %v, want %v", i.Int64(), tt.value)
			}
		})
	}
}

func TestInt56Must(t *testing.T) {
	i := MustInt56(12345678901234567)
	if i.Int64() != 12345678901234567 {
		t.Errorf("MustInt56() = %v, want %v", i.Int64(), 12345678901234567)
	}
	i = MustInt56(-12345678901234567)
	if i.Int64() != -12345678901234567 {
		t.Errorf("MustInt56() = %v, want %v", i.Int64(), -12345678901234567)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustInt56() should panic on overflow")
		}
	}()
	MustInt56(0x80000000000000)
}

func TestInt56ToBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected [7]byte
	}{
		{"positive", 0x123456789ABCDE, [7]byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}},
		{"negative", -0x123456789ABCDE, [7]byte{0xED, 0xCB, 0xA9, 0x87, 0x65, 0x43, 0x22}},
		{"zero", 0, [7]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"max", 0x7FFFFFFFFFFFFF, [7]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		{"min", -0x80000000000000, [7]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt56(tt.value)
			bytes := i.ToBytes()
			if bytes != tt.expected {
				t.Errorf("ToBytes() = %v, want %v", bytes, tt.expected)
			}
		})
	}
}

func TestInt56ToLittleEndianBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected [7]byte
	}{
		{"positive", 0x123456789ABCDE, [7]byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}},
		{"negative", -0x123456789ABCDE, [7]byte{0x22, 0x43, 0x65, 0x87, 0xA9, 0xCB, 0xED}},
		{"zero", 0, [7]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
		{"max", 0x7FFFFFFFFFFFFF, [7]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}},
		{"min", -0x80000000000000, [7]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt56(tt.value)
			bytes := i.ToLittleEndianBytes()
			if bytes != tt.expected {
				t.Errorf("ToLittleEndianBytes() = %v, want %v", bytes, tt.expected)
			}
		})
	}
}

func TestFromInt56Bytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"positive", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}, 0x123456789ABCDE, false},
		{"negative", []byte{0xED, 0xCB, 0xA9, 0x87, 0x65, 0x43, 0x22}, -0x123456789ABCDE, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, 0x7FFFFFFFFFFFFF, false},
		{"min", []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, -0x80000000000000, false},
		{"short", []byte{0x12, 0x34, 0x56}, 0, true},
		{"long", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt56Bytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt56Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt56Bytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestFromInt56LittleEndianBytes(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    int64
		wantErr bool
	}{
		{"positive", []byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789ABCDE, false},
		{"negative", []byte{0x22, 0x43, 0x65, 0x87, 0xA9, 0xCB, 0xED}, -0x123456789ABCDE, false},
		{"zero", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0, false},
		{"max", []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}, 0x7FFFFFFFFFFFFF, false},
		{"min", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}, -0x80000000000000, false},
		{"short", []byte{0xDE, 0xBC, 0x9A}, 0, true},
		{"long", []byte{0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12, 0xF0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, err := FromInt56LittleEndianBytes(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt56LittleEndianBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.Int64() != tt.want {
				t.Errorf("FromInt56LittleEndianBytes() = %v, want %v", i.Int64(), tt.want)
			}
		})
	}
}

func TestInt56String(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{"positive", 12345678901234567, "12345678901234567"},
		{"negative", -12345678901234567, "-12345678901234567"},
		{"zero", 0, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt56(tt.value)
			if i.String() != tt.expected {
				t.Errorf("String() = %v, want %v", i.String(), tt.expected)
			}
		})
	}
}

func TestInt56JSON(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected string
	}{
		{"positive", 12345678901234567, "12345678901234567"},
		{"negative", -12345678901234567, "-12345678901234567"},
		{"zero", 0, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt56(tt.value)
			data, err := json.Marshal(i)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(data), tt.expected)
			}
			var i2 Int56
			err = json.Unmarshal(data, &i2)
			if err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
			}
			if i2.Int64() != i.Int64() {
				t.Errorf("UnmarshalJSON() = %v, want %v", i2.Int64(), i.Int64())
			}
		})
	}
}

func TestInt56Binary(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		expected []byte
	}{
		{"positive", 0x123456789ABCDE, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}},
		{"negative", -0x123456789ABCDE, []byte{0xED, 0xCB, 0xA9, 0x87, 0x65, 0x43, 0x22}},
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MustInt56(tt.value)
			data, err := i.MarshalBinary()
			if err != nil {
				t.Errorf("MarshalBinary() error = %v", err)
			}
			if len(data) != len(tt.expected) {
				t.Errorf("MarshalBinary() length = %v, want %v", len(data), len(tt.expected))
			}
			for i, b := range tt.expected {
				if data[i] != b {
					t.Errorf("MarshalBinary()[%d] = %v, want %v", i, data[i], b)
				}
			}
			var i2 Int56
			err = i2.UnmarshalBinary(data)
			if err != nil {
				t.Errorf("UnmarshalBinary() error = %v", err)
			}
			if i2.Int64() != i.Int64() {
				t.Errorf("UnmarshalBinary() = %v, want %v", i2.Int64(), i.Int64())
			}
		})
	}
}
