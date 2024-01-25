import axios from 'axios';


export const voiceClone = async (formData: FormData) => {

    try {

        console.log('formData', formData)

        const response = await axios.post('http://localhost:800/api/voice/clone', formData)

        console.log('res clone', response)
    } catch (error) {
        console.error(error)
    }
}