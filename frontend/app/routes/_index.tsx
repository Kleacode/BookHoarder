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
				active="register"
				links={[
					{ id: "register", path: "", name: "本の登録" },
					{ id: "check", path: "", name: "積読確認" },
					{ id: "settings", path: "", name: "全体設定" },
				]}
			/>
		</div>
	);
}
