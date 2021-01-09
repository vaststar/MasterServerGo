import React, { Component } from 'react'
import { Link, withRouter} from 'react-router-dom'
import {Layout, Menu, Row, Col, Avatar} from 'antd'
import {
    HomeOutlined,
    SettingFilled,
  } from '@ant-design/icons';

import logo from './logo.png'

class HeaderCom extends Component {
    getLinkToUrl= (url)=>{
        //处理登陆完之后的跳转页面
        return ({
            pathname:url,
            state:{from:this.props.location.pathname}})
    }
    handleClick = (e) => {
        //处理选择菜单的问题
        if('logout'===e.key){
        }
    }

    render() {
        return (
            <Layout.Header theme='light' style={{ position: 'fixed', zIndex: 1, width: '100%' ,background:'white'}}>
                <Row justify="end" gutter={4} wrap={false}>
                <Col span={8}>
                </Col>
                <Col span={8}>
                <Menu onClick={this.handleClick}
                  mode="horizontal"
                  theme='light'
                  selectedKeys={[this.props.location.pathname]}
                  style={{ lineHeight: '64px' }}
                >
                <Menu.Item key='/weddingPage/'><Link to={this.getLinkToUrl('/weddingPage/')}/><SettingFilled />婚礼</Menu.Item>
                  <Menu.Item key="/"><Link to={'/'}/><HomeOutlined />首页</Menu.Item>
                </Menu>
                </Col>
                <Col span={2}>
                        <span style={{'fontSize':'36px','color':'black'}}>学士阁</span>
                        </Col>
                    <Col span={1}>
                        <Avatar src={logo} alt="学士" size={60}></Avatar>
                    </Col>
                </Row>
            </Layout.Header>
        );
    }
}

export default withRouter(HeaderCom)



