import React,{ Component } from 'react'
import {connect} from 'react-redux'
import { withRouter ,Redirect} from 'react-router-dom'
import {message} from 'antd'

import {get} from '../../utils/RequestREST'
import LoginForm from './LoginForm'

class Login extends Component {
  state = {redirectToReferrer: false,forgetVisiable:false}
  render() {
    const {username,password,remember} = this.props;
    let { from } = this.props.location.state || { from: "/" };
    let { redirectToReferrer } = this.state;
    
    if (redirectToReferrer)
    {
      return <Redirect to={from} />; 
    } 
    return (
      <div>
      <LoginForm userName={username} password={remember?password:null} remember={true} submitForm={this.handleSubmit} 
       forgetClick={this.clickForget} registerUrl='/register/' wrappedComponentRef={(form) => {this.formRef = form}}></LoginForm>
      {/* <Modal visible={this.state.forgetVisiable} footer={null} onCancel={()=>{this.setState({ forgetVisiable: false })}}>
      <div className="popForgetCom">
      {this.state.userid&&<ForgetPassCom userid={this.state.userid} okFunc={()=>{this.setState({ forgetVisiable: false });message.info('密码重置成功');}}></ForgetPassCom>}
      </div>
      </Modal> */}
      </div>
    );
  }
  clickForget=()=>{
    get(this.props.userUrl+"/userids/"+this.formRef.getUserName()).then(respones=>respones.json()).then(result=>{
        if(result.status){
          this.setState({userid:result.data});
          this.setState({ forgetVisiable: true})
        }
        else{
          message.error('该账户不存在');
        }
    })
  }
  handleSubmit =(form)=>{
    //请求token
    // post(this.props.userUrl+"/tokens/",{'username':form.userName,'password':form.password}).then(response => response.json()).then(result => {
    //   // 在此处写获取数据之后的处理逻辑
    //   if(result.status){
    //     this.props.ChangeUser({'username':form.userName,'password':form.password,'remember':form.remember})
    //     this.props.ChangeToken(result.data);
    //     this.props.ChangeValid(true);
    //     this.setState({ redirectToReferrer: true });
    //   }
    //   else{
    //     message.error('用户名或者密码错误');
    //     this.props.ChangeUser({'username':form.userName,'password':null,'remember':form.remember})
    //     this.setState({ redirectToReferrer: false });
    //     this.props.ChangeToken(null);
    //     this.props.ChangeValid(false);
    //   }
    // })
  }
}

const  mapStateToProps =(state)=>{
  return {
    ...state.userReducer.user,userUrl:state.userReducer.userUrl
  }
}

const mapDispatch =(dispatch)=>{
  return {
  }
}

export default withRouter(connect(mapStateToProps,mapDispatch)(Login))