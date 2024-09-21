import { type AxiosAliasWrapper, api } from "@/libs/axios/axios";
import type { PathParameters } from "@/libs/schemas/schemaHealper";

export const deleteTag = async (
	pathParam: PathParameters<"/{userId}/tags/{tagId}", "delete">,
	reqParam: AxiosAliasWrapper<"/{userId}/tags/{tagId}", "delete">,
) => {
	try {
		const data = api.delete("/{userId}/tags/{tagId}", pathParam, reqParam);
		return data;
	} catch (e) {}
};
