package main

import (
	"encoding/json"
	"fmt"
	"log"

	// Import only 24-bit types with dot import to use without prefix
	. "github.com/CVDpl/go-intx/24"
	. "github.com/CVDpl/go-intx/40"
	. "github.com/CVDpl/go-intx/48"
	. "github.com/CVDpl/go-intx/56"
)

func main() {
	fmt.Println("=== IntX Library Examples ===")

	// Uint24 example
	fmt.Println("--- Uint24 ---")
	u24 := MustUint24(123456)
	fmt.Println("Uint24 value:", u24.Uint64())
	b24 := u24.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", b24)

	// Int24 example
	fmt.Println("\n--- Int24 ---")
	i24 := MustInt24(-123456)
	fmt.Println("Int24 value:", i24.Int64())
	bi24 := i24.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", bi24)

	// Uint40 example
	fmt.Println("\n--- Uint40 ---")
	u40 := MustUint40(123456789012)
	fmt.Println("Uint40 value:", u40.Uint64())
	b40 := u40.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", b40)

	// Int40 example
	fmt.Println("\n--- Int40 ---")
	i40 := MustInt40(-123456789012)
	fmt.Println("Int40 value:", i40.Int64())
	bi40 := i40.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", bi40)

	// Uint48 example
	fmt.Println("\n--- Uint48 ---")
	u48 := MustUint48(123456789012345)
	fmt.Println("Uint48 value:", u48.Uint64())
	b48 := u48.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", b48)

	// Int48 example
	fmt.Println("\n--- Int48 ---")
	i48 := MustInt48(-123456789012345)
	fmt.Println("Int48 value:", i48.Int64())
	bi48 := i48.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", bi48)

	// Uint56 example
	fmt.Println("\n--- Uint56 ---")
	u56 := MustUint56(12345678901234567)
	fmt.Println("Uint56 value:", u56.Uint64())
	b56 := u56.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", b56)

	// Int56 example
	fmt.Println("\n--- Int56 ---")
	i56 := MustInt56(-12345678901234567)
	fmt.Println("Int56 value:", i56.Int64())
	bi56 := i56.ToBytes()
	fmt.Printf("Big-endian bytes: %x\n", bi56)

	// JSON marshal/unmarshal example with Uint48
	fmt.Println("\n--- JSON Example (Uint48) ---")
	jsonData, err := json.Marshal(u48)
	if err != nil {
		log.Fatal("Marshal JSON error:", err)
	}
	fmt.Printf("JSON: %s\n", jsonData)

	var u48Recovered Uint48
	err = json.Unmarshal(jsonData, &u48Recovered)
	if err != nil {
		log.Fatal("Unmarshal JSON error:", err)
	}
	fmt.Println("Recovered from JSON:", u48Recovered.Uint64())

	// Error handling examples
	fmt.Println("\n--- Error Handling ---")

	// Uint24 overflow
	_, err = NewUint24(1 << 24)
	if err != nil {
		fmt.Println("Uint24 overflow error:", err)
	}

	// Int24 overflow
	_, err = NewInt24(1 << 23)
	if err != nil {
		fmt.Println("Int24 overflow error:", err)
	}

	// Invalid byte length
	_, err = FromUint48Bytes([]byte{1, 2, 3})
	if err != nil {
		fmt.Println("Invalid byte length error:", err)
	}
}
