package ddk

import "github.com/justoxh/pinduoduo-sdk/common"

type DuoduoKe struct {
	*common.Context
}

func NewDuoduoKe(context *common.Context) *DuoduoKe {
	service := new(DuoduoKe)
	service.Context = context
	return service
}
