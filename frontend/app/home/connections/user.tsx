import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { User as UserType } from "@/types/user";

export function User({ user }: { user: UserType }) {
  return (
    <Card className="py-2 px-4 cursor-pointer">
      <div className="flex justify-between">
        <div>
          <b>{user.name}</b>
          <p>{user.email}</p>
        </div>
        <Badge variant="outline">#{user.id}</Badge>
      </div>
    </Card>
  );
}
