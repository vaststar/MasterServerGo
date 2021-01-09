import WeddingPage from './pages/wedding/WeddingPage'
import HomePage from './pages/homepage/HomePage'

export default [
    { path: "/", name: "App", component: HomePage ,auth:false},
    { path: "/weddingPage/", name: "WeddingPage", component: WeddingPage }
    
]

