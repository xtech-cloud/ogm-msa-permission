package handler

import (
	"context"
	"ogm-permission/model"

    "github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-permission/proto/permission"
)

type Scope struct{}


func (this *Scope) Create(_ctx context.Context, _req *proto.ScopeCreateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Scope.Create, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewScopeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Scope) Update(_ctx context.Context, _req *proto.ScopeUpdateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Scope.Update, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewScopeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Scope) Delete(_ctx context.Context, _req *proto.ScopeDeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Scope.Delete, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewScopeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Scope) Get(_ctx context.Context, _req *proto.ScopeGetRequest, _rsp *proto.ScopeGetResponse) error {
	logger.Infof("Received Scope.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewScopeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Scope) List(_ctx context.Context, _req *proto.ScopeListRequest, _rsp *proto.ScopeListResponse) error {
	logger.Infof("Received Scope.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewScopeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Scope) Search(_ctx context.Context, _req *proto.ScopeSearchRequest, _rsp *proto.ScopeListResponse) error {
	logger.Infof("Received Scope.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewScopeDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 


