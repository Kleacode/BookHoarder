import type { UnionToIntersection, Get } from "type-fest";
import type { paths } from "./tsSchema";

export type UrlPaths = keyof paths;

export type HttpMethods = keyof paths[keyof paths];

export type HttpMethodsFilteredByPath<Path extends UrlPaths> = HttpMethods &
	keyof UnionToIntersection<paths[Path]>;

export type RequestParameters<
	Path extends UrlPaths,
	Method extends HttpMethods,
> = Get<paths, `${Path}.${Method}.parameters.query`>;

export type RequestData<
	Path extends UrlPaths,
	Method extends HttpMethods,
> = Get<paths, `${Path}.${Method}.requestBody.content.application/json`>;

export type ResponseData<
	Path extends UrlPaths,
	Method extends HttpMethods,
> = Get<paths, `${Path}.${Method}.responses.200.content.application/json`>;

export type ResponseErrorData<
	Path extends UrlPaths,
	Method extends HttpMethods,
> = Get<
	paths,
	`${Path}.${Method}.responses.default.content.application/problem+json`
>;

export type PathParameters<
	Path extends UrlPaths,
	Method extends HttpMethods,
> = Get<paths, `${Path}.${Method}.parameters.path`>;

/// methodを指定した際に決まるUrlの型
export type UrlPathsFilteredByMethod<Method extends HttpMethods> = {
	[K in UrlPaths]: Method extends keyof paths[K] ? K : never;
}[keyof paths & string];
