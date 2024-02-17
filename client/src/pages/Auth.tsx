import React, { useRef, useState } from "react"
import { useForm } from "react-hook-form";
import { LoginCredentials, SignupCredentials, validEmail } from "../types/auth";
import { SignupValidation } from "../validator/auth";
import { Login, Signup } from "../services/http/post/auth";

const Auth: React.FC  = () => {
  const [Signup, setSignup] = useState<boolean>(false);

  if(Signup) return <SignupForm setSignup={setSignup} />

  return <LoginForm setSignup={setSignup} />

}


const SignupForm = ({setSignup}: any) => {

  const { register, reset, handleSubmit, formState: { errors } } = useForm<SignupCredentials>();
  const signupRef = useRef<HTMLParagraphElement | null>(null);

  const onSubmit = async (signUpCredentials: SignupCredentials) =>{
    
    const signup = await Signup(signUpCredentials)

    if(signup){

      if(signupRef.current) signupRef.current.textContent = 'Signup Successfully';
      reset();

      setTimeout(() => {
        setSignup(false)
      }, 1500)
    
    }
     
  }

  return(
      
<section className="h-full">
<div className="container h-full p-10">
  <div
    className="g-6 flex h-full flex-wrap items-center justify-center text-neutral-800 dark:text-neutral-200">
    <div className="w-full">

      <div
        className="block rounded-lg">
        <div className="g-0 lg:flex lg:flex-wrap">
          
          <div className="px-4 md:px-0 lg:w-6/12">
            <div className="md:mx-6 md:p-12 text-white bg-slate-900 rounded-md">
             
              <div className="text-center">
                <img
                  className="mx-auto w-24 rounded-full"
                  src="/img/VF_logo.png"
                  alt="logo" />
                <h4 className="mb-12 mt-1 pb-1 text-xl font-semibold">
                  Welcome to VoiceForge
                </h4>
              </div>

              <form onSubmit={handleSubmit(onSubmit)} >
                <p className="mb-4">Signup to get Started</p>

                <p ref={signupRef} className="text-green-500 text-center font-bold text-lg mb-3"></p>
                { SignupValidation(errors) }

               
                <div className="relative mb-4" data-te-input-wrapper-init>
                  <label className="font-bold">Email:</label>

                  <input
                    type="text"
                    className="peer block bg-slate-700 mt-5 min-h-[auto] w-full rounded border-0 px-3 py-[0.32rem] leading-[1.6] outline-none transition-all duration-200 ease-linear focus:placeholder:opacity-100 data-[te-input-state-active]:placeholder:opacity-100 motion-reduce:transition-none dark:placeholder:text-neutral-200 [&:not([data-te-input-placeholder-active])]:placeholder:opacity-0 focus:outline-blue-800"
                    id="exampleFormControlInput1"
                    placeholder="Email" 
                    {...register("email", { required: true, pattern: validEmail })}
                    
                    />
                    
                </div>

                
                <div className="relative mb-4" data-te-input-wrapper-init>
                  <label className="font-bold">Password:</label>

                  <input
                    type="password"
                    className="peer block bg-slate-700 mt-5 min-h-[auto] w-full rounded border-0 px-3 py-[0.32rem] leading-[1.6] outline-none transition-all duration-200 ease-linear focus:placeholder:opacity-100 data-[te-input-state-active]:placeholder:opacity-100 motion-reduce:transition-none dark:placeholder:text-neutral-200 [&:not([data-te-input-placeholder-active])]:placeholder:opacity-0 focus:outline-blue-800"
                    id="exampleFormControlInput11"
                    placeholder="Password" 
                    {...register("password", { required: true, minLength: 8})}

                    />
                 
                </div>

               
                <div className="mb-12 pb-1 pt-1 text-center">
                  <button
                    className="mb-3 inline-block w-full rounded px-6 pb-2 pt-2.5 text-xs font-medium uppercase bg-indigo-700 leading-normal text-white shadow-[0_4px_9px_-4px_rgba(0,0,0,0.2)] transition duration-150 ease-in-out hover:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)] focus:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)] focus:outline-none focus:ring-0 active:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)]"
                    type="submit"
                    data-te-ripple-init
                    data-te-ripple-color="light"
                    
                    >
                    Sign up
                  </button>

                 
                 
                </div>

               
                <div className="flex items-center justify-between pb-6">
                  <p className="mb-0 mr-2">Already have an account?</p>
                  <button
                    onClick={() => setSignup(false)}
                    type="button"
                    className="inline-block rounded border-2 border-danger px-6 pb-[6px] pt-2 text-xs font-medium uppercase leading-normal text-danger transition duration-150 ease-in-out hover:border-danger-600 hover:bg-neutral-500 hover:bg-opacity-10 hover:text-danger-600 focus:border-danger-600 focus:text-danger-600 focus:outline-none focus:ring-0 active:border-danger-700 active:text-danger-700 dark:hover:bg-neutral-100 dark:hover:bg-opacity-10"
                    data-te-ripple-init
                    data-te-ripple-color="light">
                    Login
                  </button>
                </div>

              </form>
            </div>
          </div>

          
          <div className="flex items-center rounded-b-lg lg:w-6/12 lg:rounded-r-lg lg:rounded-bl-none bg-slate-700 rounded-md">
           
            <div className="px-4 py-6 text-white md:mx-6 md:p-12">
              <h4 className="mb-6 text-4xl font-semibold">
                Company VoiceForge
              </h4>
              <p className="text-md">
                We are a dynamic force of individuals united by a common goal to drive innovation and make a positive impact. Our team is fueled by creativity, passion, and a relentless pursuit of excellence. Together, we strive to push boundaries, solve challenges, and deliver unparalleled value to our clients and partners.
                <br /><br />
                In our journey, we embrace diversity, encourage collaboration, and foster an environment where every voice is heard. Our commitment goes beyond business; it extends to building lasting relationships and contributing meaningfully to the communities we serve. We believe that by pushing the boundaries of what's possible, we can shape a future that is not only successful but also transformative.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</section>
  )
}



