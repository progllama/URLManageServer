package services

import (
	"url_manager/domain/models"
)

type AuthenticationService interface {
	Authenticate(models.Credential) (bool, error)
}
