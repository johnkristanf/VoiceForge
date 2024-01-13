import axios from "axios";


export async function getVoices() {

    try {
        const response = await axios.get('http://localhost:800/api/voices')
        if (response) return response.data.voices

    } catch (error) {
        console.error(error)
    }
}