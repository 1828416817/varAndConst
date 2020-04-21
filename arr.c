
package main

import "fmt"

func main() {
  var a1 [3]bool
  //var a2 [4]bool
  //数组的长度是数据类型的一部分
  //fmt.Printf("a1:%T a2:%T",a1,a2)
  //初始化1
  a1=[3]bool{true,true,true}
  fmt.Println(a1)
  //2
  a10:=[...]int{1,2,3,4,5,6,7,8,9}
  fmt.Println(a10)
  //根据索引初始化
  a2:=[5]int{0:1,4:2}
  fmt.Println(a2)
  //数组遍历
  cities:=[...]string{"北京","上海","深圳"}
  for i:=0;i<len(cities);i++{
    fmt.Println(cities[i])
  }
  //for range 遍历
  for i,v:=range cities{
    fmt.Println(i,v)
  }
  //多维数组
  var a11 [3][2]int
  a11=[3][2]int{
    [2]int{1,2},
    [2]int{3,4},
    [2]int{5,6},
  }
  fmt.Println(a11)
  for _,v1:=range a11{
    fmt.Println(v1)
    for _,v2:=range v1{
      fmt.Println(v2)
    }
  }   
  b1:=[3]int{1,2,3}
  b2:=b1
  b2[0]=100
  fmt.Println(b2)
  fmt.Println(b1)
  arr:=[...]int{1,3,5,7,9}
  sum:=0
  for _,v:=range arr{
      sum+=v
  }
  fmt.Println(sum)
  fmt.Println("Hello World")

}
