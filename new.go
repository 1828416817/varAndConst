package main

import "fmt"
//make��new�������������ڴ�ģ�
//newһ��������������������ڴ棬���ص�ʱ��Ӧ���͵�ָ��
//makeʱ������slice��map��channel�������ڴ�ģ����ص�ʱ��Ӧ��3�����ͱ���
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
