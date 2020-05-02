package registerTable

import (
	"gin-vue-admin/init/initlog"
	"gin-vue-admin/model/dbModel"
	"gin-vue-admin/model/sysModel"
	"gin-vue-admin/model/userCity"
	"gin-vue-admin/model/userJobs"
	"gin-vue-admin/model/userSalary"

	"github.com/jinzhu/gorm"
)

//注册数据库表专用
func RegisterTable(db *gorm.DB) {
	db.AutoMigrate(
		userSalary.SalaryUsers{},
		userSalary.Salarys{},
		userSalary.UserNews{},
		userSalary.SalaryContract{},
		userJobs.JobCollect{},
		userJobs.EnterpriseCollect{},
		userJobs.JobDutyTime{},
		userJobs.JobSalary{},
		userJobs.JobWorkExpire{},
		userJobs.EnterpriseIndust{},
		userJobs.EnterpriseType{},
		userJobs.EduProfession{},
		userJobs.Users{},
		userJobs.UserBaseInfo{},
		userJobs.UserDream{},
		userJobs.UserEducation{},
		userJobs.UserWork{},
		userJobs.EnterpriseInfo{},
		userJobs.ResumeStatus{},
		userCity.Cityname{},
		userJobs.Joblist{},
		userJobs.UserAuth{},
		sysModel.SysUser{},
		sysModel.SysAuthority{},
		sysModel.SysMenu{},
		sysModel.SysApi{},
		sysModel.SysBaseMenu{},
		sysModel.JwtBlacklist{},
		sysModel.SysWorkflow{},
		sysModel.SysWorkflowStepInfo{},
		dbModel.ExaFileUploadAndDownload{},
		dbModel.ExaFile{},
		dbModel.ExaFileChunk{},
		dbModel.ExaCustomer{},
	)
	log.L.Debug("register table success")
}
