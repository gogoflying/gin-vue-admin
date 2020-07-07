import service from '@/utils/request'

// @Tags Salarys
// @Summary 创建Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "创建Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /salary/createSalarys [post]
export const createSalarys = (data) => {
     return service({
         url: "/salary/createSalarys",
         method: 'post',
         data
     })
 }


// @Tags Salarys
// @Summary 删除Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "删除Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /salary/deleteSalarys [post]
 export const deleteSalarys = (data) => {
     return service({
         url: "/salary/deleteSalarys",
         method: 'post',
         data
     })
 }

// @Tags Salarys
// @Summary 更新Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "更新Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /salary/updateSalarys [post]
 export const updateSalarys = (data) => {
     return service({
         url: "/salary/updateSalarys",
         method: 'post',
         data
     })
 }


// @Tags Salarys
// @Summary 用id查询Salarys
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.Salarys true "用id查询Salarys"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /salary/findSalarys [post]
 export const findSalarys = (data) => {
     return service({
         url: "/salary/findSalarys",
         method: 'post',
         data
     })
 }


// @Tags Salarys
// @Summary 分页获取Salarys列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Salarys列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /salary/getSalarysList [post]
 export const getSalarysList = (data) => {
     return service({
         url: "/salary/getSalarysList",
         method: 'post',
         data
     })
 }

 export const batchdeleteSalarys = (data) => {
    return service({
        url: "/salary/batchdeleteSalarys",
        method: 'post',
        data
    })
}