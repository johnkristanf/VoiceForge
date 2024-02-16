import axios from "axios";

export const deleteAudio = async (audio_id: string, setdeletedID: any) => {

    try {
        const response = await axios.delete(`https://vf-server.onrender.com/api/audio/delete/${encodeURIComponent(audio_id)}`, {
            responseType: 'text',
            withCredentials: true
        });
        
        setdeletedID(response.data)

    } catch (error) {
        console.error(error)
    }
}