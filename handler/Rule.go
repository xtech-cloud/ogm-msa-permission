package handler

import (
	"context"
	"ogm-permission/model"

	"github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-permission/proto/permission"
)

type Rule struct{}

func (this *Rule) Add(_ctx context.Context, _req *proto.RuleAddRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Rule.Add, req is %v", _req)
	_rsp.Status = &proto.Status{}

	dao := model.NewRuleDAO(nil)

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

	if "" == _req.Scope {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "scope is required"
		return nil
	}

	rule := &model.Rule{
		UUID:  model.ToUUID(_req.Scope + _req.Key),
        Scope: _req.Scope,
		Key:   _req.Key,
		Name:  _req.Name,
		State: _req.State,
	}

	err := dao.Insert(rule)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = rule.UUID

	return nil
}

func (this *Rule) Update(_ctx context.Context, _req *proto.RuleUpdateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Rule.Update, req is %v", _req)
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

	dao := model.NewRuleDAO(nil)
	rule := &model.Rule{
		UUID:  _req.Uuid,
		Key:   _req.Key,
		Name:  _req.Name,
		State: _req.State,
	}
	err := dao.Update(rule)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = rule.UUID

	return nil
}

func (this *Rule) Delete(_ctx context.Context, _req *proto.RuleDeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Rule.Delete, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	dao := model.NewRuleDAO(nil)
	err := dao.Delete(_req.Uuid)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Uuid = _req.Uuid
	return nil
}

func (this *Rule) Get(_ctx context.Context, _req *proto.RuleGetRequest, _rsp *proto.RuleGetResponse) error {
	logger.Infof("Received Rule.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	dao := model.NewRuleDAO(nil)
	rule, err := dao.Get(_req.Uuid)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Entity = &proto.RuleEntity{
		Uuid:  rule.UUID,
		Scope: rule.Scope,
		Key:   rule.Key,
		Name:  rule.Name,
		State: rule.State,
	}

	return nil
}

func (this *Rule) List(_ctx context.Context, _req *proto.RuleListRequest, _rsp *proto.RuleListResponse) error {
	logger.Infof("Received Rule.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Scope {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "scope is required"
		return nil
	}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}
	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewRuleDAO(nil)
	total, rules, err := dao.List(offset, count, _req.Scope)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Total = uint64(total)
	_rsp.Entity = make([]*proto.RuleEntity, len(rules))
	for i, e := range rules {
		_rsp.Entity[i] = &proto.RuleEntity{
			Uuid:  e.UUID,
			Key:   e.Key,
			Name:  e.Name,
			State: e.State,
		}
	}

	return nil
}

func (this *Rule) Search(_ctx context.Context, _req *proto.RuleSearchRequest, _rsp *proto.RuleListResponse) error {
	logger.Infof("Received Rule.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

	if "" == _req.Scope {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "scope is required"
		return nil
	}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}
	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewRuleDAO(nil)
	total, rules, err := dao.Search(offset, count, _req.Scope, _req.Key, _req.Name, _req.State)
	if nil != err {
		_rsp.Status.Code = -1
		_rsp.Status.Message = err.Error()
		return nil
	}

	_rsp.Total = uint64(total)
	_rsp.Entity = make([]*proto.RuleEntity, len(rules))
	for i, e := range rules {
		_rsp.Entity[i] = &proto.RuleEntity{
			Uuid:  e.UUID,
			Key:   e.Key,
			Name:  e.Name,
			Scope: e.Scope,
			State: e.State,
		}
	}

	return nil
}
