import { useEffect, useState } from "react";
import { Footer } from "../footer/footer";
import { Navbar } from "../navbar/navbar";
import { Sidebar } from "../sidebar/sidebar";
import { Task } from "../task/task";

import axios from "axios"
import { Error } from "../error/error";

export const Dashboard = () => {
    const [isError, setIsError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");

    useEffect(() => {
        const getAllTodosOnLoad = async () => {
            try {
                setIsError(false)
                const { data } = await axios.get("http://localhost:3000/api/todo/get")
                console.log(data);
            } catch (error) {
                console.error(error);
                console.log(error.message);
                setIsError(true)
                setErrorMessage(error.message);
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
                                <Task />
                                <Task />
                            </>
                        )}
                    </div>
                </div>
            </div>
            <Footer />
        </div >
    )
}
