package main

import (

    "fmt"
    "github.com/fatih/color"
    "helper"

)

func main(){
    
    var name string
//    fmt.Print(color.Yellow("Hey Visitor!,\nTell your name : "))
    message := color.YellowString("Hey Visitor !,\nTell your name : ")
    fmt.Print(message)
    fmt.Scanf("%s",&name)
    fmt.Println(helper.Greet(name))

    }
