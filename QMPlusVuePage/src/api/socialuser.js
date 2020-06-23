import service from '@/utils/request'

// @Tags SocialUser
// @Summary 创建SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialUser.SocialUser true "创建SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /si/createSocialUser [post]
export const createSocialUser = (data) => {
     return service({
         url: "/si/createSocialUser",
         method: 'post',
         data
     })
 }


// @Tags SocialUser
// @Summary 删除SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialUser.SocialUser true "删除SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /si/deleteSocialUser [post]
 export const deleteSocialUser = (data) => {
     return service({
         url: "/si/deleteSocialUser",
         method: 'post',
         data
     })
 }

// @Tags SocialUser
// @Summary 更新SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialUser.SocialUser true "更新SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /si/updateSocialUser [post]
 export const updateSocialUser = (data) => {
     return service({
         url: "/si/updateSocialUser",
         method: 'post',
         data
     })
 }


// @Tags SocialUser
// @Summary 用id查询SocialUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialUser.SocialUser true "用id查询SocialUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /si/findSocialUser [post]
 export const findSocialUser = (data) => {
     return service({
         url: "/si/findSocialUser",
         method: 'post',
         data
     })
 }


// @Tags SocialUser
// @Summary 分页获取SocialUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /si/getSocialUserList [post]
 export const getSocialUserList = (data) => {
     return service({
         url: "/si/getSocialUserList",
         method: 'post',
         data
     })
 }