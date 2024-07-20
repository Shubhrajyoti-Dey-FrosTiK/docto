import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { useToast } from "@/components/ui/use-toast";
import { GetPatientStruct, GetPatientWithConnections } from "@/types/patient";
import { GenericResponse } from "@/types/request";
import axios, { AxiosResponse } from "axios";
import { useContext, useEffect, useState } from "react";
import { BeatLoader } from "react-spinners";
import { UserState, UserStateContext } from "../layout";

function PatientViewer({ patientId }: { patientId: string }) {
  const [patient, setPatient] = useState<GetPatientStruct>();
  const [connected, setConnected] = useState<boolean>(true);
  const [loading, setLoading] = useState<boolean>(false);
  const { toast } = useToast();

  // @ts-ignore
  const [user] = useContext<[user: UserState]>(UserStateContext);

  const handleFetchPatient = async () => {
    setLoading(true);
    // @ts-ignore
    const response: AxiosResponse<GenericResponse<GetPatientWithConnections>> =
      await axios
        .get(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/patient/populated/connection?patientId=${patientId}`,
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
      setPatient(response.data.details.patient);
      setConnected(response.data.details.connected);
      setLoading(false);
    }
  };

  const handleAssignPatient = async () => {
    if (connected) return;

    setLoading(true);

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<any>> = await axios
      .post(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/assign/patient`,
        {
          patientId,
        },
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
      setConnected(true);
      setLoading(false);
      return toast({
        title: "Success",
        description: "The doctor is assigned to you successfully.",
      });
    } else {
      return toast({
        variant: "destructive",
        description: "Some unexpected error occurred",
      });
    }
  };

  useEffect(() => {
    handleFetchPatient();
  }, [patientId]);

  return (
    <div>
      {loading && <div className="text-center py-10">{<BeatLoader />}</div>}
      {!loading && (
        <div className="text-left">
          <div className="my-2">
            <div className="my-1">
              <Badge variant="outline">#{patientId}</Badge>
            </div>
            <div className="mx-1">
              <b>{patient?.name}</b>
              <p>Email: {patient?.email}</p>
            </div>
          </div>
          <Button
            variant="secondary"
            className="w-full mt-2"
            onClick={handleAssignPatient}
          >
            {connected ? "MESSAGE" : "CONNECT"}
          </Button>
        </div>
      )}
    </div>
  );
}

function Patient({ patient }: { patient: GetPatientStruct }) {
  return (
    <Dialog>
      <DialogTrigger>
        <Card className="py-2 px-5 my-2 md:w-[420px] w-[80vw] cursor-pointer text-left">
          <div className="flex justify-between">
            <p className="font-bold">{patient.name}</p>
            <Badge variant="outline">#{patient.id}</Badge>
          </div>
          <p className="font-light">{patient.email}</p>
        </Card>
      </DialogTrigger>
      <DialogContent className="max-w-[400px]">
        <DialogHeader>
          <DialogTitle>Patient Details</DialogTitle>
          <PatientViewer patientId={patient.id} />
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}

export default Patient;
