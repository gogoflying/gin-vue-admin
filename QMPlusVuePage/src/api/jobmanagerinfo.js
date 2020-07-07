import service from '@/utils/request'

// @Tags Joblist
// @Summary 创建Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body RspJoblist.Joblist true "创建Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /joblist/createJoblist [post]
export const createJoblist = (data) => {
     return service({
         url: "/joblist/createJoblist",
         method: 'post',
         data
     })
 }


// @Tags Joblist
// @Summary 删除Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body RspJoblist.Joblist true "删除Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /joblist/deleteJoblist [post]
 export const deleteJoblist = (data) => {
     return service({
         url: "/joblist/deleteJoblist",
         method: 'post',
         data
     })
 }

// @Tags Joblist
// @Summary 更新Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body RspJoblist.Joblist true "更新Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /joblist/updateJoblist [post]
 export const updateJoblist = (data) => {
     return service({
         url: "/joblist/updateJoblist",
         method: 'post',
         data
     })
 }


// @Tags Joblist
// @Summary 用id查询Joblist
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body RspJoblist.Joblist true "用id查询Joblist"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /joblist/findJoblist [post]
 export const findJoblist = (data) => {
     return service({
         url: "/joblist/findJoblist",
         method: 'post',
         data
     })
 }


// @Tags Joblist
// @Summary 分页获取Joblist列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Joblist列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /joblist/getJoblistList [post]
 export const getJoblistList = (data) => {
     return service({
         url: "/joblist/getJoblistList",
         method: 'post',
         data
     })
 }

 export const getjoblistOptions = (data) => {
    return service({
        url: "/joblist/getjoblistOptions",
        method: 'post',
        data
    })
}

