package main 

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    welcome := "Welcome"
 
    fmt.Printf("%s\n", welcome);

    reader := bufio.NewReader(os.Stdin);

    fmt.Println("Enter some rating:")

    input, _ := reader.ReadString('\n');

    fmt.Println("Thank you for rating, ", input);

}





