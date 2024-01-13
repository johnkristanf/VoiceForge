import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMusic } from "@fortawesome/free-solid-svg-icons";

function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export function SpeechVoiceBtn({ setOpenVoiceModal }: any){

    return(
        <button onClick={() => setOpenVoiceModal(true)} className="p-2 rounded-t-md text-white w-[18%] h-[100%] bg-indigo-800">
            <FontAwesomeIcon icon={faMusic}/>&nbsp;Select a Voice 
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
