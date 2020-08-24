package main                                                                    
                                                                                
import "fmt"                                                                    
                                                                                
func main () { 
	age := 22
	if age == 21{
		fmt.Println("You may enter")
	}else if age < 21{
		fmt.Println("No entry")
	}else{
		fmt.Println("You can defo enter")
	}
}