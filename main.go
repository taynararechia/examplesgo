package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/taynararechia/curso-go/projetogo/model"
)

func main() {

	products := model.Product{}.CreateProduct()

	for id, prod := range products {
		fmt.Printf("Produto: [%d]\n%v\n", id, prod)
	}

	mapProduct := make(map[string]model.Product)
	for _, mapProd := range products {
		mapProduct[mapProd.Name] = mapProd
		fmt.Println("Map1: ", mapProduct)
		// fmt.Println("Map2: ", mapProd)
	}

	jsonProd, err := json.Marshal(products)
	if err != nil {
		return
	}
	fmt.Println("Produtos: ", string(jsonProd))

	structP := model.Product{}.CreateProduct()
	json.Unmarshal(jsonProd, &structP)

	fmt.Println("Produtos unmarshal: ", structP)

	arquivoJSON, err := os.Create("products.json")
	if err != nil {
		return
	}
	defer arquivoJSON.Close()

	write := bufio.NewWriter(arquivoJSON)
	write.WriteString("[\r\n")
	for indice, product := range products {
		productJSON, _ := json.Marshal(product)

		write.WriteString(" " + string(productJSON))
		if indice+1 < len(products) {
			write.WriteString(", \r\n")
		}
	}
	write.WriteString("\r\n]")
	write.Flush()
}
