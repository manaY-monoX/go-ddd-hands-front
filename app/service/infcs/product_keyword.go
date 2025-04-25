package infcs

import "front-exercise/service/dto"

// 商品キーワード検索インターフェイス
type ProductKeyword interface {
	Execute(keyword string) (*dto.ResponseDTO, error)
}
