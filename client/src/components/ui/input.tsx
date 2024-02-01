import { useRef } from 'react';
import '../../../public/ScrollStyle.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCloudArrowUp } from '@fortawesome/free-solid-svg-icons';


export function SpeechTextArea({Text, setText}: any){

  
    return(
        <textarea 
          name="text"
          value={Text} 
          onChange={(e) => setText(e.target.value)}
          className="bg-slate-800 rounded-b-md scrollable-container rounded-tr-md text-white resize-none h-[75%] focus:outline-none p-3 w-full" 
          placeholder="Type Here...." 
        />
    )
}


export function SearchVoices({setSearchVoice}: any){


    return (
        <input
        type="search" 
        className="p-3 text-white w-[60%] rounded-md bg-transparent focus:outline-indigo-800 focus:outline-none border border-slate-700" 
        placeholder="Search Voice"
        onChange={(e) => setSearchVoice(e.target.value)}
        />
       
    )
}


export function VoiceCloningInput({ register, setVoicefile }: any){

    const fileInputRef = useRef<HTMLInputElement | null>(null);

    const handleButtonClick = () => fileInputRef.current?.click();
      

    return(

        <div className="flex flex-col gap-8 mt-5">

            <div className="flex flex-col gap-2">
                <h1 className='font-semibold text-white'>Voice Name:</h1>
                <input 
                    type="text" 
                    className='p-4 border rounded-md font-bold' 
                    placeholder='Enter voice name'
                    {...register("voice_name", { required: true })}
                />

            </div>
           

            <div className="flex flex-col gap-2">

                <h1 className='font-semibold text-white'>Upload High Quality Audio sample</h1>

                <button 
                  type='button'
                  onClick={() => handleButtonClick()}
                  className=' bg-slate-800 p-3 rounded-md w-full flex flex-col items-center text-white'
                  >
                    <FontAwesomeIcon icon={faCloudArrowUp}/>
                    Upload file here or browse
                </button>

                <p className='font-semibold opacity-75 text-sm text-white'>
                    Minimum file size: 5kb 
                    <br /> 
                    Maximum file size: 50mb
                </p>

                <input
                    onChange={(e) => {
                      const selectedFile = e.target.files;
                      if (selectedFile) {
                        setVoicefile(selectedFile);
                      }
                    }}

                    type="file"
                    ref={fileInputRef}
                    style={{ display: 'none' }}
                />

            </div>
        
        </div>

    )
}