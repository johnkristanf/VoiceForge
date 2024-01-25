import { FieldErrors } from "react-hook-form"
import { SignupCredentials } from "../types/auth"

export const SignupValidation = (errors: FieldErrors<SignupCredentials>) => {

    if (errors.email?.type === 'required') return <p className="text-red-800 text-center text-lg font-bold mb-3">Email is Required</p>

    if(errors.email?.type === 'pattern') return <p className="text-red-800 text-center text-lg font-bold mb-3">Invalid Email Address</p>

    if(errors.password?.type === 'required') return <p className="text-red-800 text-center text-lg font-bold mb-3">Password is Required</p>

    if(errors.password?.type === 'minLength') return <p className="text-red-800 text-center text-lg font-bold mb-3">Password must have atleast 8 characters</p>
}