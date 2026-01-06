import Home from "./pages/Home";
import Login from "./pages/Login";
import "./App.css";
import React, { useEffect } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { blue, orange } from "@mui/material/colors";

const theme = createTheme({
    palette: {
        primary: blue,
        secondary: orange,
    },
});

const App: React.FC = () => {
    console.log("Frontend starting, waiting for Go...");

    useEffect(() => {
        //for testing connection with backend
        fetch("http://localhost:8000/handshake")
            .then((res) => res.json())
            .then((data) => console.log(data.message))
            .catch((err) => console.log("Connection failed: " + err));
    }, []);

    return (
        <div className="App">
            <ThemeProvider theme={theme}>
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/login" element={<Login />} />
                        {/* <Route path="/register" element={<Register />} />
                        <Route path="/topics" element={<TopicListView />} />
                        <Route path="/topics/:topicID" element={<TopicItemView />} />
                        <Route path="/topics/:topicID/posts" element={<PostListView />} />
                        <Route path="/topics/:topicID/posts/postID" element={<PostItemView />} /> */}
                    </Routes>
                </BrowserRouter>
            </ThemeProvider>
        </div>
    );
};

export default App;
