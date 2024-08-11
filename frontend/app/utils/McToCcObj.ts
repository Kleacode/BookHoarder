import _ from "lodash";
import type { Obj } from "./ITypeUtils";

/// オブジェクトのkeyをMixedCapsからCamelCaseに置き換える
export const mcToCcObj = (obj: Obj): Obj | undefined => {
	if (_.isEmpty(obj)) return obj;
	const ccObj: Obj = {};
	for (const [key, value] of Object.entries(obj)) {
		if (_.isArray(value)) ccObj[pascalToCamel(key)] = value;
		else if (_.isPlainObject(value))
			ccObj[pascalToCamel(key)] = mcToCcObj(value as unknown as Obj);
		else ccObj[pascalToCamel(key)] = value;
	}
	return ccObj;
};

/// pacal caseの文字列をcamel caseに変換する
const pascalToCamel = _.camelCase;
