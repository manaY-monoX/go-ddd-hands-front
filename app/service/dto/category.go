package dto

import "fmt"

// 商品カテゴリを保持するDTO
type CategoryDTO struct {
	Id   string `json:"categoryId"`
	Name string `json:"categoryName"`
}

func (c CategoryDTO) String() string {
	return fmt.Sprintf("CategoryDTO[Id:%s,Name:%s]", c.Id, c.Name)
}
