package ARF_Sidecar

import util "github-connector/pkg/util"

import (
	"github-connector/pkg/dto"
	"github-connector/pkg/functions"
	"github-connector/pkg/model"
	"github-connector/pkg/rule/ARF_CreateRepo"
)

type CFGContext struct {
	Request      *dto.Request
	request      string
	requestJson  *model.Requestjsonmodel
	returnString string
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Request:     pRequest,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_Sidecar(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequest, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()

	cfg.xiSubRuleNodeM12()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM12() error {
	cfg.returnString = ARF_CreateRepo.ARF_CreateRepo(cfg.requestJson).(string)
	cfg._returnValue = cfg.returnString
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.requestJson, err = functions.ConvertStringToJson(cfg.request)
	cfg._returnValue = cfg.requestJson
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.request = cfg.Request.Data
	cfg.requestJson = &model.Requestjsonmodel{}

	return nil
}
