import type { StatusType } from "@/libs/schemas/schemaHealper"
import { ConvertStatusTypeToLabel } from "@/utils/ConvertStatusType"

export interface StatusLabelProps {
    status: StatusType
};

export const StatusLabel = ({ status }: StatusLabelProps) => {
    return (
        <div>
            {ConvertStatusTypeToLabel(status)}
        </div>
    )
}