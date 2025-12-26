import Home from "./pages/Home";
import BasicThreadView from "./pages/BasicThreadView";
import StyledThreadView from "./pages/StyledThreadView";
import React, { useEffect } from "react";
import "./App.css";
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
                        <Route path="/thread/1" element={<BasicThreadView />} />
                        <Route path="/thread/1/styled" element={<StyledThreadView />} />
                        <Route path="/" element={<Home />} />
                    </Routes>
                </BrowserRouter>
            </ThemeProvider>
        </div>
    );
};

export default App;
