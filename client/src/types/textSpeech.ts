
export type TextToSpeech = {
    text: string,
    voice: string,
    output_format: string,
    speed: number
}


export type VoiceTypes = {
  name: string,
  accent: string,
  gender: string,
  language: string
  id: string
  style: string
  sample: string
}