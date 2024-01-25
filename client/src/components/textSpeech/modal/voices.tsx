import { SearchVoices } from "../../ui/input";
import { VoicesTypeBtn } from "../../ui/button";
import { getVoices } from "../../../services/http/get/voices";
import { VoiceTypes } from "../../../types/textSpeech";
import { generateRandomString, generateArrayOfNumbers } from "../../../utils/randomGenerator";
import { sortVoiceNameinParam } from "../../../utils/sort";

import { useEffect, useRef, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faX, faPlay, faPause } from "@fortawesome/free-solid-svg-icons";

import '../../../../public/ScrollStyle.css';


export function VoicesModal({selectedVoice, setselectedSpeed, setSelectedVoice, setOpenVoiceModal}: any){

  
    return(
        <>
        <div className="bg-gray-500 w-full h-screen fixed top-0 opacity-75"></div>

            <div className="w-full flex justify-center h-screen absolute top-0 py-6 ">

                <div className="bg-slate-950 h-full w-[60%] flex flex-col items-center rounded-md pt-5 px-5 relative">

                    <h1 className="text-white font-semibold text-lg">Select Your Voice</h1>
                    <FontAwesomeIcon 
                       className="absolute right-5 top-5 text-2xl font-bold text-white hover:opacity-75 hover:cursor-pointer" 
                       icon={faX}
                       onClick={() => setOpenVoiceModal(false)} 
                       />

                    <div className="flex justify-between w-[80%] mt-8">
                      <SearchVoices />
                      <VoicesTypeBtn />
                    </div>

                    <TableVoices setSelectedVoice={setSelectedVoice} />

                    <VoiceInUse selectedVoice={selectedVoice} setselectedSpeed={setselectedSpeed}  setOpenVoiceModal={setOpenVoiceModal}/>
                  
                </div>

            </div>

        </>
    )
}


function TableVoices({ setSelectedVoice }: any) {

    const [AudioLoading,setAudioLoading] = useState<boolean>(false)
    const [voicesData, setVoicesData] = useState<VoiceTypes[]>();

    const [playingSample, setPlayingSample] = useState<string | null>(null);
    const audioRef = useRef<HTMLAudioElement | null>(null);
  
    const tHead = [
      { name: "Name" },
      { name: "Gender" },
      { name: "Accent" },
      { name: "Language" },
    ];
  
    useEffect(() => {
      async function fetchVoices() {
        try {
          setVoicesData(await getVoices());
        } catch (error) {
          console.error(error);
        }
      }
  
      fetchVoices();
    }, []);
  
    const sortedArr = sortVoiceNameinParam(voicesData);
  
    const handleAudioPlayPause = (voice: VoiceTypes) => {

        const audio = audioRef.current;
      
        if (audio) {

          if (playingSample === voice.sample) {
            audio.pause();
            setPlayingSample(null);

          } else {

            if (playingSample) {
              const playingAudio = audioRef.current;
              playingAudio?.pause();
            }
      
            
            audio.src = voice.sample;
      
            audio.addEventListener('canplay', () => {
              audio.play().then(() => {
                setPlayingSample(voice.sample);
              });

            });

            audio.addEventListener('error', (error) => {
              console.error('Error loading audio:', error);
            });

          }
        }
      };
      
  
    return (
      <div className="h-[65%] w-full p-5 overflow-auto scrollable-container">

        {sortedArr ? (

          <table className="text-white font-semibold text-center w-full border-collapse border border-slate-700">

            <thead>
              <tr>

                {tHead.map((item) => (

                  <th key={item.name} className="border-b border-slate-700">
                    {item.name}
                  </th>

                ))}

              </tr>
            </thead>

            <tbody>

              {sortedArr?.map((voice) => (
                
                <tr
                  className="hover:bg-slate-800 hover:cursor-pointer"
                  key={generateRandomString()}
                  onClick={() => setSelectedVoice({voice: voice.id, name: voice.name, output_format: "mp3"})}
                >
                  <td className="py-5">

                    <button
                      className="w-[10%] bg-gray-400 p-1 text-center rounded-full mr-5 hover:opacity-75"
                      onClick={() => handleAudioPlayPause(voice)}
                    > 

                      {AudioLoading && <p className="text-red-800">Loading pa dol</p>}

                      <FontAwesomeIcon
                        icon={playingSample === voice.sample ? faPause : faPlay}
                      />

                    </button>

                    {voice.name}

                  </td>

                  <td className="py-5">{voice.gender.charAt(0).toUpperCase() + voice.gender.slice(1)}</td>
                  <td className="py-5">{voice.accent.charAt(0).toUpperCase() + voice.accent.slice(1)}</td>
                  <td className="py-5">{voice.language}</td>

                </tr>

              ))}

            </tbody>

          </table>

        ) : (
          <p className="text-white text-2xl">Loading...</p>
        )}

        <audio ref={audioRef} controls className="w-[20%] hidden" />
        
      </div>
    );
  }


function VoiceInUse({ selectedVoice, setselectedSpeed, setOpenVoiceModal }: any) {

    const numbersArray = generateArrayOfNumbers();

    return (
        <div className="flex items-end justify-around text-white h-[10%] w-full">
            
            <div className="flex gap-5 font-bold opacity-80">
                <h1>Voice in use</h1>
                <button disabled>{ selectedVoice?.name }</button>

                <select 
                    defaultValue={numbersArray[0]}
                    className="bg-transparent focus:outline-none" 
                    onChange={(e) => setselectedSpeed(e.target.value)}
                    >

                    {numbersArray.map((num) => (

                        <option
                            key={num}
                            className="bg-slate-950 hover:bg-gray-300 focus:bg-gray-300"
                            value={num}
                        >
                            {num}
                        </option>
                        
                    ))}
                </select>

            </div>

            <button 
                onClick={() => setOpenVoiceModal(false)} 
                className="p-2 rounded-md text-white font-bold w-[18%] h-[70%] bg-indigo-800 hover:opacity-75"
                >
                Confirm
            </button>
        </div>
    );
}
