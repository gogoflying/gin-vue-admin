import service from '@/utils/request'

// @Tags SocialInsurance
// @Summary 创建SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "创建SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /social/createSocialInsurance [post]
export const createSocialInsurance = (data) => {
     return service({
         url: "/social/createSocialInsurance",
         method: 'post',
         data
     })
 }


// @Tags SocialInsurance
// @Summary 删除SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "删除SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /social/deleteSocialInsurance [post]
 export const deleteSocialInsurance = (data) => {
     return service({
         url: "/social/deleteSocialInsurance",
         method: 'post',
         data
     })
 }

// @Tags SocialInsurance
// @Summary 更新SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "更新SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /social/updateSocialInsurance [post]
 export const updateSocialInsurance = (data) => {
     return service({
         url: "/social/updateSocialInsurance",
         method: 'post',
         data
     })
 }


// @Tags SocialInsurance
// @Summary 用id查询SocialInsurance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body socialInsurance.SocialInsurance true "用id查询SocialInsurance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /social/findSocialInsurance [post]
 export const findSocialInsurance = (data) => {
     return service({
         url: "/social/findSocialInsurance",
         method: 'post',
         data
     })
 }


// @Tags SocialInsurance
// @Summary 分页获取SocialInsurance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取SocialInsurance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /social/getSocialInsuranceList [post]
 export const getSocialInsuranceList = (data) => {
     return service({
         url: "/social/getSocialInsuranceList",
         method: 'post',
         data
     })
 }

 export const getSocialInsuranceOptions = (data) => {
    return service({
        url: "/social/getSocialInsuranceOptions",
        method: 'post',
        data
    })
}