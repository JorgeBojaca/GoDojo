package main

import (
  "fmt"
  "math"

)

func Sqrt(x float64) float64 {
  var z float64= (x/2)
  for i := 0; i <10; i++{
    z = z - (((z*z)-x)/(2*z));
  }
  return z
}
func SqrtDelta(x float64) float64 {
  var z float64= (x/2);
  var zant float64= 0;
  var count = 0;
  var dz float64= 0;
  for{
    zant = z;
    z = z - (((z*z)-x)/(2*z));
    dz = zant-z;
    if (dz < (1e-6)){
      break;
    }
    count = count + 1;
  }
  fmt.Println("Contador", count);
  return z
}

func main() {
  fmt.Println(Sqrt(10899))
  fmt.Println(SqrtDelta(10899))
  fmt.Println(math.Sqrt(10899))
}
