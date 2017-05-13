package main

import(
     "fmt"
     "time"
)

func main() {
    //var name ="Perapat"
   
    fmt.Println("Hello world!!!",validate("Hello",5))
    fmt.Println("Time",time.Now())

}
func validate(input string,number int)int{
    if input=="Hello"{
        return 4 * number
    }
    return 0 * number
}

 