import './App.css';
import { BrowserRouter, Route , Switch, Redirect} from 'react-router-dom'
import {Layout,BackTop} from 'antd'

import routerMap from './routerMap'
import { ConfigProvider } from 'antd';
import Header from "./components/header/Header"

function App() {
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
                              <item.component {...props} />
                          }/>
                })}
                <Redirect to='/'/>
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
  );
}

export default App;
