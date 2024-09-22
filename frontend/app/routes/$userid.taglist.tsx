import { Header } from "@/components";
import { Button } from "@/components/button/Button";
import { InputTextForm } from "@/components/form/InputTextForm";
import { Tag, type TagProps } from "@/components/tag/Tag";
import { responseGetTags, useGetTags } from "@/hooks/useGetTags";
import type { StatusType } from "@/libs/schemas/schemaHealper";
import { deleteTag } from "@/utils/api/deleteTag";
import { postTag } from "@/utils/api/postTag";
import { useParams } from "@remix-run/react";
import { useEffect, useRef, useState } from "react";

import _ from "lodash";

export interface TableElement {
	title: string;
	tags: TagProps[];
	status: StatusType;
}

export default function Index() {
	const params = useParams();
	const userID = Number.parseInt(params.userid ?? "-1", 10);

	const [tagname, setTagName] = useState<string>("");

	const ref = useRef();
	const tags = useGetTags({ userId: userID }, {}, ref);

	const onClickTag = (tagId: number) => async () => {
		await deleteTag({ userId: userID, tagId: tagId }, {});
		_.invoke(ref, "current");
	};

	const onPostTag = async () => {
		await postTag({ userId: userID }, { name: tagname });
		_.invoke(ref, "current");
	};

	return (
		<div className="font-sans p-4">
			<Header
				active="taglist"
				links={[
					{ id: "booklist", path: "booklist", name: "本の一覧" },
					{ id: "hoarder", path: "hoarder", name: "積読リスト" },
					{ id: "taglist", path: "taglist", name: "タグ一覧" },
					{ id: "register", path: "form", name: "本の登録" },
				]}
			/>
			<InputTextForm
				label="new tag"
				onChange={(e) => {
					setTagName(e.target.value);
				}}
			/>
			<Button label="submit" onClick={onPostTag} />
			<div>
				{tags?.map((e) => {
					return (
						<Tag
							key={e.tagId}
							label={e.name ?? ""}
							onClick={onClickTag(e.tagId)}
						/>
					);
				})}
			</div>
		</div>
	);
}
