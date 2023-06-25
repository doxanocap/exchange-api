import axios from 'axios'; 

const auth = axios.create({

})

const api = axios.create({
    withCredentials: true,
    baseURL: "http://localhost:8080/webchat"
})

api.interceptors.request.use((config) => {
    if (config.headers !== undefined ) {
        config.headers.Authorization = `Bearer ${localStorage.getItem('accessToken')}`
    }
    return config;
})

export default api;