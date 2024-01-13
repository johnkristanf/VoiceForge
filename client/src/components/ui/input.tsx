import '../../../public/ScrollStyle.css';

export function SpeechTextArea({setText}: any){

  
    return(
        <textarea 
          name="text" 
          onChange={(e) => setText(e.target.value)}
          className="bg-slate-800 rounded-b-md scrollable-container rounded-tr-md text-white resize-none h-[55%] focus:outline-none p-3 w-full" 
          placeholder="Type Here...." 
        />
    )
}


export function SearchVoices(){
    return <input
           type="search" 
           className="p-3 text-white w-[60%] rounded-md bg-transparent focus:outline-indigo-800 focus:outline-none border border-slate-700" 
           placeholder="Search Voice"/>
}