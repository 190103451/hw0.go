// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, а¤¦аҐЃа¤Ёа¤їа¤Їа¤ѕ!")

	fmt.Println("Fizzbuzz() test")
	fmt.Printf("Fizzbuzz(%v) = %v\n", 20, Fizzbuzz(20))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 27, Fizzbuzz(27))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 30, Fizzbuzz(30))

	fmt.Println("IsPrime() test")
	fmt.Printf("IsPrime(%v) = %v\n", 5, IsPrime(5))
	fmt.Printf("IsPrime(%v) = %v\n", 6, IsPrime(6))
	fmt.Printf("IsPrime(%v) = %v\n", 25, IsPrime(25))
	fmt.Printf("IsPrime(%v) = %v\n", 37, IsPrime(37))

	fmt.Println("IsPalindrome() test")
	fmt.Printf("IsPalindrome(%v) = %v\n", "5445", IsPalindrome("5445"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "5445", IsPalindrome("5445"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "5775", IsPalindrome("5775"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	var s string

	if n%3 == 0 {
		s += "Fizz"
	}

	if n%5 == 0 {
		s += "Buzz"
	}

	return s
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	if n <= 2 {
		return true
	}

	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
