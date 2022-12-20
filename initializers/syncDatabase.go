package initializers

import (
	"github.com/bangsyir/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
