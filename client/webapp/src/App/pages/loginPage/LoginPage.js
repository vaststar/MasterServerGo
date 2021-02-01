import React,{ Component } from 'react'
import {Row,Col} from 'antd'

import LoginForm from './LoginForm'

class LoginPage extends Component {
  render() {
    return(
    <div >
      <Row type="flex" justify="space-around" align="middle">
          <Col span={8} className="login_page_loginForm">
            <LoginForm></LoginForm>
          </Col>
      </Row> 
    </div>
    );}
}

export default LoginPage