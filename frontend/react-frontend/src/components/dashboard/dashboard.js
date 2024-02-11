import { useEffect } from "react";
import { Footer } from "../footer/footer";
import { Navbar } from "../navbar/navbar";
import { Sidebar } from "../sidebar/sidebar";
import { Task } from "../task/task";

export const Dashboard = () => {
    useEffect(() => {
        
    })

    return (
        < div className="h-screen bg-cyan-500" >
            <Navbar />
            <div className="h-full flex">
                <div className="w-full flex">
                    <Sidebar />
                    <div className="w-full mt-20">
                        <Task />
                        <Task />
                    </div>
                </div>
            </div>
            <Footer />
        </div >
    )
}
