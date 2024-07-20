"use client";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";
import { useContext } from "react";
import { UserStateContext } from "./layout";
import { UserType } from "@/types/auth";

function Page() {
  // @ts-ignore
  const [user] = useContext<[user: UserState]>(UserStateContext);

  return (
    <div className="flex flex-col justify-center items-center h-[80dvh] w-full text-center">
      <div className="flex flex-col gap-2 items-center">
        <h1 className="text-4xl my-2 font-extrabold tracking-tight lg:text-5xl">
          Welcome to Docto !
        </h1>
        <h2 className="text-3xl font-mediumbold tracking-tight lg:text-4xl">
          You have successfully logged in !
        </h2>
        <Image
          alt="Doctor Illustration"
          src={"/images/doctor2.jpg"}
          width={400}
          height={400}
        />
        <div className="flex items-center gap-5">
          <Link href="/home/search">
            <Button>
              Search {user.userType === UserType.DOCTOR ? "Patient" : "Doctor"}
            </Button>
          </Link>
          <Link href="/home/connections">
            <Button variant="secondary">
              View connected{" "}
              {user.userType === UserType.DOCTOR ? "Patient" : "Doctor"}
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
}

export default Page;
