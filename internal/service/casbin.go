package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"sync"
)

type sCasbin struct{}
type adapterCasbin struct {
	Enforcer    *casbin.SyncedEnforcer
	EnforcerErr error
}

var (
	casbinCtx = gctx.New()
	insCasbin = sCasbin{}
	once      sync.Once
	ac        *adapterCasbin
)

// Casbin 访问控制服务
func Casbin() *sCasbin {
	return &insCasbin
}

// GetEnforcer 获取adapter单例对象
func (s *sCasbin) GetEnforcer() (enforcer *casbin.SyncedEnforcer, err error) {
	once.Do(func() {
		ac = s.newAdapter()
	})
	enforcer = ac.Enforcer
	err = ac.EnforcerErr
	return
}

// SavePolicy 将策略保存到数据库
func (a *adapterCasbin) SavePolicy(model model.Model) (err error) {
	err = a.dropTable()
	if err != nil {
		return
	}
	err = a.createTable()
	if err != nil {
		return
	}

	for pType, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(pType, rule)
			_, err = dao.CasbinRule.Ctx(casbinCtx).Data(line).Insert()
			if err != nil {
				return
			}
		}
	}
	for pType, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(pType, rule)
			_, err = dao.CasbinRule.Ctx(casbinCtx).Data(line).Insert()
			if err != nil {
				return
			}
		}
	}
	return
}

// LoadPolicy 从数据库加载策略
func (a *adapterCasbin) LoadPolicy(model model.Model) (err error) {
	var lines []*entity.CasbinRule
	err = dao.CasbinRule.Ctx(casbinCtx).Scan(&lines)
	if err != nil {
		return
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return
}

// AddPolicy 添加策略规则
func (a *adapterCasbin) AddPolicy(sec string, pType string, rule []string) (err error) {
	line := savePolicyLine(pType, rule)
	_, err = dao.CasbinRule.Ctx(casbinCtx).Data(line).Insert()
	return
}

// RemovePolicy 移除策略规则
func (a *adapterCasbin) RemovePolicy(sec string, pType string, rule []string) (err error) {
	line := savePolicyLine(pType, rule)
	err = rawDelete(line)
	return
}

// RemoveFilteredPolicy 移除匹配筛选器的策略规则
func (a *adapterCasbin) RemoveFilteredPolicy(sec string, pType string, fieldIndex int, fieldValues ...string) (err error) {
	line := &entity.CasbinRule{}
	line.PType = pType
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err = rawDelete(line)
	return
}

// 初始化adapter操作
func (s *sCasbin) newAdapter() (a *adapterCasbin) {
	a = &adapterCasbin{}
	a.initPolicy()
	return
}

func (a *adapterCasbin) dropTable() (err error) {
	return
}

func (a *adapterCasbin) createTable() (err error) {
	return
}

func (a *adapterCasbin) initPolicy() {
	// 因为DB一开始是空的，所以我们需要从文件适配器.csv文件加载策略
	e, err := casbin.NewSyncedEnforcer(
		g.Cfg().MustGet(casbinCtx, "casbin.modelFile").String(),
		g.Cfg().MustGet(casbinCtx, "casbin.policyFile").String(),
	)
	if err != nil {
		a.EnforcerErr = err
		return
	}
	// 设置适配器
	e.SetAdapter(a)
	// 清除当前的策略
	e.ClearPolicy()
	a.Enforcer = e
	// 从DB加载策略
	err = a.LoadPolicy(e.GetModel())
	if err != nil {
		a.EnforcerErr = err
		return
	}
}

func loadPolicyLine(line *entity.CasbinRule, model model.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	persist.LoadPolicyLine(lineText, model)
}

func savePolicyLine(pType string, rule []string) (line *entity.CasbinRule) {
	line = &entity.CasbinRule{}
	line.PType = pType
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return
}

func rawDelete(line *entity.CasbinRule) (err error) {
	dbModel := dao.CasbinRule.Ctx(casbinCtx).Where(do.CasbinRule{PType: line.PType})
	if line.V0 != "" {
		dbModel = dbModel.Where(do.CasbinRule{V0: line.V0})
	}
	if line.V1 != "" {
		dbModel = dbModel.Where(do.CasbinRule{V1: line.V1})
	}
	if line.V2 != "" {
		dbModel = dbModel.Where(do.CasbinRule{V2: line.V2})
	}
	if line.V3 != "" {
		dbModel = dbModel.Where(do.CasbinRule{V3: line.V3})
	}
	if line.V4 != "" {
		dbModel = dbModel.Where(do.CasbinRule{V4: line.V4})
	}
	if line.V5 != "" {
		dbModel = dbModel.Where(do.CasbinRule{V5: line.V5})
	}
	_, err = dbModel.Delete()
	return
}
