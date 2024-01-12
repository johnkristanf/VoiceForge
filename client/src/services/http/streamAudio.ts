import axios from "axios";
import { TextToSpeech } from "../../types/textSpeech";

export const streamAudio = async (data : TextToSpeech) => {
    
    try {
        const res = await axios.post("http://localhost:800/api/stream/voices", data, {
          responseType: 'arraybuffer'
        })
  
        const blob = new Blob([res.data], { type: 'audio/mpeg' });

        console.log('audio blob', blob)

    } catch (error) {
        console.error(error)
    }
}

