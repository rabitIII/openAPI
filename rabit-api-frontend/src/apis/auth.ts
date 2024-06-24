// import {useAxios} from "@/utils/request";

import { useAxios } from "@/utils/request";


export function getJwtToken() {
    return localStorage.getItem("jwtToken");
}

export function saveUser(user: any) {
    return localStorage.setItem("user", JSON.stringify(user));
}

export function getUser() {
    return JSON.parse(localStorage.getItem("user") || "");
}

interface AuthRequest {
    userAccount: string,
    userPassword: string
}

export function register({userAccount, userPassword}: AuthRequest) {
    return useAxios.post("/api/login", 
        
        {
        userAccount,
        userPassword,
    })
}

// export function loginApi(form: AuthRequest) {
//     return useAxios.post("/api/login", form)
// }
