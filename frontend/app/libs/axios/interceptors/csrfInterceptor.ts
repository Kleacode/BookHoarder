import type { InternalAxiosRequestConfig } from "axios";

// biome-ignore lint/suspicious/noExplicitAny: <explanation>
export const interceptCsrf = (config: InternalAxiosRequestConfig<any>) => {
	if (!sessionStorage.getItem("x-csrf-token")) return config;
	if (["post", "put", "patch", "delete"].includes(config.method ?? "")) {
		config.headers["X-CSRF-Token"] = sessionStorage.getItem("x-csrf-token");
	}
	return config;
};
