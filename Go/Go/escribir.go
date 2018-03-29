package main

import(
  "fmt"
  "os"
)

func main(){
  fichero,err := os.Create("fichero3.txt")
  if err != nil {
    fmt.Println(err)
  }else{
    fmt.Fprintln(fichero,567)
    fichero.Close()
  }

}
