export const Header = () => {
	return  <header className='flex py-2 px-4 sm:px-10 bg-white font-[sans-serif] min-h-[70px] tracking-wide relative z-50'>
				<div className='flex flex-wrap items-center justify-between gap-5 w-full'>
				<a href="javascript:void(0)"><img src="book_icon.svg" alt="logo" className='w-12' />
				</a>
			
				<div id="collapseMenu"
					className='max-lg:hidden lg:!block max-lg:before:fixed max-lg:before:bg-black max-lg:before:opacity-50 max-lg:before:inset-0 max-lg:before:z-50'>
			
					<ul
					className='lg:flex gap-x-5 max-lg:space-y-3 max-lg:fixed max-lg:bg-white max-lg:w-1/2 max-lg:min-w-[300px] max-lg:top-0 max-lg:left-0 max-lg:p-6 max-lg:h-full max-lg:shadow-md max-lg:overflow-auto z-50'>
					<li className='max-lg:border-b border-gray-300 max-lg:py-3 px-3'>
						<a href='javascript:void(0)'
						className='hover:text-[#007bff] text-[#007bff] block font-semibold text-[15px]'>本の登録</a>
					</li>
					<li className='max-lg:border-b border-gray-300 max-lg:py-3 px-3'><a href='javascript:void(0)'
						className='hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]'>積読確認</a>
					</li>
					<li className='max-lg:border-b border-gray-300 max-lg:py-3 px-3'><a href='javascript:void(0)'
						className='hover:text-[#007bff] text-gray-500 block font-semibold text-[15px]'>設定</a>
					</li>
					</ul>
				</div>
			
				<div className='flex max-lg:ml-auto space-x-4'>
					<button
					className='px-4 py-2 text-sm rounded-lg font-bold text-white border-2 border-[#007bff] bg-[#007bff] transition-all ease-in-out duration-300 hover:bg-transparent hover:text-[#007bff]'>ログイン</button>
					<button
					className='px-4 py-2 text-sm rounded-lg font-bold text-white border-2 border-[#007bff] bg-[#007bff] transition-all ease-in-out duration-300 hover:bg-transparent hover:text-[#007bff]'>ユーザー登録</button>
				</div>
				</div>
			</header>;
};
