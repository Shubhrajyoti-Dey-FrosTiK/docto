"use client";

import { Input as InputComponent, InputProps } from "@material-tailwind/react";

import { ThemeProvider } from "@material-tailwind/react";

export { ThemeProvider };

export function Input(props: InputProps): React.ReactNode {
  // @ts-ignore
  return <InputComponent {...props} />;
}
