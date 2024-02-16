import axios from "axios";
import { TextToSpeech } from "../../../types/textSpeech";

export const streamAudio = async (data: TextToSpeech): Promise<string | undefined> => {
    
    try {

        const response = await axios.post("https://vf-server.onrender.com/api/stream/voices", data, {
           responseType: 'text',
           withCredentials: true
        })

        console.log('response stream', response)

        if(response) return response.data

    } catch (error) {
        console.error(error)
        return undefined
    }
}

