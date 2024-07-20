"use client";

import { Input } from "@/components/ui/input";
import { z } from "zod";
import { CreateDoctorRequest, CreateDoctorResponse } from "@/types/doctor";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { FormEvent, FormEventHandler, useState } from "react";
import PasswordStrengthBar from "react-password-strength-bar";
import { Card } from "@/components/ui/card";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { CreatePatientRequest } from "@/types/patient";
import { Label } from "@/components/ui/label";
import { Toaster } from "@/components/ui/toaster";
import { useToast } from "@/components/ui/use-toast";
import { Spinner } from "@material-tailwind/react";
import axios, { AxiosResponse } from "axios";
import { GenericResponse } from "@/types/request";
import useLocalStorage from "use-local-storage";
import { LOCALSTORAGE_DOCTO_TOKEN } from "@/constants/local";
import { useRouter } from "next/navigation";
import { Loader2 } from "lucide-react";

interface SignupFormElemet {
  placeholder: string;
  formStateKey: string;
  label: string;
}

const doctorSignupForm: SignupFormElemet[] = [
  {
    formStateKey: "name",
    placeholder: "Name",
    label: "Name",
  },
  {
    formStateKey: "designation",
    placeholder: "Designation eg MBBS from ....",
    label: "Designation",
  },
  {
    formStateKey: "headline",
    placeholder: "Headline eg Chief of surgery at ..",
    label: "Headline",
  },
  {
    formStateKey: "email",
    placeholder: "Email",
    label: "Email",
  },
  {
    formStateKey: "password",
    placeholder: "Password",
    label: "Password",
  },
];

const patientSignupForm: SignupFormElemet[] = [
  {
    formStateKey: "name",
    placeholder: "Name",
    label: "Name",
  },
  {
    formStateKey: "email",
    placeholder: "Email",
    label: "Email",
  },
  {
    formStateKey: "password",
    placeholder: "Password",
    label: "Password",
  },
];

enum SignupType {
  PATIENT = "PATIENT",
  DOCTOR = "DOCTOR",
}

const CreateDoctorValidator = z
  .object({
    name: z.string().min(1, { message: "Name is required" }),
    designation: z.string().min(1, { message: "Designation is required" }),
    headline: z.string().min(1, { message: "Headline is required" }),
    email: z
      .string()

      .min(1, { message: "Email is required" })
      .email({ message: "Email not valid" }),
    password: z.string().min(1, { message: "Password is required" }),
  })
  .required({
    name: true,
    designation: true,
    email: true,
    headline: true,
    password: true,
  });

const CreatePatientValidator = z
  .object({
    name: z.string().min(1, { message: "Name is required" }),
    email: z
      .string()

      .min(1, { message: "Email is required" })
      .email({ message: "Email not valid" }),
    password: z.string().min(1, { message: "Password is required" }),
  })
  .required({
    name: true,
    email: true,
    password: true,
  });

