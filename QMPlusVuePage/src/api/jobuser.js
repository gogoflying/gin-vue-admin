import service from '@/utils/request'


export const createUsers = (data) => {
     return service({
         url: "/users/createUsers",
         method: 'post',
         data
     })
 }

 export const deleteUsers = (data) => {
     return service({
         url: "/users/deleteUsers",
         method: 'post',
         data
     })
 }

 export const updateUsers = (data) => {
     return service({
         url: "/users/updateUsers",
         method: 'post',
         data
     })
 }


 export const findJobUserById = (data) => {
     return service({
         url: "/users/findJobUserById",
         method: 'post',
         data
     })
 }

 export const getUsersList = (data) => {
     return service({
         url: "/users/getUsersList",
         method: 'post',
         data
     })
 }

 export const findUsersByOpenid = (data) => {
    return service({
        url: "/users/findUsersByOpenid",
        method: 'post',
        data
    })
}

export const getUserOptions = (data) => {
    return service({
        url: "/users/getUserOptions",
        method: 'post',
        data
    })
}

export const getUserResumeList = (data) => {
    return service({
        url: "/users/getUserResumeList",
        method: 'post',
        data
    })
}

export const createUserBaseInfo = (data) => {
    return service({
        url: "/bf/createUserBaseInfo",
        method: 'post',
        data
    })
}

export const updateUserBaseInfo = (data) => {
    return service({
        url: "/bf/updateUserBaseInfo",
        method: 'post',
        data
    })
}

export const createUserDream = (data) => {
    return service({
        url: "/dm/createUserDream",
        method: 'post',
        data
    })
}

export const updateUserDream = (data) => {
    return service({
        url: "/dm/updateUserDream",
        method: 'post',
        data
    })
}

export const createUserWork = (data) => {
    return service({
        url: "/wk/createUserWork",
        method: 'post',
        data
    })
}

export const deleteUserWork = (data) => {
    return service({
        url: "/wk/deleteUserWork",
        method: 'post',
        data
    })
}

export const updateUserWork = (data) => {
    return service({
        url: "/wk/updateUserWork",
        method: 'post',
        data
    })
}

export const createUsereducation = (data) => {
    return service({
        url: "/ed/createUsereducation",
        method: 'post',
        data
    })
}

export const deleteUsereducation = (data) => {
    return service({
        url: "/ed/deleteUsereducation",
        method: 'post',
        data
    })
}

export const updateUsereducation = (data) => {
    return service({
        url: "/ed/updateUsereducation",
        method: 'post',
        data
    })
}