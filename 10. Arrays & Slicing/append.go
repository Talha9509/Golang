package main


func main(){
  var a []int
  b := append(a, 3,4)
  printSlicee(b)
  c := append(b, 3)
  printSlicee(b)
  printSlicee(c)
}
