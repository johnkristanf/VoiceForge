import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faX } from "@fortawesome/free-solid-svg-icons";
import { VoiceCloningInput } from "../ui/input";
import { SubmitCloneBtn } from "../ui/button";
import { SubmitHandler, useForm } from "react-hook-form";
import { VoiceCloneInput } from "../../types/voiceClone";
import { useEffect, useRef, useState } from "react";
import { voiceClone } from "../../services/http/post/voiceClone";


function VoiceCloningModal({setOpenCloningModal}: any){

    const [VoiceClone, setVoiceClone] = useState<boolean>();
    const { register, reset, handleSubmit } = useForm<VoiceCloneInput>();
    const [Voicefile, setVoicefile] = useState<FileList>();
    const [isSubmitting, setisSubmitting] = useState<boolean>(false)

    const errRef = useRef<HTMLParagraphElement>(null)

    const onSubmit: SubmitHandler<VoiceCloneInput> = async (cloneData: VoiceCloneInput) => {
        setisSubmitting(true)

        const formData = new FormData();
        
        formData.append('voice_name', cloneData.voice_name);
        if (Voicefile) formData.append('sample_file', Voicefile[0]);

        const clone = await voiceClone(formData);
        
        setVoiceClone(clone)
        setisSubmitting(false)
    };

    useEffect(() => {
        
        if(VoiceClone === true) {
            reset();
            setOpenCloningModal(false);
        }
        
        if (VoiceClone === false) {
            if (errRef.current) errRef.current.textContent = 'You can only clone 1 voice';
        }

    }, [VoiceClone]);
   

    return(
        <>
            <div className="bg-gray-500 w-full h-screen fixed top-0 opacity-75"></div>

            <div className="w-full flex justify-center h-screen absolute top-0 py-6">

                <div className="bg-slate-950 h-[95%] w-[40%] flex flex-col items-center rounded-md pt-5 px-12 relative">


                    <h1 className="text-white font-semibold text-lg">Let's create your Instant Voice Clone</h1>
                    <FontAwesomeIcon 
                       className="absolute right-5 top-5 text-2xl font-bold text-white hover:opacity-75 hover:cursor-pointer" 
                       icon={faX}
                       onClick={() => setOpenCloningModal(false)} 
                    />

                    <form onSubmit={handleSubmit(onSubmit)} className="h-full w-full" encType="multipart/form-data">

                        <p ref={errRef} className="text-red-800 text-center text-xl font-bold mt-3"></p>

                        <VoiceCloningInput register={register} setVoicefile={setVoicefile} />
                        <SubmitCloneBtn isSubmitting={isSubmitting} />
                    </form> 
                  
                </div>

            </div>

        </>
    )
}

export default VoiceCloningModal;