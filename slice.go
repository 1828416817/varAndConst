package main

import "fmt"

func main() {
  //1自定义切片
  var s1[]int
  var s2[]string
  fmt.Println(s1,s2)
  fmt.Println(s1==nil)
  fmt.Println(s2==nil)
  s1=[]int{1,2,3}
  s2=[]string{"A","B","C"}
  fmt.Println(s1==nil)
  fmt.Println(s2==nil)//nil相当于其他语言的NULL
  fmt.Println(len(s1),cap(s1))
  fmt.Println(len(s2),cap(s2))
  //2数组得到切片
  //切片的容量是指底层数组的容量
  //底层数组从切片的第一个元素到最后的元素数量
  a1:=[...]int{1,3,5,7,9,11,13}
  s3:=a1[0:4]
  fmt.Println(s3)
  s4:=a1[2:]
  fmt.Println(s4)
  s5:=a1[3:]
  fmt.Println(s5)
   s6:=a1[:]
  fmt.Println(s6)
  fmt.Println("Hello World")
  fmt.Println(len(s3),cap(s3))
  fmt.Println(len(s4),cap(s4))
  fmt.Println(len(s5),cap(s5))
  fmt.Println(len(s6),cap(s6))
  //切片再切割;切片是个引用值，都指向了一个底层数组
  s7:=s5[3:]
  fmt.Println(s7)
  fmt.Println(len(s7),cap(s7))
  a1[6]=1300
  fmt.Println(s5)
   fmt.Println(s7)
   //make函数创建切片
   

}
