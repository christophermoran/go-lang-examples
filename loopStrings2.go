package main                                                                    
                                                                                
import "fmt"                                                                    
                                                                                
func main () { 
	// 3 operators
	// 1: start 
	// 2: number to reach
	// 3: increase number if statement 2 is not true
	
	names := []string{"Chris", "Ant","Trev"}
	for x := 0; x < len(names); x ++{
		if names[x] == "Ant"{
			fmt.Println("hi",names[x])
		}

	}

}
