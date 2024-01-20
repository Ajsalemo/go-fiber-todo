import { Footer } from "../footer/footer";
import { Navbar } from "../navbar/navbar";
import { Sidebar } from "../sidebar/sidebar";
import { Task } from "../task/task";

export const Dashboard = () => (
    <div className="h-screen bg-cyan-500">
        <Navbar />
        <div className="h-full flex">
            <div className="w-full flex">
                <Sidebar />
                <div className="w-full">
                    <Task />
                    <Task />

                </div>
            </div>
        </div>
        <Footer />
    </div>
)