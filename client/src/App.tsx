import React, { Suspense } from "react";

import { BrowserRouter, Routes, Route } from "react-router-dom";

const TextSpeech = React.lazy(() => import("./pages/TextSpeech"))

const App = () => {

  return(
    <BrowserRouter basename="/">
        
    <Suspense fallback={<div>Loading...</div>}>

         <Routes>
                <Route path="/text-speech" Component={TextSpeech} />
         </Routes>

     </Suspense>

 </ BrowserRouter>
  )
  
}

export default App