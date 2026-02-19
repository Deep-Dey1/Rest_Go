package greeting

import (

    "errors" // to use the error package    
    "fmt"

)

func Hello(name string) (string, error) {
        if name == "" {

            return "", errors.New("empty name")

            }
        message := fmt.Sprintf("Hi " + name + ", Welcome to the interface")
        return message , nil
        

}
