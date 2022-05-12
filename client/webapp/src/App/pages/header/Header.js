import React, { Component } from 'react'
import {Link} from 'react-router-dom'
import {Layout, Menu, Avatar, Row} from 'antd'
import {
    HomeOutlined,
    HeartOutlined,
    VideoCameraOutlined,
  } from '@ant-design/icons';
import logo from './logo.png'
import withRouter from '../../utils/WithRouter';
import { RoutePath } from '../../define/dataDefine';

class HeaderCom extends Component {
    render() {
        console.log("test:",this.props.location)
        return (
            <Layout.Header align="center" style={{ position: 'fixed', zIndex: 1, width: '100%' }} className="app-header">
                <Row>
                <Avatar src={logo} alt="学士" size={64}></Avatar>
                <Menu onClick={this.handleClick}
                  mode="horizontal"
                  theme='dark'
                  selectedKeys={[this.props.location.pathname]}
                >
                <Menu.Item key={RoutePath.home}><Link to={RoutePath.home}/><HomeOutlined /> 首页</Menu.Item>
                <Menu.Item key={RoutePath.wedding}><Link to={RoutePath.wedding}/><HeartOutlined /> 婚礼</Menu.Item>
                <Menu.Item key={RoutePath.video}><Link to={RoutePath.video}/><VideoCameraOutlined /> 视频</Menu.Item>
                </Menu></Row>
            </Layout.Header>
        );
    }
}

export default withRouter(HeaderCom)



