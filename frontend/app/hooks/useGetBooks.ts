import { type AxiosAliasWrapper, api } from "@/libs/axios/axios";
import type {
	PathParameters,
	ResponseData,
} from "@/libs/schemas/schemaHealper";
import { useEffect, useState } from "react";

export const useGetBooks = (
	pathParam: PathParameters<"/books", "get">,
	reqParam: AxiosAliasWrapper<"/books", "get">,
) => {
	const [books, setBooks] = useState<ResponseData<"/books", "get">>([]);
	const getBooks = async () => {
		try {
			const data = await api.get("/books", pathParam, reqParam);
			setBooks(data.response?.data ?? []);
		} catch (e) {}
	};

	// biome-ignore lint/correctness/useExhaustiveDependencies: TODO
	useEffect(() => {
		getBooks();
	}, []);

	return books;
};
