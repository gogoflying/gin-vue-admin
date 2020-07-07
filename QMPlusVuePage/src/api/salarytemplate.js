import service from '@/utils/request'

// @Tags SalaryTemplates
// @Summary 创建SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "创建SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /salary/createSalaryTemplates [post]
export const createSalaryTemplates = (data) => {
     return service({
         url: "/salary/createSalaryTemplates",
         method: 'post',
         data
     })
 }


// @Tags SalaryTemplates
// @Summary 删除SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "删除SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /salary/deleteSalaryTemplates [post]
 export const deleteSalaryTemplates = (data) => {
     return service({
         url: "/salary/deleteSalaryTemplates",
         method: 'post',
         data
     })
 }

// @Tags SalaryTemplates
// @Summary 更新SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "更新SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /salary/updateSalaryTemplates [post]
 export const updateSalaryTemplates = (data) => {
     return service({
         url: "/salary/updateSalaryTemplates",
         method: 'post',
         data
     })
 }


// @Tags SalaryTemplates
// @Summary 用id查询SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "用id查询SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /salary/findSalaryTemplates [post]
 export const findSalaryTemplates = (data) => {
     return service({
         url: "/salary/findSalaryTemplates",
         method: 'post',
         data
     })
 }


// @Tags SalaryTemplates
// @Summary 分页获取SalaryTemplates列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SalaryTemplates列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /salary/getSalaryTemplatesList [post]
 export const getSalaryTemplatesList = (data) => {
     return service({
         url: "/salary/getSalaryTemplatesList",
         method: 'post',
         data
     })
 }

 export const getSalaryTemplatesListBySearch = (params) => {
    return service({
        url: "/salary/getSalaryTemplatesListBySearch",
        method: 'post',
        params
    })
}