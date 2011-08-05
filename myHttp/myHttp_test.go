package myHttp
import (
	"testing"
	"fmt"
)

func TestHttp(t *testing.T){
}

func TestRot13Mod(t *testing.T){
	translate := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	expected := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	actual := ""

	temp := []byte(translate)
	Rot13Mod(temp)
	actual = string(temp)

	fmt.Printf("Translating: %s\n", translate)
	fmt.Printf("Expecting: %s\n", expected)
	fmt.Printf("Actual: %s\n", actual)
	if actual == expected {
		fmt.Printf("The Mod Version Passed\n")
	}
}

func TestRot13Table(t *testing.T){
	translate := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	expected := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	actual := ""

	temp := []byte(translate)
	Rot13Table(temp)
	actual = string(temp)
	
	fmt.Printf("Translating: %s\n", translate)
	fmt.Printf("Expecting: %s\n", expected)
	fmt.Printf("Actual: %s\n", actual)
	if actual == expected {
		fmt.Printf("The Table Version Passed\n")
	}

}

func BenchmarkRot13Mod(b *testing.B){
	translate := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	temp := []byte(translate)
	for i:=0; i < b.N; i++ {
		Rot13Mod(temp)
	}
}

func BenchmarkRot13Table(b *testing.B){
	translate := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	temp := []byte(translate)
	for i:=0; i < b.N; i++ {
		Rot13Table(temp)
	}
}
