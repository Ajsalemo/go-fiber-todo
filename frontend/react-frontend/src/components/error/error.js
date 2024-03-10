export const Error = ({ error }) => {
    return (
        <div className="h-full flex justify-center items-center">
            <p className="mr-2 font-bold">An error occured while loading tasks | </p>
            <p className="text-red-500">error: {error}</p>
        </div>
    )
}