import { type AxiosAliasWrapper, api } from "@/libs/axios/axios";
import type {
	PathParameters,
	ResponseData,
} from "@/libs/schemas/schemaHealper";
import { useEffect, useState } from "react";

export const useGetHoarderBooks = (
	pathParam: PathParameters<"/{userId}/hoarder", "get">,
	reqParam: AxiosAliasWrapper<"/{userId}/hoarder", "get">,
) => {
	const [hoarderBooks, setHoarderBooks] = useState<
		ResponseData<"/{userId}/hoarder", "get">
	>([]);

	const getHoarderBooks = async () => {
		try {
			const data = await api.get("/{userId}/hoarder", pathParam, reqParam);
			setHoarderBooks(data.response?.data ?? []);
		} catch (e) {}
	};

	useEffect(() => {
		getHoarderBooks();
	}, []);

	return hoarderBooks;
};
