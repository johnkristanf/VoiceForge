import axios from "axios";


export async function getVoices() {

    try {
        const response = await axios.get('http://localhost:800/api/voices')
        console.log("res voice", response.data)
        if (response) return response.data.voices

    } catch (error) {
        console.error(error)
    }
}