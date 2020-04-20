
package main

import "fmt"

func main() {

  var a int =6
  fmt.Printf("%d\n", a)
  fmt.Printf("%b\n", a)
  var b int =077
  fmt.Printf("%o\n", b)
  var c int =0xff
  fmt.Printf("%x\n", c)
  fmt.Printf("%T\n", c)
  fmt.Println("Hello World")
  i:=int8(9)
  fmt.Println(i)
  fmt.Printf("%T\n", i)
  f1:=1.23456
  fmt.Printf("%T\n", f1)
  f2:=float32(1.23456)
  fmt.Printf("%T\n", f2)//float32类型的数据不可以直接赋值给float64类型的数据
  b1:=true
  fmt.Printf("%T\n", b1)
  var b2 bool//默认false
  fmt.Println(b2)
  fmt.Printf("%v", b2)//%v打印变量的值

}
