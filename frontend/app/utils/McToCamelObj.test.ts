import { describe, expect, it } from "vitest";
import { mcToCcObj } from "./McToCcObj";

describe("mcToCcObj", () => {
	it("test", () => {
		expect(mcToCcObj({ test: "test" })).toStrictEqual({ test: "test" });
		expect(
			mcToCcObj({
				AaaBbb: "aaaBBB",
				CccDdd: {
					EeeFff: "Drop",
				},
			}),
		).toStrictEqual({
			aaaBbb: "aaaBBB",
			cccDdd: {
				eeeFff: "Drop",
			},
		});
		expect(mcToCcObj({ ArrayTest: [0, 1, 2] })).toStrictEqual({
			arrayTest: [0, 1, 2],
		});
	});
});
