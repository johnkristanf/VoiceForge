import axios from "axios";

export const RefreshToken = async () => {

    try {
            const response = await axios.post('http://locahost:800/token/refresh', {}, {
                withCredentials: true,
            });

            console.log("response in refresh token endpoint", response)


    } catch (error) {
        console.error(error)
    }
}