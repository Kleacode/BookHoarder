import Axios, {
	type AxiosResponse,
	type AxiosError,
	type AxiosInstance,
} from "axios";
import _ from "lodash";

import { interceptCsrf } from "./interceptors/csrfInterceptor";
import type * as SchemaHelper from "@/libs/schemas/schemaHealper";

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

/// Method関連の型
export type Get = "get";
export type Delete = "delete";
export type Post = "post";
export type Put = "put";

export type AxiosConfigWrapper<
	Path extends SchemaHelper.UrlPaths,
	Method extends SchemaHelper.HttpMethods,
> = {
	url: Path;
	method: Method & SchemaHelper.HttpMethodsFilteredByPath<Path>;
	params?: SchemaHelper.RequestParameters<Path, Method>;
	data?: SchemaHelper.RequestData<Path, Method>;
};

type Response<
	Path extends SchemaHelper.UrlPaths,
	Method extends SchemaHelper.HttpMethods,
> = AxiosResponse<SchemaHelper.ResponseData<Path, Method>>;

type Error<
	Path extends SchemaHelper.UrlPaths,
	Method extends SchemaHelper.HttpMethods,
> = AxiosError<SchemaHelper.ResponseErrorData<Path, Method>>;

/// get, delete, post, put などの alias 関数で利用する型
/// Method トリガーで Url を変えるため AxiosConfigWrapper とは別で用意する
export type AxiosAliasWrapper<
	Path extends SchemaHelper.UrlPaths,
	Method extends SchemaHelper.HttpMethods,
> = {
	params?: SchemaHelper.RequestParameters<Path, Method>;
	data?: SchemaHelper.RequestData<Path, Method>;
};

/// Responseを受け取るクラス

export class APIResponse<
	Paths,
	Path extends SchemaHelper.UrlPaths,
	Method extends SchemaHelper.HttpMethods,
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

class Api {
	private static instance: Api | null = null;
	private axios: AxiosInstance;

	private constructor() {
		this.axios = Axios.create({
			baseURL: API_URL,
		});
		// request
		this.axios.interceptors.request.use(interceptCsrf);
	}

	public static getInstance(): Api {
		if (!Api.instance) {
			Api.instance = new Api();
		}
		return Api.instance;
	}

	async request<
		Path extends SchemaHelper.UrlPaths,
		Method extends SchemaHelper.HttpMethods,
	>(
		config: AxiosConfigWrapper<Path, Method>,
		pathParams: SchemaHelper.PathParameters<Path, Method>,
	) {
		const url = this.#replacePathParams(config.url, pathParams);
		return this.axios
			.request<
				SchemaHelper.ResponseData<Path, Method>,
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
	#replacePathParams<
		Path extends SchemaHelper.UrlPaths,
		Method extends SchemaHelper.HttpMethods,
	>(
		baseURL: string,
		pathParams: SchemaHelper.PathParameters<Path, Method>,
	): string {
		if (_.isEmpty(pathParams)) return baseURL;
		let url = baseURL;
		for (const [paramKey, paramValue] of Object.entries(pathParams)) {
			const paramPlaceholder = `{${paramKey}}`;
			url = url.replace(paramPlaceholder, String(paramValue));
		}
		return url;
	}

	/// Alias
	/// Axiosの方ではprototypeを使って定義しているが、Genericsが必要な関係上、各aliasをベタ書きしている
	async get<Path extends SchemaHelper.UrlPathsFilteredByMethod<Get>>(
		url: Path,
		pathParams: SchemaHelper.PathParameters<Path, Get>,
		config?: AxiosAliasWrapper<Path, Get>,
	) {
		return this.request<Path, Get>(
			{ ...config, url, method: "get", data: config?.data },
			pathParams,
		);
	}

	async delete<Path extends SchemaHelper.UrlPathsFilteredByMethod<Delete>>(
		url: Path,
		pathParams: SchemaHelper.PathParameters<Path, Delete>,
		config?: AxiosAliasWrapper<Path, Delete>,
	) {
		return this.request<Path, Delete>(
			{ ...config, url, method: "delete", data: config?.data },
			pathParams,
		);
	}

	async post<Path extends SchemaHelper.UrlPathsFilteredByMethod<Post>>(
		url: Path,
		pathParams: SchemaHelper.PathParameters<Path, Post>,
		data: SchemaHelper.RequestData<Path, Post>,
		config?: AxiosAliasWrapper<Path, Post>,
	) {
		return this.request<Path, Post>(
			{ ...config, url, method: "post", data },
			pathParams,
		);
	}

	async put<Path extends SchemaHelper.UrlPathsFilteredByMethod<Put>>(
		url: Path,
		pathParams: SchemaHelper.PathParameters<Path, Put>,
		data: SchemaHelper.RequestData<Path, Put>,
		config?: AxiosAliasWrapper<Path, Put>,
	) {
		return this.request<Path, Put>(
			{ ...config, url, method: "put", data },
			pathParams,
		);
	}
}

export const api = Api.getInstance();
