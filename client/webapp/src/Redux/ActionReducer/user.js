const UserData = {
    accessToken:window.localStorage.getItem('accessToken'),
    refreshToken:window.localStorage.getItem('refreshToken'),
    requestAccessTokenUrl:"/rest/authenticate/requestAccessToken",
    requestRefreshTokenUrl:"/rest/authenticate/requestRefreshToken",
    userid:window.localStorage.getItem('userid'),
    valid:window.sessionStorage.getItem('valid') ==='true'
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

export const userReducer = (state=UserData,action)=>{
    if (typeof state === 'undefined') {
      return UserData
    }
    switch (action.type) {
        case UpdateAccessToken:
            window.localStorage.setItem('accessToken',action.token)
            if(typeof action.token == "undefined" || action.token == null || action.token == ""){
                window.sessionStorage.setItem('valid', 'false')
            }else{
                window.sessionStorage.setItem('valid', 'true')
            }
            return {...state,accessToken:action.token};
        case UpdateRefreshToken:
            window.localStorage.setItem('refreshToken',action.token)
            window.localStorage.setItem('userid',action.userid)
            return {...state, refreshToken:action.token, userid:action.userid};
        default:
            return {...state};
    }
}
