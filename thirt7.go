package main
import "fmt"
func main(){
  Map := map[int]string{
    	90: "Dog",
	91: "Cat",
	92: "Cow",
	93: "Bird",
	94: "Rabbit",
  }
	fmt.Println("Original map: ", Map)
	value1 := Map[90]
	value2 := Map[94]
	fmt.Println("Value of key[90]: ", value1)
	fmt.Println("Value of key[94]: ", value2)
}
