import React, {  useEffect, useState } from "react";
import { SpeechVoiceBtn, GenerateSpeechBtn, StreamAudioActionsBtn } from "../ui/button";
import { SpeechTextArea } from "../ui/input";

import { fetchAudioData } from "../../services/http/get/audioData";

import { TextToSpeech, AudioDataTypes } from "../../types/textSpeech";
import { streamAudio } from "../../services/http/post/streamAudio";
import { VoicesModal } from "./modal/voices";

import { getBlobUrl } from "../../utils/stream";

import '../../../public/ScrollStyle.css';


export function SpeechForm() {

  const [isSubmiting, setisSubmiting] = useState<boolean>(false);
  const [OpenVoiceModal, setOpenVoiceModal] = useState<boolean>(false);
  const [audioURL, setaudioURL] = useState<string>('')
  
  const [selectedVoice, setSelectedVoice] = useState<{voice: string, name: string, output_format: string;}>({
    voice: "s3://mockingbird-prod/charlotte_vo_narrative_9290be17-ccea-4700-a7fd-a8fe5c49fb20/voices/speaker/manifest.json",
    name: "Charlotte (Narrative)",
    output_format: "mp3"
  });


  const [selectedSpeed, setselectedSpeed] = useState<string>("1");
  const [Text, setText] = useState<string>('');

  const textToSpeechData: TextToSpeech = {
    text: Text,
    voice: selectedVoice?.voice,
    output_format: selectedVoice?.output_format,
    speed: parseFloat(selectedSpeed),
  };

  const onSubmit = async (e: React.ChangeEvent<HTMLFormElement>) => {
    e.preventDefault();

    setisSubmiting(true)

    const streamRes = await streamAudio(textToSpeechData);

    if (streamRes) setaudioURL(streamRes)

    setisSubmiting(false)
    setText('')

  };

  useEffect(() => {

    if (Text === '') {
      setisSubmiting(true);
      return
    }

    setisSubmiting(false);

  }, [Text]);



  return (
    <>
        <div className="flex flex-col w-[60%] mt-12 ml-24 max-lg:w-[90%] max-lg:ml-0 max-lg:mt-24 ">

            <div className="flex w-full">
              <SpeechVoiceBtn selectedVoice={selectedVoice} setOpenVoiceModal={setOpenVoiceModal} />
            </div>

            <form onSubmit={onSubmit} className="w-full h-[40%] flex items-start gap-3 mb-3">
              <SpeechTextArea Text={Text} setText={setText} />

              <GenerateSpeechBtn isSubmiting={isSubmiting} /> 

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


function AudioPlayer({ audioURL }: any) {

  const [audioDataArray, setAudioDataArray] = useState<AudioDataTypes>();
  const [deletedID, setDeletedID] = useState<Number>();
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    async function fetchData() {

      setIsLoading(true);

      try {
        const audioDataArray = await fetchAudioData();
        setAudioDataArray(audioDataArray);

      } catch (error) {
        console.error("Error fetching audio data:", error);

      } finally {
        setIsLoading(false);
      }

    }

    fetchData();

  }, [audioURL, deletedID]);

  if (isLoading) {
    return <div className="text-white text-2xl font-bold">Loading Speech...</div>;
  }

  if (!audioDataArray || audioDataArray.length === 0) {
    return <div className="text-white text-2xl font-bold">No generated speech available</div>;
  }

  return (
    <div className="flex flex-col gap-3 h-[40%] text-white">
      <h1 className="text-white font-bold text-2xl">Generated Speech</h1>
      <p className="text-white font-bold text-md">12,500 characters limit</p><br />
      <p className="text-white font-bold text-md">1 Instant voice clone only</p><br />
      
      <div className="overflow-auto scrollable-container">
        {audioDataArray?.map((item) => (

          <div
            key={item.audioStream} 
            className="flex justify-around items-center bg-slate-900 p-3 mb-4 rounded-md gap-5"
          >
            <h1 className="w-[50%] truncate font-bold">{item.audioText}</h1>

            <audio
              src={getBlobUrl(item.audioStream)}
              controls
              className="h-[40px]"
              controlsList="nodownload nofullscreen noremoteplayback"
            />
            <StreamAudioActionsBtn setdeletedID={setDeletedID} audio_id={item.audio_id} base64StreamBinary={item.audioStream} />
          </div>
        ))}
      </div>
    </div>
  );
}
