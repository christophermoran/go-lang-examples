package main                                                                    
                                                                                
import "fmt"                                         

//Must declare function input type                           
//Must put string after () so it knows what to output				
func printName(name , age string)string{
	return "Hello," + name + "my age is " + age
}


func main () {                                                                                                         
	fmt.Println(printName("Jimmy","31"))
} 