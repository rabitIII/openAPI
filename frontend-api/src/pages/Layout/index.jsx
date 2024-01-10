import React from "react";
import "./index.css";
import {
  HomeOutlined,
  SettingOutlined,
  ApartmentOutlined,
  LaptopOutlined,
  LogoutOutlined,
  NotificationOutlined,
  UserOutlined,
  ApiOutlined,
} from "@ant-design/icons";
import { Breadcrumb, Layout, Menu, theme } from "antd";
const { Header, Content, Sider } = Layout;
const items1 = [
  {
    label: "settings",
    key: "SubMenu",
    icon: <SettingOutlined />,
    children: [
      {
        label: "Exit",
        key: "exit",
        icon: <LogoutOutlined />,
      },
    ],
  },
];
const onClick = ({ key }) => {
  console.log(key);
};
const items2 = [
  {
    label: "Home",
    key: "home",
    icon: <HomeOutlined />,
  },
  {
    label: "状态管理",
    key: "role",
    icon: <ApartmentOutlined />,
    children: [
      {
        label: "用户管理",
        key: "user",
        icon: <UserOutlined />,
      },
      {
        label: "api管理",
        key: "api",
        icon: <ApiOutlined />,
      },
    ],
  },
];
const App = () => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  return (
    <Layout className="layout">
      <Header
        className="header"
        style={{
          display: "flex",
          //   alignItems: "center",
        }}
      >
        <div className="demo-logo">OpenAPI</div>
        <Menu
          className="menu"
          //   theme="dark"
          mode="horizontal"
          defaultSelectedKeys={["2"]}
          items={items1}
          style={{
            // flex: 1,
            minWidth: 0,
          }}
          onClick={onClick}
        />
      </Header>
      <Layout>
        <Sider
          width={200}
          style={{
            background: colorBgContainer,
          }}
        >
          <Menu
            mode="inline"
            defaultSelectedKeys={["home"]}
            // defaultOpenKeys={["sub1"]}
            style={{
              height: "100%",
              borderRight: 0,
            }}
            items={items2}
            onClick={onClick}
          />
        </Sider>
        <Layout
          style={{
            padding: "0 24px 24px",
          }}
        >
          <Breadcrumb
            style={{
              margin: "16px 0",
            }}
          >
            <Breadcrumb.Item>Home</Breadcrumb.Item>
            <Breadcrumb.Item>List</Breadcrumb.Item>
            <Breadcrumb.Item>App</Breadcrumb.Item>
          </Breadcrumb>
          <Content
            style={{
              padding: 24,
              margin: 0,
              minHeight: 280,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            Content
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
};
export default App;
