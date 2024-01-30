import axios from 'axios';


export const voiceClone = async (formData: FormData) => {

    try {

        console.log('formData', formData)

        const response = await axios.post('http://localhost:800/api/voice/clone', formData , {
            withCredentials: true,
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })

        console.log('res clone', response)
    } catch (error) {
        console.error(error)
    }
}