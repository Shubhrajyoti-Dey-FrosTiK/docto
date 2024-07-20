"use client";

import { Input } from "@/components/ui/input";
import { z } from "zod";
import {
  CreateDoctorRequest,
  CreateDoctorResponse,
  LoginUserRequest,
} from "@/types/doctor";
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

interface LoginFormElement {
  placeholder: string;
  formStateKey: string;
  label: string;
}

const userLoginForm: LoginFormElement[] = [
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

enum LoginType {
  PATIENT = "PATIENT",
  DOCTOR = "DOCTOR",
}

const LoginUserValidator = z
  .object({
    email: z
      .string()

      .min(1, { message: "Email is required" })
      .email({ message: "Email not valid" }),
    password: z.string().min(1, { message: "Password is required" }),
  })
  .required({
    email: true,
    password: true,
  });

function Page() {
  const router = useRouter();
  const { toast } = useToast();

  const [token, setToken] = useLocalStorage(LOCALSTORAGE_DOCTO_TOKEN, "");
  const [loading, setLoading] = useState<boolean>(false);
  const [userLogin, setUserLogin] = useState<LoginUserRequest>({
    password: "",
    email: "",
  });
  const [loginMode, setLoginMode] = useState<LoginType>(LoginType.DOCTOR);

  const handleDoctorLogin = async () => {
    const createDoctorValidation = LoginUserValidator.safeParse(userLogin);

    if (!createDoctorValidation.success) {
      return toast({
        variant: "destructive",
        description: createDoctorValidation.error.errors[0].message,
      });
    }

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<CreateDoctorResponse>> =
      await axios
        .post(`${process.env.NEXT_PUBLIC_BACKEND_URL}/doctor/login`, userLogin)
        .catch(() => {
          return toast({
            variant: "destructive",
            title: "Error",
            description: "Incorrect Credentials",
          });
        });

    if (response.data && response.data.success) {
      toast({
        title: "Success",
        description: "Logged in successfully",
      });

      setToken(response.data.details.token);
      router.push(`/home`);
    }
  };

  const handlePatientLogin = async () => {
    const createDoctorValidation = LoginUserValidator.safeParse(userLogin);

    if (!createDoctorValidation.success) {
      return toast({
        variant: "destructive",
        description: createDoctorValidation.error.errors[0].message,
      });
    }

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<CreateDoctorResponse>> =
      await axios
        .post(`${process.env.NEXT_PUBLIC_BACKEND_URL}/patient/login`, userLogin)
        .catch(() => {
          return toast({
            variant: "destructive",
            title: "Error",
            description: "Incorrect Credentials",
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

    if (loginMode === LoginType.DOCTOR) {
      await handleDoctorLogin();
    }

    if (loginMode === LoginType.PATIENT) {
      await handlePatientLogin();
    }

    setLoading(false);
  };

  return (
    <div className="flex gap-5 flex-wrap justify-center items-center min-h-[70dvh]">
      <Toaster />

      <div className="min-h-[70vh] w-[400px] my-20 flex flex-col">
        <Card className=" w-full p-5">
          <h2 className="text-1xl my-3 font-bold tracking-tight lg:text-2xl my-2">
            Login
          </h2>

          <div>
            <form onSubmit={handleSubmit}>
              <Tabs
                defaultValue={LoginType.DOCTOR}
                onValueChange={(tab) => {
                  setLoginMode(tab as LoginType);
                }}
              >
                <TabsList>
                  <TabsTrigger value={LoginType.DOCTOR}>Doctor</TabsTrigger>
                  <TabsTrigger value={LoginType.PATIENT}>Patient</TabsTrigger>
                </TabsList>
              </Tabs>

              {userLoginForm.map((formElement) => (
                <div
                  key={`UserLogin_${formElement.formStateKey}`}
                  className="my-2"
                >
                  <Label htmlFor={formElement.formStateKey}>
                    {formElement.label}
                  </Label>
                  <Input
                    value={
                      userLogin[
                        formElement.formStateKey as keyof LoginUserRequest
                      ]
                    }
                    required
                    placeholder={formElement.placeholder}
                    onChange={(e) => {
                      setUserLogin({
                        ...userLogin,
                        [formElement.formStateKey]: e.currentTarget.value,
                      });
                    }}
                  />
                </div>
              ))}

              <Button disabled={loading} type="submit" className="w-full my-2">
                {/* @ts-ignore */}
                {loading ? (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  "Login"
                )}
              </Button>
              <Link href="/signup">
                <Button variant="secondary" className="w-full my-2">
                  Create Account
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
