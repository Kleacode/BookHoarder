import { api } from "@/libs/axios/axios";
import type { PathParameters, RequestData } from "@/libs/schemas/schemaHealper";

export const postTag = async (
	param: PathParameters<"/{userId}/tags", "post">,
	body: RequestData<"/{userId}/tags", "post">,
) => {
	try {
		const data = api.post("/{userId}/tags", param, body);
		return data;
	} catch (e) {}
};
