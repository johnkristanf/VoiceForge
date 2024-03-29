import axios from "axios";
import { LoginCredentials, SignupCredentials } from "../../../types/auth";

export const Signup = async (signupCredentials: SignupCredentials) => {

    try {
        const response = await axios.post("https://vf-server.onrender.com/auth/signup", signupCredentials, {
            responseType: 'json'
        })

        if(response.status === 200) return true

    } catch (error) {
        console.error(error)
    }
}


export const Login = async (loginCredentials: LoginCredentials) => {

    try {

        const response = await axios.post("https://vf-server.onrender.com/auth/login", loginCredentials, {
            responseType: 'json',
            withCredentials: true
        })

        if(response.status === 200) return true

    } catch (error: any) {
        if (error.response) return error.response.data;
        console.error(error)
    }
}


export const Verify = async (verification_code :string) => {

    try {

        const response = await axios.post("https://vf-server.onrender.com/auth/verification", { verification_code: verification_code }, {
            responseType: 'json',
            withCredentials: true
        })

        if(response.status === 200) return true

    } catch (error: any) {
        if (error.response) return error.response.data;
        console.error(error)
    }
}


export const Logout = async () => {

    try {

        const response = await axios.post("https://vf-server.onrender.com/logout", {}, {
            withCredentials: true
        })

        if(response.data) window.location.href = "/"

    } catch (error: any) {
        if (error.response) return error.response.data;
        console.error(error)
    }
}