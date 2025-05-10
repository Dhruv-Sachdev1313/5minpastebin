package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Dhruv-Sachdev1313/5minpastebin.com/cmd/web"
	"github.com/Dhruv-Sachdev1313/5minpastebin.com/internal/pkg"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))
	e.GET("/", echo.WrapHandler(templ.Handler(web.HomeTextArea())))
	e.POST("/create_link", s.createLinkHandler)
	e.GET("/:id", s.serveContentWithIdHandler)
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.redis.Health())
}

func (s *Server) createLinkHandler(c echo.Context) error {
	content := c.FormValue("content")
	url_id := pkg.RandomStringGenerator(10)
	s.redis.SetKeyWithExpiration(url_id, content)
	component := web.Link(fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), url_id))
	err := component.Render(c.Request().Context(), c.Response().Writer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Rendering failed")
	}
	return nil
}

func (s *Server) serveContentWithIdHandler(c echo.Context) error {
	url_id := c.Param("id")
	content, err := s.redis.GetData(url_id)
	if err != nil {
		err = web.NotFoundPage().Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Rendering failed")
		}
		return nil
	}
	component := web.ViewPaste(content)
	err = component.Render(c.Request().Context(), c.Response().Writer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Rendering failed")
	}
	return nil
}
