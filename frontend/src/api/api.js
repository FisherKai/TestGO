import { post } from './axios.js';

const login = async (username, password) => {
    try {
        const response = await post('/login', { username, password });
        const token = response.data.token;
        localStorage.setItem('token', token); // 存储 token 到 localStorage
        return token;
    } catch (error) {
        throw new Error('Login failed: ' + error.message);
    }
};

export default {
    login,
};