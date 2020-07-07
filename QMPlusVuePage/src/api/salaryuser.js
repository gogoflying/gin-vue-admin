import service from '@/utils/request'

// @Tags SalaryUsers
// @Summary 创建SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "创建SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /salary/createSalaryUsers [post]
export const createSalaryUsers = (data) => {
     return service({
         url: "/salary/createSalaryUsers",
         method: 'post',
         data
     })
 }


// @Tags SalaryUsers
// @Summary 删除SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "删除SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /salary/deleteSalaryUsers [post]
 export const deleteSalaryUsers = (data) => {
     return service({
         url: "/salary/deleteSalaryUsers",
         method: 'post',
         data
     })
 }

// @Tags SalaryUsers
// @Summary 更新SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "更新SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /salary/updateSalaryUsers [post]
 export const updateSalaryUsers = (data) => {
     return service({
         url: "/salary/updateSalaryUsers",
         method: 'post',
         data
     })
 }


// @Tags SalaryUsers
// @Summary 用id查询SalaryUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.SalaryUsers true "用id查询SalaryUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /salary/findSalaryUsers [post]
 export const findSalaryUsers = (data) => {
     return service({
         url: "/salary/findSalaryUsers",
         method: 'post',
         data
     })
 }


// @Tags SalaryUsers
// @Summary 分页获取SalaryUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SalaryUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /salary/getSalaryUsersList [post]
 export const getSalaryUsersList = (data) => {
     return service({
         url: "/salary/getSalaryUsersList",
         method: 'post',
         data
     })
 }

 export const uploadUserContract = (data) => {
    return service({
        url: "/salary/uploadUserContract",
        method: 'post',
        data
    })
}

export const downloadContractList = (data) => {
    return service({
        url: "/salary/downloadContractList",
        method: 'post',
        data
    })
}