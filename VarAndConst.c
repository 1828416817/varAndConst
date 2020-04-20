package main

import "fmt"
//变量声明
var(
name string
age  int
isOk bool
)
//常量声明
const pi=12
const(
  n1=iota//0
  n2//1
  n3//2
)
const(
c1=iota//0
c2=100//1
c3=iota//2
c4//3

)
//第一次遇到const，iota会初始化为0，,每新增一行const声明iota会加1；
const(

  d1,d2=iota+1,iota+2//iota=0;这一行iota都为0
  d3,d4=iota+1,iota+2//iota=1;
  )
  //定义数量级
  const(
      _=iota
      KB=1<<(10*iota)
      MB=1<<(10*iota)
      GB=1<<(10*iota)
      TB=1<<(10*iota)
      PB=1<<(10*iota)


  )

func main()  {

	name="林冲"
	age=12
	isOk=true
	fmt.Println(isOk)//打印并换行
	fmt.Println(age)
	fmt.Printf("name: %s\n",name)//格式化打印
  var s1 string ="abc"
  fmt.Printf("%s\n",s1)
	a:=2
	b:=1//简短声明
  s2:="lvy"
  fmt.Println(s2)
  var c int
  c=a*pi
  var d int
  d=n3*2
  fmt.Println(d)
 fmt.Println(c)
	fmt.Println(a,b)//在终端打印
  fmt.Println(c1)
  fmt.Println(c2)
  fmt.Println(c3)
  fmt.Println(c4)
  fmt.Println(d1)
  fmt.Println(d2)
  fmt.Println(d3)
  fmt.Println(d4)
  fmt.Println(KB)


}


