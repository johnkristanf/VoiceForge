import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faComments, faClone, faRightFromBracket } from '@fortawesome/free-solid-svg-icons';

import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { FetchUserData } from '../services/http/get/userData';
import { RefreshToken } from '../services/http/post/refreshToken';
import { UserData } from '../types/userData';
import { Logout } from '../services/http/post/auth';

export function SideBar(){

    return(

        <aside className="bg-slate-900 h-screen w-[20%] p-5 flex flex-col items-center gap-10 fixed left-0">
           <Logo />
           <SideBarLinks />

        </aside>
    )
}


function Logo(){

    return(
            <div className="flex items-center gap-3">
                <img src="https://voiceforge.vercel.app/public/img/VF_logo.png" className="rounded-full" width={50}/>
                <h1 className="text-3xl text-white font-semibold">VoiceForge</h1>
            </div>       
    )
}


const link = [
    {name: "Text to Speech", to: "text-speech", icon: <FontAwesomeIcon icon={faComments}/> },
    {name: "Voice Cloning", to: "voice-cloning", icon: <FontAwesomeIcon icon={faClone}/> },
]



function SideBarLinks(){

    const [UserData, setUserData] = useState<UserData>()
    const [Unauthorized, setUnauthorized] = useState<boolean>(false);


    useEffect(() => {

        if(Unauthorized) RefreshToken().catch((err) => console.error(err))

        FetchUserData(setUnauthorized).then((data: UserData) => {
            setUserData(data)
        })

    }, [Unauthorized])


    
    return(

    <>
        <ul className="text-white font-bold text-lg">

        {
                link.map((item) => (

                    <Link key={item.name} to={`/${item.to}`}><li className="mt-10 hover:opacity-100 hover:cursor-pointer opacity-75">
                        {item.icon} {item.name}
                    </li></Link>
                    ))
            }
            
        </ul>

        <div className="flex bg-slate-700 absolute bottom-5 right-4 w-[90%] rounded-md text-white p-4 flex justify-around items-center gap-4">

            <div className="flex items-center gap-3 text-lg font-bold" >
               <img src="https://voiceforge.vercel.app/public/img/user.jpg" className="rounded-full" width={40}/>
               <h1 className='truncate w-[80%]' >{UserData?.email}</h1> 
            </div>
           
            <FontAwesomeIcon onClick={() => Logout()} className='hover:opacity-75 hover:cursor-pointer text-xl' icon={faRightFromBracket}/>
        </div>

    </>

    )
}