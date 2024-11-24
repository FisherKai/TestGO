import {get, post} from './axios.js';

const login = async (username, password) => {
    try {
        const response = await post('/login', {username, password});
        const token = response.data.token;
        localStorage.setItem('token', token); // 存储 token 到 localStorage
        return token;
    } catch (error) {
        throw new Error('Login failed: ' + error.message);
    }
};

const createContent = async (formData, status) => {
    formData.status = status
    const resp = await post('/content/create', formData)
    return resp;
}

const updateContent = async (formData, status) => {
    formData.status = status
    const resp = await post('/content/update', formData)
    return resp;
}

const deleteContent = async (id) => {
    const resp = await post('/content/delete', {id})
    return resp;
}

const getContentList = async (id) => {
    const resp = await get('/content/getList', {id})
    return resp;
}

const getResourceInfo = async () => {
    const resp = await get("/resource/getInfo")
    return resp
}

export default {
    login,
    createContent,
    updateContent,
    deleteContent,
    getContentList,
    getResourceInfo
};