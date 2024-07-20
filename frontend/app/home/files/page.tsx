"use client";
import { Card, CardHeader } from "@/components/ui/card";
import { GenericResponse } from "@/types/request";
import { IconFile3d } from "@tabler/icons-react";
import axios, { AxiosResponse } from "axios";
import React, { useContext, useEffect, useState } from "react";
import { UserState, UserStateContext } from "../layout";
import { UserType } from "@/types/auth";
import { useToast } from "@/components/ui/use-toast";
import { Progress } from "@/components/ui/progress";
import { Toaster } from "@/components/ui/toaster";
import { File as FileType } from "@/types/file";
import { BeatLoader } from "react-spinners";
import { File as FileComponent } from "./file";

function Page() {
  // @ts-ignore
  const [user] = useContext<[user: UserState]>(UserStateContext);
  const [progress, setProgress] = useState<number>(0);
  const [submitting, setSubmitting] = useState<boolean>(false);
  const [files, setFiles] = useState<FileType[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const { toast } = useToast();

  const handleFileUpload = async (filesChosen: File[]) => {
    if (!filesChosen.length)
      return toast({
        variant: "destructive",
        description: "Please select a file to upload",
      });

    setSubmitting(true);
    const formData = new FormData();

    // @ts-ignore
    filesChosen.forEach((file) => formData.append("file", file));

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<FileType[]>> = await axios
      .post(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/${
          user.userType === UserType.DOCTOR ? "doctor" : "patient"
        }/upload`,
        formData,
        {
          headers: {
            Authorization: `Bearer ${user.token}`,
            "Content-Type": "multipart/form-data",
          },
          onUploadProgress(currentProgress) {
            console.log(currentProgress);
            setProgress(
              Math.round(
                (100 * currentProgress.loaded) / (currentProgress.total ?? 1)
              )
            );
          },
        }
      )
      .catch(() =>
        toast({
          variant: "destructive",
          description: "Some unexpected error occurred",
        })
      );

    if (response.data && response.data.success) {
      setFiles([...files, ...response.data.details]);
      toast({
        description: "File uploaded successfully",
      });
    }

    setSubmitting(false);
  };

  const handleFetchFiles = async () => {
    setLoading(true);

    // @ts-ignore
    const response: AxiosResponse<GenericResponse<FileType[]>> = await axios
      .get(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/${
          user.userType === UserType.DOCTOR ? "doctor" : "patient"
        }/files`,
        {
          headers: {
            Authorization: `Bearer ${user.token}`,
          },
        }
      )
      .catch(() =>
        toast({
          variant: "destructive",
          description: "Some unexpected error occurred",
        })
      );

    if (response.data && response.data.success) {
      setFiles(response.data.details);
      setLoading(false);
    } else
      return toast({
        variant: "destructive",
        description: "Some unexpected error occurred",
      });
  };

  useEffect(() => {
    handleFetchFiles();
  }, []);

  return (
    <div className="pt-5">
      <Toaster />
      <h2 className="text-3xl font-mediumbold tracking-tight lg:text-4xl mb-5">
        Files
      </h2>
      <div>
        {/* @ts-ignore */}
        <label for="uploadFile">
          <Card className="flex px-4 py-2 gap-2 items-center cursor-pointer max-w-[400px]">
            <IconFile3d />
            <h4 className="text-md font-mediumbold tracking-tight lg:text-xl">
              Upload Files
            </h4>
          </Card>
        </label>
        <input
          onChange={(e) =>
            handleFileUpload(Array.from(e.currentTarget.files ?? []))
          }
          type="file"
          required
          id="uploadFile"
          className="hidden"
        />
      </div>

      {submitting && (
        <div className="my-2">
          <Progress value={progress} className="max-w-[400px] h-1" />
        </div>
      )}

      {loading && <div className="text-center py-10">{<BeatLoader />}</div>}

      {!loading && (
        <div className="flex gap-4 flex-wrap mt-4">
          {files.map((file) => (
            <FileComponent key={`File_${file.id}`} file={file} />
          ))}
        </div>
      )}
    </div>
  );
}

export default Page;
