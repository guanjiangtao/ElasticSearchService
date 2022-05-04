package main

import (
	"ElasticSearchService/tars-protocol/ElasticSearchServiceApp"
	"context"
	"github.com/TarsCloud/TarsGo/tars"
)

// WebApiImp servant implementation
type WebApiImp struct {
}

var logger = tars.GetLogger("ElasticSearchWebApi") // 注册日志API

type ElasticSearchWebApiImp struct {
}

func (imp *WebApiImp) Init() (ret int32, err error) {
	return 1, nil
}

func (imp *WebApiImp) Update(tarsCtx context.Context, id int32, docType string, data []ElasticSearchServiceApp.Data) (ret int32, err error) {

	return 0, nil
}

func (imp *WebApiImp) Insert(tarsCtx context.Context, data []ElasticSearchServiceApp.Data) (ret int32, err error) {

	return 0, nil
}

func (imp *WebApiImp) Delete(tarsCtx context.Context, id int32, docType string) (ret int32, err error) {
	logger.Info("删除成功！")
	return 999, nil
}

func (imp *WebApiImp) BulkInsert(tarsCtx context.Context, data map[string]ElasticSearchServiceApp.Data) (ret int32, err error) {

	return 0, nil
}

func (imp *WebApiImp) Search(tarsCtx context.Context, data map[string]string) (ret []ElasticSearchServiceApp.Data, err error) {

	return []ElasticSearchServiceApp.Data{}, nil
}
