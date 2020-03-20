package partner

import "github.com/kekhuay/echo-installation-service/model"

type Store interface {
	CreatePartner(*model.Partner) error
}
