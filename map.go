package main

import "fmt"

func main() {
  var m1 map[string]int
  m1= make(map[string]int,10)
  m1["����"]=18
  m1["������"]=35
  fmt.Println(m1)
  fmt.Println("Hello World")
  fmt.Println(m1["����"])
  value,ok:=m1["������"]
  if !ok{
    fmt.Println("NO")
  }else{
    fmt.Println(value)
  }
  //map����
  for k,v:=range m1{
    fmt.Println(k,v)
  }
  for k:=range m1{//ֻ����key
    fmt.Println(k)
  }
  for _,v:= range m1{//ֻ����ֵ
    fmt.Println(v)
  }
  delete(m1,"������")
  fmt.Println(m1)
  delete(m1,"A")
}
