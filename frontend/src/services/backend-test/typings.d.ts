// @ts-ignore
/* eslint-disable */

declare namespace MYAPI {
  // 登录json参数
  type UserLoginRequest = {
    userName?: string;
    password?: string;
    currentAuthority?: string;
  };
  type BaseResponseUser = {
    code?: number;
    data?: MyUser;
    msg?: string;
  };
  type InterfaceInfo = {
    createTime?: string;
    description?: string;
    id?: number;
    isDelete?: number;
    method?: string;
    name?: string;
    requestHeader?: string;
    requestParams?: string;
    responseHeader?: string;
    status?: number;
    updateTime?: string;
    url?: string;
    userId?: number;
  };

  type BaseResponseUserBO = {
    code?: number;
    data?: UserVO;
    msg?: string;
  };
  type UserVO = {
    createdAt?: string;
    // gender?: number;
    id?: number;
    updatedAt?: string;
    // userAccount?: string;
    // userAvatar?: string;
    userName?: string;
    roleID?: number;
    token?: string;
  };
  type MyUser = {
    nickName?: string;
    roleID?: number;
    id?: number;
    token?: string;
  };
}
