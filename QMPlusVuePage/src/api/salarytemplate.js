import service from '@/utils/request'

// @Tags SalaryTemplates
// @Summary 创建SalaryTemplates
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryTemplates true "创建SalaryTemplates"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /un/createSalaryTemplates [post]
export const createSalaryTemplates = (data) => {
     return service({
         url: "/un/createSalaryTemplates",
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
// @Router /un/deleteSalaryTemplates [post]
 export const deleteSalaryTemplates = (data) => {
     return service({
         url: "/un/deleteSalaryTemplates",
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
// @Router /un/updateSalaryTemplates [post]
 export const updateSalaryTemplates = (data) => {
     return service({
         url: "/un/updateSalaryTemplates",
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
// @Router /un/findSalaryTemplates [post]
 export const findSalaryTemplates = (data) => {
     return service({
         url: "/un/findSalaryTemplates",
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
// @Router /un/getSalaryTemplatesList [post]
 export const getSalaryTemplatesList = (data) => {
     return service({
         url: "/un/getSalaryTemplatesList",
         method: 'post',
         data
     })
 }

 export const getSalaryTemplatesListBySearch = (params) => {
    return service({
        url: "/un/getSalaryTemplatesListBySearch",
        method: 'post',
        params
    })
}