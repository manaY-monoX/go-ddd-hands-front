package infcs

import "front-exercise/service/dto"

// 商品登録インターフェイス
type ProductRegister interface {
	Execute(product *dto.ProductDTO) (*dto.ResponseDTO, error)
}
