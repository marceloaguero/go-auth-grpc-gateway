package delivery

import (
	"github.com/marceloaguero/go-auth-grpc-gateway/users/pkg/user"
)

type delivery struct {
	usecase user.Usecase
}

func newDelivery(uc user.Usecase) *delivery {
	return &delivery{
		usecase: uc,
	}
}
