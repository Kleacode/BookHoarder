import { assert, describe, expect, it } from "vitest";
import { ConvertStatusTypeToLabel } from "./ConvertStatusType";

describe("isObject", () => {
	it("ok", () => {
		expect(ConvertStatusTypeToLabel("todo")).toBe("積読");
		expect(ConvertStatusTypeToLabel("wip")).toBe("読書中");
		expect(ConvertStatusTypeToLabel("done")).toBe("読了");
	});
	it("ng", () => {
		assert.throws(
			// biome-ignore lint: ng test case
			() => ConvertStatusTypeToLabel("aaaaaa" as any),
			"invalid status type",
		);
	});
});
