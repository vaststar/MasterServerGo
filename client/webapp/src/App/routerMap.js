import WeddingPage from './pages/weddingPage/WeddingPage'
import HomePage from './pages/homePage/HomePage'
import LoginPage from './pages/loginPage/LoginPage'
import VideoPage from './pages/videoPage/VideoPage'
import {RoutePath} from "./define/dataDefine"


let pathMap = [
    { path: RoutePath.home, name: "App", component: HomePage, auth:true},
    { path: RoutePath.login, name: "LoginPage", component: LoginPage, auth:false},
    { path: RoutePath.wedding, name: "WeddingPage", component: WeddingPage, auth:true},
    { path: RoutePath.video, name: "VideoPage", component: VideoPage, auth:false}
]

export default pathMap

