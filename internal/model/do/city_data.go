// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CityData is the golang structure of table city_data for DAO operations like Where/Data.
type CityData struct {
	g.Meta    `orm:"table:city_data, do:true"`
	Id        interface{} //
	Name      interface{} // 名字
	Code      interface{} // 编码
	ParentId  interface{} // 父ID
	Sort      interface{} // 排序
	Status    interface{} // 状态;0:禁用;1:正常
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建日期
	UpdatedBy interface{} // 更新者
	UpdatedAt *gtime.Time // 修改日期
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}