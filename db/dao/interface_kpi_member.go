package dao

import "github.com/xionghengheng/ff_plib/db/model"

// KPIMemberInterface KPI 人员配置接口，参数校验和时间字段由业务层负责。
type KPIMemberInterface interface {
	// DeleteKPIMember 删除人员，保留历史月度配置。
	DeleteKPIMember(id int64) error
	// AddKPIMember 新增人员，自增 ID 回填到 record。
	AddKPIMember(record *model.KPIMemberModel) error
	// UpdateKPIMember 按人员 ID 更新指定字段，支持写入零值。
	UpdateKPIMember(id int64, updates map[string]interface{}) error
	// GetKPIMemberList 查询人员列表，memberType=0 查全部，否则按类型过滤；按 ID 升序。
	GetKPIMemberList(memberType int) ([]model.KPIMemberModel, error)
}

type KPIMemberInterfaceImp struct{}

var ImpKPIMember KPIMemberInterface = &KPIMemberInterfaceImp{}

var _ KPIMemberInterface = (*KPIMemberInterfaceImp)(nil)
