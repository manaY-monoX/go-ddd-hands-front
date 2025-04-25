package dto

// プレゼンテーション層と相互に利用するDTO
type ResponseDTO struct {
	Categories *[]CategoryDTO // 商品カテゴリのスライス
	Products   *[]ProductDTO  // 商品のスライス
	Product    *ProductDTO    // 商品
	Error      *ErrorDTO      // エラーメッセージ
}

func NewResponseDTO() *ResponseDTO {
	return &ResponseDTO{
		Categories: nil,
		Products:   nil,
		Product:    nil,
		Error:      nil,
	}
}
