import { api } from "@/libs/axios/axios";
import type { PathParameters, RequestData } from "@/libs/schemas/schemaHealper";

export const postHoarder = async (
	param: PathParameters<"/{userId}/hoarder", "post">,
	body: RequestData<"/{userId}/hoarder", "post">,
) => {
	try {
		const data = api.post("/{userId}/hoarder", param, body);
		return data;
	} catch (e) {}
};
