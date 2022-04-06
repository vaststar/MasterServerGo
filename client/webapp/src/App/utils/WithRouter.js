import {useLocation, useNavigate} from "react-router"
import React from "react"
export default function withRouter(Child){
    return (props)=>{
        return <Child {...props} navigate={useNavigate()} location={useLocation()}/>
    }
}