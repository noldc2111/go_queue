//Casey Nold
// Queue Implementation in go

package main

import (
   "fmt"
   "mygocode/my_queue"
)

/*************| MAIN |*****************/

func main() {

   q := new(my_queue.Queue)

   if q.Is_Empty(){ fmt.Println("Empty Queue")}

   q.Enqueue(4)
   q.Enqueue("Gopher")
   q.Enqueue(true)
   q.Enqueue(7)


   for n:= q.Head(); n!=nil; n=n.Next() {
      fmt.Printf("%v\n", n.Value)
   }

   x:=q.FrontOf()
   //fmt.Println(q.FrontOf())
   fmt.Printf("Front of queue: %v\n",x.Value)
   q.Remove(5)
   pVal := q.Dequeue()

   fmt.Printf("Dequeued: %v\n",pVal.Value)

   for n:= q.Head(); n!=nil; n=n.Next() {
      fmt.Printf("%v\n", n.Value)
   }

   //q.Empty()

   if q.Head() != nil{
      for n:= q.Head(); n!=nil; n=n.Next() {
         fmt.Printf("%v\n", n.Value)
      }
   }else{fmt.Println("Queue has been emptied")}

}
