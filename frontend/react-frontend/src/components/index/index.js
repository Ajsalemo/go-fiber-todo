import { Link } from "react-router-dom";
import checkboxIcon from "../../assets/images/checkbox.png"

export const Index = () => (
    <div className="bg-cyan-500 flex justify-center w-screen items-center">
        <div className="w-1/2 h-1/4 border-solid border-black rounded shadow-xl bg-cyan-700 justify-center items-center flex flex-col font-sans">
            <div><img src={checkboxIcon} alt="checkbox icon" /></div>
            <h1 className="text-white text-6xl my-6">Todo</h1>
            <button className="rounded-sm bg-cyan-500 px-8 py-2 text-center text-white font-medium hover:bg-cyan-600 ease-in-out duration-300 mb-8">
                <Link to={"/dashboard"}>
                    Continue
                </Link>
            </button>
        </div>
    </div>
)