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

	if "" == _req.Key {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "key is required"
		return nil
	}

	if "" == _req.Name {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "name is required"
		return nil
	}

	scope := &model.Scope{
		UUID: model.ToUUID(_req.Key),
		Key:  _req.Key,
		Name: _req.Name,
	}
	err := dao.Insert(scope)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = scope.UUID

	return nil
}

func (this *Scope) Update(_ctx context.Context, _req *proto.ScopeUpdateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Scope.Update, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	if "" == _req.Key {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "key is required"
		return nil
	}

	if "" == _req.Name {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "name is required"
		return nil
	}

	dao := model.NewScopeDAO(nil)
	scope := &model.Scope{
		UUID: _req.Uuid,
		Key:  _req.Key,
		Name: _req.Name,
	}
	err := dao.Update(scope)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = scope.UUID

	return nil
}

func (this *Scope) Delete(_ctx context.Context, _req *proto.ScopeDeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Scope.Delete, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	dao := model.NewScopeDAO(nil)
	err := dao.Delete(_req.Uuid)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = _req.Uuid

	return nil
}

func (this *Scope) Get(_ctx context.Context, _req *proto.ScopeGetRequest, _rsp *proto.ScopeGetResponse) error {
	logger.Infof("Received Scope.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	dao := model.NewScopeDAO(nil)
	scope, err := dao.Get(_req.Uuid)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Entity = &proto.ScopeEntity{
		Uuid: scope.UUID,
		Key:  scope.Key,
		Name: scope.Name,
	}

	return nil
}

func (this *Scope) List(_ctx context.Context, _req *proto.ScopeListRequest, _rsp *proto.ScopeListResponse) error {
	logger.Infof("Received Scope.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}
	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewScopeDAO(nil)
	total, scopes, err := dao.List(offset, count)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Total = uint64(total)
	_rsp.Entity = make([]*proto.ScopeEntity, len(scopes))
	for i, e := range scopes {
		_rsp.Entity[i] = &proto.ScopeEntity{
			Uuid: e.UUID,
			Key:  e.Key,
			Name: e.Name,
		}
	}

	return nil
}

func (this *Scope) Search(_ctx context.Context, _req *proto.ScopeSearchRequest, _rsp *proto.ScopeListResponse) error {
	logger.Infof("Received Scope.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}
	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewScopeDAO(nil)
	total, scopes, err := dao.Search(offset, count, _req.Key, _req.Name)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Total = uint64(total)
	_rsp.Entity = make([]*proto.ScopeEntity, len(scopes))
	for i, e := range scopes {
		_rsp.Entity[i] = &proto.ScopeEntity{
			Uuid: e.UUID,
			Key:  e.Key,
			Name: e.Name,
		}
	}

	return nil
}
