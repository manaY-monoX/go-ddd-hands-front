package session

import (
	"embed"
	"encoding/gob"
	"front-exercise/service/dto"

	"log"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

//go:embed secret.key
var secretKey embed.FS

// セッション管理用の構造体
type SessionManager struct {
	store *sessions.CookieStore
}

// セッション管理機能を追加するミドルウェア
// リクエストごとにセッションを取得または初期化をする
// セッションを後続のハンドラやミドルウェアで利用可能にする
func (s *SessionManager) SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	// ミドルウェアとして機能する関数を返す
	// Echoのミドルウェアは、echo.HandlerFuncを引数に取り、echo.HandlerFuncを返す必要がある
	return func(context echo.Context) error {
		// HTTPリクエストからセッションを取得する
		session, err := s.store.Get(context.Request(), "app-session")
		// セッションの取得中にエラーが発生した場合は、エラーを返して処理を中断する
		if err != nil {
			return err
		}
		// セッションオブジェクトをEchoコンテキストの 'session' というキーに関連付けする
		// リクエストを処理する他のハンドラやミドルウェアからセッションへのアクセスを可能にする
		context.Set("session", session)
		// 次のハンドラまたはミドルウェアを呼び出す
		// ミドルウェアチェーンのこの部分で次の処理を呼び出すことにより、リクエスト処理が連続する
		return next(context)
	}
}

// コンストラクタ
// セッションストアの初期化を行う
func NewSessionManager() *SessionManager {
	// dto.CategoryDTOをgobに登録する
	gob.Register([]dto.CategoryDTO{})
	// 埋め込まれたファイルを読み取る
	secretKey, err := secretKey.ReadFile("secret.key")
	if err != nil {
		log.Fatalf("埋め込まれたシークレットキーの読み取りに失敗しました: %v", err)
	}
	// CookieStoreを生成する
	store := sessions.NewCookieStore([]byte(secretKey))
	store.Options = &sessions.Options{
		Path:     "/",  // クッキーのパス
		MaxAge:   3600, // 有効期限(秒):1時間
		HttpOnly: true, // JavaScriptからアクセス不可
	}
	return &SessionManager{store: store}
}
