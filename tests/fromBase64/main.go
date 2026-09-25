package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	data := "SwM/qmG/9D8=" // 1.296724

	decoded, err := base64.StdEncoding.DecodeString(data)
	
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}

	if len(decoded) != 8 {
		fmt.Println("Error: decoded data is not 8 bytes long")
		return
	}

	bits := binary.LittleEndian.Uint64(decoded)
	
	value := math.Float64frombits(bits)

	fmt.Println("Converted double:", value)
}
