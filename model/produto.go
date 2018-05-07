package model

//Product representa um tipo de produto
type Product struct {
	ID    int    `json:"index"`
	Name  string `json:"nome"`
	Type  string `json:"tipo"`
	value float64
}

//SetValue recebe o valor do produto
func (p *Product) SetValue(v float64) {
	p.value = v
}

//GetValue retorna o valor do produto
func (p Product) GetValue() float64 {
	return p.value
}

//CreateProduct cria um slice de produtos
func (p Product) CreateProduct() (prod []Product) {
	product1 := Product{}
	product1.ID = 1
	product1.Name = "Caneta"
	product1.Type = "Utilitarios"
	product1.SetValue(1.00)

	product2 := Product{}
	product2.ID = 2
	product2.Name = "Lapis"
	product2.Type = "Utilitarios"
	product2.SetValue(1.50)

	product3 := Product{}
	product3.ID = 3
	product3.Name = "Caderno"
	product3.Type = "Utilitarios"
	product3.SetValue(3.75)

	prod = append(prod, product1)
	prod = append(prod, product2)
	prod = append(prod, product3)

	return
}
