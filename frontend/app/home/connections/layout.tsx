import React from "react";

function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex flex-col justify-center items-center pt-10 px-2">
      <div>
        <h2 className="text-3xl font-mediumbold tracking-tight lg:text-4xl mb-5">
          Connections
        </h2>
      </div>
      {children}
    </div>
  );
}

export default Layout;
