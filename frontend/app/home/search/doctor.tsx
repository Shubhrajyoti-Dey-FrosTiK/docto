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
import {
  DoctorSearchStruct,
  GetDoctorStruct,
  GetDoctorWithConnectionsResponse,
} from "@/types/doctor";
import { GenericResponse } from "@/types/request";
import axios, { AxiosResponse } from "axios";
import { useContext, useEffect, useState } from "react";
import { BeatLoader } from "react-spinners";
import { UserState, UserStateContext } from "../layout";

function DoctorViewer({ doctorId }: { doctorId: string }) {
  const [doctor, setDoctor] = useState<GetDoctorStruct>();
  const [connected, setConnected] = useState<boolean>(true);
  const [loading, setLoading] = useState<boolean>(false);
  const { toast } = useToast();

  // @ts-ignore
  const [user] = useContext<[user: UserState]>(UserStateContext);

  const handleFetchDoctor = async () => {
    setLoading(true);
    // @ts-ignore
    const response: AxiosResponse<
      GenericResponse<GetDoctorWithConnectionsResponse>
    > = await axios
      .get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/doctor/populated/connection?doctorId=${doctorId}`,
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
      setDoctor(response.data.details.doctor);
      setConnected(response.data.details.connected);
      setLoading(false);
    }
  };

  const handleAssignDoctor = async () => {
    if (connected) return;

    setLoading(true);

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<any>> = await axios
      .post(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/assign/doctor`,
        {
          doctorId,
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
    handleFetchDoctor();
  }, [doctorId]);

  return (
    <div>
      {loading && <div className="text-center py-10">{<BeatLoader />}</div>}
      {!loading && (
        <div className="text-left">
          <div className="my-2">
            <div className="my-1">
              <Badge variant="outline">#{doctorId}</Badge>
            </div>
            <div className="mx-1">
              <b>{doctor?.name}</b>
              <p>{doctor?.headline}</p>
              <p>{doctor?.designation}</p>
              <p>Email: {doctor?.email}</p>
            </div>
          </div>
          <Button
            variant="secondary"
            className="w-full mt-2"
            onClick={handleAssignDoctor}
          >
            {connected ? "MESSAGE" : "CONNECT"}
          </Button>
        </div>
      )}
    </div>
  );
}

function Doctor({ doctor }: { doctor: DoctorSearchStruct }) {
  return (
    <Dialog>
      <DialogTrigger>
        <Card className="py-2 px-5 my-2 md:w-[420px] w-[80vw] cursor-pointer text-left">
          <div className="flex justify-between">
            <p className="font-bold">{doctor.name}</p>
            <Badge variant="outline">#{doctor.id}</Badge>
          </div>
          <p className="font-light">{doctor.headline}</p>
        </Card>
      </DialogTrigger>
      <DialogContent className="max-w-[400px]">
        <DialogHeader>
          <DialogTitle>Doctor Details</DialogTitle>
          <DoctorViewer doctorId={doctor.id} />
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}

export default Doctor;
