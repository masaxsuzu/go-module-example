package infra

import (
	"fmt"

	"github.com/masa-suzu/go-module-example/core/domain"
)

type Db struct {
}

func (db *Db) Save(u domain.User) {
	fmt.Printf("[save]%s\n", u)
}
