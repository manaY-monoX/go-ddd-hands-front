package impls

import (
	"encoding/json"
	"front-exercise/config"
	"front-exercise/service/dto"
	"front-exercise/service/infcs"

	"io"
	"net/http"
	"strings"
)

// 商品カテゴリ取得インターフェイスの実装
type categoryListImpl struct {
	*BaseService // BaseServiceをエンベデッドする
}

// 商品カテゴリを取得する
func (c *categoryListImpl) Execute() (*dto.ResponseDTO, error) {
	// サーバーとの接続を確認する
	err := c.BaseService.Ping()
	if err != nil {
		return nil, err
	}
	// 商品カテゴリ取得URLを生成する
	url, err := c.BaseService.Config.BuildServerURL("categoryList")
	if err != nil {
		return nil, err
	}
	// 商品カテゴリ取得リクエストを送信する
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 取得したレスポンスのボディをクローズする
	defer resp.Body.Close()
	// プレゼンテーション層に返すResponseDTOを生成する
	return c.CreateResponse(resp)
}

// レスポンス生成メソッドをcategoryListImplでシャドーイング(隠蔽)する
func (c *categoryListImpl) CreateResponse(resp *http.Response) (*dto.ResponseDTO, error) {
	response := dto.NewResponseDTO()
	// レスポンスボディを取得する
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 商品カテゴリ取得成功
	if resp.StatusCode == http.StatusOK {
		var categories []dto.CategoryDTO
		// JSONをCategoryDTOのスライスに変換する
		err = json.Unmarshal(body, &categories)
		if err != nil {
			return nil, err
		}
		response.Categories = &categories
	} else {
		// エラーメッセージをErrorDTOに格納する
		response.Error = &dto.ErrorDTO{Message: strings.Trim(string(body), "\"")}
	}
	return response, nil
}

// コンストラクタ
func NewcategoryListImpl(config *config.Config) infcs.CategoryList {
	return &categoryListImpl{
		BaseService: NewBaseService(config),
	}
}
