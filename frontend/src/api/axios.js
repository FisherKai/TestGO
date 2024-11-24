import axios from 'axios';

// 创建 Axios 实例
const instance = axios.create({
    baseURL: 'http://localhost:8080/api', // 替换为你的 API 地址
    timeout: 5000, // 请求超时时间
});

// 请求拦截器
instance.interceptors.request.use(
    config => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers.Authorization = `${token}`;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// 响应拦截器
instance.interceptors.response.use(
    response => {
        return response;
    },
    error => {
        if (error.response.status === 401) {
            // 处理未授权的情况，例如重新登录
            console.error('Unauthorized, please log in again.');
        }
        return Promise.resolve(error.response);
    }
);

// 导出常用的 HTTP 方法
export const get = (url, params) => instance.get(url, { params });
export const post = (url, data) => instance.post(url, data);
export const put = (url, data) => instance.put(url, data);
export const del = (url, data) => instance.delete(url, { data });

// 导出 Axios 实例
export default instance;