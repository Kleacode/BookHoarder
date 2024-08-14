import _ from "lodash"

export type Obj = { [x: string]: unknown };

export const isObject = (value: unknown): value is object => {
	const type = typeof value;
	return value !== null && (type === "object" || type === "function") && !isArray(value);
};

export const isArray =  (value: unknown): value is Array<unknown> => {
	return _.isArray(value)
};