import React from 'react';
import ReactDOM from 'react-dom/client';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import ErrorPage from "./error-page";
import SignInSide from "./component/signIn/SignInside";
import Paperbase from "./component/chatRoom/paperbase";
import Register from "./component/register/register";
import ContactPage from "./component/contact/contactPage";
import GroupPage from "./component/groupMsg/groupPage";
import PersonalPage from "./component/personalMsg/personalPage";
import ContentPage from "./component/chatRoom/ContentPage";
const router = createBrowserRouter([
    {
        path: "/",
        element: <SignInSide/>,
        errorElement:<ErrorPage/>,
        children: [
            {
                path: "/register",
                element: <Register/>,
            },
        ],
    },
    {
        path:"/chatroom",
        element:<Paperbase/>,
        children:[
            {
                path: "/chatroom/contact",
                element: <ContactPage/>,
            },

            {
                path:"/chatroom/group",
                element: <GroupPage/>,
            },
            {
                path:"/chatroom/personal",
                element: <PersonalPage/>,
            },
            {
                path:"/chatroom/content",
                element: <ContentPage/>
            }
        ],
    },

]);
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);

