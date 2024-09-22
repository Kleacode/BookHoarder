import { Header } from "@/components/Header/Header";
import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
	return [
		{ title: "New Remix App" },
		{ name: "description", content: "Welcome to Remix!" },
	];
};

export default function Index() {
	return (
		<div className="font-sans p-4">
			<Header
				active="booklist"
				links={[
					{ id: "booklist", path: "1/booklist", name: "本の一覧" },
					{ id: "hoarder", path: "1/hoarder", name: "積読リスト" },
					{ id: "taglist", path: "1/taglist", name: "タグ一覧" },
					{ id: "register", path: "1/form", name: "本の登録" },
				]}
			/>
		</div>
	);
}
