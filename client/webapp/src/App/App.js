import './App.css'
import React,{ Component } from 'react'
import { BrowserRouter, Route , Switch, Redirect} from 'react-router-dom'
import {Layout,BackTop} from 'antd'
import { ConfigProvider } from 'antd'
import {connect} from 'react-redux'

import routerMap from './routerMap'
import Header from "./pages/header/Header"

class App extends Component{
  render(){
    return (
    <div className="App">
        <ConfigProvider>
          <Layout>
            <BrowserRouter >
              <Header/>
              <div className="contentPage">
                <Switch>
                  {routerMap.map((item,index) => {
                      return <Route key={index} path={item.path} exact render={props =>
                                (!item.auth? (<item.component {...props} />) : 
                                  ( this.props.valid ? <item.component {...props} /> : <Redirect to='/login' />))
                            }/>
                  })}
                  <Redirect to='/home'from="/"/>
                </Switch>
              </div>
              <Layout.Footer style={{ textAlign: 'center',background: 'rgba(255,255,255,0)'  }}>
                大学士阁 ©2019 Created by Thomas Zhu
              </Layout.Footer>
              <BackTop />
            </BrowserRouter>
          </Layout>
        </ConfigProvider>
    </div>
  );}
}

const  mapStateToProps =(state)=>{
  return {
    ...state.userReducer
  }
}
const A =connect(mapStateToProps)(App);
export default A;
