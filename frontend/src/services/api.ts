import axios, { type AxiosInstance } from 'axios'

const api: AxiosInstance = axios.create({
    baseURL: `http://${window.location.hostname}:8080/api`,
})

export default api
