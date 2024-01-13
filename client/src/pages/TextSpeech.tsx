import { SideBar } from "../components/sidebar";
import { SpeechForm } from "../components/textSpeech/speechForm";

function TextSpeech(){

    return(
        <div className="flex justify-center w-full h-screen">
          <SideBar />
          <SpeechForm />
        </div>
       
    )
}

export default TextSpeech