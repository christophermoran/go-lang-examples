package main                                                                    
                                                                                
import "fmt"                                                                    
                                                                                
func main () {                                                                  
		fmt.Println(5 == 5)
		fmt.Println(5 == 6)
		fmt.Println(5 != 6)
		fmt.Println(5 > 6)
		fmt.Println(5 < 6)
		fmt.Println(5 >= 6)
		fmt.Println(5 <= 6)
		fmt.Println("This is equal" =="This is equal" )
		fmt.Println("This is not equal" =="This is equal" )
		num := 6						   
		fmt.Println(num == 6)
		fmt.Println(num == 6 || num == 7) // || is used for OR
} 
