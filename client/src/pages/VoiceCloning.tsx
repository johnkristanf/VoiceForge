import React, { useState } from "react";
import { SideBar } from "../components/sidebar";
import { CreateNewCloneBtn } from "../components/ui/button";

const VoiceCloningModal = React.lazy(() => import('../components/voiceCloning/voiceCloningFormModal'))

function VoiceCloning(){
  const [OpenCloningModal, setOpenCloningModal] = useState<boolean>(false)

    return(
        <div className="flex justify-center w-full h-screen">
          <SideBar />
          <CreateNewCloneBtn setOpenCloningModal={setOpenCloningModal} />

          { OpenCloningModal && <VoiceCloningModal setOpenCloningModal={setOpenCloningModal} /> }
        </div>
       
    )
}

export default VoiceCloning