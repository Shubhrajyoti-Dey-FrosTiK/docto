"use client";
import { Navbar, NavigationControl } from "@/components/navbar/Navbar";
import { useViewportSize } from "@mantine/hooks";
import React, { createContext, useEffect, useState } from "react";
import { usePathname } from "next/navigation";
import useLocalStorage from "use-local-storage";
import { LOCALSTORAGE_DOCTO_TOKEN } from "@/constants/local";
import { TokenVerifyResponse, UserType } from "@/types/auth";
import axios, { AxiosResponse } from "axios";
import { GenericResponse } from "@/types/request";
import { useRouter } from "next/navigation";
import { ScaleLoader } from "react-spinners";

export interface UserState {
  userType: UserType;
  token?: string;
  userId?: string;
}

export const UserStateContext = createContext<UserState>({
  userType: UserType.UNKNOWN,
});

function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter();
  const pathname = usePathname();
  const { width } = useViewportSize();
  const [token, setToken] = useLocalStorage(LOCALSTORAGE_DOCTO_TOKEN, "");
  const [loading, setLoading] = useState<boolean>(true);

  const userState = useState<UserState>({
    userType: UserType.UNKNOWN,
  });

  const verifyToken = async () => {
    setLoading(true);
    const response: AxiosResponse<GenericResponse<TokenVerifyResponse>> =
      await axios.get(`${process.env.NEXT_PUBLIC_BACKEND_URL}/token/verify`, {
        headers: { Authorization: `Bearer ${token}` },
      });

    if (response.data && response.data.success) {
      setToken(response.data.details.token);
      userState[1]({
        userType: response.data.details.role as UserType,
        token: response.data.details.token,
        userId: response.data.details.userId,
      });
    } else {
      router.push("/");
    }
    setLoading(false);
  };

  useEffect(() => {
    verifyToken();
  }, [token]);

  return (
    <div className="p-5 w-full h-[100dvh]">
      {/* @ts-ignore */}
      <UserStateContext.Provider value={userState}>
        <Navbar
          loggedIn
          navigationControl={pathname != "/home" && width < 700}
          selectedType={
            pathname.split("/").length > 2
              ? (pathname.split("/")[2] as any)
              : "home"
          }
        />

        {!loading && (
          <div className="flex">
            {pathname != "/home" && width >= 700 && (
              <div className="py-10">
                <NavigationControl
                  vertical
                  selectedType={
                    pathname.split("/").length > 2
                      ? (pathname.split("/")[2] as any)
                      : "home"
                  }
                />
              </div>
            )}
            <div
              className={`w-full h-full ${width >= 700 ? "ml-10 mt-10" : ""}`}
            >
              {children}
            </div>
          </div>
        )}

        {loading && (
          <div className="h-full w-full flex justify-center items-center">
            <ScaleLoader />
          </div>
        )}
      </UserStateContext.Provider>
    </div>
  );
}

export default Layout;
