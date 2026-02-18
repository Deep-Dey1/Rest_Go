package helper

import (

//    "fmt"
    "github.com/fatih/color"

)

func Greet(name string) string {
    
    //message := or.Red("Hello,") + color.Cyan(name) + color.Red("\nWelcome to the GOland")
    message := color.RedString("Hello,") + color.CyanString(name) + color.RedString("\nWelcome to Goland")
    return message

    }
