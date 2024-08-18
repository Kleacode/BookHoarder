import Axios, { type AxiosResponse, type AxiosError } from "axios";
import _ from "lodash";

import { interceptCsrf } from "./interceptors/csrfInterceptor";
import type * as schemaHelper from "@/libs/schemas/schemaHealper";

const API_URL = "";

/// Axiosのインストラクタ用意

Axios.defaults.withCredentials = true;

const axios = Axios.create({
	baseURL: API_URL,
});

/// Axiosのinterceptor設定

// request
axios.interceptors.request.use(interceptCsrf);

/// typeの設定
/// openapi-typescriptで生成した型のschemaをaxiosのrequest関数にあてる

export type AxiosConfigWrapper<
	Path extends schemaHelper.UrlPaths,
	Method extends schemaHelper.HttpMethods,
> = {
	url: Path;
	method: Method & schemaHelper.HttpMethodsFilteredByPath<Path>;
	params?: schemaHelper.RequestParameters<Path, Method>;
	data?: schemaHelper.RequestData<Path, Method>;
};

type Response<
	Path extends schemaHelper.UrlPaths,
	Method extends schemaHelper.HttpMethods,
> = AxiosResponse<schemaHelper.ResponseData<Path, Method>>;

type Error<
	Path extends schemaHelper.UrlPaths,
	Method extends schemaHelper.HttpMethods,
> = AxiosError<schemaHelper.ResponseErrorData<Path, Method>>;

class APIResponse<
	Path extends schemaHelper.UrlPaths,
	Method extends schemaHelper.HttpMethods,
> {
	private _response?: Response<Path, Method>;
	private _error?: Error<Path, Method>;

	constructor(data: {
		response?: Response<Path, Method>;
		error?: Error<Path, Method>;
	}) {
		// responseとerrorの一方のみ存在する場合設定する
		if (
			(data.response !== undefined && data.error !== undefined) ||
			(data.response === undefined && data.error === undefined)
		) {
			throw new Error("responseとerrorは一方のみ存在する必要があります。");
		}
		this._response = data.response;
		this._error = data.error;
	}

	get response() {
		return this._response;
	}

	get error() {
		return this._error;
	}
}

export async function request<
	Path extends schemaHelper.UrlPaths,
	Method extends schemaHelper.HttpMethods,
>(
	config: AxiosConfigWrapper<Path, Method>,
	pathParams?: schemaHelper.PathParameters<Path, Method>,
) {
	const url = _.isEmpty(pathParams)
		? config.url
		: _replacePathParams(config.url, pathParams);
	return axios
		.request<
			schemaHelper.ResponseData<Path, Method>,
			Response<Path, Method>,
			AxiosConfigWrapper<Path, Method>["data"]
		>({ ...config, url })
		.then((response) => {
			return new APIResponse({ response });
		})
		.catch((error: Error<Path, Method>) => {
			// 通信エラー発生時はresponseが返却されないため個別にエラー情報を返却
			console.error(error);
			return new APIResponse({ error });
		});
}

/// baseURL内の`{パスパラメータ}`となっている箇所を与えた値で置き換える
function _replacePathParams<
	Path extends schemaHelper.UrlPaths,
	Method extends schemaHelper.HttpMethods,
>(
	baseURL: string,
	pathParams: schemaHelper.PathParameters<Path, Method>,
): string {
	let url = baseURL;
	for (const [paramKey, paramValue] of Object.entries(pathParams || {})) {
		const paramPlaceholder = `{${paramKey}}`;
		url = url.replace(paramPlaceholder, String(paramValue));
	}
	return url;
}
