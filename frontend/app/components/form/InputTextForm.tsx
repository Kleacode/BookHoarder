
export interface InputTextFormProps {
    label: string
    placeholder?: string
    onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void
};

export const InputTextForm = ({
    label,
    placeholder,
    ...props
}: InputTextFormProps) => {
    return (
        <div>
            <div className="text-gray-700 text-xs font-bold">{label}</div>
            <input className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                type="text"
                placeholder={placeholder}
                {...props}></input>
        </div>
    )
}