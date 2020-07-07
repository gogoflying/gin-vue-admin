import service from '@/utils/request'

// @Tags UserNews
// @Summary 创建UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "创建UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/createUserNews [post]
export const createUserNews = (data) => {
     return service({
         url: "/news/createUserNews",
         method: 'post',
         data
     })
 }


// @Tags UserNews
// @Summary 删除UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "删除UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteUserNews [post]
 export const deleteUserNews = (data) => {
     return service({
         url: "/news/deleteUserNews",
         method: 'post',
         data
     })
 }

// @Tags UserNews
// @Summary 更新UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "更新UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /news/updateUserNews [post]
 export const updateUserNews = (data) => {
     return service({
         url: "/news/updateUserNews",
         method: 'post',
         data
     })
 }


// @Tags UserNews
// @Summary 用id查询UserNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userSalary.UserNews true "用id查询UserNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /news/findUserNews [post]
 export const findUserNews = (data) => {
     return service({
         url: "/news/findUserNews",
         method: 'post',
         data
     })
 }


// @Tags UserNews
// @Summary 分页获取UserNews列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取UserNews列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/getUserNewsList [post]
 export const getUserNewsList = (data) => {
     return service({
         url: "/news/getUserNewsList",
         method: 'post',
         data
     })
 }