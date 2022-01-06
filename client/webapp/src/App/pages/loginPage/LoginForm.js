import React,{ Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import { withRouter ,Redirect} from 'react-router-dom'
import {message} from 'antd'
import {post} from '../../utils/RequestREST'
import {updateAccessToken, updateRefreshToken, updateLoginInfo, updateValidState} from '../../../Redux/ActionReducer/user'

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

///////////////////////////////////////////////////////////////////////////////////
//loginForm
//////////////////////////////////////////////////////////////////////////////////
class LoginForm extends Component {
    onFinish = (values) => {
      this.props[SUBMIT_FORM](values)
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
          <Input prefix={<UserOutlined/>} placeholder="用户名" />
        </Form.Item>
        <Form.Item
          name={PASSWORD}
          rules={[
            { required: true, message: '请输入密码!' }
          ]}
          initialValue={this.props[PASSWORD]}
        >
          <Input prefix={<LockOutlined/>} type="password" placeholder="密码" />
        </Form.Item>
        <Form.Item>
          <Form.Item className="login-form-remember"
            name={REMEMBER}
            valuePropName='checked'
            initialValue={this.props[REMEMBER]}
          >
            <Checkbox>记住密码</Checkbox>
          </Form.Item>
          <div className="login-form-forgot"  onClick={this.props[FORGOT_FUNC]}>忘记密码</div>
        </Form.Item>
        <Form.Item>
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

/////////////////////////////////////////////////////////////////////////////////////////////////
//////   LoginComponent
/////////////////////////////////////////////////////////////////////////////////////////////////
class LoginComponent extends Component {
  state = {redirectToReferrer: false,forgetVisiable:false, from:this.props.location.state || { from: "/" }}
  render(){
    let { redirectToReferrer } = this.state;
    if (redirectToReferrer)
    {
      return <Redirect to={this.state.from} />; 
    } 
    return(
      <div>
        <LoginForm userName={this.props.loginInfo.username} password={this.props.loginInfo.remember?this.props.loginInfo.password:null} remember={true} submitForm={this.handleSubmit} 
        forgetClick={this.clickForget} registerUrl='/register/' wrappedComponentRef={(form) => {this.formRef = form}}></LoginForm>
      </div>
  );}
  componentDidMount(){
    //如果有access token，则去刷新下
    if(this.props.valid){
      return;
    }
    if(typeof this.props.accessToken !== "undefined" || this.props.accessToken !== null || this.props.accessToken !== ""){
      post(this.props.validAccessTokenUrl).then(response => response.json()).then(result => {
        if(result.code === 0){
          this.props.setValidState(true)
          this.setState({ redirectToReferrer: true });
        }else if(typeof this.props.refreshToken !== "undefined" || this.props.refreshToken !== null || this.props.refreshToken !== ""){ 
          postMessage(this.props.requestAccessTokenUrl,{'refreshToken':this.props.refreshToken,'userid':this.props.userid}).then(resp => resp.json()).then(res =>{
            if(res.code === 0){
              this.props.setAccessToken(res.data)
            }else{
              this.props.setRefreshToken(null,null)
              this.props.setAccessToken(null)
            }
          })
        }else{
          this.props.setAccessToken(null)
        }
      })
    }
  }
  handleSubmit =(form)=>{
    //请求token
    post(this.props.requestRefreshTokenUrl,{'username':form.userName,'password':form.password}).then(response => response.json()).then(result => {
      // 在此处写获取数据之后的处理逻辑
      if(result.code === 0){
        this.props.setRefreshToken(result.data.token,result.data.userid)
        this.props.setLoginInfo({'username':form.userName,'password':form.password,'remember':form.remember})
        post(this.props.requestAccessTokenUrl,{'refreshToken':result.data.token,'userid':result.data.userid}).then(resp => resp.json()).then(res =>{
          if(res.code === 0){
            this.props.setAccessToken(res.data)
            this.setState({ redirectToReferrer: true });
          }else{
            message.error('获取accessToken失败');
            this.props.setAccessToken(null)
            this.setState({ redirectToReferrer: false });
          }
        })
      }
      else{
        message.error('获取refreshToken失败');
        this.props.setRefreshToken(null,null)
        this.props.setLoginInfo({'username':form.userName,'password':null,'remember':form.remember})
        this.props.setAccessToken(null)
        this.setState({ redirectToReferrer: false });
      }
    })
  }
}

const  mapStateToProps =(state)=>{
  return {
    ...state.userReducer
  }
}

const mapDispatch =(dispatch)=>{
  return {
    setRefreshToken:(tokenStr,uid)=>{
      dispatch(updateRefreshToken(tokenStr,uid))
    },
    setAccessToken:(tokenStr)=>{
      dispatch(updateAccessToken(tokenStr))
    },
    setLoginInfo:(userInfo)=>{
      dispatch(updateLoginInfo(userInfo))
    },
    setValidState:(isValid)=>{
      dispatch(updateValidState(isValid))
    }
  }
}
export default withRouter(connect(mapStateToProps,mapDispatch)(LoginComponent))


   