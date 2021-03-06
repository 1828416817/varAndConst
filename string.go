package main

import "fmt"
import "strings"

func main() {
  s:="Hello 流沙"
  c1:='h'
  c2:='1'
  c3:='沙'
  fmt.Println(s)
  fmt.Println(c1)
  fmt.Println(c2)
  fmt.Println(c3)
  fmt.Printf("%c\n", c1)
  fmt.Printf("%c\n", c2)
  fmt.Printf("%c\n", c3)

  //一个字节=8Bit（8个2进制位）
  //1个字符'A'=1个字节
  //1个utf8编码的汉字'沙'=一般占3个字节
  path:="\"E:\\gits\\sqlist\""
  fmt.Println(path)
  //多行字符串
  s1:=`
  世情薄
        人情恶
      雨送黄昏花易落
  
  `
  s2:=`"E:\gits\sqlist"`//反引号原样输出，不需要借助转义字符输出双引号
  fmt.Println(s1)
  fmt.Println(s2)
  s3:="hello"
  length:=len(s3)
  fmt.Println(length)
  name:="I love you"
  world:="lvy"
  //ss:=name+world
  //fmt.Println(ss)
  //字符串拼接
  ss1:=fmt.Sprintf("%s%s",name,world)
  fmt.Println(ss1)
  //字符串分割
  ret:=strings.Split(path,`\`)
  fmt.Println(ret)
  fmt.Println(strings.Contains(ss1,"lvy"))//包含
  fmt.Println(strings.HasPrefix(ss1,"lvy"))//前缀
  fmt.Println(strings.HasSuffix(ss1,"lvy"))//后缀
  s4:="abcdefc"
  fmt.Println(strings.Index(s4,"c"))
   fmt.Println(strings.LastIndex(s4,"c"))
  fmt.Println("Hello World")
  fmt.Println(strings.Join(ret,"+"))
}
