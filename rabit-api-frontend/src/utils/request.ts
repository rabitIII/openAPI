import { getJwtToken } from "@/apis/auth";
import { useUserStore } from "@/stores/user";
import axios from "axios";

export const useAxios = axios.create({
    baseURL: "http://localhost:8080",
    timeout: 1000,
    

});

export interface baseResponse<T> {
    code: number,
    data: T,
    msg: string
};

useAxios.interceptors.request.use(
    (config) => {
        const userStore = useUserStore();
        // config.headers["token"] = store.userInfo.token
        return config
    },
    (error) => {
        return Promise.reject(error);
    }
);

useAxios.interceptors.response.use(
    (response) => {
        if(response.status !== 200) {
            // 失败信息
            console.log("服务失败", response.status);
            return Promise.reject(response.statusText)
        }
        return response.data
    },
    (error) => {
        console.log(error)
        return Promise.reject(error);
    }
)

