package main

import "fmt"

// decimalToBinary converts a decimal number to a binary string.
func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		n = n / 2
	}
	return binary
}

func main() {
	var a, b int = 10, 200
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	binaryNumber := decimalToBinary(a)
	fmt.Printf("The binary representation of %d   is %s\n", a, binaryNumber)

	binaryNumber = decimalToBinary(b)
	fmt.Printf("The binary representation of %d   is %s\n", b, binaryNumber)

	fmt.Println("a & b: ", a&b)   // binary AND
	fmt.Println("a | b: ", a|b)   // binary OR
	fmt.Println("a ^ b: ", a^b)   //binary XOR
	fmt.Println("a << b:", a<<b)  // binary left shift
	fmt.Println("a >> b: ", a>>b) // binary right shift
}
