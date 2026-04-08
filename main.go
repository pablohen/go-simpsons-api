package main

import (
	"log"
	"net/url"
	"os"

	"go-simpsons-api/internal/handler"
	"go-simpsons-api/internal/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-simpsons-api/docs"
)

// @title The Simpsons API (proxied)
// @version 1.0
// @description Mirrors https://thesimpsonsapi.com/api. JSON responses are proxied from the upstream service; pagination next/prev URLs reference the upstream host.
// @termsOfService https://thesimpsonsapi.com/

// @contact.name API support
// @contact.url https://thesimpsonsapi.com/

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// Swag references models.* in route comments; keep a compile-time link to the package.
var _ = models.APIIndex{}

// Regenerate API docs after changing annotations or models:
//
//	go run github.com/swaggo/swag/cmd/swag@latest init --parseDependency -g main.go -o docs

func main() {
	upstream := os.Getenv("UPSTREAM")
	if upstream == "" {
		upstream = "https://thesimpsonsapi.com"
	}
	origin, err := url.Parse(upstream)
	if err != nil {
		log.Fatalf("UPSTREAM: %v", err)
	}

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	proxy := handler.NewUpstreamProxy(origin)
	srv := &apiServer{proxy: proxy}

	r := gin.Default()
	r.GET("/health", health)

	srv.registerAPI(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("listening on %s, upstream %s", addr, origin.String())
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

type apiServer struct {
	proxy gin.HandlerFunc
}

func (s *apiServer) registerAPI(r *gin.Engine) {
	r.GET("/api", s.getAPIIndex)
	r.GET("/api/characters", s.listCharacters)
	r.GET("/api/characters/:id", s.getCharacter)
	r.GET("/api/episodes", s.listEpisodes)
	r.GET("/api/episodes/:id", s.getEpisode)
	r.GET("/api/locations", s.listLocations)
	r.GET("/api/locations/:id", s.getLocation)
}

// getAPIIndex godoc
// @Summary API index
// @Description Returns links to the characters, episodes, and locations collection endpoints.
// @Tags meta
// @Produce json
// @Success 200 {object} models.APIIndex
// @Router /api [get]
func (s *apiServer) getAPIIndex(c *gin.Context) {
	s.proxy(c)
}

// listCharacters godoc
// @Summary List characters
// @Description Paginated list of characters (20 per page). Use the page query parameter.
// @Tags characters
// @Produce json
// @Param page query int false "Page number (1-based)"
// @Success 200 {object} models.PaginatedCharacters
// @Router /api/characters [get]
func (s *apiServer) listCharacters(c *gin.Context) {
	s.proxy(c)
}

// getCharacter godoc
// @Summary Get character by ID
// @Description Full character record including description and optional nested first-appearance episode.
// @Tags characters
// @Produce json
// @Param id path int true "Character ID"
// @Success 200 {object} models.CharacterDetail
// @Failure 404 {string} string "Not found"
// @Router /api/characters/{id} [get]
func (s *apiServer) getCharacter(c *gin.Context) {
	s.proxy(c)
}

// listEpisodes godoc
// @Summary List episodes
// @Description Paginated list of episodes (20 per page).
// @Tags episodes
// @Produce json
// @Param page query int false "Page number (1-based)"
// @Success 200 {object} models.PaginatedEpisodes
// @Router /api/episodes [get]
func (s *apiServer) listEpisodes(c *gin.Context) {
	s.proxy(c)
}

// getEpisode godoc
// @Summary Get episode by ID
// @Description Episode detail including long description.
// @Tags episodes
// @Produce json
// @Param id path int true "Episode ID"
// @Success 200 {object} models.EpisodeDetail
// @Failure 404 {string} string "Not found"
// @Router /api/episodes/{id} [get]
func (s *apiServer) getEpisode(c *gin.Context) {
	s.proxy(c)
}

// listLocations godoc
// @Summary List locations
// @Description Paginated list of Springfield locations (20 per page).
// @Tags locations
// @Produce json
// @Param page query int false "Page number (1-based)"
// @Success 200 {object} models.PaginatedLocations
// @Router /api/locations [get]
func (s *apiServer) listLocations(c *gin.Context) {
	s.proxy(c)
}

// getLocation godoc
// @Summary Get location by ID
// @Description Location detail including description and optional nested appearances.
// @Tags locations
// @Produce json
// @Param id path int true "Location ID"
// @Success 200 {object} models.LocationDetail
// @Failure 404 {string} string "Not found"
// @Router /api/locations/{id} [get]
func (s *apiServer) getLocation(c *gin.Context) {
	s.proxy(c)
}
