
export interface TagProps {
    backgroundColor?: string;
    label: string;
    onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
}

export const Tag = ({
    backgroundColor = "#1976d2",
    label,
    onClick
}: TagProps) => {
    return (
        <span
            className="gap-1 w-fit whitespace-nowrap flex items-center justify-center min-w-[50px] max-w-[200px] rounded-full py-0 px-3 text-sm/6 font-semibold text-white"
            style={{ backgroundColor }}>
            <span className="truncate">{label}</span>
            {onClick && <button
                className="rounded-full w-5 h-5 hover:bg-gray-100 hover:bg-opacity-50 flex items-center justify-center"
                onClick={onClick}>x</button>}
        </span>
    )
}