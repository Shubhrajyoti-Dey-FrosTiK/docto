"use client";
import { IconFiles, IconSearch, IconUsersGroup } from "@tabler/icons-react";
import Link from "next/link";
import { useState } from "react";

export type SelectedType = "search" | "connections" | "files" | "home";

export function Navbar({
  navigationControl,
  selectedType,
  loggedIn,
}: {
  navigationControl?: boolean;
  selectedType?: SelectedType;
  loggedIn?: boolean;
}) {
  return (
    <div className="flex items-center justify-between">
      <Link href={loggedIn ? "/home" : "/"}>
        <h1 className="text-3xl my-2 font-extrabold tracking-tight lg:text-4xl">
          Docto
        </h1>
      </Link>
      {navigationControl && <NavigationControl selectedType={selectedType} />}
    </div>
  );
}

export function NavigationControl({
  vertical,
  selectedType,
}: {
  vertical?: boolean;
  selectedType?: SelectedType;
}) {
  const [selected, setSelected] = useState<SelectedType>(
    selectedType || "search"
  );

  const iconStyles = (selected: boolean) => {
    return `flex items-center justify-center cursor-pointer select-none font-sans font-medium text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none w-10 max-w-[40px] h-10 max-h-[40px] rounded-lg text-xs ${
      selected ? "bg-gray-800" : "bg-gray-100"
    } text-white shadow-md shadow-gray-900/10 hover:shadow-lg hover:shadow-gray-900/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none`;
  };

  return (
    <div
      className={`flex ${
        vertical && "flex-col"
      } items-center justify-center gap-5`}
    >
      <Link href="/home/search">
        <div
          className={iconStyles(selected == "search")}
          onClick={() => setSelected("search")}
        >
          <IconSearch color={selected == "search" ? "white" : "black"} />
        </div>
      </Link>
      <Link href="/home/connections">
        <div
          className={iconStyles(selected == "connections")}
          onClick={() => setSelected("connections")}
        >
          <IconUsersGroup
            color={selected == "connections" ? "white" : "black"}
          />
        </div>
      </Link>
      <Link href="/home/files">
        <div
          className={iconStyles(selected == "files")}
          onClick={() => setSelected("files")}
        >
          <IconFiles color={selected == "files" ? "white" : "black"} />
        </div>
      </Link>
    </div>
  );
}
