import service from '@/utils/request'


export const createUsers = (data) => {
     return service({
         url: "/jobuser/createUsers",
         method: 'post',
         data
     })
 }

 export const deleteUsers = (data) => {
     return service({
         url: "/jobuser/deleteUsers",
         method: 'post',
         data
     })
 }

 export const updateUsers = (data) => {
     return service({
         url: "/jobuser/updateUsers",
         method: 'post',
         data
     })
 }


 export const findJobUserById = (data) => {
     return service({
         url: "/jobuser/findJobUserById",
         method: 'post',
         data
     })
 }

 export const getUsersList = (data) => {
     return service({
         url: "/jobuser/getUsersList",
         method: 'post',
         data
     })
 }

export const findUsersResumeInfo = (data) => {
    return service({
        url: "/jobuser/findUsersResumeInfo",
        method: 'post',
        data
    })
}

export const getUserOptions = (data) => {
    return service({
        url: "/jobuser/getUserOptions",
        method: 'post',
        data
    })
}

export const getUserResumeList = (data) => {
    return service({
        url: "/jobuser/getUserResumeList",
        method: 'post',
        data
    })
}

export const getContactInfo = (data) => {
    return service({
        url: "/jobuser/getContactInfo",
        method: 'get',
        data
    })
}

export const createUserBaseInfo = (data) => {
    return service({
        url: "/jobuser/createUserBaseInfo",
        method: 'post',
        data
    })
}

export const updateUserBaseInfo = (data) => {
    return service({
        url: "/jobuser/updateUserBaseInfo",
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
        url: "/jobuser/updateUserDream",
        method: 'post',
        data
    })
}

export const createUserWork = (data) => {
    return service({
        url: "/jobuser/createUserWork",
        method: 'post',
        data
    })
}

export const deleteUserWork = (data) => {
    return service({
        url: "/jobuser/deleteUserWork",
        method: 'post',
        data
    })
}

export const updateUserWork = (data) => {
    return service({
        url: "/jobuser/updateUserWork",
        method: 'post',
        data
    })
}

export const createUsereducation = (data) => {
    return service({
        url: "/jobuser/createUsereducation",
        method: 'post',
        data
    })
}

export const deleteUsereducation = (data) => {
    return service({
        url: "/jobuser/deleteUsereducation",
        method: 'post',
        data
    })
}

export const updateUsereducation = (data) => {
    return service({
        url: "/jobuser/updateUsereducation",
        method: 'post',
        data
    })
}