package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  // x := 2 //type is inferred
  var a string
  var x int = 2
  var b int
  var y = "yes"
  var c bool
  if x > 0 {
    fmt.Println(y)
}

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
// Thursday: Error Types, Error Behaviors