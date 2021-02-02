import WeddingPage from './pages/weddingPage/WeddingPage'
import HomePage from './pages/homePage/HomePage'
import LoginPage from './pages/loginPage/LoginPage'


let pathMap = [
    { path: "/home", name: "App", component: HomePage ,auth:true},
    { path: "/login", name: "LoginPage", component: LoginPage ,auth:false},
    { path: "/weddingPage", name: "WeddingPage", component: WeddingPage, auth:true}
]

export default pathMap

