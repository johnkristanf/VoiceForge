import axios from "axios";


export const fetchAudioData = async () => {


    try {
        const response = await axios.get(`http://locahost:800/api/audio/data`, {
            responseType: 'json',
            withCredentials: true
        })

        if (response) return response.data
        
    } catch (error) {
        console.error(error)
    }
}