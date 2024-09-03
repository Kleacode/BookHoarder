import type { StatusType } from "@/libs/schemas/schemaHealper";

export const ConvertStatusTypeToLabel = (type: StatusType) => {
	switch (type) {
		case "todo":
			return "積読";
		case "wip":
			return "読書中";
		case "done":
			return "読了";
		default:
			throw Error("invalid status type");
	}
};
