import axios from "axios";


export const fetchAudioData = async () => {


    try {
        const response = await axios.get(`http://localhost:800/api/audio/data`, {
            responseType: 'json'
        })

        if (response) return response.data
        
    } catch (error) {
        console.error(error)
    }
}