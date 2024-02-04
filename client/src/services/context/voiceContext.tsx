import { createContext, useContext, useState, ReactNode, Dispatch, SetStateAction } from 'react';

interface VoiceCloneContextType {
    EmailVerfication: string; // Change 'any' to the type of your voiceCloneData
    setEmailVerfication: Dispatch<SetStateAction<string>>; // Change 'any' to the type of your voiceCloneData
}

const defaultVoiceCloneContext: VoiceCloneContextType = {
    EmailVerfication: '',
    setEmailVerfication: () => {},
};

const UserDataContext = createContext<VoiceCloneContextType>(defaultVoiceCloneContext);

export const useUserData = () => useContext(UserDataContext);

export const UserDataProvider = ({ children }: { children: ReactNode }) => {
    const [EmailVerfication, setEmailVerfication] = useState<string>(''); // Change 'any' to the type of your voiceCloneData

    return (
        <UserDataContext.Provider value={{ EmailVerfication, setEmailVerfication }}>
            {children}
        </UserDataContext.Provider>
    );
};
