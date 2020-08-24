package main                                                                    
                                                                                
import "fmt"                                                                    
				//Must put string after () so it knows what to output
func printName()string{
	return "Hello, world!"
}


func main () {                                                                                                         
	fmt.Println(printName())
} 