function Page() {
  const router = useRouter();
  const { toast } = useToast();

  const [token, setToken] = useLocalStorage(LOCALSTORAGE_DOCTO_TOKEN, "");
  const [loading, setLoading] = useState<boolean>(false);
  const [doctorSignup, setDoctorSignup] = useState<CreateDoctorRequest>({
    designation: "",
    headline: "",
    name: "",
    password: "",
    email: "",
  });
  const [patientSignup, setPatientSignup] = useState<CreatePatientRequest>({
    name: "",
    password: "",
    email: "",
  });
  const [signupMode, setSignupMode] = useState<SignupType>(SignupType.DOCTOR);

  const handleDoctorSignup = async () => {
    const createDoctorValidation =
      CreateDoctorValidator.safeParse(doctorSignup);

    if (!createDoctorValidation.success) {
      return toast({
        variant: "destructive",
        description: createDoctorValidation.error.errors[0].message,
      });
    }

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<CreateDoctorResponse>> =
      await axios
        .post(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/doctor/create`,
          doctorSignup
        )
        .catch(() => {
          return toast({
            variant: "destructive",
            title: "Error",
            description:
              "Email already exists with another account. Use another email address",
          });
        });

    if (response.data && response.data.success) {
      toast({
        title: "Success",
        description: "Account Created",
      });

      setToken(response.data.details.token);
      router.push("/home");
    }
  };

  const handlePatientSignup = async () => {
    const createPatientValidator =
      CreatePatientValidator.safeParse(doctorSignup);

    if (!createPatientValidator.success) {
      return toast({
        variant: "destructive",
        description: createPatientValidator.error.errors[0].message,
      });
    }

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<CreateDoctorResponse>> =
      await axios
        .post(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/patient/create`,
          patientSignup
        )
        .catch(() => {
          return toast({
            variant: "destructive",
            title: "Error",
            description:
              "Email already exists with another account. Use another email address",
          });
        });

    if (response.data && response.data.success) {
      toast({
        title: "Success",
        description: "Account Created",
      });

      setToken(response.data.details.token);
      router.push("/home");
    }
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);

    if (signupMode === SignupType.DOCTOR) {
      await handleDoctorSignup();
    }

    if (signupMode === SignupType.PATIENT) {
      await handlePatientSignup();
    }

    setLoading(false);
  };

  return (
    <div className="flex gap-5 flex-wrap justify-center items-center min-h-[70dvh]">
      <Toaster />

      <div className="min-h-[70vh] w-[400px] my-20 flex flex-col">
        <Card className=" w-full p-5">
          <h2 className="text-1xl my-3 font-bold tracking-tight lg:text-2xl my-2">
            Signup
          </h2>

          <div>
            <form onSubmit={handleSubmit}>
              <Tabs
                defaultValue={SignupType.DOCTOR}
                onValueChange={(tab) => {
                  setSignupMode(tab as SignupType);
                }}
              >
                <TabsList>
                  <TabsTrigger value={SignupType.DOCTOR}>Doctor</TabsTrigger>
                  <TabsTrigger value={SignupType.PATIENT}>Patient</TabsTrigger>
                </TabsList>
                <TabsContent value={SignupType.DOCTOR}>
                  {doctorSignupForm.map((formElement) => (
                    <div
                      key={`DoctorSignup_${formElement.formStateKey}`}
                      className="my-2"
                    >
                      <Label htmlFor={formElement.formStateKey}>
                        {formElement.label}
                      </Label>
                      <Input
                        value={
                          doctorSignup[
                            formElement.formStateKey as keyof CreateDoctorRequest
                          ]
                        }
                        required
                        placeholder={formElement.placeholder}
                        onChange={(e) => {
                          setDoctorSignup({
                            ...doctorSignup,
                            [formElement.formStateKey]: e.currentTarget.value,
                          });
                        }}
                      />
                    </div>
                  ))}
                </TabsContent>
                <TabsContent value={SignupType.PATIENT}>
                  {patientSignupForm.map((formElement) => (
                    <div
                      key={`DoctorSignup_${formElement.formStateKey}`}
                      className="my-2"
                    >
                      <Label htmlFor={formElement.formStateKey}>
                        {formElement.label}
                      </Label>
                      <Input
                        value={
                          patientSignup[
                            formElement.formStateKey as keyof CreatePatientRequest
                          ]
                        }
                        required
                        placeholder={formElement.placeholder}
                        aria-errormessage="wfberg"
                        onChange={(e) => {
                          setPatientSignup({
                            ...doctorSignup,
                            [formElement.formStateKey]: e.currentTarget.value,
                          });
                        }}
                      />
                    </div>
                  ))}
                </TabsContent>
              </Tabs>

              <PasswordStrengthBar password={doctorSignup.password} />

              <Button disabled={loading} type="submit" className="w-full my-2">
                {/* @ts-ignore */}
                {loading ? (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  "Create Account"
                )}
              </Button>
              <Link href="/login">
                <Button variant="secondary" className="w-full my-2">
                  Login
                </Button>
              </Link>
            </form>
          </div>
        </Card>
      </div>
    </div>
  );
}

export default Page;
