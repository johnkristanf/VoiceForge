import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import '../public/tailwind.css'

import { VoiceCloneDataProvider } from './services/context/voiceContext.tsx'
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
         <VoiceCloneDataProvider>
                <App />
          </VoiceCloneDataProvider>
  </React.StrictMode>,
)
