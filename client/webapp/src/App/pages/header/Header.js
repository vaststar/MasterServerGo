import React, { Component } from 'react'
import {Link} from 'react-router-dom'
import {Layout, Menu, Avatar, Row} from 'antd'
import {
    HomeOutlined,
    HeartOutlined,
  } from '@ant-design/icons';
import logo from './logo.png'
import withRouter from '../../utils/WithRouter';

class HeaderCom extends Component {
    render() {
        const HomePath = "/home"
        const WeddingPath = "/weddingPage"
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
                <Menu.Item key={HomePath}><Link to={HomePath}/><HomeOutlined />首页</Menu.Item>
                <Menu.Item key={WeddingPath}><Link to={WeddingPath}/><HeartOutlined />婚礼</Menu.Item>
                </Menu></Row>
            </Layout.Header>
        );
    }
}

export default withRouter(HeaderCom)



