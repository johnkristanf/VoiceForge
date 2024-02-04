import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { Verify } from "../services/http/post/auth";
import { useRef } from "react";

import { useUserData } from "../services/context/voiceContext"; 

function classNames(...classes: any) {
   return classes.filter(Boolean).join(' ');
}


const codes = [
    { name: "first_Code" },
    { name: "second_Code" },
    { name: "third_Code" },
    { name: "fourth_Code" },
    { name: "fifth_Code" },
];


type CodeTypes = {
    first_Code: string,
    second_Code: string,
    third_Code: string,
    fourth_Code: string,
    fifth_Code: string,
}

const Verification = () => {

    const { register, reset, handleSubmit} = useForm<CodeTypes>();
    const navigate = useNavigate()
    const errorMsgRef = useRef<HTMLParagraphElement | null>(null);
    const { EmailVerfication } = useUserData()

    console.log("EmailVerfication", EmailVerfication)

    const onSubmit = async (codes: CodeTypes) => {

        let verification_code = ""
        Object.values(codes).forEach(([value]) => {
            verification_code += value
        });


        const verfied = await Verify(verification_code)
        reset();

        if(verfied.ERROR){
            if(errorMsgRef.current) errorMsgRef.current.textContent = verfied.ERROR;
            return
        }

        navigate("/text-speech")       
    }

    return (

        <div className="relative flex min-h-screen flex-col justify-center overflow-hidden py-12">
                
        <div className="relative bg-indigo-900 px-6 pt-10 pb-9 shadow-xl mx-auto w-full max-w-lg rounded-2xl">
            <div className="mx-auto flex w-full max-w-md flex-col space-y-16">
            <div className="flex flex-col items-center justify-center text-center space-y-2">
                <div className="font-semibold text-3xl">
                <p>Email Verification</p>
                </div>
                <div className="flex flex-row text-sm font-medium text-gray-400">
                <p>We have sent a code to your email {EmailVerfication}</p>
                </div>
            </div>

            <div>

                <p ref={errorMsgRef} className="text-red-800 text-center font-bold text-lg mb-5"></p>

                <form onSubmit={handleSubmit(onSubmit)} >

                <div className="flex flex-col space-y-16">
                    
                    <div className="flex flex-row items-center justify-between mx-auto w-full max-w-xs">
                        
                        {codes.map((data) => (

                            <div className="w-16 h-16 mr-3" key={data.name}>

                            <input
                                className={classNames("w-full h-full flex flex-col items-center justify-center text-center px-5 outline-none rounded-xl border border-gray-200 text-lg bg-white focus:bg-gray-50 focus:ring-1 ring-blue-700")}
                                type="text" 
                                inputMode="numeric"
                                maxLength={1}
                                {...register(data.name as keyof CodeTypes, { required: true })}
                            />
                            </div>

                        ))}
                    </div>

                    <div className="flex flex-col space-y-5">
                    <div>
                        <button type="submit" className="font-bold flex flex-row items-center justify-center text-center w-full border rounded-xl outline-none py-5 bg-blue-700 border-none text-white text-sm shadow-sm hover:opacity-75">
                        Verify Account
                        </button>
                    </div>

                    <div className="flex flex-row items-center justify-center text-center text-sm font-medium space-x-1 text-gray-500">
                        <p>Didn't receive code?</p> <a className="flex flex-row items-center text-blue-600" href="http://" target="_blank" rel="noopener noreferrer">Resend</a>
                    </div>
                    </div>
                </div>

                </form>


            </div>
            </div>
        </div>
        </div>
    );
    };

    export default Verification;
