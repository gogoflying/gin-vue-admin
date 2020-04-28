import service from '@/utils/request'

// @Tags UserAuth
// @Summary 创建UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "创建UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userauth/createUserAuth [post]
export const createUserAuth = (data) => {
     return service({
         url: "/userauth/createUserAuth",
         method: 'post',
         data
     })
 }


// @Tags UserAuth
// @Summary 删除UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "删除UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userauth/deleteUserAuth [post]
 export const deleteUserAuth = (data) => {
     return service({
         url: "/userauth/deleteUserAuth",
         method: 'post',
         data
     })
 }

// @Tags UserAuth
// @Summary 更新UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "更新UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userauth/updateUserAuth [post]
 export const updateUserAuth = (data) => {
     return service({
         url: "/userauth/updateUserAuth",
         method: 'post',
         data
     })
 }


// @Tags UserAuth
// @Summary 用id查询UserAuth
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userauth.UserAuth true "用id查询UserAuth"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userauth/findUserAuth [post]
 export const findUserAuth = (data) => {
     return service({
         url: "/userauth/findUserAuth",
         method: 'post',
         data
     })
 }


// @Tags UserAuth
// @Summary 分页获取UserAuth列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UserAuth列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userauth/getUserAuthList [post]
 export const getUserAuthList = (data) => {
     return service({
         url: "/userauth/getUserAuthList",
         method: 'post',
         data
     })
 }