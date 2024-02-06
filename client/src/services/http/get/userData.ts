import axios from "axios";

export const FetchUserData = async (setUnauthorized: any) => {

    try {

        const response = await axios.get("https://voiceforge-server.onrender.com/user/data", {
            responseType: 'json',
            withCredentials: true
        })

        return response.data

    } catch (error: any) {
        console.error(error)

        if (error.response && error.response.status === 401) setUnauthorized(true)
        
    }
}

