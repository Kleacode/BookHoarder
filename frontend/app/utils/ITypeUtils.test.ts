import { describe, expect, it } from "vitest";
import { isObject } from "./ITypeUtils";

describe("isObject", () => {
	it("ok", () => {
		expect(isObject({ test: "test" })).toBe(true);
		expect(
			isObject({
				AaaBbb: "aaaBBB",
				CccDdd: {
					EeeFff: "Drop",
				},
			}),
		).toBe(true);
		expect(isObject({ ArrayTest: [0, 1, 2] })).toBe(true);
		expect(isObject({})).toBe(true);
	});

	it("ng", () => {
		expect(isObject(1)).toBe(false);
		expect(isObject(0)).toBe(false);
		expect(isObject(true)).toBe(false);
		expect(isObject(false)).toBe(false);
		expect(isObject("string")).toBe(false);
		expect(isObject("")).toBe(false);
		expect(isObject([])).toBe(false);
		expect(isObject([1, "string"])).toBe(false);
	});
});
