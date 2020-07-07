import service from '@/utils/request'

// @Tags ResumeStatus
// @Summary 创建ResumeStatus
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resumeStatus.ResumeStatus true "创建ResumeStatus"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jst/createResumeStatus [post]
export const createResumeStatus = (data) => {
     return service({
         url: "/jst/createResumeStatus",
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
// @Router /jst/deleteResumeStatus [post]
 export const deleteResumeStatus = (data) => {
     return service({
         url: "/jst/deleteResumeStatus",
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
// @Router /jst/updateResumeStatus [post]
 export const updateResumeStatus = (data) => {
     return service({
         url: "/jst/updateResumeStatus",
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
// @Router /jst/findResumeStatus [post]
 export const findResumeStatus = (data) => {
     return service({
         url: "/jst/findResumeStatus",
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
// @Router /jst/getResumeStatusList [post]
 export const getResumeStatusList = (data) => {
     return service({
         url: "/jst/getResumeStatusList",
         method: 'post',
         data
     })
 }

 export const updateResumeMemo = (data) => {
    return service({
        url: "/jst/updateResumeMemo",
        method:'post',
        data
    })
}

export const updateResumeRemark = (data) => {
    return service({
        url: "/jst/updateResumeRemark",
        method:'post',
        data
    })
}

export const findUsersByOpenid = (data) => {
    return service({
        url: "/jst/findUsersByOpenid",
        method: 'post',
        data
    })
}