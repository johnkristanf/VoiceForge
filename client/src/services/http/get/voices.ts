import axios from "axios";


export async function getVoices(SearchVoice: string) {

    try {
        const response = await axios.get(`https://voiceforge-server.onrender.com/api/voices/${encodeURIComponent(SearchVoice)}`, {
            withCredentials: true
        })

        if (response) return response.data.voices

    } catch (error) {
        console.error(error)
    }
}


export async function FetchVoiceClone() {

    try {
        const response = await axios.get(`https://voiceforge-server.onrender.com/api/get/voice/clone`, {
            withCredentials: true
        })
        
        if (response) return response.data

    } catch (error) {
        console.error(error)
    }
}