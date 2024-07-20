import { Navbar } from "@/components/navbar/Navbar";
import React from "react";

function layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="p-5">
      <Navbar /> {children}
    </div>
  );
}

export default layout;
