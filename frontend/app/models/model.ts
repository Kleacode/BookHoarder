export class Tag {
	id: number;
	name: string;

	private constructor(id: number, name: string) {
		this.id = id;
		this.name = name;
	}

	isEqual(target: Tag): boolean {
		return this.id === target.id;
	}

	static fromJson(obj: { id: number; name: string }) {
		return new Tag(obj.id, obj.name);
	}
}
