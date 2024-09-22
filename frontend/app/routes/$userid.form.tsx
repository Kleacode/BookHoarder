import { Button, Header } from "@/components";
import { DropDownForm } from "@/components/form/DropDownForm";
import { InputTextForm } from "@/components/form/InputTextForm";
import { type SuggestItem, TagForm } from "@/components/form/tagform/TagForm";
import { useGetSuggestTags } from "@/hooks/useGetSuggestTags";
import type { StatusType } from "@/libs/schemas/schemaHealper";
import { ConvertStatusTypeToLabel } from "@/utils/ConvertStatusType";
import { postHoarder } from "@/utils/api/postHoarder";
import { useParams } from "@remix-run/react";
import { useState } from "react";

export default function Index() {
	const [title, setTitle] = useState<string>("");
	const [searchTerm, setSearchTerm] = useState<string>("");
	const params = useParams();
	const userID = Number.parseInt(params.userid ?? "-1", 10);

	// TODO tagform refactor
	const tags = useGetSuggestTags(
		{ userId: userID },
		{ params: { name: searchTerm } },
		searchTerm,
	);
	const suggestItems: SuggestItem[] = tags?.map((e) => ({
		key: e.tagId,
		label: e.name,
	}));
	const postNewHoarder = () => {
		postHoarder(
			{ userId: userID },
			{
				book: { title: title },
				tags: [],
				status: "wip",
			},
		);
	};

	const statusTypes: StatusType[] = ["todo", "wip", "done"];
	const statusOptions = statusTypes.map((e) => ConvertStatusTypeToLabel(e));

	return (
		<div className="font-sans p-4">
			<Header
				active="register"
				links={[
					{ id: "booklist", path: "booklist", name: "本の一覧" },
					{ id: "hoarder", path: "hoarder", name: "積読リスト" },
					{ id: "taglist", path: "taglist", name: "タグ一覧" },
					{ id: "register", path: "form", name: "本の登録" },
				]}
			/>
			<div className="flex flex-col gap-16">
				<InputTextForm
					label="タイトル"
					onChange={(e) => {
						setTitle(e.target.value);
					}}
				/>
				<DropDownForm<string> label="状態" options={statusOptions} />
				<TagForm
					SetSearchTerm={(e) => {
						setSearchTerm(e.target.value);
					}}
					SuggestItems={suggestItems}
				/>
				<Button label="Submit" onClick={postNewHoarder} />

				<div>isbnから自動入力する</div>
				<InputTextForm label="isbnコード" />
				<Button label="コードから自動入力" />
			</div>
		</div>
	);
}
