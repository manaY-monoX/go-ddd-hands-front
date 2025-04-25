package impls

import (
	"bytes"
	"encoding/json"
	"front-exercise/config"
	"front-exercise/service/dto"
	"front-exercise/service/infcs"

	"io"
	"net/http"
	"strings"
)

// 商品登録インターフェイスの実装
type productRegisterImpl struct {
	*BaseService // BaseServiceをエンベデッドする
}

// 商品を登録する
func (p *productRegisterImpl) Execute(product *dto.ProductDTO) (*dto.ResponseDTO, error) {
	// サーバーとの接続を確認する
	err := p.BaseService.Ping()
	if err != nil {
		return nil, err
	}
	// 登録データをJSON形式に変換する
	jsonData, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	// リクエストボディを作成する
	requestBody := bytes.NewBuffer(jsonData)
	// 商品登録URLを生成する
	url, err := p.BaseService.Config.BuildServerURL("productRegister")
	if err != nil {
		return nil, err
	}
	// 商品登録リクエストを送信する
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return nil, err
	}
	// 取得したレスポンスのボディをクローズする
	defer resp.Body.Close()
	// プレゼンテーション層に返すResponseDTOを生成する
	return p.CreateResponse(resp)
}

// レスポンス生成メソッドを*productRegisterImplでシャドーイング(隠蔽)する
func (p *productRegisterImpl) CreateResponse(resp *http.Response) (*dto.ResponseDTO, error) {
	response := dto.NewResponseDTO()
	// レスポンスボディを取得する
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 商品登録成功
	if resp.StatusCode == http.StatusOK {
		var product dto.ProductDTO
		// JSONをProductDTOに変換する
		err = json.Unmarshal(body, &product)
		if err != nil {
			return nil, err
		}
		response.Product = &product
	} else {
		// エラーメッセージをErrorDTOに格納する
		response.Error = &dto.ErrorDTO{Message: strings.Trim(string(body), "\"")}
	}
	return response, nil
}

// コンストラクタ
func NewproductRegisterImpl(config *config.Config) infcs.ProductRegister {
	return &productRegisterImpl{
		BaseService: NewBaseService(config),
	}
}
