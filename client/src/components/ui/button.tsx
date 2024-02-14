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



export function VoicesTypeBtn({ setVoiceCloneTable }: any){

    const vType = [
        {name: "Pre-made", clicked: false, onClick: () => setVoiceCloneTable(false)},
        {name: "Cloned", clicked: false, onClick: () => setVoiceCloneTable(true)}
    ]

    return(

        <div className='flex gap-6'>
            {
                vType.map((item) => (

                    <button key={item.name}
                      className={
                        classNames(
                            item.clicked ? 'opacity-75' : 'bg-indigo-800 p-3 text-white font-bold rounded-md'
                        )}
                        
                        onClick={() => item.onClick()}
                      >

                      {item.name}

                    </button>
                ))
            }
        </div>
    )
}



export function GenerateSpeechBtn({isSubmiting}: any){


    return(
        <div className="flex flex-col gap-3 w-[40%] h-[75%] rounded-md">
            <button 
                disabled={isSubmiting}
                className={ isSubmiting ? "text-white font-semibold bg-gray-400 w-full h-full rounded-md p-5 bg-gray-400 cursor-no-drop opacity-75 text-xl"  
                                        : "text-white font-semibold bg-indigo-800 w-full h-full rounded-md p-5 hover:opacity-75 text-xl"}>
                <FontAwesomeIcon icon={faMusic}/> Generate Speech
            </button>

        </div>
    )
}


export function StreamAudioActionsBtn({setdeletedID, audio_id, base64StreamBinary}: any){

    return(
        <div className="flex gap-5">
            <FontAwesomeIcon className='hover:opacity-75 hover:cursor-pointer' icon={faDownload} onClick={() => downloadAudio(base64StreamBinary)}/>
            <FontAwesomeIcon className='hover:opacity-75 hover:cursor-pointer' icon={faTrash} onClick={() => deleteAudio(audio_id, setdeletedID)} />
        </div>
    )
}



export function CreateNewCloneBtn({setOpenCloningModal}: any){
    
    return(
        <div className="flex flex-col gap-3 w-[58%] h-1/2 mt-5">
            <h1 className='text-white font-semibold text-xl'>Voice Cloning</h1>

            <button 
                onClick={() => setOpenCloningModal(true)}
                className='p-5 w-[35%] bg-transparent text-white rounded-md border bg-gray-800 flex justify-center items-center gap-3 font-semibold'
                >
                <FontAwesomeIcon className='bg-indigo-500 p-2 rounded-md' icon={faPlus}/> 
                Create a Voice Clone
            </button>

        </div>
    )
}

export function SubmitCloneBtn({isSubmitting}: any){
    return <button 
               disabled={isSubmitting} 
               type='submit' 
               className={isSubmitting ? 'p-3 w-full rounded-md mt-6 bg-gray-400 hover:opacity-75 text-white font-semibold cursor-no-drop opacity-75' 
                                       : 'p-3 w-full rounded-md mt-6 bg-indigo-700 hover:opacity-75 text-white font-semibold'}>
                Create
            </button>
}