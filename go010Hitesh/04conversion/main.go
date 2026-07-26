package main 

import(
    "fmt"
    "bufio"
    "os"
)

func main(){
    fmt.Printf("Please rate our pizza: ")

    reader := bufio.NewReader(os.Stdin);

    input, _ := reader.ReadString('\n')

    numRating := int(input) + 1;

    println(numRating)

    

}   
