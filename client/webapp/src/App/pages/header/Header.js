import React, { Component } from 'react'
import { Link, withRouter} from 'react-router-dom'
import {Layout, Menu, Avatar, Row} from 'antd'
import {
    HomeOutlined,
    HeartOutlined,
  } from '@ant-design/icons';

import logo from './logo.png'

class HeaderCom extends Component {
    getLinkToUrl= (url)=> {
        return ({
            pathname:url,
            state:{from:this.props.location.pathname}})
    };
    handleClick = (e) => {
    }
    render() {
        return (
            <Layout.Header align="center" style={{ position: 'fixed', zIndex: 1, width: '100%' }} className="app-header">
                <Row>
                <Avatar src={logo} alt="学士" size={64}></Avatar>
                <Menu onClick={this.handleClick}
                  mode="horizontal"
                  theme='dark'
                  selectedKeys={[this.props.location.pathname]}
                >
                <Menu.Item key="/home"><Link to={this.getLinkToUrl('/home')}/><HomeOutlined />首页</Menu.Item>
                <Menu.Item key='/weddingPage'><Link to={this.getLinkToUrl('/weddingPage')}/><HeartOutlined />婚礼</Menu.Item>
                </Menu></Row>
            </Layout.Header>
        );
    }
}

export default withRouter(HeaderCom)



