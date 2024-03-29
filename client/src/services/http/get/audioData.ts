import axios from "axios";


export const fetchAudioData = async () => {


    try {
        const response = await axios.get(`https://vf-server.onrender.com/api/audio/data`, {
            responseType: 'json',
            withCredentials: true
        })

        if (response) return response.data
        
    } catch (error) {
        console.error(error)
    }
}