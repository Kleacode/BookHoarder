import { type AxiosAliasWrapper, api } from "@/libs/axios/axios";
import type {
	PathParameters,
	ResponseData,
} from "@/libs/schemas/schemaHealper";
import { useEffect, useState } from "react";

export type responseGetTags = ResponseData<"/{userId}/tags", "get">;

export const useGetSuggestTags = (
	pathParam: PathParameters<"/{userId}/tags", "get">,
	reqParam: AxiosAliasWrapper<"/{userId}/tags", "get">,
	searchTerm: string,
) => {
	const [tags, setTags] = useState<responseGetTags>([]);

	const getTags = async () => {
		try {
			const data = await api.get("/{userId}/tags", pathParam, reqParam);
			setTags(data.response?.data ?? []);
		} catch (e) {}
	};

	// biome-ignore lint/correctness/useExhaustiveDependencies: TODO
	useEffect(() => {
		getTags();
	}, [searchTerm]);

	return tags;
};
