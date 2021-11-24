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
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Rule) Update(_ctx context.Context, _req *proto.RuleUpdateRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Rule.Update, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewRuleDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Rule) Delete(_ctx context.Context, _req *proto.RuleDeleteRequest, _rsp *proto.UuidResponse) error {
	logger.Infof("Received Rule.Delete, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewRuleDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Rule) Get(_ctx context.Context, _req *proto.RuleGetRequest, _rsp *proto.RuleGetResponse) error {
	logger.Infof("Received Rule.Get, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewRuleDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Rule) List(_ctx context.Context, _req *proto.RuleListRequest, _rsp *proto.RuleListResponse) error {
	logger.Infof("Received Rule.List, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewRuleDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 

func (this *Rule) Search(_ctx context.Context, _req *proto.RuleSearchRequest, _rsp *proto.RuleListResponse) error {
	logger.Infof("Received Rule.Search, req is %v", _req)
	_rsp.Status = &proto.Status{}

    dao := model.NewRuleDAO(nil)
    _, err := dao.Count()
    if nil != err {
        _rsp.Status.Code = -1
        _rsp.Status.Message = err.Error()
        return nil
    }

	return nil
} 


