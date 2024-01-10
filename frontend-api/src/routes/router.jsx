import {createBrowserRouter} from "react-router-dom";
import ErrorPage from "../error-page.jsx";
import Login from "../pages/Login/index.jsx";
import Layout from "../pages/Layout/index.jsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Login />,
        errorElement: <ErrorPage />,

    },
    {
        path: "/layout",
        element: <Layout />,
    }
]);

export default router;