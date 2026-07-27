package main 

import(
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main(){
    fmt.Printf("Please rate our pizza: ")

    reader := bufio.NewReader(os.Stdin);

    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    numRating, err := strconv.ParseInt(input, 10, 64);

    fmt.Println(numRating)
    fmt.Println(err)

    

}   
