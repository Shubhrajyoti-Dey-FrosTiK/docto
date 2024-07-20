"use client";
import { ScrollArea } from "@/components/ui/scroll-area";
import { useToast } from "@/components/ui/use-toast";
import { UserType } from "@/types/auth";
import { GenericResponse } from "@/types/request";
import { GetAssociatedUsersResponse, User } from "@/types/user";
import axios, { AxiosResponse } from "axios";
import { useContext, useEffect, useState } from "react";
import { BeatLoader } from "react-spinners";
import { UserState, UserStateContext } from "../layout";
import { User as UserComponent } from "./user";

function Page() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const { toast } = useToast();

  const [user] =
    // @ts-ignore
    useContext<[user: UserState]>(UserStateContext);

  const fetchUsers = async () => {
    setLoading(true);
    // @ts-ignore
    const response: AxiosResponse<GenericResponse<GetAssociatedUsersResponse>> =
      await axios
        .get(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/${
            user.userType === UserType.DOCTOR
              ? "doctor/connectedPatients"
              : "patient/connectedDoctors"
          }`,
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
      setUsers(response.data.details.users);
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div>
      {loading && <div className="text-center py-10">{<BeatLoader />}</div>}
      {!loading && (
        <div>
          <div className="w-[400px]">
            <ScrollArea
              className="rounded-md border p-1"
              style={{
                height: "calc(100dvh - 180px)",
              }}
            >
              {users.map((user) => (
                <UserComponent key={user.id} user={user} />
              ))}
            </ScrollArea>
          </div>
        </div>
      )}
    </div>
  );
}

export default Page;
