import axios from "axios";

export const deleteAudio = async (audio_id: string, setdeletedID: any) => {

    try {
        const response = await axios.delete(`http://localhost:800/api/audio/delete/${encodeURIComponent(audio_id)}`, {
            withCredentials: true
        });
        
        setdeletedID(response.data.DELETED)

    } catch (error) {
        console.error(error)
    }
}