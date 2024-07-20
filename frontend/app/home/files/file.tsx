import { Card } from "@/components/ui/card";
import { File as FileType } from "@/types/file";
import { IconFile3d } from "@tabler/icons-react";
import React from "react";

export function File({ file }: { file: FileType }) {
  return (
    <div className="max-w-[500px]">
      <a target="__blank" rel="noopener noreferrer" href={file.url}>
        <Card className="px-4 py-2">
          <div className="flex items-center gap-2">
            <IconFile3d />
            <div>
              <b>{file.fileName}</b>
              <p>
                Updated At:{" "}
                {new Date(file.updatedAt * 1000).toLocaleDateString()}
              </p>
            </div>
          </div>
          <iframe className="rounded-md my-2" src={file.url} />
        </Card>
      </a>
    </div>
  );
}
