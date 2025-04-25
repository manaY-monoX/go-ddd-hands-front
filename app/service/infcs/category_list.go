package infcs

import "front-exercise/service/dto"

// 商品カテゴリ一覧取得インターフェイス
type CategoryList interface {
	Execute() (*dto.ResponseDTO, error)
}
