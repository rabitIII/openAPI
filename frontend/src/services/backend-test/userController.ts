/* eslint-disable */
import { request } from '@umijs/max';

/** 获取当前的用户 GET /api/currentUser */
export async function getCurrentUser(options?: { [key: string]: any }) {
  return request<MYAPI.BaseResponseUserBO>('/api/user/currentUser', {
    headers: {
      'Content-Type': 'application/json',
    },
    method: 'GET',
    ...(options || {}),
  });
}

/** userLogin POST /api/user/login */
export async function userLoginUsingPOST(
  body: MYAPI.UserLoginRequest,
  options?: { [key: string]: any },
) {
  return request<MYAPI.BaseResponseUser>('/api/user/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
