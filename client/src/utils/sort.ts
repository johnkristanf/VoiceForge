import { VoiceTypes } from "../types/textSpeech";

export function sortVoiceNameinParam(voiceArr: VoiceTypes[] | undefined){

    console.log('voiceArr', voiceArr)

    function customSort(a: VoiceTypes, b: VoiceTypes){

        if (a.name.length > b.name.length) return -1;
        if (a.name.length < b.name.length) return 1;
        
        return 0;
    }

    return voiceArr?.slice().sort(customSort)
    
}