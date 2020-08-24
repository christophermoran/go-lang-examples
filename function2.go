package main                                                                    
                                                                                
import "fmt"                                         

//Must declare function input type                           
//Must put string after () so it knows what to output				
func printName(name string)string{
	return "Hello," + name
}


func main () {                                                                                                         
	fmt.Println(printName("Jimmy"))
} 