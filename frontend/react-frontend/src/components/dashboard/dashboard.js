import { useEffect, useState } from "react";
import { Footer } from "../footer/footer";
import { Navbar } from "../navbar/navbar";
import { Sidebar } from "../sidebar/sidebar";
import { Task } from "../task/task";
import { Error } from "../error/error";

import { axiosInstance } from "../../utils/utils";

export const Dashboard = () => {
    const [isError, setIsError] = useState(false)
    const [errorMessage, setErrorMessage] = useState("")
    const [tasks, setTasks] = useState([])

    useEffect(() => {
        const getAllTodosOnLoad = async () => {
            try {
                setIsError(false)
                const { data: { data } } = await axiosInstance.get("http://localhost:3000/api/todo/get")
                console.log(data)
                setTasks(data)
            } catch (error) {
                console.error(error)
                console.log(error.message)
                setIsError(true)
                setErrorMessage(error.message)
            }
        }

        getAllTodosOnLoad()
    }, [])

    return (
        < div className="h-screen bg-cyan-500" >
            <Navbar />
            <div className="h-full flex">
                <div className="w-full flex">
                    <Sidebar />
                    <div className="w-full mt-20">
                        {isError ? (
                            <Error error={errorMessage} />
                        ) : (
                            <>
                                {/* Check if the array returned for /api/todo/get is not zero (0) - since this would indicate no tasks were returned */}
                                {tasks && tasks.length > 0 ? (
                                    tasks.map((task, i) => (
                                        <Task taskName={task.name} key={i} />
                                    ))
                                )  : "n"}
                            </>
                        )}
                    </div>
                </div>
            </div>
            <Footer />
        </div >
    )
}
