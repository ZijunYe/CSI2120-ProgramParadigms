package main
import "fmt"

type dog struct {
  name string
  race string
  female bool
}

//passing value + function name

func(d *dog) rename(new string){
  d.name = new
}
func main() {
    fido := dog {"Fido", "Poodle", false }
    fido.rename("Cocotte")
    //using the object call, it is not only passing the name to the function
    //it should passing the object
    fmt.Println(fido.name)
}
