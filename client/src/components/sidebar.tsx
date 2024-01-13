import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faComments, faClone, faGear, faRightFromBracket } from '@fortawesome/free-solid-svg-icons';

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
                <img src="../../public/img/VF_logo.png" className="rounded-full" width={50}/>
                <h1 className="text-3xl text-white font-semibold">VoiceForge</h1>
            </div>       
    )
}


const link = [
    {name: "Text to Speech", icon: <FontAwesomeIcon icon={faComments}/> },
    {name: "Voice Cloning", icon: <FontAwesomeIcon icon={faClone}/> },
    {name: "Settings", icon: <FontAwesomeIcon icon={faGear}/> }
]



function SideBarLinks(){
    
    return(

    <>
        <ul className="text-white font-bold text-lg">

            {
                link.map((item) => (

                    <li key={item.name} className="mt-10 hover:opacity-100 hover:cursor-pointer opacity-75">
                        {item.icon} {item.name}
                    </li>
                ))
            }
            
        </ul>

        <div className="flex bg-slate-700 absolute bottom-5  w-[80%] rounded-md text-white p-4 flex justify-around items-center ">

            <div className="flex items-center gap-3 text-lg font-bold" >
               <img src="../../public/img/VF_logo.png" className="rounded-full" width={40}/>
               <h1>John</h1> 
            </div>
           
            <FontAwesomeIcon className='hover:opacity-75 hover:cursor-pointer text-xl' icon={faRightFromBracket}/>
        </div>

    </>

    )
}