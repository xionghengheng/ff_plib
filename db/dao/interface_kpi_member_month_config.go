package dao

import "github.com/xionghengheng/ff_plib/db/model"

// KPIMemberMonthConfigInterface 每人每月指标配置接口。
// 月份、数值、人员校验以及快照和时间字段由业务层负责。
type KPIMemberMonthConfigInterface interface {
	// AddKPIMemberMonthConfig 新增配置，ID 回填；重复月份+人员由数据库唯一键报错。
	AddKPIMemberMonthConfig(record *model.KPIMemberMonthConfigModel) error
	// UpdateKPIMemberMonthConfig 按月份、人员更新，map 值为 nil 可清空可空指标，0 可写入零值。
	UpdateKPIMemberMonthConfig(statMonth string, memberID int64, updates map[string]interface{}) error
	// GetKPIMemberMonthConfig 查询某人某月配置，未找到返回 gorm.ErrRecordNotFound。
	GetKPIMemberMonthConfig(statMonth string, memberID int64) (*model.KPIMemberMonthConfigModel, error)
	// GetKPIMemberMonthConfigList 查询当月全部人员配置，按 member_id 升序。
	GetKPIMemberMonthConfigList(statMonth string) ([]model.KPIMemberMonthConfigModel, error)
}

type KPIMemberMonthConfigInterfaceImp struct{}

var ImpKPIMemberMonthConfig KPIMemberMonthConfigInterface = &KPIMemberMonthConfigInterfaceImp{}

var _ KPIMemberMonthConfigInterface = (*KPIMemberMonthConfigInterfaceImp)(nil)
