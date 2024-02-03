import React, { useEffect, useState } from "react";
import { SideBar } from "../components/sidebar";
import { CreateNewCloneBtn } from "../components/ui/button";
import { VoiceCloneType } from "../types/voiceClone";
import { FetchVoiceClone } from "../services/http/get/voices";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleCheck, faTrashCan } from '@fortawesome/free-solid-svg-icons';
import { DeletevoiceClone } from "../services/http/post/voiceClone";

const VoiceCloningModal = React.lazy(() => import('../components/voiceCloning/voiceCloningFormModal'))

function VoiceCloning(){
  
  const [OpenCloningModal, setOpenCloningModal] = useState<boolean>(false)
  const [CloneData, setCloneData] = useState<VoiceCloneType[]>();
  const [CloneDelete, setCloneDelete] = useState<boolean>(false);

  useEffect(() => {
    FetchVoiceClone().then((data: VoiceCloneType[]) => {
      setCloneData(data)
    })

    setCloneDelete(false)

  }, [CloneDelete, OpenCloningModal])

    return(
        <div className="flex flex-col items-center justify-start w-full h-screen">
          <SideBar />
          <CreateNewCloneBtn setOpenCloningModal={setOpenCloningModal} />

          { OpenCloningModal && <VoiceCloningModal setOpenCloningModal={setOpenCloningModal} /> }

          <CloneVoiceData setCloneDelete={setCloneDelete} CloneData={CloneData} />
        
        </div>
       
    )
}

function CloneVoiceData({setCloneDelete, CloneData}: any){

    return(
        <div className="w-[60%]">
              {
                CloneData?.map((data: VoiceCloneType) => (
                  <div key={data.id} className="w-full flex items-center justify-around text-white font-semibold">

                    <div className="text-md flex items-center gap-2">
                      {data.name} 
                      <div className="rounded-md p-1 bg-indigo-700">PlayHT2.0</div>
                    </div>

                    <div className="flex text-lg items-center gap-2"> 
                        <FontAwesomeIcon icon={faCircleCheck}/> 
                        Cloning Completed
                    </div>

                    <button onClick={async () => setCloneDelete(await DeletevoiceClone(data.id)) } className="bg-gray-700 rounded-md p-2 flex items-center gap-2">
                      <FontAwesomeIcon icon={faTrashCan}/> 
                      Delete
                    </button>

                  </div>

                ))
              }
        </div>
    )
}

export default VoiceCloning