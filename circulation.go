package main

import "fmt"

func main() {
  s1:="白萝卜"
  s2:=[]rune(s1)//将字符串转化趁一个rune类型的切片，类似列表
  s2[0]='红'
  fmt.Println(string(s2))//将rune类型的切片强制转化成字符串
  c1:="a"
  c2:='a'
  fmt.Printf("%T %T\n",c1,c2)
  fmt.Println(c2)
  fmt.Println("Hello World")
  if age:=19;age>18{//if语句的特殊表达else必须紧靠}后,age只在if语句中有效，是个局部变量
    fmt.Println("Love")
  }else {
    fmt.Println("回家写作业")
  }
  //for循环
  for i:=0;i<10;i++{
   
    fmt.Println(i)
  }
  //变式1
  var i=5
  for ;i<10;i++{
    fmt.Println(i)
  }
  //变式2
  var j=5
  for j<10{
    fmt.Println(j)
    j++
  }
  //无限循环
  for{
    fmt.Println("123")
  }
  //for range 循环
  s:="Hello Lvy"
  for i,v:=range s{
    fmt.Printf("%d %c\n",i,v)
  }
}
