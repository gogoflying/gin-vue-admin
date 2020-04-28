import service from '@/utils/request'

// @Tags Cityname
// @Summary 创建Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "创建Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityname/createCityname [post]
export const createCityname = (data) => {
     return service({
         url: "/cityname/createCityname",
         method: 'post',
         data
     })
 }


// @Tags Cityname
// @Summary 删除Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "删除Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cityname/deleteCityname [post]
 export const deleteCityname = (data) => {
     return service({
         url: "/cityname/deleteCityname",
         method: 'post',
         data
     })
 }

// @Tags Cityname
// @Summary 更新Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "更新Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cityname/updateCityname [post]
 export const updateCityname = (data) => {
     return service({
         url: "/cityname/updateCityname",
         method: 'post',
         data
     })
 }


// @Tags Cityname
// @Summary 用id查询Cityname
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userCity.Cityname true "用id查询Cityname"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cityname/findCityname [post]
 export const findCityname = (data) => {
     return service({
         url: "/cityname/findCityname",
         method: 'post',
         data
     })
 }


// @Tags Cityname
// @Summary 分页获取Cityname列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取Cityname列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityname/getCitynameList [post]
 export const getCitynameList = (data) => {
     return service({
         url: "/cityname/getCitynameList",
         method: 'post',
         data
     })
 }