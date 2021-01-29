import React,{ Component } from 'react'
import PropTypes from 'prop-types'

import {
    Form, Input, Button, Checkbox,
  } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';



const USER_NAME = 'userName';
const PASSWORD = 'password';
const REMEMBER = 'remember';
const SUBMIT_FORM = 'submitForm';
const FORGOT_FUNC = 'forgetClick';
const REGISTER_URL = 'registerUrl';

class LoginForm extends Component {
    onFinish = (values) => {
      console.log('Received values of form: ', values);
    };
    render(){
      return (
      <Form onFinish={this.onFinish} className="login-form">
        <Form.Item
          name={USER_NAME}
          rules= {[
            { required: true, message: '请输入用户名!'}
          ]}
          initialValue={this.props[USER_NAME]}
        >
            <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="用户名" />
        </Form.Item>
        <Form.Item
          name={PASSWORD}
          rules={[
            { required: true, message: '请输入密码!' }
          ]}
          initialValue={this.props[PASSWORD]}
        >
          <Input prefix={<LockOutlined className="site-form-item-icon" />} type="password" placeholder="密码" />
        </Form.Item>
        <Form.Item
          name={REMEMBER}
          valuePropName='checked'
          initialValue={this.props[REMEMBER]}
        >
          <Checkbox>记住密码</Checkbox>
          <div className="login-form-forgot"  onClick={this.props[FORGOT_FUNC]}>忘记密码</div>
          <Button type="primary" htmlType="submit" className="login-form-button">
            登陆
          </Button>
          <a href={this.props[REGISTER_URL]}>注册!</a>
        </Form.Item>
      </Form>
    );}
}

//类型检查，需要传入 userName,password,remember以及回调函数
LoginForm.propTypes={
  [USER_NAME]:PropTypes.string.isRequired,
  [REMEMBER]:PropTypes.bool.isRequired,
  [SUBMIT_FORM]:PropTypes.func.isRequired,
  [FORGOT_FUNC]:PropTypes.func.isRequired,
  [REGISTER_URL]:PropTypes.string.isRequired
};

export default LoginForm;

   