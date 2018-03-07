//Casey Nold
// queue Implementation in go

package my_queue

import "fmt"


type Element struct {
   Value interface{}
   next, prev *Element
}

type Queue struct {
   head, tail *Element
}


func (q *Queue) Head() *Element {
   return q.head
}

func (q *Queue) Tail() *Element {
   return q.tail
}

func (e* Element) Next() *Element{
  return e.next
}

func (e* Element) Prev() *Element {
  return e.prev
}

func (q * Queue) Empty() *Queue {
   if q.head != nil{
      for e :=q.Head(); e.Next()!=nil; e=e.Next(){
          e.prev.next = e.next
          e.next.prev = e.prev
	  e.next = nil
          e.prev = nil
      return q
      }
   }
   return q
}

func (q *Queue) FrontOf() *Element{
   if q.head != nil{
      return q.head
   }else{
      q.head = nil}
   return q.head

}


func (q *Queue) Is_Empty() bool{
   if q.head == nil{
      return true
   }else{ return false}
}


func (q *Queue) Enqueue(v interface{}) *Queue {
   e := &Element{Value: v}

   if q.head == nil {
      q.head = e
   }else{
      q.tail.next = e
      e.prev = q.tail
   }
   q.tail = e
   return q
}

func (q *Queue) Dequeue() *Element {
   if q.head == nil{
      return nil
   }else{
      temp := q.head
      q.head = q.head.Next()
      temp.next.prev = nil
      temp.next = nil
      return temp
   }
}


func (q *Queue) Remove(v int) bool {

   if q.head == nil {
      fmt.Println("Empty list")
      return false
   }else {
      for e:=q.Head(); e!=nil; e=e.Next() {
         if e.Value == v {
            e.prev.next = e.next
	    e.next.prev = e.prev
	    e.next = nil
            e.prev = nil
	    return true
         }
      }
   }
   return false
}


