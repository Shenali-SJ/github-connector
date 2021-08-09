package ARF_CreateRepo

import util "github-connector/pkg/util"

import (
	"github-connector/pkg/functions"
	"github-connector/pkg/model"
)

type CFGContext struct {
	Requestjsonmodel *model.Requestjsonmodel
	repoData         *model.Createrepodata
	returnString     string
	_context         *util.CustomContext
	_ruleConfig      map[string]interface{}
	_returnValue     interface{}
	_errorValue      interface{}
}

func NewCFGContext(pRequestjsonmodel *model.Requestjsonmodel, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Requestjsonmodel: pRequestjsonmodel,
		_context:         pContext,
		_ruleConfig:      pRuleConfig,
	}
}
func ARF_CreateRepo(pRequestjsonmodel *model.Requestjsonmodel, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequestjsonmodel, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.returnString, err = functions.CreateGithubRepo(cfg.Requestjsonmodel)
	cfg._returnValue = cfg.returnString
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.repoData = &model.Createrepodata{}

	return nil
}
