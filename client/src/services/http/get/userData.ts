import axios from "axios";

export const FetchUserData = async (setUnauthorized: any) => {

    try {

        const response = await axios.get("http://localhost:800/user/data", {
            responseType: 'json',
            withCredentials: true
        })
        console.log("user data:", response.data)

        return response.data

    } catch (error: any) {
        console.error(error)

        if (error.response && error.response.status === 401) setUnauthorized(true)
        
    }
}


export const RefreshToken = async () => {

    try {
            const response = await axios.post('http://localhost:800/token/refresh', {}, {
                withCredentials: true,
            });

            console.log("response in refresh token endpoint", response)

    } catch (error) {
        console.error(error)
    }
}