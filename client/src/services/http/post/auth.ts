import axios from "axios";
import { SignupCredentials } from "../../../types/auth";

export const Signup = async (signupCredentials: SignupCredentials) => {

    try {
        const response = await axios.post("http://localhost:800/auth/signup", signupCredentials, {
            responseType: 'json'
        })

        if(response.status === 200) return true

    } catch (error) {
        console.error(error)
    }
}