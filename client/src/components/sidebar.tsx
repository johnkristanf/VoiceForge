import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faComments, faClone, faRightFromBracket, faTimes } from '@fortawesome/free-solid-svg-icons';

import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { FetchUserData } from '../services/http/get/userData';
import { RefreshToken } from '../services/http/post/refreshToken';
import { UserData } from '../types/userData';
import { Logout } from '../services/http/post/auth';
import { faBars } from '@fortawesome/free-solid-svg-icons/faBars';

export function SideBar(){

    const [navToggleOn, setnavToggleOn] = useState(false);
    const [UserData, setUserData] = useState<UserData>()

    return(

        <>

        { navToggleOn && <MenuPopUpModal UserData={UserData} setnavToggleOn={setnavToggleOn} /> }

        <FontAwesomeIcon 
            onClick={() => setnavToggleOn(true)}
            className="max-lg:block text-5xl absolute top-2 left-3 text-white font-bold hidden hover:opacity-40 hover:cursor-pointer" 
            icon={faBars} />

        <aside className="max-lg:hidden bg-slate-900 h-screen w-[20%] p-5 flex flex-col items-center gap-10 fixed left-0">
           <Logo />
           <SideBarLinks UserData={UserData} setUserData={setUserData} />

        </aside> 
        </>

       
    )
}


function Logo(){

    return(
            <div className="flex items-center gap-3">
                <img src="/img/VF_logo.png" className="rounded-full" width={50}/>
                <h1 className="text-3xl text-white font-semibold">VoiceForge</h1>
            </div>       
    )
}


const link = [
    {name: "Text to Speech", to: "text-speech", icon: <FontAwesomeIcon icon={faComments}/>, current: false },
    {name: "Voice Cloning", to: "voice-cloning", icon: <FontAwesomeIcon icon={faClone}/>, current: false },
]


const MenuPopUpModal = ({ UserData, setnavToggleOn }: any) => {

    return(
  
      <div className='max-lg:block max-lg:fixed hidden w-full h-screen bg-slate-900 absolute top-0 z-50'>
  
            <FontAwesomeIcon 
              className="text-5xl hover:opacity-75 text-white cursor-pointer absolute right-12 top-2"
              icon={faTimes} 
              onClick={() => setnavToggleOn(false)}
              />
  
  
            <div className="flex flex-col items-center justify-start p-20 h-full gap-10 inset-0">
  
                <h1 className="text-white text-5xl">
                   <Logo />
                </h1>
  

                <ul className="flex flex-col gap-8 text-white">
  
                    { link.map((item) => (
  
                      <li key={item.name}>
  
                            <Link
                              key={item.name} 
                              to={`/${item.to}`}
  
                                className={ item.current ? 'current' 
                                 : 'text-lg font-bold p-2 rounded-md hover:bg-black hover:opacity-100'
                                } >
  
                              {item.name}
  
                            </Link>
  
                     </li>
  
                    ))}
  
                </ul>

                <UserDataDisplay UserData={UserData}  />
    
            </div>
            
        </div>
    )
  }
  


function SideBarLinks({UserData, setUserData}: any){

    const [Unauthorized, setUnauthorized] = useState<boolean>(false);
    const [Token, setToken] = useState<boolean | undefined>(false)

    useEffect(() => {

        async function token() {
            const token = await RefreshToken()
            setToken(token)
        }

        if (Unauthorized) token()

        FetchUserData(setUnauthorized).then((data: UserData) => {
            setUserData(data)
        })

    }, [Unauthorized])

    if(Token) window.location.href = '/text-speech'


    
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

        <UserDataDisplay UserData={UserData} />

    </>

    )
}


function UserDataDisplay({ UserData }: any){
    
    return(
        <div className="flex bg-slate-700 absolute bottom-5 right-4 w-[90%] rounded-md text-white p-4 flex justify-around items-center gap-4">

            <div className="flex items-center gap-3 text-lg font-bold" >
               <img src="/img/user.jpg" className="rounded-full" width={40}/>
               <h1 className='truncate w-[80%]' >{UserData?.email}</h1> 
            </div>
           
            <FontAwesomeIcon onClick={() => Logout()} className='hover:opacity-75 hover:cursor-pointer text-xl' icon={faRightFromBracket}/>
        </div>
    )
}