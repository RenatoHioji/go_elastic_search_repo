package app

import (
	"database/sql"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/service"
)

type App struct {
	user_service *service.UserService
}

func NewApp(db *sql.DB) {

}
