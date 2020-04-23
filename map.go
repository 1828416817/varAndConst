package main

import "fmt"

func main() {
  var m1 map[string]int
  m1= make(map[string]int,10)
  m1["理想"]=18
  m1["姬无命"]=35
  fmt.Println(m1)
  fmt.Println("Hello World")
  fmt.Println(m1["理想"])
  value,ok:=m1["姬无命"]
  if !ok{
    fmt.Println("NO")
  }else{
    fmt.Println(value)
  }
  //map遍历
  for k,v:=range m1{
    fmt.Println(k,v)
  }
  for k:=range m1{//只遍历key
    fmt.Println(k)
  }
  for _,v:= range m1{//只遍历值
    fmt.Println(v)
  }
  delete(m1,"姬无命")
  fmt.Println(m1)
  delete(m1,"A")
}
