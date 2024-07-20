"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import React, {
  FormEvent,
  FormEventHandler,
  useContext,
  useState,
} from "react";
import { UserState, UserStateContext } from "../layout";
import { UserType } from "@/types/auth";
import { Toaster } from "@/components/ui/toaster";
import axios, { AxiosResponse } from "axios";
import { GenericResponse } from "@/types/request";
import { DoctorSearchResponse, DoctorSearchStruct } from "@/types/doctor";
import { useToast } from "@/components/ui/use-toast";
import { Loader2 } from "lucide-react";
import { BeatLoader } from "react-spinners";
import Doctor from "./doctor";
import { ScrollArea } from "@/components/ui/scroll-area";
import Patient from "./patient";
import { GetPatientStruct, SearchPatientResponse } from "@/types/patient";

function Page() {
  const [user] =
    // @ts-ignore
    useContext<[user: UserState]>(UserStateContext);
  const { toast } = useToast();
  const [userQuery, setUserQuery] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [doctors, setDoctors] = useState<DoctorSearchStruct[]>([]);
  const [patients, setPatients] = useState<GetPatientStruct[]>([]);

  const handleFetchDoctors = async () => {
    // @ts-ignore
    const response: AxiosResponse<GenericResponse<DoctorSearchResponse>> =
      await axios
        .get(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/doctor/search?searchQuery=${userQuery}`,
          {
            headers: { Authorization: `Bearer ${user.token}` },
          }
        )
        .catch(() => {
          return toast({
            variant: "destructive",
            description: "Some unexpected error occurred",
          });
        });

    if (response.data && response.data.success) {
      setDoctors(response.data.details.doctors ?? []);
    }
  };

  const handleFetchPatients = async () => {
    // @ts-ignore
    const response: AxiosResponse<GenericResponse<SearchPatientResponse>> =
      await axios
        .get(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/patient/search?searchQuery=${userQuery}`,
          {
            headers: { Authorization: `Bearer ${user.token}` },
          }
        )
        .catch(() => {
          return toast({
            variant: "destructive",
            description: "Some unexpected error occurred",
          });
        });

    if (response.data && response.data.success) {
      setPatients(response.data.details.patients ?? []);
    }
  };

  const handleFetchUsers = async (e: FormEvent<any>) => {
    e.preventDefault();

    if (userQuery === "") {
      return toast({
        variant: "destructive",
        description: "Search field can't be empty",
      });
    }

    setLoading(true);

    if (user.userType === UserType.DOCTOR) {
      await handleFetchPatients();
    } else {
      await handleFetchDoctors();
    }

    setLoading(false);
  };

  return (
    <div className="pt-10">
      <Toaster />
      <form
        onSubmit={handleFetchUsers}
        className="w-full text-center flex flex-col w-full items-center "
      >
        <h2 className="text-3xl py-10 font-mediumbold tracking-tight lg:text-4xl">
          Search {user.userType === UserType.DOCTOR ? "patients" : "doctors"}
        </h2>
        <div className="flex w-full max-w-[480px] items-center space-x-2">
          <Input
            value={userQuery}
            onChange={(e) => {
              setUserQuery(e.currentTarget.value);
            }}
            placeholder="Search"
          />
          <Button type="submit" onClick={handleFetchUsers}>
            Search
          </Button>
        </div>
      </form>
      {loading && <div className="text-center py-10">{<BeatLoader />}</div>}

      {!loading && (
        <div>
          {user.userType === UserType.PATIENT && (
            <div className="flex justify-center">
              {!doctors.length ? (
                <div>
                  <p className="text-xl py-10 font-mediumbold tracking-tight text-center">
                    No doctors found
                  </p>
                </div>
              ) : (
                <div className="max-w-[480px] mt-2">
                  <ScrollArea
                    className="rounded-md border p-1"
                    style={{
                      maxHeight: "calc(100dvh - 320px)",
                    }}
                  >
                    <div className="inline-flex flex-col justify-center items-center">
                      {doctors.map((doctor) => (
                        <Doctor key={doctor.id} doctor={doctor} />
                      ))}
                    </div>
                  </ScrollArea>
                </div>
              )}
            </div>
          )}

          {user.userType === UserType.DOCTOR && (
            <div className="flex justify-center">
              {!patients.length ? (
                <div>
                  <p className="text-xl py-10 font-mediumbold tracking-tight text-center">
                    No patients found
                  </p>
                </div>
              ) : (
                <div className="max-w-[480px] mt-2">
                  <ScrollArea
                    className="rounded-md border p-1"
                    style={{
                      maxHeight: "calc(100dvh - 320px)",
                    }}
                  >
                    <div className="inline-flex flex-col justify-center items-center">
                      {patients.map((patient) => (
                        <Patient key={patient.id} patient={patient} />
                      ))}
                    </div>
                  </ScrollArea>
                </div>
              )}
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default Page;
