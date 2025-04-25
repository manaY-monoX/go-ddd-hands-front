package handler

// Handlers: Echoに登録するハンドラーをグループ化
type Handlers struct {
	Top              *topPageHandler
	Search           *searchPageHandler
	RegisterEntry    *registerEnterPageHandler
	RegisterComplete *registerCompletePageHandler
	PathChecker      *pathChecker
	Err              *errorHandler
}

// Handlers構造体のインスタンスをfxのコンテナに提供するための関数
func ProvideHandlers(
	top *topPageHandler,
	search *searchPageHandler,
	registerEntry *registerEnterPageHandler,
	registerComplete *registerCompletePageHandler,
	err *errorHandler,
	pathChecker *pathChecker,
) Handlers {
	return Handlers{
		Top:              top,
		Search:           search,
		RegisterEntry:    registerEntry,
		RegisterComplete: registerComplete,
		PathChecker:      pathChecker,
		Err:              err,
	}
}
