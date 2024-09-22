import { type AxiosAliasWrapper, api } from "@/libs/axios/axios";
import type {
	PathParameters,
	ResponseData,
} from "@/libs/schemas/schemaHealper";
import { useEffect, useState } from "react";

export type responseGetTags = ResponseData<"/{userId}/tags", "get">;

export const useGetTags = (
	pathParam: PathParameters<"/{userId}/tags", "get">,
	reqParam: AxiosAliasWrapper<"/{userId}/tags", "get">,
	// biome-ignore lint: TODO 型情報直す
	ref: any,
) => {
	const [tags, setTags] = useState<responseGetTags>([]);

	const getTags = async () => {
		try {
			const data = await api.get("/{userId}/tags", pathParam, reqParam);
			setTags(data.response?.data ?? []);
		} catch (e) {}
	};

	ref.current = getTags;

	// biome-ignore lint/correctness/useExhaustiveDependencies(getTags): TODO
	useEffect(() => {
		getTags();
	}, []);

	return tags;
};
