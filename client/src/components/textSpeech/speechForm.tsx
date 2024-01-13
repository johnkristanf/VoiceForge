import React, {  useState } from "react";
import { SpeechVoiceBtn } from "../ui/button";
import { SpeechTextArea } from "../ui/input";

import ReactPlayer from "react-player";

import { TextToSpeech } from "../../types/textSpeech";
import { streamAudio } from "../../services/http/post/streamAudio";
import { VoicesModal } from "./modal/voices";

export function SpeechForm() {

  const [OpenVoiceModal, setOpenVoiceModal] = useState<boolean>(false);
  const [audioURL, setaudioURL] = useState<string>('')
  const [selectedVoice, setSelectedVoice] = useState<{voice: string, name: string, output_format: string;}>({
    voice: "",
    name: "",
    output_format: ""

  });

  const [selectedSpeed, setselectedSpeed] = useState<string>("0.5");
  const [Text, setText] = useState<string>('');

  const textToSpeechData: TextToSpeech = {
    text: Text,
    voice: selectedVoice?.voice,
    output_format: selectedVoice?.output_format,
    speed: parseFloat(selectedSpeed)

  };

  console.log('selectedVoice', selectedVoice);
  console.log('selectedSpeed', selectedSpeed);

  const onSubmit = async (e: React.ChangeEvent<HTMLFormElement>) => {
    e.preventDefault();

    console.log("textToSpeechData", textToSpeechData);
    const streamBlob = await streamAudio(textToSpeechData);

    if (streamBlob instanceof Blob) {
      const audioURL = URL.createObjectURL(streamBlob);
      setaudioURL(audioURL)
    }

  };



  return (
    <>
        <div className="flex flex-col w-[60%] mt-24 ml-24">

            <div className="flex w-full">
              <SpeechVoiceBtn setOpenVoiceModal={setOpenVoiceModal} />
            </div>

            <form onSubmit={onSubmit} className="w-full h-[60%] flex items-start gap-3">
              <SpeechTextArea setText={setText} />

              <div className="flex flex-col p-4 gap-3 bg-slate-900 w-[40%] h-[55%] rounded-md">
                <button type="submit" className="text-white font-semibold bg-indigo-800 w-full rounded-md p-5 hover:opacity-75">Generate Speech</button>
                <p className="text-white font-bold opacity-75 text-[12px]">Unleash the potential of text-to-speech innovation and share your message in a unique and engaging way</p>
              </div>

            </form>

            <AudioPlayer audioURL={audioURL} />

        
        </div>

        { 
          OpenVoiceModal && (<VoicesModal
           selectedVoice={selectedVoice}
           setselectedSpeed={setselectedSpeed}
           setSelectedVoice={setSelectedVoice}
           setOpenVoiceModal={setOpenVoiceModal}
           />) 
        }

    </>

  );

}


function AudioPlayer({audioURL}: any){

    return(
        <div className="flex">

            {
              audioURL ? (<ReactPlayer height="80px" url={audioURL} controls config={{ file: { forceAudio: true } }}/>) 
              : (<p>Loading...</p>)
            }

        </div>
    )
}