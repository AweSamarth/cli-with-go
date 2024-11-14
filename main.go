package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("hello, what do you want to do?")
    
    for scanner.Scan() {
        input := scanner.Text()
        
        switch input {
        case "greet":
            fmt.Println("wasssup")
        
        case "add":
            fmt.Println("enter two numbers")
            
            scanner.Scan()
            num1, _ := strconv.Atoi(scanner.Text())
            
            scanner.Scan()
            num2, _ := strconv.Atoi(scanner.Text())
            
            fmt.Printf("Result: %d\n", num1 + num2)
		
		case "subtract":
			fmt.Println("enter two numbers")
			scanner.Scan()
			num1, _ := strconv.Atoi(scanner.Text())

			scanner.Scan()
			num2, _:= strconv.Atoi(scanner.Text())

			fmt.Printf("Result: %d\n", num1 - num2)
            
        case "bye":
            fmt.Println("bye!")
            return
        }
        
        fmt.Println("hello, what do you want to do?")
    }
}