const LoginForm = ({setSignup}: any) => {

  const { register, reset, handleSubmit, formState: { errors } } = useForm<LoginCredentials>();
  const loginRef = useRef<HTMLParagraphElement | null>(null);
  const [isSubmitting, setisSubmitting] = useState<boolean>(false)

  const onSubmit = async (loginCredentials: LoginCredentials) => {

    setisSubmitting(true)
    
    const login = await Login(loginCredentials)

    if(login.Need_Verification){
      window.location.href = '/verify'
      reset()
      setisSubmitting(false)
      return
    }

    if(login.Invalid_Credentials){
      if(loginRef.current) loginRef.current.textContent = login.Invalid_Credentials;
      setisSubmitting(false)
      return
    }

    window.location.href = '/text-speech';
    reset()
    setisSubmitting(false)
  }


  return(
      
<section className="h-full">
<div className="container h-full p-10">
  <div
    className="g-6 flex h-full flex-wrap items-center justify-center text-neutral-800 dark:text-neutral-200">
    <div className="w-full">

      <div
        className="block rounded-lg">
        <div className="g-0 lg:flex lg:flex-wrap">
          
          <div className="px-4 md:px-0 lg:w-6/12">
            <div className="md:mx-6 md:p-12 text-white bg-slate-900 rounded-md">
             
              <div className="text-center">
                <img
                  className="mx-auto w-24 rounded-full"
                  src="/img/VF_logo.png"
                  alt="logo" />
                <h4 className="mb-12 mt-1 pb-1 text-xl font-semibold">
                  Welcome to VoiceForge
                </h4>
              </div>

              <form  onSubmit={handleSubmit(onSubmit)} >
                <p className="mb-4">Please login to your account</p>
               

                <p ref={loginRef} className="text-red-800 text-center text-lg font-bold mb-3"></p>
                { SignupValidation(errors) }

                <div className="relative mb-4" data-te-input-wrapper-init>
                  <label className="font-bold">Email:</label>
                  <input
                    type="text"
                    className="peer block mt-5 min-h-[auto] w-full rounded border-0 bg-slate-700 px-3 py-[0.32rem] leading-[1.6] outline-none transition-all duration-200 ease-linear focus:placeholder:opacity-100 data-[te-input-state-active]:placeholder:opacity-100 motion-reduce:transition-none dark:placeholder:text-white [&:not([data-te-input-placeholder-active])]:placeholder:opacity-0 focus:outline-blue-800"
                    id="exampleFormControlInput1"
                    placeholder="Email" 
                    {...register("email", { required: true })}
                    />
                 
                </div>

                
                <div className="relative mb-4" data-te-input-wrapper-init>
                   <label className="font-bold">Password:</label>
                  <input
                    type="password"
                    className="peer mt-5 block min-h-[auto] w-full rounded border-0 bg-slate-700 px-3 py-[0.32rem] leading-[1.6] outline-none transition-all duration-200 ease-linear focus:placeholder:opacity-100 data-[te-input-state-active]:placeholder:opacity-100 motion-reduce:transition-none dark:placeholder:text-neutral-200 [&:not([data-te-input-placeholder-active])]:placeholder:opacity-0 focus:outline-blue-800"
                    id="exampleFormControlInput11"
                    placeholder="Password" 
                    {...register("password", { required: true })}/>
                 
                </div>

               
                <div className="mb-12 pb-1 pt-1 text-center">
                  <button
                    disabled={isSubmitting}
                    className={ isSubmitting ? "mb-3 inline-block w-full rounded px-6 pb-2 pt-2.5 text-xs font-medium uppercase bg-gray-400 cursor-no-drop opacity-75 leading-normal text-white shadow-[0_4px_9px_-4px_rgba(0,0,0,0.2)] transition duration-150 ease-in-out hover:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)] focus:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)] focus:outline-none focus:ring-0 active:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)]"
                                 : "mb-3 inline-block w-full rounded px-6 pb-2 pt-2.5 text-xs font-medium uppercase bg-indigo-700 hover:opacity-75 leading-normal text-white shadow-[0_4px_9px_-4px_rgba(0,0,0,0.2)] transition duration-150 ease-in-out hover:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)] focus:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)] focus:outline-none focus:ring-0 active:shadow-[0_8px_9px_-4px_rgba(0,0,0,0.1),0_4px_18px_0_rgba(0,0,0,0.2)]"
                    }
                    type="submit"
                    data-te-ripple-init
                    data-te-ripple-color="light"
                    
                    >
                    Log in
                  </button>

                
                </div>

               
                <div className="flex items-center justify-between pb-6">
                  <p className="mb-0 mr-2">Don't have an account?</p>
                  <button
                    onClick={() => setSignup(true)}
                    type="button"
                    className="inline-block rounded border-2 border-danger px-6 pb-[6px] pt-2 text-xs font-medium uppercase leading-normal text-danger transition duration-150 ease-in-out hover:border-danger-600 hover:bg-neutral-500 hover:bg-opacity-10 hover:text-danger-600 focus:border-danger-600 focus:text-danger-600 focus:outline-none focus:ring-0 active:border-danger-700 active:text-danger-700 dark:hover:bg-neutral-100 dark:hover:bg-opacity-10"
                    data-te-ripple-init
                    data-te-ripple-color="light">
                    Sign up
                  </button>
                </div>

              </form>
            </div>
          </div>

          
          <div className="flex items-center rounded-b-lg lg:w-6/12 lg:rounded-r-lg lg:rounded-bl-none bg-slate-700 rounded-md">
           
            <div className="px-4 py-6 text-white md:mx-6 md:p-12">
              <h4 className="mb-6 text-4xl font-semibold">
                Company VoiceForge
              </h4>
              <p className="text-md">
                We are a dynamic force of individuals united by a common goal to drive innovation and make a positive impact. Our team is fueled by creativity, passion, and a relentless pursuit of excellence. Together, we strive to push boundaries, solve challenges, and deliver unparalleled value to our clients and partners.
                <br /><br />
                In our journey, we embrace diversity, encourage collaboration, and foster an environment where every voice is heard. Our commitment goes beyond business; it extends to building lasting relationships and contributing meaningfully to the communities we serve. We believe that by pushing the boundaries of what's possible, we can shape a future that is not only successful but also transformative.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</section>
  )
}

export default Auth