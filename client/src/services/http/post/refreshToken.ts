import axios from "axios";

export const RefreshToken = async (): Promise<boolean | undefined> => {
    
    try {
        const response = await axios.post('https://vf-server.onrender.com/refresh', {}, {
            withCredentials: true,
        });


        return response.data.New_Access_Token_Generated; // Return the boolean value directly

    } catch (error) {
        console.error(error);
        return undefined; // Return undefined in case of an error
    }
};
