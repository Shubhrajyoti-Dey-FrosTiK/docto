import { Button } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <main className="flex items-center h-[90dvh] justify-center">
      <div className="flex gap-5 flex-wrap">
        <div className="max-w-[400px]">
          <div className="my-5">
            <h1 className="text-4xl my-2 font-extrabold tracking-tight lg:text-5xl">
              Docto
            </h1>
            <h2 className="text-3xl font-mediumbold tracking-tight lg:text-4xl">
              One platform to for all the doctors and the patients
            </h2>
          </div>

          <div className="flex items-center gap-5">
            <Link href="/login">
              <Button>Login</Button>
            </Link>
            <Link href="/signup">
              <Button variant="secondary">Create Account</Button>
            </Link>
          </div>
        </div>
        <div className="flex items-center">
          <Image
            alt="Doctor Illustration"
            src={"/images/doctor.jpg"}
            width={400}
            height={400}
          />
        </div>
      </div>
    </main>
  );
}
