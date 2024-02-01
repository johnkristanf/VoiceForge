import { createContext, useContext, useState, ReactNode, Dispatch, SetStateAction } from 'react';

interface VoiceCloneContextType {
    voiceCloneData: any; // Change 'any' to the type of your voiceCloneData
    setVoiceCloneData: Dispatch<SetStateAction<any>>; // Change 'any' to the type of your voiceCloneData
}

const defaultVoiceCloneContext: VoiceCloneContextType = {
    voiceCloneData: null,
    setVoiceCloneData: () => {},
};

const VoiceCloneDataContext = createContext<VoiceCloneContextType>(defaultVoiceCloneContext);

export const useVoiceCloneData = () => useContext(VoiceCloneDataContext);

export const VoiceCloneDataProvider = ({ children }: { children: ReactNode }) => {
    const [voiceCloneData, setVoiceCloneData] = useState<any>(null); // Change 'any' to the type of your voiceCloneData

    return (
        <VoiceCloneDataContext.Provider value={{ voiceCloneData, setVoiceCloneData }}>
            {children}
        </VoiceCloneDataContext.Provider>
    );
};
