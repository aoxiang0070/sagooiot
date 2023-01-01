// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/sagoo-cloud/sagooiot/internal/dao/internal"
)

// internalSysOrganizationDao is internal type for wrapping internal DAO implements.
type internalSysOrganizationDao = *internal.SysOrganizationDao

// sysOrganizationDao is the data access object for table sys_organization.
// You can define custom methods on it to extend its functionality as you wish.
type sysOrganizationDao struct {
	internalSysOrganizationDao
}

var (
	// SysOrganization is globally public accessible object for table sys_organization operations.
	SysOrganization = sysOrganizationDao{
		internal.NewSysOrganizationDao(),
	}
)

// Fill with you ideas below.