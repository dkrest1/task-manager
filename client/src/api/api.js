import axios from "axios";

const API_URL = "http://localhost:9090"

export const signup = async (userData) => {
    try {
        const response = await axios.post(`${API_URL}/auth/signup`, userData, {headers: {'Content-Type': 'application/json'}});
        return response.data

    }catch(error) {
        throw error.response.data
    }
}

export const login = async(userData) => {
    try {
        const response = await axios.post(`${API_URL}/auth/login`, userData, {headers: {'Content-Type': 'application/json'}})
        return response.data
    }catch(error) {
        throw error.response.data
    }
}