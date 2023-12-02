import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import SignInSide from "./component/signIn/SignInside";
import Page from "./component/chatRoom/chatRoom";
function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<SignInSide />} />
                <Route path="/chatroom" element={<Page/>} />
            </Routes>
        </Router>
    );
}

export default App;