package impls

import (
	"encoding/json"
	"fmt"
	"front-exercise/config"
	"front-exercise/service/dto"
	"front-exercise/service/infcs"

	"io"
	"net/http"
	"strings"
)

// 商品キーワード検索インターフェイスの実装
type productKeywordImpl struct {
	*BaseService // BaseServiceをエンベデッドする
}

// 商品キーワード検索
func (p *productKeywordImpl) Execute(keyword string) (*dto.ResponseDTO, error) {
	// サーバーとの接続を確認する
	err := p.BaseService.Ping()
	if err != nil {
		return nil, err
	}
	// 商品キーワード検索URLを生成する
	url, err := p.BaseService.Config.BuildServerURL("productKeyword")
	if err != nil {
		return nil, err
	}
	// URLに検索パラメータを接続する
	url = fmt.Sprintf(url, keyword)
	// 検索リクエストを送信する
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 取得したレスポンスのボディをクローズする
	defer resp.Body.Close()
	// プレゼンテーション層に返すResponseDTOを生成する
	return p.CreateResponse(resp)
}

// レスポンス生成メソッドをproductSearchImplでシャドーイング(隠蔽)する
func (p *productKeywordImpl) CreateResponse(resp *http.Response) (*dto.ResponseDTO, error) {
	response := dto.NewResponseDTO()
	// レスポンスボディを取得する
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 商品キーワード検索成功
	if resp.StatusCode == http.StatusOK {
		var products []dto.ProductDTO
		// JSONをProductDTOのスライスに変換する
		err = json.Unmarshal(body, &products)
		if err != nil {
			return nil, err
		}
		response.Products = &products
	} else {
		// エラーメッセージをErrorDTOに格納する
		response.Error = &dto.ErrorDTO{Message: strings.Trim(string(body), "\"")}
	}
	return response, nil
}

// コンストラクタ
func NewproductKeywordImpl(config *config.Config) infcs.ProductKeyword {
	return &productKeywordImpl{
		BaseService: NewBaseService(config),
	}
}
