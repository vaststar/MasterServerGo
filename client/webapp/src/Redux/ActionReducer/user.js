const UserData = {
    accessToken:window.localStorage.getItem('accessToken'),
    refreshToken:window.localStorage.getItem('refreshToken'),
    requestAccessTokenUrl:"/rest/authenticate/requestAccessToken",
    requestRefreshTokenUrl:"/rest/authenticate/requestRefreshToken",
    userid:window.localStorage.getItem('userid'),
    valid:window.sessionStorage.getItem('valid') ==='true',
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
    type:UpdateRefreshToken,
    userInfo
}); 

export const userReducer = (state=UserData,action)=>{
    if (typeof state === 'undefined') {
      return UserData
    }
    switch (action.type) {
        case UpdateAccessToken:
            window.localStorage.setItem('accessToken',action.token)
            if(typeof action.token == "undefined" || action.token == null || action.token === ""){
                window.sessionStorage.setItem('valid', 'false')
            }else{
                window.sessionStorage.setItem('valid', 'true')
            }
            return {...state,accessToken:action.token};
        case UpdateRefreshToken:
            window.localStorage.setItem('refreshToken',action.token)
            window.localStorage.setItem('userid',action.userid)
            return {...state, refreshToken:action.token, userid:action.userid};
        case UpdateLoginInfo:
            window.localStorage.setItem('loginInfo',JSON.stringify(action.userInfo))
            return {...state,loginInfo:action.userInfo};
        default:
            return {...state};
    }
}
