import axios from "axios";


export const SearchVoice = async (voice_name: string) => {


    try {
        const response = await axios.get(`https://voiceforge-server.onrender.com/search/voice/${encodeURIComponent(voice_name)}`, {
            responseType: 'json',
            withCredentials: true
        })

        if (response) return response.data
        
    } catch (error) {
        console.error(error)
    }
}