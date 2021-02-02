const UserData = {
    accessToken:window.localStorage.getItem('accessToken'),
    refreshToken:window.localStorage.getItem('refreshToken'),
    requestAccessTokenUrl:"/rest/authenticate/requestAccessToken",
    requestRefreshTokenUrl:"/rest/authenticate/requestRefreshToken",
    validAccessTokenUrl:"/rest/authenticate/isAccessTokenValid",
    userid:window.localStorage.getItem('userid'),
    valid:window.sessionStorage.getItem('valid')==='true',
    loginInfo:JSON.parse(window.localStorage.getItem('loginInfo'))||{
        username:'',
        password:'',
        remember:true
    },
}

export const UpdateAccessToken = "updateAccessToken";
export const updateAccessToken = (tokenStr)=>({
    type:UpdateAccessToken,
    token:tokenStr
});

export const UpdateRefreshToken = 'updateRefreshToken';
export const updateRefreshToken = (tokenStr,uid)=>({
    type:UpdateRefreshToken,
    token:tokenStr,
    userid:uid
});

export const UpdateLoginInfo = 'updateLoginInfo';
export const updateLoginInfo = (userInfo)=>({
    type:UpdateLoginInfo,
    userInfo
}); 

export const UpdateValidState = "updateValidState";
export const updateValidState = (isValid)=>({
    type:UpdateValidState,
    isValid
}); 


export const userReducer = (state=UserData,action)=>{
    if (typeof state === 'undefined') {
      return UserData
    }
    switch (action.type) {
        case UpdateAccessToken:
            let isValid = 'false'
            if(typeof action.token === "undefined" || action.token === null || action.token === ""){
                isValid = 'false'
            }else{
                isValid = 'true'
            }
            window.localStorage.setItem('accessToken',action.token)
            window.sessionStorage.setItem('valid', isValid)
            return {...state,accessToken:action.token, valid:isValid};
        case UpdateRefreshToken:
            window.localStorage.setItem('refreshToken',action.token)
            window.localStorage.setItem('userid',action.userid)
            return {...state, refreshToken:action.token, userid:action.userid};
        case UpdateLoginInfo:
            window.localStorage.setItem('loginInfo',JSON.stringify(action.userInfo))
            return {...state,loginInfo:action.userInfo};
        case UpdateValidState:
            console.log("param isvalid:",action.isValid)
            window.sessionStorage.setItem('valid',action.isValid)
            console.log("isvalid:",window.sessionStorage.getItem('valid'))
            return {...state, valid:action.isValid};
        default:
            return {...state};
    }
}
