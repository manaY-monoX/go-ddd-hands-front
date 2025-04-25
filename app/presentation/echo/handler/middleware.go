package handler

import (
	"front-exercise/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 一致するパス以外の場合にリダイレクトするミドルウェア
type pathChecker struct {
	pathTop      string
	allowedPaths map[string]bool
}

// 一致するパス以外の場合にリダイレクトするミドルウェア
func (p *pathChecker) RedirectMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if !p.allowedPaths[path] {
			// ここで指定したURLにリダイレクト
			return c.Redirect(http.StatusMovedPermanently, p.pathTop)
		}
		return next(c)
	}
}

// コンストラクタ
func NewpathChecker(config *config.Config) *pathChecker {
	allowedPaths := make(map[string]bool)
	allowedPaths[config.App.Group+config.App.Paths.Top] = true
	allowedPaths[config.App.Group+config.App.Paths.Search] = true
	allowedPaths[config.App.Group+config.App.Paths.Register] = true
	allowedPaths[config.App.Group+config.App.Paths.Error] = true

	return &pathChecker{
		pathTop:      config.App.Group + config.App.Paths.Top,
		allowedPaths: allowedPaths,
	}
}
