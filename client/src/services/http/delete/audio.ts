import axios from "axios";

export const deleteAudio = async (audio_id: string, setdeletedID: any) => {

    console.log('audio_id', audio_id)

    try {
        const response = await axios.delete(`http://localhost:800/api/audio/delete/${encodeURIComponent(audio_id)}`, {
            withCredentials: true
        });
        console.log('res', response)
        setdeletedID(response.data.DELETED)

    } catch (error) {
        console.error(error)
    }
}