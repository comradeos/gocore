package main 

import "fmt"

func main(){
    user := User { Name: "Iaroslav" }
    print_info(user)
    print_info_p(&user)
}


type User struct {
    Name string
}


func print_info(u User){
    fmt.Printf("User %s\n", u.Name)
}

func print_info_p(u * User){
    fmt.Printf("User %s\n", (*u).Name)
}

