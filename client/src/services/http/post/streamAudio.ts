import axios from "axios";
import { TextToSpeech } from "../../../types/textSpeech";

export const streamAudio = async (data: TextToSpeech): Promise<Blob | undefined> => {
    
    try {

        const res = await axios.post("http://localhost:800/api/stream/voices", data, {
          responseType: 'blob'
        })

        console.log('response stream', res)

        if(res) return res.data

    } catch (error) {
        console.error(error)
        return undefined
    }
}

