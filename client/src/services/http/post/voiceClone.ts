import axios from 'axios';


export const voiceClone = async (formData: FormData) => {

    try {

        const response = await axios.post('https://voiceforge-server.onrender.com/api/voice/clone', formData , {
            withCredentials: true,
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })

        if(response.data.id === '') return false
        
        return true

    } catch (error) {
        console.error(error)
    }
}



export const DeletevoiceClone = async (voice_id: string) => {

    try {

        const response = await axios.post('https://voiceforge-server.onrender.com/voice/clone/delete', { voice_id: voice_id } , {
            withCredentials: true,
            responseType: 'json'
        })

        if(response.data) return true

        return false
        
    } catch (error) {
        console.error(error)
    }
}
 