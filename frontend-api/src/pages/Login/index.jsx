import { Button, Form, Input, message } from "antd";
import "./index.css";
import { login } from "../../api/adminApi.js";
// import { parseToken } from "../../utils/jwts.js";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

// const onFinishFailed = (errorInfo) => {
//   console.log("Failed:", errorInfo);
// };
const Login = () => {
  const [form] = Form.useForm();
  const navigate = useNavigate();

  useEffect(() => {
    if (localStorage.getItem("token")) {
      console.log("token 已存在!");
      navigate("/layout");
    }
  }, [navigate]);

  const onFinish = async (values) => {
    try {
      const res = await login(values);
      if (res.data.data) {
        navigate("/layout");
      }

      // console.log(res)
    } catch (error) {
      console.log(error);
      message.error("登录失败, 请重试");
      navigate("/");
    }
    // console.log(res);
    // // console.log(res.data.token)
    // const token = res.data
    // console.log(token)
    // const payload = parseToken(token)
    // console.log(payload)
  };

  return (
    <div className="bg_login">
      <div className="content">
        <h2>OpenAPI</h2>
        <Form
          name="basic"
          form={form}
          labelCol={{
            span: 8,
          }}
          wrapperCol={{
            span: 16,
          }}
          style={{
            maxWidth: 450,
          }}
          initialValues={{
            userName: "",
            password: "",
          }}
          onFinish={onFinish}
          //   onFinishFailed={onFinishFailed}
          autoComplete="off"
        >
          <Form.Item
            label="Username"
            name="userName"
            rules={[
              {
                required: true,
                message: "Please input your username!",
              },
            ]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label="Password"
            name="password"
            rules={[
              {
                required: true,
                message: "Please input your password!",
              },
            ]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item
            wrapperCol={{
              offset: 8,
              span: 16,
            }}
          >
            <Button type="primary" htmlType="submit">
              Submit
            </Button>

            <Button
              type="primary"
              htmlType="submit"
              style={{ marginLeft: "30px" }}
            >
              register
            </Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  );
};
export default Login;
