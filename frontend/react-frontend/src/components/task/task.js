import { useState } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCircleCheck } from '@fortawesome/free-regular-svg-icons'
import { faCircleCheck as faSolidCircleCheck, faTrash } from '@fortawesome/free-solid-svg-icons'
import { axiosInstance } from "../../utils/utils";

export const Task = ({ taskName, id, completed, loading, getAllTodosOnLoad }) => {
    const [isError, setIsError] = useState(false)
    const [showDeleteError, setShowDeleteError] = useState(false)
    const [isLoading, setIsLoading] = useState(false)

    const deleteTask = async (id) => {
        setIsLoading(true)
        try {
            console.log(`id:  ${id}`);
            setIsError(false)
            setShowDeleteError(false)
            const data = await axiosInstance.delete(`http://localhost:3000/api/todo/delete/${id}`)
            console.log(data)
            // Reload data after an operation
            if (data.status === 204) {
                getAllTodosOnLoad()
            }
            setIsLoading(true)
        } catch (error) {
            console.error(error)
            console.log(error.message)
            setShowDeleteError(true)
            setIsLoading(false)
            setIsError(true)
        }
    }

    return (
        <>
            <div className="bg-slate-800 w-4/5 min-h-16 rounded-md p-4 mt-8 mx-8 shadow-2xl flex flex-row">
                <div className="flex w-full">
                    <div>
                        <button disabled={loading || isLoading ? true : false}>
                            <FontAwesomeIcon icon={completed ? faSolidCircleCheck : faCircleCheck} className={completed ? "hover:text-blue-300 text-white mr-2" : "hover:text-green-300 text-white mr-2"} />
                        </button>
                    </div>
                    <div className="break-all text-white">
                        {taskName}
                    </div>
                </div>
                <div className="flex self-center">
                    <button disabled={loading || isLoading? true : false} onClick={() => deleteTask(id)}>
                        <FontAwesomeIcon icon={faTrash} className="text-white mr-2 hover:text-red-300" />
                    </button>
                    {isError && (
                        <div className={showDeleteError ? "bg-slate-800 flex flex-col text-red-300 absolute mt-7 p-4 rounded-md border-solid border-2 border-sky-500 items-end" : "hidden"}>
                            <button className="text-white w-fit" onClick={() => setShowDeleteError(false)}>x</button>
                            <span>An error occurred while deleting a task</span>
                        </div>
                    )}
                </div>
            </div>
        </>
    )
}