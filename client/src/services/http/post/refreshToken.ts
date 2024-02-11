import axios from "axios";

export const RefreshToken = async (): Promise<boolean | undefined> => {
    
    try {
        const response = await axios.post('http://localhost:800/token/refresh', {}, {
            withCredentials: true,
        });

        console.log("response in refresh token endpoint", response)

        return response.data.New_Access_Token_Generated; // Return the boolean value directly

    } catch (error) {
        console.error(error);
        return undefined; // Return undefined in case of an error
    }
};
