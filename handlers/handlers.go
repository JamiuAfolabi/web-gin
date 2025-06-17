package handlers

import (
	"jamiuafolabi/web-service-gin/config"
	"jamiuafolabi/web-service-gin/databasedriver"
	"jamiuafolabi/web-service-gin/repository"
	"jamiuafolabi/web-service-gin/repository/dbrepo"

	"github.com/gin-gonic/gin"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(app *config.AppConfig, db *databasedriver.DB) *Repository {
	return &Repository{
		App: app,
		DB:  dbrepo.NewPostgresRepo(db.SQL, app),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HelloWorld(context *gin.Context) {
	products, err := m.DB.GetAllRows()
	// m.DB.GetAllRows()

	if err != nil {
		context.JSON(500, gin.H{"error": "Unable to fetch products"})
		return
	}

	context.IndentedJSON(200, products)
}
