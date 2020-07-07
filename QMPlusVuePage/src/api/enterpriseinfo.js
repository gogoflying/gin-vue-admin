import service from '@/utils/request'

// @Tags EnterpriseInfo
// @Summary 创建EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "创建EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enterprise/createEnterpriseInfo [post]
export const createEnterpriseInfo = (data) => {
     return service({
         url: "/enterprise/createEnterpriseInfo",
         method: 'post',
         data
     })
 }


// @Tags EnterpriseInfo
// @Summary 删除EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "删除EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enterprise/deleteEnterpriseInfo [post]
 export const deleteEnterpriseInfo = (data) => {
     return service({
         url: "/enterprise/deleteEnterpriseInfo",
         method: 'post',
         data
     })
 }

// @Tags EnterpriseInfo
// @Summary 更新EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "更新EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /enterprise/updateEnterpriseInfo [post]
 export const updateEnterpriseInfo = (data) => {
     return service({
         url: "/enterprise/updateEnterpriseInfo",
         method: 'post',
         data
     })
 }


// @Tags EnterpriseInfo
// @Summary 用id查询EnterpriseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userJobs.EnterpriseInfo true "用id查询EnterpriseInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /enterprise/findEnterpriseInfo [post]
 export const findEnterpriseInfo = (data) => {
     return service({
         url: "/enterprise/findEnterpriseInfo",
         method: 'post',
         data
     })
 }


// @Tags EnterpriseInfo
// @Summary 分页获取EnterpriseInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取EnterpriseInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enterprise/getEnterpriseInfoList [post]
 export const getEnterpriseInfoList = (data) => {
     return service({
         url: "/enterprise/getEnterpriseInfoList",
         method: 'post',
         data
     })
 }

//  export const getEnterpriseAllInfo = (data) => {
//     return service({
//         url: "/enterprise/getEnterpriseAllInfo",
//         method: 'post',
//         data
//     })
// }

export const getEnterpriseListBySearch = (params) => {
    return service({
        url: "/enterprise/getEnterpriseListBySearch",
        method: 'post',
        params
    })
}


export const getEnterpriseOptions = (data) => {
    return service({
        url: "/enterprise/getEnterpriseOptions",
        method: 'post',
        data
    })
}