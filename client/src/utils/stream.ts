import pako from 'pako'

export const getBlobUrl = (data: string) => {

  const base64StringtoBinary = atob(data);
  const bytesArray = []
  
  for(let index = 0; index < base64StringtoBinary.length; index++){
    bytesArray.push(base64StringtoBinary.charCodeAt(index))
  }
  
  
  const bufferArray = new Uint8Array(bytesArray);
  const decompressed = pako.inflate(bufferArray)

  const blob = new Blob([decompressed], { type: 'audio/mpeg' });
  return URL.createObjectURL(blob);
  
};



export const downloadAudio = (base64StreamBinary: string) => {
  
  const audioURL = getBlobUrl(base64StreamBinary)

  const a = document.createElement('a')

  a.href = audioURL
  a.download = "audio.mp3"

  a.click()

  URL.revokeObjectURL(audioURL);

}

