import service from '@/utils/request'

// @Tags ResumeStatus
// @Summary 创建ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resumeStatus.ResumeStatus true "创建ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobstatus/createResumeStatus [post]
export const createResumeStatus = (data) => {
     return service({
         url: "/jobstatus/createResumeStatus",
         method: 'post',
         data
     })
 }


// @Tags ResumeStatus
// @Summary 删除ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resumeStatus.ResumeStatus true "删除ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jobstatus/deleteResumeStatus [post]
 export const deleteResumeStatus = (data) => {
     return service({
         url: "/jobstatus/deleteResumeStatus",
         method: 'post',
         data
     })
 }

// @Tags ResumeStatus
// @Summary 更新ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resumeStatus.ResumeStatus true "更新ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jobstatus/updateResumeStatus [post]
 export const updateResumeStatus = (data) => {
     return service({
         url: "/jobstatus/updateResumeStatus",
         method: 'post',
         data
     })
 }


// @Tags ResumeStatus
// @Summary 用id查询ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resumeStatus.ResumeStatus true "用id查询ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jobstatus/findResumeStatus [post]
 export const findResumeStatus = (data) => {
     return service({
         url: "/jobstatus/findResumeStatus",
         method: 'post',
         data
     })
 }


// @Tags ResumeStatus
// @Summary 分页获取ResumeStatus列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取ResumeStatus列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jobstatus/getResumeStatusList [post]
 export const getResumeStatusList = (data) => {
     return service({
         url: "/jobstatus/getResumeStatusList",
         method: 'post',
         data
     })
 }

 export const updateResumeMemo = (data) => {
    return service({
        url: "/jobstatus/updateResumeMemo",
        method:'post',
        data
    })
}