package main

import "fmt"
//make和new都是用来申请内存的，
//new一般给基本数据类型申请内存，返回的时对应类型的指针
//make时用来给slice，map，channel，申请内存的，返回的时对应的3个类型本身
func main() {
  a1:=16
  p:=&a1
  fmt.Println(&a1)
  fmt.Printf("%p\n",p)
  fmt.Println(*p)
  fmt.Println("Hello World")
    var a0 *int
    fmt.Println(a0)
  var a2 =new(int)
  fmt.Println(a2)
  *a2=10
  fmt.Println(*a2)
}
