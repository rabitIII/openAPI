
export const parseToken = (token) => {
    let payload = token.split(".")[1];
    return JSON.parse(decodeURIComponent(encodeURI(window.atob(payload.replace(/-/g, "+").replace(/_/g, "/")))))
}