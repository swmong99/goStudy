/*
구조체를 이용한 단일 링크드 리스트 구현

1. 링크드 리스트의 노드로 구조체를 사용할 것
2. 구조체는 다음 노드의 주소를 저장할 멤버와 데이터를 저장할 멤버로 나뉜다.
3. 주소를 저장할 변수는 포인터로 사용한다.
4. 링크드 리스트에 추가/삭제/검색 기능을 별도의 함수로 작성한다.

끝
*/
package main

import "fmt"

type linkedlist struct {
   next *int
   data string
}

func main() {
 var single_list *linkedlist= new(linkedlist)
 var temp_addr *int = new(int)

 single_list.next=temp_addr
 single_list.data="default"
 append(single_list,"11")
 append(single_list,"12")

 fmt.Println(single_list.data)
}

func append(src_list *linkedlist, aa string  ){
 var dst_list *linkedlist= new(linkedlist)
 var addr *int = new(int) 

 dst_list.next=addr
 dst_list.data=aa

 src_list.next=&dst_list

}
