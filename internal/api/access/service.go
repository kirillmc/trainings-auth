package access

import (
	"github.com/kirillmc/trainings-auth/internal/service"
	desc "github.com/kirillmc/trainings-auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
