import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMusic,faTrash,faDownload, faPlus } from "@fortawesome/free-solid-svg-icons";

import { downloadAudio } from '../../utils/stream';
import { deleteAudio } from '../../services/http/delete/audio';

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export function SpeechVoiceBtn({ selectedVoice, setOpenVoiceModal }: any){

    return(
        <button onClick={() => setOpenVoiceModal(true)} className="p-2 rounded-t-md text-white w-[25%] h-[100%] bg-indigo-800">
            <FontAwesomeIcon icon={faMusic}/>&nbsp;{selectedVoice.name} 
        </button>
    )
}

const vType = [
    {name: "Pre-made", clicked: false},
    {name: "Cloned", clicked: false}
]


export function VoicesTypeBtn(){

    return(

        <div className='flex gap-6'>
            {
                vType.map((item) => (

                    <button key={item.name}
                      className={
                        classNames(
                            item.clicked ? 'opacity-75' : 'bg-indigo-800 p-3 text-white font-bold rounded-md'
                            )} 
                      >

                      {item.name}

                    </button>
                ))
            }
        </div>
    )
}



export function GenerateSpeechBtn(){


    return(
        <div className="flex flex-col p-4 gap-3 bg-slate-900 w-[40%] h-[75%] rounded-md">
            <button className="text-white font-semibold bg-indigo-800 w-full rounded-md p-5 hover:opacity-75">Generate Speech</button>
            <p className="text-white font-bold opacity-75 text-[12px]">Unleash the potential of text-to-speech innovation and share your message in a unique and engaging way</p>
        </div>
    )
}


export function StreamAudioActionsBtn({setdeletedID, audio_id, base64StreamBinary}: any){

    return(
        <div className="flex gap-5">
            <FontAwesomeIcon icon={faDownload} onClick={() => downloadAudio(base64StreamBinary)}/>
            <FontAwesomeIcon icon={faTrash} onClick={() => deleteAudio(audio_id, setdeletedID)} />
        </div>
    )
}



export function CreateNewCloneBtn({setOpenCloningModal}: any){
    
    return(
        <div className="flex flex-col gap-3 w-[58%] h-1/2 mt-5">
            <h1 className='text-white font-semibold text-xl'>Voice Cloning</h1>

            <button 
                onClick={() => setOpenCloningModal(true)}
                className='p-5 w-[35%] bg-transparent rounded-md border bg-gray-800 flex justify-center items-center gap-3 font-semibold'
                >
                <FontAwesomeIcon className='bg-indigo-500 p-2 rounded-md' icon={faPlus}/> 
                Create a New Clone
            </button>

        </div>
    )
}

export function SubmitCloneBtn(){
    return <button type='submit' className='p-3 w-full rounded-md mt-6 bg-indigo-700 hover:opacity-75'>Create</button>
}