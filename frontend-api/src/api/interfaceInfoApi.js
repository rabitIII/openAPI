import axios from "../utils/request.js";

export const listInterfaceInfoGet = async (params) => {
  let url = "api/interface/list";
  if (params) {
    url += `?${new URLSearchParams({ key: params })}`;
  }
  const res = await axios.get(url);
  return res.data;
};

export const addInterfaceInfoPost = async (values) => {
  const res = await axios.post("api/interface/create", values);
  console.log("createAPI:", res);
};

export const updateInterfaceInfoPut = async (values) => {
  const res = await axios.post("api/interface/update", values);
  console.log("createAPI:", res);
};
