import { Header } from "@/components";
import { Button } from "@/components/button/Button";
import { InputTextForm } from "@/components/form/InputTextForm";
import { Tag, type TagProps } from "@/components/tag/Tag";
import { deleteTag } from "@/hooks/deleteTag";
import { postTag } from "@/hooks/postTag";
import { responseGetTags, useGetTags } from "@/hooks/useGetTags";
import type { StatusType } from "@/libs/schemas/schemaHealper";
import { useParams } from "@remix-run/react";
import { useEffect, useState } from "react";

export interface TableElement {
	title: string;
	tags: TagProps[];
	status: StatusType;
}

export default function Index() {
	const params = useParams();
	const userID = Number.parseInt(params.userid ?? "-1", 10);

	const [tagname, setTagName] = useState<string>("");
	// TODO post, delete等がある場合、fetchはどうやる？
	const tags = useGetTags({ userId: userID }, {});

	const onClickTag = (tagId: number) => {
		deleteTag({ userId: userID, tagId: tagId }, {});
	};

	const onPostTag = () => {
		postTag({ userId: userID }, { name: tagname });
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
						// TODO onClickアンチパターン？
						<Tag
							key={e.tagId}
							label={e.name ?? ""}
							onClick={() => onClickTag(e.tagId)}
						/>
					);
				})}
			</div>
		</div>
	);
}
