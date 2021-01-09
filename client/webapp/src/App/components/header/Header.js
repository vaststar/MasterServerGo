import React, { Component } from 'react'
import { Link, withRouter} from 'react-router-dom'
import {Layout, Menu, Avatar, Row} from 'antd'
import {
    HomeOutlined,
    HeartOutlined,
  } from '@ant-design/icons';

import logo from './logo.png'

class HeaderCom extends Component {
    state = {
        current: '/',
      };
    handleClick = e => {
        this.setState({ current: e.key });
    };
    render() {
        const { current } = this.state;
        return (
            <Layout.Header align="center" style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
                <Row>
                <Avatar src={logo} alt="学士" size={64}></Avatar>
                <Menu onClick={this.handleClick}
                  mode="horizontal"
                  theme='dark'
                  selectedKeys={[current]}
                >
                <Menu.Item key="/"><Link to={'/'}/><HomeOutlined />首页</Menu.Item>
                <Menu.Item key='/weddingPage/'><Link to={'/weddingPage'}/><HeartOutlined />婚礼</Menu.Item>
                </Menu></Row>
            </Layout.Header>
        );
    }
}

export default HeaderCom



