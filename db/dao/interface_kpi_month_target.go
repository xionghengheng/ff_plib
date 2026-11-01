package dao

import "github.com/xionghengheng/ff_plib/db/model"

// KPIMonthTargetInterface 月度目标配置接口，参数校验、操作人及时间字段由业务层负责。
type KPIMonthTargetInterface interface {
	// SaveKPIMonthTarget 事务保存目标及变更日志，保留原创建信息。
	SaveKPIMonthTarget(record *model.KPIMonthTargetModel) error
	// AddKPIMonthTarget 新增月度目标，ID 回填；重复月份由数据库唯一键报错。
	AddKPIMonthTarget(record *model.KPIMonthTargetModel) error
	// UpdateKPIMonthTarget 按月份更新指定字段，支持写入零值。
	UpdateKPIMonthTarget(statMonth string, updates map[string]interface{}) error
	// GetKPIMonthTarget 按月份查询目标，未找到返回 gorm.ErrRecordNotFound。
	GetKPIMonthTarget(statMonth string) (*model.KPIMonthTargetModel, error)
}

type KPIMonthTargetInterfaceImp struct{}

var ImpKPIMonthTarget KPIMonthTargetInterface = &KPIMonthTargetInterfaceImp{}

var _ KPIMonthTargetInterface = (*KPIMonthTargetInterfaceImp)(nil)
