import service from '@/utils/request'

// @Tags SocialOrder
// @Summary 创建SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialOrder.SocialOrder true "创建SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /si/createSocialOrder [post]
export const createSocialOrder = (data) => {
     return service({
         url: "/si/createSocialOrder",
         method: 'post',
         data
     })
 }


// @Tags SocialOrder
// @Summary 删除SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialOrder.SocialOrder true "删除SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /si/deleteSocialOrder [post]
 export const deleteSocialOrder = (data) => {
     return service({
         url: "/si/deleteSocialOrder",
         method: 'post',
         data
     })
 }

// @Tags SocialOrder
// @Summary 更新SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialOrder.SocialOrder true "更新SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /si/updateSocialOrder [post]
 export const updateSocialOrder = (data) => {
     return service({
         url: "/si/updateSocialOrder",
         method: 'post',
         data
     })
 }


// @Tags SocialOrder
// @Summary 用id查询SocialOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body SocialOrder.SocialOrder true "用id查询SocialOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /si/findSocialOrder [post]
 export const findSocialOrder = (data) => {
     return service({
         url: "/si/findSocialOrder",
         method: 'post',
         data
     })
 }


// @Tags SocialOrder
// @Summary 分页获取SocialOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /si/getSocialOrderList [post]
 export const getSocialOrderList = (data) => {
     return service({
         url: "/si/getSocialOrderList",
         method: 'post',
         data
     })
 }