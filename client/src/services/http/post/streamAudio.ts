import axios from "axios";
import { TextToSpeech } from "../../../types/textSpeech";

export const streamAudio = async (data: TextToSpeech): Promise<Blob | undefined> => {
    
    try {

        const res = await axios.post("https://voiceforge-server.onrender.com/api/stream/voices", data, {
          responseType: 'blob',
          withCredentials: true
        })

        console.log('response stream', res)

        if(res) return res.data

    } catch (error) {
        console.error(error)
        return undefined
    }
}

