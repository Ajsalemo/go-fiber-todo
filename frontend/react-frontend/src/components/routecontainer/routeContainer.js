import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Index } from "../index";

export const RouteContainer = () => (
    <BrowserRouter>
        <Routes>
            <Route path="/" Component={Index} />
            <Route path="/test" Component={Index} />
        </Routes>
    </BrowserRouter>
)