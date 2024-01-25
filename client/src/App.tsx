import React, { Suspense } from "react";

import { BrowserRouter, Routes, Route } from "react-router-dom";

const TextSpeech = React.lazy(() => import("./pages/TextSpeech"))
const VoiceCloning = React.lazy(() => import('./pages/VoiceCloning'))

const App = () => {

  return(
    <BrowserRouter basename="/">
        
    <Suspense fallback={<div>Loading...</div>}>

         <Routes>
                <Route path="text-speech" Component={TextSpeech} />
                <Route path="voice-cloning" Component={VoiceCloning} />
         </Routes>

     </Suspense>

 </ BrowserRouter>
  )
  
}

export default App