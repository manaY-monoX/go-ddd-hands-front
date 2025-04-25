package dto

import "fmt"

// 商品を保持するDTO
type ProductDTO struct {
	Id       string      `json:"productId"`
	Name     string      `json:"productName"`
	Price    string      `json:"productPrice"`
	Category CategoryDTO `json:"category"`
}

func (p ProductDTO) String() string {
	return fmt.Sprintf("ProductDTO[Id:%s,Name:%s,Price:%s,%s]",
		p.Id, p.Name, p.Price, p.Category.String())
}
