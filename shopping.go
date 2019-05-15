package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type shopping struct {
	ig *instagram
}

func newShopping(i *instagram) *shopping {
	return &shopping{ig: i}
}

func (sh *shopping) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = sh.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}