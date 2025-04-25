package impls

import (
	"fmt"
	"front-exercise/config"
	"front-exercise/service/dto"
	"front-exercise/service/errortype"
	"net"
	"net/http"
	"time"
)

// 共通サービス構造体
type BaseService struct {
	Config *config.Config // アプリケーション環境情報
}

// サーバ接続確認
func (b *BaseService) Ping() error {
	// 接続を試みるホストとポート
	host := fmt.Sprintf("%s:%d", b.Config.Server.Hostname, b.Config.Server.Port)
	timeout := 3 * time.Second // タイムアウト設定
	// TCPで接続を試みる
	conn, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		return errortype.NewInternalError("サーバ接続ができませんでした。")
	}
	defer conn.Close() // 接続を閉じる
	return nil
}

// レスポンス生成
func (b *BaseService) CreateResponse(resp *http.Response) (*dto.ResponseDTO, error) {
	return nil, nil
}

// コンストラクタ
func NewBaseService(config *config.Config) *BaseService {
	return &BaseService{
		Config: config,
	}
}
