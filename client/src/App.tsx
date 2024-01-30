import React, { Suspense } from "react";

import { BrowserRouter, Routes, Route } from "react-router-dom";

const Auth = React.lazy(() => import('./pages/Auth'))
const TextSpeech = React.lazy(() => import("./pages/TextSpeech"))
const VoiceCloning = React.lazy(() => import('./pages/VoiceCloning'))
const Verification = React.lazy(() => import('./pages/Verification'))

const App = () => {

  return(
    <BrowserRouter basename="/">
        
    <Suspense fallback={<div>Loading...</div>}>

         <Routes>
                <Route path="/" Component={Auth} />
                <Route path="text-speech" Component={TextSpeech} />
                <Route path="voice-cloning" Component={VoiceCloning} />
                <Route path="verify" Component={Verification} />
         </Routes>

     </Suspense>

 </ BrowserRouter>
  )
  
}

export default App