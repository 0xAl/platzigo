//Range 

package main 

import "fmt"

func main(){
	
	//arreglo
	numeros := []int{2,4,3,6}

	suma := 0

	//Recorre el arreglo y va sumando sus valores en la variable suma
	for _, num := range numeros{
		suma +=  num
	} 

	//Muestra la suma del arreglo
	fmt.Println("suma: ", suma)

	//Muestra resultado si es que existe un 3 en el arreglo
	for i, numero := range numeros{
		  if numero == 3 {
		  	fmt.Println("index: ", i)
		  }
	}	

	//Mapa
	algo := map[string]string{"a":"auto","b":"bebe"}

	//Muestra los resultados del mapa
	for key, value := range algo {
	fmt.Printf("La %s es de %s\n",key,value)
	}

}

