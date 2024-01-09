import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Index } from "../index/index";
import { Dashboard } from "../dashboard/dashboard";

export const RouteContainer = () => (
    <BrowserRouter>
        <Routes>
            <Route path="/" Component={Index} />
            <Route path="/dashboard" Component={Dashboard} />
        </Routes>
    </BrowserRouter>
)