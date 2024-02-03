import axios from "axios";


export async function getVoices(SearchVoice: string) {

    try {
        const response = await axios.get(`http://localhost:800/api/voices/${encodeURIComponent(SearchVoice)}`, {
            withCredentials: true
        })

        console.log('SearchVoice', SearchVoice)
        
        console.log("res voice", response.data)
        if (response) return response.data.voices

    } catch (error) {
        console.error(error)
    }
}


export async function FetchVoiceClone() {

    try {
        const response = await axios.get(`http://localhost:800/api/get/voice/clone`, {
            withCredentials: true
        })
        
        if (response) return response.data

    } catch (error) {
        console.error(error)
    }
}