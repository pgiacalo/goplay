package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binary string
	fmt.Print("Enter Binary Number:")
	fmt.Scanln(&binary)
	fmt.Println(binary)
	fmt.Println(BinaryToInt(binary))
	hex, err := BinaryToHex(binary)
	fmt.Println(hex, err)
	fmt.Println(HexToBin(hex))

}

func BinaryToInt(binary string) (int64, error) {
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return output, err

	//    s := fmt.Sprintf("Output %x", output)
	//    return s, nil
}

func BinaryToHex(binary string) (string, error) {
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	s := fmt.Sprintf("%x", output)
	return s, nil
}

func HexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	fmt.Println(ui)
	if err != nil {
		return "", err
	}

	// %016b indicates base 2, zero padded, with 16 characters
	return fmt.Sprintf("%b", ui), nil
}
