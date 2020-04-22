package main

import ("fmt"
        "sort"
)

func main() {
  s1:=make([]int,5,10)
  
  fmt.Println(s1,len(s1),cap(s1))
  s1[0]=12
  fmt.Println(s1)
  s2:=[]string{"北京","上海","深圳"}
  fmt.Println(s2)
  //s2[3]="广州"错误的写法，索引越界
  //调用append函数必须用原来的切片接收返回值
  s2=append(s2, "广州")//append追加元素时，原来的底层数组放不下时，Go 语言会把底层数组换一个
  fmt.Println(s2)
  fmt.Println(len(s2),cap(s2))
  ss:=[]string{"杭州","苏州","天津"}
  s2=append(s2, ss...)//...表示将切片拆开
  fmt.Println(s2)
  a1:=[]int{1,3,5}
  a2:=a1//赋值
  a3:=make([]int,3,3)
  copy(a3, a1)
  fmt.Println(a2)
  fmt.Println(a3)
  a1[0]=12
  fmt.Println(a2,a3)
  a1=append(a1[:1],a1[2:]...)
  fmt.Println(a1)
  fmt.Println(cap(a1))
  fmt.Println("Hello World")
  x1:=[...]int{1,3,5}
  s3:=x1[:]
  fmt.Println(s3,len(s3),cap(s3))
  s3=append(s3[:1],s3[2:]...)
  fmt.Println(s3,len(s3),cap(s3))
  fmt.Println(x1)
   a:=make([]int,5 ,10)

  for i:=0;i<10;i++{
    a=append(a,i)
  }
fmt.Println(a,cap(a))
var aa=[...]int{1,4,7,9,3,5,8,6}
sort.Ints(aa[:])
fmt.Println(aa)
}

