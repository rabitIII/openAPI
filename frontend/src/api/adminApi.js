import axios from "../utils/request.js";
import { parseToken } from "../utils/jwts.js";

export const login = async (params) => {
  const res = await axios.post("/api/users/login", params);
  if (res.data.data) {
    sessionStorage.setItem("token", JSON.stringify(parseToken(res.data.data)));
  }
  return res;
};
